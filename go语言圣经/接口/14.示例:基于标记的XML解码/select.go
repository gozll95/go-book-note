// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	//fmt.Println(dec)

	var stack []string // stack of element names
	for {
		tok, err := dec.Token()
		fmt.Println("==============")
		fmt.Println(tok)
		fmt.Println("==============")
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			fmt.Println("//////")
			fmt.Println(tok.Name.Local)
			fmt.Println("//////")
			stack = append(stack, tok.Name.Local) // push
			fmt.Println("push stack is", stack)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
			fmt.Println("pop stack is", stack)
		case xml.CharData:
			fmt.Println("到了CharData了")
			if containsAll(stack, os.Args[1:]) {
				fmt.Println("reading~~~~")
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
				fmt.Println("finish~~~~")
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
