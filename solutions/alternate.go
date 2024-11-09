package solutions

func AlternateBlackAndWhite(n int) (cnt int) {
  list := make([]bool, n << 1)
  for i := range list {
    // true false true false ...
    list[i] = i & 1 == 0
  }

  for i := 0; i < n; i++ {
    if list[i] {
      continue
    }

    for j := 1; j < n; j++ {
      // find next true
      if list[i + j] {
        cnt += j // swap j times
        list[i], list[i + j] = list[i + j], list[i]
        break
      }
    }
  }

  return
}
