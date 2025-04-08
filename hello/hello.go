package main 

import (
    "fmt"
    "math/rand"
)

func main() {
    // Seed the random number generator
    // rand.Seed(42)

    // Generate a random number between 0 and 100
    randomNumber := rand.Intn(100)

    // Print the random number
    fmt.Println("Hello, World!")
    fmt.Println("Random number:", randomNumber)
}
