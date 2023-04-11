package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// {
// 	"month": "12",
// 	"day": "15",
// 	"year": "2008",
// 	"num": "571",
// 	"transcript": "[[Someone is in bed ... long int.]]",
// 	"img": "https://imgs.xkcd.com/comics/marshmallow_gun.png",
// 	"title": "Marshmallow Gun",
// 	}

type xkcd struct {
	Num        int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no file is given")
		os.Exit(-1)
	}

	fn := os.Args[1]

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "no search term is given")
		os.Exit(0)
	}

	var (
		items []xkcd
		terms []string
		input io.ReadCloser
		cnt   int
		err   error
	)

	// decode the file
	input, err = os.Open(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		os.Exit(-1)
	}
	defer input.Close()

	if err = json.NewDecoder(input).Decode(&items); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %v\n", err)
		os.Exit(-1)
	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", len(items))
	/* data, err := ioutil.ReadAll(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %v\n", err)
		os.Exit(-1)
	}


	err = json.Unmarshal(data, &items)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse JSON: %v\n", err)
		os.Exit(-1)
	} */

	// get serch items
	for _, t := range os.Args[2:] {
		terms = append(terms, strings.ToLower(t))
	}

	// search
outer:
	for _, item := range items {
		title := strings.ToLower(item.Title)
		transcript := strings.ToLower(item.Title)

		for _, term := range terms {
			if !strings.Contains(title, term) && !strings.Contains(transcript, term) {
				continue outer
			}

		}

		fmt.Printf("http://xkcd.com/%d/ %s/%s/%s %q\n",
			item.Num, item.Month, item.Day, item.Year, item.Title)
		cnt++
	}

	fmt.Fprintf(os.Stdout, "Found %d comics\n", cnt)
}
