package main

import (
    "fmt"
)

// parameters of different types  string and int
func print(msg, msg2 string, repeat int) {
    for repeat > 0 {
        fmt.Printf("%s\n", msg)
        fmt.Printf("%s\n", msg2)
        repeat -= 1
    }
}


func main() {
    print("Hello", "  World!", 5)
}