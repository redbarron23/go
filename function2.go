package main

import (
    "fmt"
)

// two parameters
//func print(msg string, msg2 string) {
func print(msg, msg2 string) {
    fmt.Printf("%s\n", msg)
    fmt.Printf("%s\n", msg2)
}


func main() {
    //fmt.Printf("Hello World")
    print("Hello, World!")
}