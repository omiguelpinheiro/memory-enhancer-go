package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gopasspw/gopass/pkg/pwgen"
	"github.com/icza/backscanner"
)

type Exercise struct {
	token string
	title string
	verse []string
	author string
}

func main() {
	InitLog()

	randomizer := randomizer()
	installFolder := InstallFolder(false)

	history := OpenHistory(installFolder)
	defer history.Close()

	shouldRun := ShouldRun(history)
	if shouldRun {
		conf := OpenConf(installFolder)
		defer conf.Close()
		opt := TreatOpt(ReadOpt(conf), conf)

		token := pwgen.GeneratePassword(opt["tokenLen"], true)

		poemPath := getPoemPath(installFolder, randomizer)
		title, author := TitleAuthorFromPath(poemPath)
		verse := Verse(randomizer, opt, poemPath)

		ex := Exercise {token, title, verse, author}

		WriteToHistory(history, ex)
		ShowEx(token, title, author, verse)
	}
}

func InstallFolder(dev bool) (home string) {
	homeDir, err := os.UserHomeDir()
	check(err)
	if !dev {
		return fmt.Sprintf("%s/.memory-enhancer", homeDir)
	} else {
		return fmt.Sprintf("%s/projects/memory-enhancer", homeDir)
	}
}

func ShouldRun(history *os.File) bool {
	fi, err := history.Stat()
	check(err)

	bs := backscanner.New(history, int(fi.Size()))
	search := []byte("Date:")
	pos, line, err := ReverseSearch(bs, search)
	check(err)
	if pos == -1 {
		return true
	}
	lastDate := strings.Split(line, ":")[1]
	curDate := time.Now().Format("2006-01-02")
	if lastDate < curDate {
		return true
	} else {
		return false
	}
}

func ShowEx(token, title, author string, verse []string) {
	fmt.Printf("Todays token is:%s\n", token)
	fmt.Println("Have fun remembering this piece of a poem")
	fmt.Printf("Title:%s\n", title)
	for _, line := range verse {
		fmt.Printf("%s", line)
	}
	fmt.Printf("Author:%s\n", author)
}
