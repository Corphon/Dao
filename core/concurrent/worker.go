package concurrent

import (
    "context"
    "sync"
)

// Task 定义任务结构
type Task struct {
    ID string
    Fn func() error
}

// WorkerPool 定义工作池结构
type WorkerPool struct {
    workers int
    tasks   chan Task
    wg      sync.WaitGroup
    ctx     context.Context
    cancel  context.CancelFunc
}

// NewWorkerPool 创建新的工作池
func NewWorkerPool(workers int) *WorkerPool {
    ctx, cancel := context.WithCancel(context.Background())
    pool := &WorkerPool{
        workers: workers,
        tasks:   make(chan Task, workers*2),
        ctx:     ctx,
        cancel:  cancel,
    }
    pool.Start()
    return pool
}

// Start 启动工作池
func (p *WorkerPool) Start() {
    for i := 0; i < p.workers; i++ {
        go p.worker()
    }
}

// worker 工作协程
func (p *WorkerPool) worker() {
    for {
        select {
        case <-p.ctx.Done():
            return
        case task, ok := <-p.tasks:
            if !ok {
                return
            }
            if err := task.Fn(); err != nil {
                // 错误处理可以根据需要扩展
                continue
            }
            p.wg.Done()
        }
    }
}

// AddTask 添加任务到工作池
func (p *WorkerPool) Submit(task Task) error {
    p.wg.Add(1)
    select {
    case <-p.ctx.Done():
        p.wg.Done()
        return p.ctx.Err()
    case p.tasks <- task:
        return nil
    }
}

// Wait 等待所有任务完成
func (p *WorkerPool) Wait() {
    p.wg.Wait()
}

// Stop 停止工作池
func (p *WorkerPool) Stop() {
    p.cancel()
    close(p.tasks)
}
