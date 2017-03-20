# Golang Notes

These are my personal nodes for using Google Go. 

## Sorting

```
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
