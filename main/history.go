package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func OpenHistory(installFolder string) (hist *os.File) {
	histPath := fmt.Sprintf("%s/history", installFolder)
	hist, err := os.OpenFile(histPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	check(err)
	return hist
}

func WriteToHistory(history *os.File, ex Exercise) {
	writer := bufio.NewWriter(history)

	date := time.Now().Format("2006-01-02")
	writer.WriteString(fmt.Sprintf("Date:%s\n", date))
	writer.WriteString(fmt.Sprintf("Token:%s\n", ex.token))
	writer.WriteString(fmt.Sprintf("Title:%s\n", ex.title))
	for _, line := range ex.verse {
		writer.WriteString(line)
	}
	writer.WriteString(fmt.Sprintf("Author:%s\n\n", ex.author))
	writer.Flush()
}
