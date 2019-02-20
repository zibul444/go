package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

func main() {
	//f, err := os.Open("go/src/go/Lalala/config.yml")
	f, err := os.Open("Lalala/config.yml") //, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	data := make([]byte, 64)
	t := T{}

	for {
		n, err := f.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
}
