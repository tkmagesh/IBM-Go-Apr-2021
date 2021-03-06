package main

import "fmt"

func main() {
	//Array
	var nos [10]int
	fmt.Println(nos)

	for i := 0; i < 10; i++ {
		nos[i] = i * 2
	}
	fmt.Println(nos)

	//initialization while creation
	var fullname [2]string = [2]string{"Magesh", "Kuppan"}
	fmt.Println(fullname)

	//multidimentional array
	var matrix [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = fmt.Sprintf("row - %d col - %d", i, j)
		}
	}
	/* fmt.Printf("%q\n", matrix) */
	for _, row := range matrix {
		for _, col := range row {
			fmt.Println(col)
		}
	}

	//Slice
	/*
		dynamic
		uses an array internally
		highlevel apis for manipulation
	*/
	fmt.Println()
	fmt.Println("Slice")
	oddNos := []int{1, 3, 5, 7, 9}
	fmt.Println(oddNos)

	fmt.Printf("Size of the oddNos = %d\n", len(oddNos))
	/*
		slice[lo : hi] => return from lo to (hi - 1)
		slice[lo : lo] => empty
		slice[lo : lo+1] => one element (at lo)
		slice[lo :] => all the values from lo
	*/
	fmt.Println(oddNos[1:4])
	fmt.Println(oddNos[1:1])
	fmt.Println(oddNos[1:2])
	fmt.Println(oddNos[1:])
	fmt.Println(oddNos[:4])
	oddNos = append(oddNos, 11)
	fmt.Printf("After adding a new no, oddNos = %v\n", oddNos)

	//Use of 'make' to create instances applied only for slice, map & channel
	cities := make([]string, 0)
	cities = append(cities, "Bengaluru")
	cities = append(cities, "Mysuru")
	cities = append(cities, "Mangaluru")
	cities = append(cities, "Udupi")

	fmt.Println(cities)

	/*
		for i := 0; i < len(cities); i++ {
			fmt.Println(cities[i])
		}
	*/

	for _, city := range cities {
		fmt.Println(city)
	}

	//map
	cityRanks := map[string]int{
		"Udupi":     2,
		"Mysuru":    1,
		"Mangaluru": 3,
		"Bengaluru": 4,
	}

	fmt.Printf("Rank of Bengaluru = %d\n", cityRanks["Bengaluru"])
	cityRanks["Tumkuru"] = 5
	fmt.Println(cityRanks)

	for city, rank := range cityRanks {
		fmt.Printf("City = %s, Rank=%d\n", city, rank)
	}

	//removing an item from map
	delete(cityRanks, "Bengaluru")
	fmt.Println("After deleting Bengaluru")
	fmt.Println(cityRanks)

	//check for the existance of a key
	_, exists := cityRanks["Bengaluru"]
	fmt.Println("Is Bengaluru in the list? ", exists)
}
