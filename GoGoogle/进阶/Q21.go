// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// )

// func CheckErr(err error) {
// 	if nil != err {
// 		panic(err)
// 	}
// }

// func getStrings(filename string) {
// 	f, err := os.Open(filename)
// 	CheckErr(err)
// 	defer f.Close()

// 	reader := bufio.NewReader(f)
// 	line := 0
// 	for {
// 		line_context, err := reader.ReadString('\n') //以'\n'为结束符读入一行

// 		if err != nil || io.EOF == err {
// 			break
// 		}
// 		line++
// 		fmt.Println(line, line_context)
// 	}

// }

// func main() {
// 	getStrings("a.txt")
// }

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var numberFlag = flag.Bool("n", false, "number each line")

func cat(r *bufio.Reader) {
	i := 1
	for {
		buf, e := r.ReadBytes('\n')
		if e == io.EOF {
			break
		}
		if *numberFlag {
			fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, e := os.Open(flag.Arg(i))
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s:error readingfrom %s: %s\n", os.Args[0], flag.Arg(i), e.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
