# Golang Notes

These are my personal nodes for using Google Go. 

## input 

### read an int from stdin

```go
var i int
fmt.Scan(&i)
```

## read from a file

```go
import "os"
import "bufio"

file, err := os.Open("filename")
defer file.Close()
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
}
```

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

## arrays

initialize

```go
a := make([]int, 10)
a := []int{}
```

operations
```go
a = append(a, v)                    // push
v, a = a[len(a)-1], a[:len(a)-1]    // pop
a = append([]int{v}, a)             // prepend
v, a = a[0], a[1:]                  // head

a = append(a[:i], a[i+1:])          // delete a[i]
b := make([]int,len(a))
copy(b,a)                           // copy
```

2-D arrays
```go
temp := make([]int, n * m)
x := make([][]int, n)
for i := 0; i < n; i++ {
    x[i] = temp[i * m : (i+1) * m]
}
```

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

Here is another example for arrays of int64s. 

```go
type int64arr []int64
func (a int64arr) Len() int { return len(a) }
func (a int64arr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }

n := 100
a := make(int64arr, n)
sort.Sort(a)
```
