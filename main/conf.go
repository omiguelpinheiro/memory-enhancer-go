package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func OpenConf(installFolder string) (conf *os.File) {
	confPath := fmt.Sprintf("%s/memory.conf", installFolder)
	conf, err := os.OpenFile(confPath, os.O_CREATE|os.O_RDWR, 0600)
	check(err)
	return conf
}

func ReadOpt(conf *os.File) (map[string]int) {
	opt := make(map[string]int)
	opt["tokenLen"] = -1
	opt["lineMax"] = -1
	opt["lineTol"] = -1

	confS := bufio.NewScanner(conf)
	for confS.Scan() {
		line := strings.ToUpper(confS.Text())
		if strings.Contains(line, "TOKEN_LENGHT") {
			opt["tokenLen"] = ExtractVal(line)
		}
		if strings.Contains(line, "LINE_MAX") {
			opt["lineMax"] = ExtractVal(line)
		}
		if strings.Contains(line, "LINE_TOL") {
			opt["lineTol"] = ExtractVal(line)
		}
	}
	if err := confS.Err(); err != nil {
		log.Fatal(err)
	}

	return opt
}

func ExtractVal(line string) (value int) {
	sLine := strings.Split(line, "=")
	sVal := sLine[1]
	val, err := strconv.Atoi(sVal)
	check(err)
	return val
}

func TreatOpt(opt map[string]int, conf *os.File) map[string]int {
	if opt["tokenLen"] == -1 || opt["lineMax"] == -1 || opt["lineTol"] == -1 {
		writer := bufio.NewWriter(conf)
		if opt["tokenLen"] == -1 {
			writer.WriteString("TOKEN_LENGHT=4\n")
			opt["tokenLen"] = 4
		}
		if opt["lineMax"] == -1 {
			writer.WriteString("LINE_MAX=3\n")
			opt["lineMax"] = 2
		}
		if opt["lineTol"] == -1 {
			writer.WriteString("LINE_TOL=1\n")
			opt["lineTol"] = 1
		}
		writer.Flush()
	}
	return opt
}
