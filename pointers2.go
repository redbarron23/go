 package main
import "fmt"

func main() {
    sum := 0
    var  double_sum *int //a pointer to int
    for i:=0; i<10; i++{
        sum += i
    }
    double_sum = new(int) //allocate memory for an int and make double_sum point to it
    *double_sum = sum*2 //use the allocated memory, by dereferencing double_sum
    fmt.Println("The sum of numbers from 0 to 10 is: ", sum)
    fmt.Println("The double of this sum is: ", *double_sum)
}
