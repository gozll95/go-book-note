package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	t := 0
	word := 0
	char := 0
	for {
		s, e := r.ReadString('\n')
		if e == io.EOF {
			break
		}
		char = char + len(s)
		word = word + len(strings.Fields(s))
		t++
	}
	fmt.Printf("line is %d,word is %d,char is %d,", t, word, char)
}

//for-if-err<--> for switch-case-default
