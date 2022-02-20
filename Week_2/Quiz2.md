# Quiz 2

## Questions
1. Which of the following expressions does NOT compute the average of two integers a and b?
2. What is printed when the following program is executed?
    ```
    func main() {
    i, _ := strconv.Atoi("10")
    y := i * 2
    fmt.Println(y)
    }
    ```
3. What is printed when the following program is executed?
    ```
    func main() {
    s := strings.Replace("ianianian", "ni", "in", 2)
    fmt.Println(s)
    }
    ```
4. What is printed by this code?
    ```
    func main() {
    x:=7
    switch {
        case x>3:
        fmt.Printf("1")
        case x>5:
        fmt.Printf("2")
        case x==7:
        fmt.Printf("3")
        default: 
        fmt.Printf("4")
    }
    }
    ```
5. What is printed by this code?
    ```
    func main() {
    var xtemp int
    x1 := 0
    x2 := 1
    for x:=0; x<5; x++ {
        xtemp = x2
        x2 = x2 + x1
        x1 = xtemp
    }
    fmt.Println(x2)
    }
    ```
6. True or False: This code compiles correctly.
    ```
    func main() {
    var x int
    var y *int
    z := 3
    y = &z
    x = &y
    }
    ```
7. Which integer type provides higher accuracy?

## Answers
1. Generates executable machine code from source code in a high-level language
2. 20
3. iainainan
4. 1
5. 8
6. False
7. All of these types(int, int16, int32) provide the same accuracy