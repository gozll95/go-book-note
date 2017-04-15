package main

// written by zhulilei

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var result map[string]int
	result = make(map[string]int)

	r := bufio.NewReader(os.Stdin)
	for {
		line, e := r.ReadString('\n')
		if e == io.EOF {
			break
		}
		newLine := strings.Trim(line, "\n")
		result[newLine]++
	}
	fmt.Println("---")
	for i, v := range result {
		fmt.Printf("%s\t%d\n", i, v)
	}
}

/*
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
*/
