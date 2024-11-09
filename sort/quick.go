package sort

import (
  "cmp"
)

func quick[Tp cmp.Ordered](s []Tp, first, last int) {
  if first >= last {
    return
  }
  // partition
  pivot := s[last]
  // if var in s[:i], var <= pivot
  i := first
  for j := first; j < last; j++ {
    if s[j] <= pivot {
      s[i], s[j] = s[j], s[i]
      i ++
    } 
  }
  s[i], s[last] = s[last], s[i]

  quick(s, first, i - 1)
  quick(s, i + 1, last)
}

func quickSort[Tp cmp.Ordered](s []Tp) {
  quick(s, 0, len(s) - 1)
}

func NewQuick[Tp cmp.Ordered]() SortStrategy[Tp] {
  return sortStrategy[Tp](quickSort[Tp])
}

func parallelQuick[Tp cmp.Ordered](p *parallelSystem, s []Tp, first, last int) {
  if first >= last {
    return
  }
  // partition
  pivot := s[last]
  // if var in s[:i], var <= pivot
  i := first
  for j := first; j < last; j++ {
    if s[j] <= pivot {
      s[i], s[j] = s[j], s[i]
      i ++
    } 
  }
  s[i], s[last] = s[last], s[i]

  if p != nil {
    p.putTask(func() {
      parallelQuick(nil, s, first, i - 1)
    })
  } else {
    parallelQuick(nil, s, first, i - 1)
  }
  parallelQuick(nil, s, i + 1, last)
  // TODO: using context to fix
  // p.putTask(func() {
  //   parallelQuick(p, s, first, i - 1)
  // })
  // p.putTask(func() {
  //   parallelQuick(p, s, i + 1, last)
  // })
  // p.doTasksSync()
}

func parallelQuickSort[Tp cmp.Ordered](s []Tp) {
  p := newParallelSystem(5)
  p.doTasksAsync()
  parallelQuick(p, s, 0, len(s) - 1)
  p.closeTaskq()
  p.wg.Wait()
}

func NewParallelQuick[Tp cmp.Ordered]() SortStrategy[Tp] {
  return sortStrategy[Tp](parallelQuickSort[Tp])
}
