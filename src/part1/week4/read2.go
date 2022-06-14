package main
import ("fmt"
		"os"
		"bufio"
)
func main(){
	type person struct{
		fName string
		lName string
	}
	var fileName string
	fmt.Println("Please enter a file name:")
	fmt.Scanln(&fileName)

    readFile, err := os.Open("data.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
        fmt.Println(fileScanner.Text())
    }
  
    readFile.Close()
}