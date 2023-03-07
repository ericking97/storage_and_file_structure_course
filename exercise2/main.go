package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var option int
	fmt.Printf("Please select an storage option: (Type option number)\n1. Length\n2. Separator\n3. Description\n\n")
	fmt.Scan(&option)

	switch option {
	case 1:
		lengthRegister()
	case 2:
		separatorRegister()
	case 3:
		descriptionRegister()
	default:
		fmt.Print("Invalid selection")
	}
}

func lengthRegister() {
	fmt.Println("Not implemented")
}

func separatorRegister() {
	records := make([]string, 0)
	// We open the file and delay the closing of the file
	f, err := os.Open("./separador.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		records = append(records, fileScanner.Text())
	}

	dump := strings.Join(records, "&")
	fw, err := os.Create("./DBSeparador.txt")
	if err != nil {
		fmt.Println(err)
	}

	fw.WriteString(dump)
}

func descriptionRegister() {
	fmt.Println("Not implemented")
}
