package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/icza/backscanner"
)

func RunToday(history *os.File) bool {
	fi, err := history.Stat()
	check(err)

	s := backscanner.New(history, int(fi.Size()))
	search := []byte("Date:")
	pos, line, err := ReverseSearch(s, search)
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

func ShowTodays(token, title, author string, verse []string) {
	fmt.Printf("Todays token is:%s\n", token)
	fmt.Println("Have fun remembering this piece of a poem")
	fmt.Printf("Title:%s\n", title)
	for _, line := range verse {
		fmt.Printf("%s", line)
	}
	fmt.Printf("Author:%s\n", author)
}

func ReverseSearch(scanner *backscanner.Scanner, search []byte) (position int, line string, err error) {
	for {
		line, pos, err := scanner.LineBytes()
		if err != nil {
			if err == io.EOF {
				return -1, "", errors.New(fmt.Sprintf("%q is not found in file.\n", search))
			} else {
				return -1, "", err
			}
		}
		if bytes.Contains(line, search) {
			return pos, string(line), nil
		}
	}
}

func OpenHistory() (history *os.File) {
	memoryDir := MemoryFolder()
	history, err := os.OpenFile(fmt.Sprintf("%s/history", memoryDir), os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
	check(err)
	return history
}

func WriteToHistory(history *os.File, token, title, author string, verse []string) {
	writer := bufio.NewWriter(history)

	date := time.Now().Format("2006-01-02")
	writer.WriteString(fmt.Sprintf("Date:%s\n", date))
	writer.WriteString(fmt.Sprintf("Token:%s\n", token))
	writer.WriteString(fmt.Sprintf("Title:%s\n", title))
	for _, line := range verse {
		writer.WriteString(fmt.Sprintf("%s", line))
	}
	writer.WriteString(fmt.Sprintf("Author:%s\n\n", author))
	writer.Flush()
	ShowTodays(token, title, author, verse)
}
