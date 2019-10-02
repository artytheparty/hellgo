package main

import (
	"fmt"
	"os"
)

func main() {
	var locale string
	var translations = make(map[string]string)
	translations["en"] = "Hello"
	translations["de"] = "Guten tag"
	translations["fr"] = "Bonjour"

	if len(os.Args) == 1 {
		fmt.Println("Enter language code:")
		fmt.Scanf("%s", &locale)
	} else {
		locale = os.Args[1]
	}

	output := translations[locale]
	if output == "" {
		output = "Yo"
	}

	fmt.Println(output, "Go!")
}
