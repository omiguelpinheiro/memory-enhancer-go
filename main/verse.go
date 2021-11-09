package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

func getPoemPath(r *rand.Rand) (path string) {
	memoryDir := MemoryFolder()
	langs, err := ioutil.ReadDir(fmt.Sprintf("%s/poems", memoryDir))
	check(err)
	i := r.Intn(len(langs))
	lang := langs[i].Name()
	langPath := fmt.Sprintf("%s/poems/%s", memoryDir, lang)

	poems, err := ioutil.ReadDir(langPath)
	check(err)
	poem_index := r.Intn(len(poems))
	poem := poems[poem_index].Name()
	return fmt.Sprintf("%s/%s", langPath, poem)
}

func TitleAuthorFromPath(path string) (title string, author string) {
	sPath := strings.Split(path, "/")
	pPath := sPath[len(sPath)-1]
	sPoem := strings.Split(pPath, "-")
	title = sPoem[0]
	author = sPoem[1][1:]
	return title, author
}

func Verse(r *rand.Rand, lineMax int, lineTol int, poemPath string) (verse []string) {
	lines, verses := LinesVersesFromFile(poemPath)
	indexes := r.Perm(len(verses))
	lower := lineMax - lineTol
	upper := lineMax + lineTol
	for _, i := range indexes {
		current_verse := verses[i]
		size := len(current_verse)
		if size >= lower && size <= upper {
			return verses[i]
		}
	}
	randI := r.Intn(len(lines) - lineMax)
	for suffix := 0; suffix < lineMax; suffix++ {
		verse = append(verse, lines[randI+suffix])
	}
	if verse == nil {
		log.Fatal("Could not get poem from file")
	}
	return verse
}

func LinesVersesFromFile(path string) (lines []string, verses [][]string) {
	poemF, err := os.Open(path)
	check(err)
	defer poemF.Close()

	var verse []string
	poemScan := bufio.NewScanner(poemF)
	var lineNum int = 0
	for poemScan.Scan() {
		line := poemScan.Text()
		if len(line) > 1 {
			lineLn := fmt.Sprintf("%d-%s\n", lineNum, line)
			lines = append(lines, lineLn)
			verse = append(verse, lineLn)
		} else if len(line) == 0 {
			lineLn := fmt.Sprintf("%d-%s\n", lineNum, line)
			lines = append(lines, lineLn)
			verses = append(verses, verse)
			verse = nil
		}
		lineNum += 1
	}
	verses = append(verses, verse)
	return lines, verses
}
