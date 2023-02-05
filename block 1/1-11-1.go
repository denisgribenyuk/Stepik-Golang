var workArray [10]uint8
	for i := 0; i < len(workArray); i++ {
		var x uint8
		fmt.Scan(&x)
		workArray[i] = x
	}
	for a := 0; a < 3; a++ {
		var first int
		var second int
		fmt.Scan(&first)
		fmt.Scan(&second)
		workArray[first], workArray[second] = workArray[second], workArray[first]
	}
	for _, el := range workArray {
		fmt.Printf("%d ", el)
	}




