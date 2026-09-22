package jobmgr

import (
	"context"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-meridian/logger"
)

// PanicHandler panic 堆栈处理函数，使用方注入
// 参数为 panic 堆栈字符串
type PanicHandler func(stack string)

// Config JobManager 配置
type Config struct {
	PanicHandler PanicHandler // panic 处理，可选（nil 则使用 logger 打日志，logger 未初始化则 log.Printf 兜底）
}

var (
	mgr  *JobManager
	once sync.Once
)

// JobManager 协程任务管理器
type JobManager struct {
	ctxStop      context.Context
	cancel       context.CancelFunc
	wg           *sync.WaitGroup
	stopOnce     sync.Once
	panicHandler PanicHandler
	running      atomic.Int64
}

// Init 初始化全局 JobManager（单例，仅首次调用生效）
// cfg 为 nil 时使用默认配置
func Init(cfg *Config) *JobManager {
	once.Do(func() {
		ctx, cancel := context.WithCancel(context.Background())
		var ph PanicHandler
		if cfg != nil {
			ph = cfg.PanicHandler
		}
		mgr = &JobManager{
			ctxStop:      ctx,
			cancel:       cancel,
			wg:           new(sync.WaitGroup),
			panicHandler: ph,
		}
	})
	return mgr
}

// Mgr 获取全局 JobManager 实例（需先调用 Init）
func Mgr() *JobManager {
	return mgr
}

// AddJob 启动 goroutine 执行任务，带 panic 恢复
func (m *JobManager) AddJob(job func()) {
	m.launch(func() {
		defer func() {
			if e := recover(); e != nil {
				stack := captureStack()
				if m.panicHandler != nil {
					m.panicHandler(stack)
				} else if l := logger.Get(); l != nil {
					l.Error("[jobmgr] panic recovered", logger.String("stack", stack))
				} else {
					log.Printf("[jobmgr] panic recovered:\n%s", stack)
				}
			}
		}()
		job()
	})
}

// AddJobNaked 启动 goroutine 执行任务，不捕获 panic（panic 会导致进程崩溃）
func (m *JobManager) AddJobNaked(job func()) {
	m.launch(job)
}

// launch 启动 goroutine，统一管理 WaitGroup 和运行计数
func (m *JobManager) launch(job func()) {
	m.wg.Add(1)
	m.running.Add(1)

	go func() {
		defer m.wg.Done()
		defer m.running.Add(-1)
		job()
	}()
}

// Running 返回当前运行中的任务数量
func (m *JobManager) Running() int64 {
	return m.running.Load()
}

// StopEvent 返回停止信号 channel，供长驻任务监听
func (m *JobManager) StopEvent() <-chan struct{} {
	return m.ctxStop.Done()
}

// StopAll 通知所有任务停止并等待完成（带超时），多次调用安全
func (m *JobManager) StopAll(timeout time.Duration) {
	m.stopOnce.Do(func() {
		m.cancel()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		done := make(chan struct{})
		go func() {
			m.wg.Wait()
			close(done)
		}()

		select {
		case <-done:
		case <-ctx.Done():
			if l := logger.Get(); l != nil {
				l.Error("[jobmgr] stop timeout",
					logger.Duration("timeout", timeout),
					logger.Int64("remaining", m.Running()))
			}
		}
	})
}

// captureStack 获取 panic 堆栈
func captureStack() string {
	buf := make([]byte, 64*1024)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}
