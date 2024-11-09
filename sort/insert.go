package sort 

import "cmp"

func insertion[Tp cmp.Ordered](s []Tp) {
  n := len(s)
  for i := 1; i < n; i++ {
    for j := i; j > 0; j-- {
      if (s[j] >= s[j - 1]) {
        break
      } else {
        s[j], s[j - 1] = s[j - 1], s[j]
      }
    }
  }
}

func NewInsertion[Tp cmp.Ordered]() SortStrategy[Tp] {
  return sortStrategy[Tp](insertion[Tp])
}
