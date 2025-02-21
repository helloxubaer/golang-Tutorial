package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type response struct {
	Item   string `json:"item"`
	Album  string
	Title  string
	Artist string
}
type responseWrapper struct {
	response
}

var jdata1 = `{"item": "album",
			"album": {"title":"The Dark Side of the Moon"}
			}`
var jdata2 = `{"item": "song",
			"song":{"title":"bella Donna", "artist":"Steave Nicks"}
			}`

func main() {
	var resp1, resp2 responseWrapper
	var err error
	if err = json.Unmarshal([]byte(jdata1), &resp1); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp1.response)

	if err = json.Unmarshal([]byte(jdata2), &resp2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp2.response)
}

func (r *responseWrapper) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}

	if err := json.Unmarshal(b, &r.response); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return fmt.Errorf("error unmarshalling raw: %w", err)
	}

	switch r.Item {
	case "album":
		inner, ok := raw["album"].(map[string]interface{})
		if ok {
			if album, ok := inner["title"].(string); ok {
				r.Album = album
			}
		}
	case "song":
		inner, ok := raw["song"].(map[string]interface{})
		if ok {
			if title, ok := inner["title"].(string); ok {
				r.Title = title
			}
			artist, ok := inner["artist"].(string)
			if ok {
				r.Artist = artist
			}
		}
	}

	return nil
}
