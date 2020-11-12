package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const filePath = "/Users/leonardollobato/Desktop/uploadts-1539677320-sec_Google_Shopping_Feed"

//Reading chunks
func file_chunck_main() {
	var start int64 = 5 * 1000 // Offset from where the reading starts
	var size int64 = 1 * 1000  //In MB

	// Open the File
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Define the buffer that will keep the chunk in buffer
	buffer := make([]byte, size)
	//for {

	// Read the file and put the content into the buffer (ReadAt(buffer, offset))
	if _, err := file.ReadAt(buffer, start); err == io.EOF {
		log.Fatal(err)
	}

	// do something with a buffer of 4Kb
	fmt.Print(string(buffer))
	//}
}
