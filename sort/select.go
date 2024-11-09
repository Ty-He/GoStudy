package sort 

import (
  "cmp"
)

func selection[Tp cmp.Ordered](s []Tp) {
  n := len(s)
  for i := 0; i < n; i++ {
    k := i // k \in [i, n-1] && s[k] <= s[j]
    for j := i + 1; j < n; j++ {
      if s[j] < s[i] {
        k = j
      }
    }
    if k != i {
      s[i], s[k] = s[k], s[i]
    }
  }
}

func NewSelection[Tp cmp.Ordered]() SortStrategy[Tp] {
  return sortStrategy[Tp](selection[Tp])
} 
