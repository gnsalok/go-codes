package main

import "fmt"

func main() {

	/*
		- Maps pass to function by reference / its reference type.
		- Made chagnes will refect to function caller / other callers.
		- Maps are not Thread safe
		- Try to specify size of larger Maps (when size is large its costly)
	*/

	emp := make(map[int]string)

	/*
		OR,
		emp1 := map[int]string{
			1: "Alok",
			2: "Rio",
			3: "Tokyo",
		}

		fmt.Println(emp1)
	*/

	emp[1] = "Tokyo"
	emp[2] = "Professor"
	emp[3] = "Rio"

	//updating map
	emp[3] = "Alok"

	//delete element in map
	delete(emp, 1)

	// it will come in order by key
	fmt.Println(emp)

	// it will not come in order
	for k, v := range emp {
		fmt.Println(k, v)
	}

	// Method-1 :Check if value in the Map
	if val, ok := emp[5]; ok {
		fmt.Println("Key in Present, Value ", val)
	} else {
		fmt.Println("Key in not Present")
	}

	// Method-2 : Check if value present in the Map

	val, exist := emp[3]
	if exist {
		fmt.Println("exist ", val)
	}
}

/*
nameMap := make(map[string]int)
nameMap["river"] = 33
v ,exist := nameMap["river"]
if exist {
    fmt.Println("exist ",v)
}
*/
