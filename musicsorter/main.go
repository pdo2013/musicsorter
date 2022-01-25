package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	sourcedir string
	targetdir string
	ascii     = [...]rune{47, 92, 58, 42, 63, 34, 60, 62, 124}
	files     = []string{}
)

func main() {
	// l := list.New()
	currentpath, _ := os.Getwd()
	var input string
	var Musicdir = flag.String("m", currentpath, "音乐目录")
	var albdir = flag.String("t", "", "专辑目录")
	flag.Parse()

	sourcedir = *Musicdir
	if *albdir == "" {
		targetdir = sourcedir
	} else {
		targetdir = *albdir
	}
	fmt.Printf("音乐目录(Music):%s\n专辑目录(Album):%s\n开始整理(start)?(y/n)", sourcedir, targetdir)
	fmt.Scanf("%s", &input)
	if input == "y" || input == "Y" {
		fmt.Println("start....")
		// GetMusicFiles(sourcedir)
		Genlist(sourcedir)
	}
	// fmt.Println(sourcedir, targetdir)

	for _, value := range files {
		handlefile(value, targetdir)
	}
	fmt.Println("done!")
}
