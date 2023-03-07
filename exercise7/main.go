package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	records := read_binary_db()
	dump := read_origin_file(records)
	write_encoded_content(dump)
	reversed_map := reverse_map(records)
	read_encoded_content(reversed_map)
}

func read_binary_db() map[string]string {
	records := map[string]string{}
	f, err := os.Open("./binario.txt")
	if err != nil {
		panic("File couldn't be opened")
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), ",")
		letter := line[0]
		binary := line[1]
		records[letter] = binary
	}

	return records
}

func read_origin_file(dict map[string]string) string {
	buffer, err := ioutil.ReadFile("./origen.txt")
	if err != nil {
		panic("File couldn't be opened")
	}
	text_buffer := string(buffer)
	text := bufio.NewScanner(strings.NewReader(text_buffer))
	text.Split(bufio.ScanRunes)

	var b bytes.Buffer
	for text.Scan() {
		result := fmt.Sprintf("%s ", dict[text.Text()])
		b.WriteString(result)
	}

	return b.String()
}

func write_encoded_content(dump string) {
	f, err := os.Create("./destino.txt")
	if err != nil {
		panic("File couldn't be created")
	}
	defer f.Close()
	f.WriteString(dump)
	fmt.Println("Content Encoded")
}

func read_encoded_content(dict map[string]string) {
	f, err := os.Open("./destino.txt")
	if err != nil {
		panic("File couldn't be opened")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		bs := strings.Split(scanner.Text(), " ")
		for _, v := range bs {
			fmt.Printf("%s", dict[v])
		}
	}
}

func reverse_map(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}
