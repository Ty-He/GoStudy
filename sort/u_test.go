package sort

import (
  "testing"
  "math/rand"
  "time"
  "fmt"
  "runtime"
)

const (
  MAX_VALUE = 1 << 30
  N = 10000000
  FRONT_SIZE = 50
)

func randomSlice() []int {
  source := rand.NewSource(time.Now().UnixNano())
  r := rand.New(source)
  s := make([]int, N)
  for i := range s {
    s[i] = r.Intn(MAX_VALUE)
  }
  return s
}

func printFront(s []int) {
  m := min(FRONT_SIZE, len(s))
  for i := 0; i < m; i++ {
    fmt.Printf("%d ", s[i])
  }
  fmt.Println("...")
} 

func isOrdered(s []int) bool {
  n := len(s)
  for i := 1; i < n; i++ {
    if s[i - 1] > s[i] {
      return false
    }
  }
  return true
}

func sortRandomtByStartegy(strategy SortStrategy[int]) bool {
  fmt.Println("CPU Cores = ", runtime.NumCPU())
  // s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
  s := randomSlice()
  fmt.Println("Before Sort: ->")
  printFront(s)
  sorter := &Sorter[int]{}
  sorter.SetStrategy(strategy)
  sorter.Sort(s)
  fmt.Println("After Sort: ->")
  printFront(s)
  return isOrdered(s)
}

func TestSort(t *testing.T) {
  // t.Run("bubble", func(t *testing.T) {
  //   if !sortRandomtByStartegy(NewBubble[int]()) {
  //     t.Errorf("sort failed!!")
  //   }
  // })
  // t.Run("selection", func(t *testing.T) {
  //   if !sortRandomtByStartegy(NewSelection[int]()) {
  //     t.Errorf("sort failed!!")
  //   }
  // })
  // t.Run("insertion", func(t *testing.T) {
  //   if !sortRandomtByStartegy(NewInsertion[int]()) {
  //     t.Errorf("sort failed!!")
  //   }
  // })
  // shell sort
  t.Run("shell1: ", func(t *testing.T) {
    if !sortRandomtByStartegy(NewShell[int]()) {
      t.Errorf("sort failed!!")
    }
  })
  t.Run("shell2: ", func(t *testing.T) {
    if !sortRandomtByStartegy(NewParallelShell[int]()) {
      t.Errorf("sort failed!!")
    }
  })
  // t.Run("shell3: ", func(t *testing.T) {
  //   if !sortRandomtByStartegy(NewParallelEmplaceShell[int]()) {
  //     t.Errorf("sort failed!!")
  //   }
  // })

  // merge sort
  t.Run("merge1", func(t *testing.T) {
    if !sortRandomtByStartegy(NewMerge[int]()) {
      t.Errorf("sort failed")
    }
  })
  t.Run("merge2", func(t *testing.T) {
    if !sortRandomtByStartegy(NewParallelMerge[int]()) {
      t.Errorf("sort failed")
    }
  })

  // quick sort
  t.Run("quick1", func(t *testing.T) {
    if !sortRandomtByStartegy(NewQuick[int]()) {
      t.Errorf("sort failed")
    }
  })
  t.Run("quick2", func(t *testing.T) {
    if !sortRandomtByStartegy(NewParallelQuick[int]()) {
      t.Errorf("sort failed")
    }
  })

  // heap sort
  t.Run("heap1", func(t *testing.T) {
    if !sortRandomtByStartegy(NewHeap[int]()) {
      t.Errorf("sort failed")
    }
  })
} 
