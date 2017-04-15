package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(2)
	log.SetPrefix("hehe:")
	log.Print("nihao")
	fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", nil)
	fmt.Println("zhulilkei")
}
