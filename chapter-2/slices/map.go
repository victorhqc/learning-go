package main

import "fmt"

func main() {
	dict := make(map[string]string)
	dict["go"] = "Golang"
	dict["cs"] = "CSharp"
	dict["rb"] = "Ruby"
	dict["py"] = "Python"
	dict["js"] = "JavaScript"

	for key, value := range dict {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	// lan, ok := dict["js"]
	if lan, ok := dict["js"]; ok {
		fmt.Println(lan, ok)
	}
}
