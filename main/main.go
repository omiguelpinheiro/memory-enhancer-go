package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gopasspw/gopass/pkg/pwgen"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	history := OpenHistory()
	defer history.Close()
	runToday := RunToday(history)
	tokenLen, lineMax, lineTol := ParseCfgFile("./memory.cfg")
	if runToday {
		poemPath := getPoemPath(r)
		token := pwgen.GeneratePassword(tokenLen, true)
		title, author := TitleAuthorFromPath(poemPath)
		verse := Verse(r, lineMax, lineTol, poemPath)
		WriteToHistory(history, token, title, author, verse)
	}
}
