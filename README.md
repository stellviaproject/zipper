# Documentation
## Introduction

You will see now how easy is to compress a folder in golang.
This package has only one function, call to ZipFolder(folder string) ([]byte, error) to get the bytes of folder compressed.

## Instalation
go get github.com/stellviaproject/zipper

## Example
```golang
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
```