package main

import (
	"fmt"
	"os"
	"io"
	"bytes"
)

func countLines(reader io.Reader) (int) {
	buffer := make([]byte, 32*1024)
	var countLines int = 0

	for {
		symbols, err := reader.Read(buffer)
		countLines += bytes.Count(buffer[:symbols], []byte{'\n'})

		switch {
		case err == io.EOF: return countLines + 1
		case err != nil: return countLines + 1
		}
	}
}


func main()  {
	argv := os.Args[1:] //var argv []string = os.Args[1:]
	if len(argv) < 1 {fmt.Println("Enter the file name"); return}
	file, err := os.Open(argv[0])
	if err != nil {fmt.Println("File error"); return} //can use log.Fatal
	defer file.Close()

	fmt.Println("Count of lines in file: ", countLines(file))
}
