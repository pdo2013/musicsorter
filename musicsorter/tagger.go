package main

import (
	"log"
	"os"

	"github.com/dhowden/tag"
)

// type music struct {
// 	musicmeta *tag.Metadata
// 	path      string
// 	filename  string
// }

func readtag(f *os.File) (musicmeta tag.Metadata) {
	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Println(err)
	}
	// log.Print(m)
	return m
}
