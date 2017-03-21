# Golang Notes

These are my personal nodes for using Google Go. 

## maps

```go
m := make(map[string]int)
m := map[string]int {
  "hello" : 5,
  "world" : 6,
  "times up" : -1,
}

m["hello"] = 5
m["hello"]++      // zero value allows even if m["hello"] not already set

val, ok := m[key]

for key, val := range m {
}

delete(m, "yes")   // noop if "yes" does not exist, not an error 

```

maps are not write thread safe.  Use a goroutine with input channel. 

With for loop on range of a map, order will vary!

## Sorting

```go
import "sort"

type Matrix [][2]int

func (m Matrix) Len() int {
    return len(m)
}

func (m Matrix) Less(i, j int) bool {
    return m[i][0] < m[j][0]
}

func (m Matrix) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

func main() {
    m := make(Matrix, 100)
    ...
    sort.Sort(m)
    ...
}
```

