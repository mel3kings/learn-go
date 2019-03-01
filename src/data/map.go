package data

import "fmt"

func TestMap() {
	fmt.Println("==== Testing Mapping function")

	mapper := map[string]int{
		"Charlie": 11,
		"Amy":     21,
	}

	fmt.Println(mapper)

	for name := range mapper {
		fmt.Println(name)
	}

	fmt.Println("==== Testing Duplicate Key")
	mapper["Charlie"] = 29 // overrides existing key same as other languages
	fmt.Println(mapper)

	if age, ok := mapper["Charlie"]; ok {
		fmt.Println("Charlie's age is ", age)
	}

	_, ok := mapper["New Person"]
	if !ok {
		fmt.Println(" not found in map")
	}
}
