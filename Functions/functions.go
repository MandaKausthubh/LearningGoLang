package main 

import (
    "fmt"
    "math/rand"
)

func GenerateRandoms() (x int) {
    x = rand.Intn(100)
    return
}

func split(sum int) (x, y int) {
    x = sum * 4/9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(GenerateRandoms()))
}
