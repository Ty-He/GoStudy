package sort

import (
  "sync"
  "runtime"
)

type task_t func()

type parallelSystem struct {
  wg sync.WaitGroup
  taskq chan task_t
}

// n is chan buffer size
func newParallelSystem(n int) *parallelSystem {
  p := new(parallelSystem)
  p.taskq = make(chan task_t, n)
  return p
}

func (p *parallelSystem) putTask(task task_t) {
  p.wg.Add(1)
  p.taskq <- task
}

func (p *parallelSystem) doTasksSync() {
  for task := range p.taskq {
    task()
    p.wg.Done()
  }
}

func (p *parallelSystem) doTasksAsync() {
  cpuCores := runtime.NumCPU() - 1
  for i := 0; i < cpuCores; i++ {
    go p.doTasksSync()
  } // for go
}

func (p *parallelSystem) closeTaskq() {
  close(p.taskq)
}

