// Simple switch statement example:
package main
import "fmt"

func main() {
        //switch example with fallthrough
	integer := 6
	switch integer {
	case 4:
	    fmt.Println("The integer was <= 4")
	    fallthrough
	case 5:
	    fmt.Println("The integer was <= 5")
	    fallthrough
	case 6:
	    fmt.Println("The integer was <= 6")
	    fallthrough
	case 7:
	    fmt.Println("The integer was <= 7")
	    fallthrough
	case 8:
	    fmt.Println("The integer was <= 8")
	    fallthrough
	default:
	    fmt.Println("default case")
	}
}
