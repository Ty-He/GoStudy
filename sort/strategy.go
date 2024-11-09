package sort 

import "cmp"

// abstract sorter api
type Sorter[T any] struct {
  s SortStrategy[T]
}

func (st *Sorter[T]) Sort(s []T) {
  st.s.Sort(s)
} 

func (st *Sorter[T]) SetStrategy(strategy SortStrategy[T]) {
  st.s = strategy
}

// abstract sort strategy
type SortStrategy[T any] interface {
  Sort(s []T)
}

// sort strategy for cmp.Ordered
type sortStrategy[T cmp.Ordered] func(s []T)

func (sorts sortStrategy[T]) Sort(s []T) {
  sorts(s)
}

// TODO sort for any
type sortStrategyForAny[T any] struct{}
