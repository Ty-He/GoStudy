package solutions

// true is locking, and false is unlocking
func LockGate(n int) []bool {
  gates := make([]bool, n + 1)
  for i := 1; i <= n; i++ {
    for k := 1; k * i <= n; k++ {
      gates[k * i] = !gates[k * i]
    }
  }
  return gates[1:]
}
