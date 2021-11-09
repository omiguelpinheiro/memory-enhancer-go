package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
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
	memoryDir := MemoryFolder()
	tokenLen, lineMax, lineTol := ParseCfgFile(fmt.Sprintf("%s/memory.cfg", memoryDir))
	if runToday {
		poemPath := getPoemPath(r)
		token := pwgen.GeneratePassword(tokenLen, true)
		title, author := TitleAuthorFromPath(poemPath)
		verse := Verse(r, lineMax, lineTol, poemPath)
		WriteToHistory(history, token, title, author, verse)
	}
}

func MemoryFolder() (home string) {
	homeDir, err := os.UserHomeDir()
	check(err)
	return fmt.Sprintf("%s/.memory-enhancer", homeDir)
}
