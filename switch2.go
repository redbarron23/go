// Simple switch statement example:
package main
import "fmt"

func main() {
    index := 10
    switch {
        case index < 10:
            fmt.Println("The index is less than 10")
        case index > 10, index < 0:
            fmt.Println("The index is either bigger than 10 or less than 0")
        case index == 10:
            fmt.Println("The index is equal to 10")
        default:
            fmt.Println("This won't be printed anyway")
    }
}
