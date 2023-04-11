package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func getCommic(index int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", index)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "stoped working: %s\n", err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	// check if status code 404
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("skippin 404: got %d\n", resp.StatusCode)
		return nil
	}
	// read the response body in byte slice
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "bad body: %s", err)
		os.Exit(-1)
	}

	return body

}

func main() {
	var (
		output      io.WriteCloser = os.Stdout
		err         error
		errCount    int
		comicCount  int
		commicsData []byte
	)

	if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		defer output.Close()
	}

	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	for i := 1; errCount < 2; i++ {
		commicsData = getCommic(i)

		if commicsData == nil {
			errCount++
			continue
		}

		if comicCount > 0 {
			fmt.Fprint(output, ",") // to io writer
		}
		// Write the JSON data to the standard output stream
		_, err = io.Copy(output, bytes.NewBuffer(commicsData))
		if err != nil {
			fmt.Fprintf(os.Stderr, "stopped: %s\n", err)
			os.Exit(-1)
		}
		errCount = 0
		comicCount++

	}

}
