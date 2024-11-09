package solutions

import (
  "testing"
  "fmt"
)

func Test3Kettle(t *testing.T) {
  ThreeKettlePuzzle()
}

func TestAlternate(t *testing.T) {
  for i := 0; i < 10; i++ {
    fmt.Printf("when n = %d, need swap %d times.\n", i, AlternateBlackAndWhite(i))
  }
}

func TestLockGate(t *testing.T) {
  for i := 1; i <= 10; i++ {
    fmt.Println("When gates size = ", i)
    res := LockGate(i)
    for j, v := range res {
      if v {
        fmt.Printf("Gate[%d] is opened...\n", j + 1)
      } else {
        fmt.Printf("Gate[%d] is closed...\n", j + 1)
      }
    }
  }
}
