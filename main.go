package main

import (
	"fmt"

	"github.com/otiai10/gosseract"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("/mnt/d/go/ocr_test/sheet.jpg")
	client.SetLanguage("chi_sim")
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!
}
