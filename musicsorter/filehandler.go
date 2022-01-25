package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func handlefile(path string, targetdir string) {
	music, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	// defer music.Close()
	m := readtag(music)
	music.Close()
	lrcpath := strings.TrimSuffix(path, filepath.Ext(path)) + ".lrc"

	if m.Album() != "" {
		finaldir := genpath(m.Album())
		filename := filepath.Base(path)
		lrcfinal := finaldir + "\\" + strings.TrimSuffix(filename, filepath.Ext(path)) + ".lrc"
		move(path, finaldir+"\\"+filename)
		move(lrcpath, lrcfinal)
	} else {
		log.Println("err" + path)
	}
}

func move(path string, dest string) error {
	if path != dest {
		log.Println(path, "->", dest)
		log.Printf("移动:%s -> %s\n", path, dest)
		err2 := os.Rename(path, dest)
		if err2 != nil {
			log.Println(err2)
			return err2
		}
	}
	return nil
}

func genpath(album string) string {
	albumdir := strings.TrimFunc(album, trimfilename)
	str3 := targetdir + "\\" + albumdir
	_, err := os.Stat(str3)
	if os.IsNotExist(err) {
		log.Printf("创建专辑目录:%s\n", albumdir)
		os.Mkdir(str3, os.ModePerm)
	}
	return str3
}

func getsupportsuffix(suffix string) bool {
	c := []string{".mp3", ".flac", ".ape", ".wma", ".wav", ".aac", ".ogg"}
	for _, value := range c {
		if suffix == value {
			return true
		}
	}
	return false
}

func trimfilename(r rune) bool {
	for _, num := range ascii {
		if r == num {
			return true
		}
	}
	return false
}
