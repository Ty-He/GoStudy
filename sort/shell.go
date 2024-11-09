package sort

import (
  "cmp"
  "sync"
)

func shell[Tp cmp.Ordered](s []Tp) {
  n := len(s)
  for gap := n / 2; gap > 0; gap /= 2 {
    for i := 0; i < gap; i++ { // every group
      // sort a group
      for j := i + gap; j < n; j += gap {
        for k := j; k >= i + gap; k -= gap {
          if s[k] >= s[k - gap] {
            break
          } else {
            s[k], s[k - gap] = s[k - gap], s[k]
          }
        }
      } // sort a group
    } // for i group
  } // for gap
}

func NewShell[Tp cmp.Ordered]() SortStrategy[Tp] {
  return sortStrategy[Tp](shell[Tp])
}

func parallelShell[Tp cmp.Ordered](s []Tp) {
  n := len(s)
  for gap := n / 2; gap > 0; gap /= 2 {
    p := newParallelSystem(gap)
    p.doTasksAsync()
    for t := 0; t < gap; t++ { // every group
      i := t
      sortGroup := func() {
        // sort a group
        for j := i + gap; j < n; j += gap {
          for k := j; k >= i + gap; k -= gap {
            if s[k] >= s[k - gap] {
              break
            } else {
              s[k], s[k - gap] = s[k - gap], s[k]
            }
          }
        } // sort a group
      } // func()
      p.putTask(sortGroup)
    } // for i group
    p.closeTaskq()
    p.doTasksSync()
    p.wg.Wait()
  } // for gap
}

func NewParallelShell[Tp cmp.Ordered]() SortStrategy[Tp] {
  return sortStrategy[Tp](parallelShell[Tp])
}

// bad func, should not be called
func parallelEmplaceShell[Tp cmp.Ordered](s []Tp) {
  n := len(s)
  var wg sync.WaitGroup
  for gap := n / 2; gap > 0; gap /= 2 {
    wg.Add(gap)
    for t := gap - 1; t >= 0; t-- { // every group
      i := t
      sortGroup := func() {
        // sort a group
        for j := i + gap; j < n; j += gap {
          for k := j; k >= i + gap; k -= gap {
            if s[k] >= s[k - gap] {
              break
            } else {
              s[k], s[k - gap] = s[k - gap], s[k]
            }
          }
        } // sort a group
        wg.Done()
      } // func()
      if t > 0 {
        go sortGroup()
      } else {
        sortGroup()
      }
    } // for i group
    wg.Wait()
  } // for gap
}

func NewParallelEmplaceShell[Tp cmp.Ordered]() SortStrategy[Tp] {
  return sortStrategy[Tp](parallelEmplaceShell[Tp])
}
