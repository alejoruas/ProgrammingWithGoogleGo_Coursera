package main
import "fmt"

func main() {
var input_num float32
fmt.Println("please type a float number: ")
fmt.Scan(&input_num)
var trunc_int int = int(input_num)
fmt.Printf("The truncated number is: %d \n", trunc_int)
}
