package sort

import "cmp"

func heapSort[Tp cmp.Ordered](s []Tp) {
  n := len(s)
  for i := n / 2 - 1; i >= 0; i-- {
    make_heap(s, i, n)
  }

  for i := n - 1; i > 0; i-- {
    s[0], s[i] = s[i], s[0]
    make_heap(s, 0, i)
  }
}

// max heap
func make_heap[Tp cmp.Ordered](s []Tp, begin, end int) {
  left, right := (begin << 1) + 1, (begin << 1) + 2
  root := begin

  if left < end && s[left] > s[root] {
    root = left
  }

  if right < end && s[right] > s[root] {
    root = right
  }

  if root != begin {
    s[root], s[begin] = s[begin], s[root]
    make_heap(s, root, end)
  }
} 

func NewHeap[Tp cmp.Ordered]() SortStrategy[Tp] {
  return sortStrategy[Tp](heapSort[Tp])
}
