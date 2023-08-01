package main

import (
	"log"
	"os"

	"github.com/stellviaproject/zipper"
)

func main() {
	compressedBytes, err := zipper.ZipFolder("./folder_to_zip")
	if err != nil {
		log.Fatalln(err)
	}
	if err := os.WriteFile("compressedFolder.zip", compressedBytes, os.ModePerm|os.ModeDevice); err != nil {
		log.Fatalln(err)
	}
}
