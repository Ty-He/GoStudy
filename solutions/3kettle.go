package solutions 

import (
  "fmt"
  "bytes"
)

func ThreeKettlePuzzle() {
  // from a kettle pour into another kettle
  pour := func(cur, tarcur, tarcap int) (int, int) {
    if cur <= tarcap - tarcur {
      return 0, tarcur + cur
    }
    return cur - (tarcap - tarcur), tarcap
  }
  memo := map[int]struct{}{}
  buf := bytes.NewBuffer([]byte{})
  cap := []int{8, 5, 3}
  var dfs func([3]int) bool
  dfs = func(vec3 [3]int) bool {
    if vec3[0] == 4 || vec3[1] == 4 {
      return true
    }
    mask := (vec3[0] << 8) | (vec3[1] << 4) | vec3[2]
    if _, ok := memo[mask]; ok {
      return false
    }
    memo[mask] = struct{}{}
    for i := range vec3 {
      for j := 0; j < 3; j++ {
        if i == j {
          continue
        }
        oldi, oldj := vec3[i], vec3[j]
        vec3[i], vec3[j] = pour(vec3[i], vec3[j], cap[j])
        if dfs(vec3) {
          fmt.Fprintf(buf, "Pour %d to %d s=%v\n", i, j, vec3)
          return true
        }
        vec3[i], vec3[j] = oldi, oldj
      }
    }
    return false
  } 

  dfs([3]int{8, 0, 0})
  ans := bytes.Split(buf.Bytes(), []byte{'\n'})
  for i := len(ans) - 2; i >= 0; i-- {
    fmt.Println(string(ans[i]))
  }
}
