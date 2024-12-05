package main

import "fmt"

// map with type aliases
type floatMap map[string]float64

// map with type aliases and function
func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://amazon.com",
	} // map[type of keys]values datatype to be used{"key": "value"} // empty map {}

	fmt.Println(websites)

	// access values using keys
	fmt.Println(websites["Google"])
	fmt.Println(websites["Amazon"])

	// add new key value pairs in map[type]type
	websites["Github"] = "https://github.com"
	fmt.Println(websites)

	// delete key value pairs in map[type]type
	delete(websites, "Google")
	fmt.Println(websites)

	// create map[type]type
	courseRating := map[string]float64{} // map of string type keys and float type values
	courseRating["go"] = 4.7
	courseRating["python"] = 5

	fmt.Println(courseRating)
	fmt.Println(courseRating["go"])

	// create map() using make() function
	animals := make(map[string]string, 3) // create map of keys string and values string has capacity of 3
	animals["cat"] = "simple and cute"
	animals["tiger"] = "dangerous and bad"
	animals["dog"] = "honest and cute"
	fmt.Println(animals)

	// map with type aliases
	newTypeAliasesMap := make(floatMap, 3)
	newTypeAliasesMap["movie"] = 9
	newTypeAliasesMap["programming"] = 10
	newTypeAliasesMap.output()
}
