# Quiz 3

## Questions
1. What is printed when the following program is executed?
    ```
    func main() {
    x := []int {4, 8, 5}
    y := -1
    for _, elt := range x {
        if elt > y {
        y = elt
        }
    }
    fmt.Print(y)
    }
    ```
2. What is printed when the following program is executed?
    ```
    func main() {
    x := [...]int {4, 8, 5}
    y := x[0:2]
    z := x[1:3]
    y[0] = 1
    z[1] = 3
    fmt.Print(x)
    }
    ```
3. What is printed when the following program is executed?
    ```
    func main() {
    x := [...]int {1, 2, 3, 4, 5}
    y := x[0:2]
    z := x[1:4]
    fmt.Print(len(y), cap(y), len(z), cap(z))
    }
    ```
4. What is printed when the following program is executed?
    ```
    func main() {
    x := map[string]int {
        "ian": 1, "harris": 2}
    for i, j := range x {
        if i == "harris" {
        fmt.Print(i, j)
        }
    }
    }
    ```
5. What is printed when the following program is executed?
    ```
    type P struct {
        x string
        y int
    } 
    func main() {
    b := P{"x", -1}
    a := [...]P{P{"a", 10}, 
            P{"b", 2},
            P{"c", 3}}
    for _, z := range a {
        if z.y > b.y {
        b = z
        }
    }
    fmt.Println(b.x)
    }
    ```
6. What is printed when the following program is executed?
    ```
    func main() {
    s := make([]int, 0, 3)
    s = append(s, 100)
    fmt.Println(len(s), cap(s))
    }
    ```

## Answers
1. 8
2. [1 8 3]
3. 2 5 3 4
4. harris2
5. a
6. 1 3