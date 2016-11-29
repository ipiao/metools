package demo

import (
	"fmt"
	"log"
	"os"
	"testing"
	"text/scanner"
)

func TextScanner() {
	var path = "../file/test_scanner.txt"
	file, err := os.Open(path)
	log.Println(err)
	scan := scanner.Scanner{}
	scan.Scan()
	fmt.Println(scanner.TokenString())

}

func TestTextScanner(t *testing.T) {
	TextScanner()
}
