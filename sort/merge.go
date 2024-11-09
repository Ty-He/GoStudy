package sort

import (
  "cmp"
  "runtime"
)

func mergeSort[Tp cmp.Ordered](s []Tp) {
  divide(s)
}


func divide[Tp cmp.Ordered](s []Tp) []Tp {
  n := len(s)
  if n <= 1 {
    return s
  }

  mid := n / 2
  left := divide(s[:mid])
  right := divide(s[mid:])
  return merge(s, left, right)
}

func merge[Tp cmp.Ordered](dst, left, right []Tp) []Tp {
  m, n := len(left), len(right)
  var i, j, k int
  for i < m && j < n {
    if left[i] < right[j] {
      dst[k] = left[i]
      i++
    } else {
      dst[k] = right[j]
      j++
    }
    k++
  }
  for i < m {
    dst[k] = left[i]
    k++
    i++
  }
  for j < n {
    dst[k] = right[j]
    k++
    j++
  }
  return dst
}


func NewMerge[Tp cmp.Ordered]() SortStrategy[Tp] {
  return sortStrategy[Tp](mergeSort[Tp])
}

type parallelMerge[Tp cmp.Ordered] struct {
  n int
  ps *parallelSystem
}

func (p *parallelMerge[Tp]) divide(s []Tp) []Tp {
  n := len(s)
  if n <= 1 {
    return s
  }

  mid := n / 2
  left := s[:mid]
  var ch chan []Tp
  if p.n > 0 {
    p.n --
    ch = make(chan []Tp)
    p.ps.putTask(func() {
      ch <- p.divide(left)
    })
  } else {
    left = p.divide(left)
  }

  right := p.divide(s[mid:])
  if left == nil {
    p.ps.closeTaskq()
    p.ps.wg.Wait()
    left = <-ch
    close(ch)
  }

  return merge(s, left, right)
}

func parallelMergeSort[Tp cmp.Ordered](s []Tp) {
  p := &parallelMerge[Tp]{
    n: runtime.NumCPU() + 1,
    ps: newParallelSystem(runtime.NumCPU() - 1),
  }
  p.ps.doTasksAsync()
  p.divide(s)
}

func NewParallelMerge[Tp cmp.Ordered]() SortStrategy[Tp] {
  return sortStrategy[Tp](parallelMergeSort[Tp])
}
