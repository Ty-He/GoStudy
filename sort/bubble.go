package sort

import (
  "cmp"
)

// using operator<
func bubble[Tp cmp.Ordered](s []Tp) {
  n := len(s)
  for i := 0; i < n - 1; i++ {
    isSorted := true
    for j := 0; j < n - 1 - i; j++ {
      if s[j] > s[j + 1] {
        s[j], s[j + 1] = s[j + 1], s[j]
        isSorted = false
      }
    }
    if isSorted {
      return
    }
  } // for i
}

// default bubble strategy
func NewBubble[Tp cmp.Ordered]() SortStrategy[Tp] {
  return sortStrategy[Tp](bubble[Tp])
} 
