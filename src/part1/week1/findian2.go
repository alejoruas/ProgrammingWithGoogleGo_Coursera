package main
import (
	"os"
	"bufio"
	"fmt"
	"strings"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter a string: ")
	scanner.Scan()
	input := strings.ToLower(scanner.Text())
	if strings.HasPrefix(input, "i") &&
		strings.Contains(input, "a") &&
		strings.HasSuffix(input, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}