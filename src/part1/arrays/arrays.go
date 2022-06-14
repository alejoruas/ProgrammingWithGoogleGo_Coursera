package main

import "fmt"

func main() {
	x := [...]int{1, 454, 232, 5, 0}
	for i, v := range x {
		fmt.Printf("In the index %d the value is %d\n", i, v)
	}

	//Slices

	fmt.Printf("Sliceof x\n")
	s1 := x[2:5]
	for i, v := range s1 {
		fmt.Printf("In de index %d the value is %d\n", i, v)
	}

	fmt.Printf("len ofr slice of x %d\n", len(s1))
	fmt.Printf("capacity of slice of x %d\n", cap(s1))

	sli := make([]int, 10, 15)
	fmt.Printf("Slice sli\n")
	for i, v := range sli {
		fmt.Printf("In de index %d the value is %d\n", i, v)
	}

	sli = append(sli, 5)

	fmt.Printf("Slice sli\n")
	for i, v := range sli {
		fmt.Printf("In de index %d the value is %d\n", i, v)
	}

	//var idMap map[string]int
	idMap2 := map[string]int{"joe": 122}

	fmt.Println("MAP")
	fmt.Println(idMap2["joe"])

	idMap2["joe"] = 563456
	fmt.Println(idMap2["joe"])

	id, p := idMap2["joe"]

	if p == true {
		fmt.Println(id)
	} else {
		fmt.Println("doesn't exist")
	}

	for key, val := range idMap2 {
		fmt.Printf("key %s value %d", key, val)
	}

	delete(idMap2, "joe")

}
