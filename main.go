package main

import (
	"fmt"

	"github.com/pixelbender/go-matroska/matroska"
)

func main() {

	var doc, err = matroska.Decode("filePath")

	if err == nil {
		return
	}

	var a = doc.EBML.Version

	fmt.Println(a)
}
