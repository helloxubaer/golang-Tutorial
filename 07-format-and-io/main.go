package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		//_, err := io.Copy(os.Stdout, file);

		/* 	data, err := ioutil.ReadAll(file)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		fmt.Println("The file has ", len(data), " bytes")
		*/

		var lc, wc, cc int

		scan := bufio.NewScanner(file)
		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++

		}
		fmt.Printf("%7d %7d %7d %s\n", lc, wc, cc, fname)
		file.Close()
	}
}
