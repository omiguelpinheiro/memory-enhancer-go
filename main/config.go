package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseCfgFile(cfgName string) (tokenLen int, lineMax int, lineTol int) {
	exists, err := exists(cfgName)
	check(err)
	if !exists {
		m, err := os.OpenFile(cfgName, os.O_WRONLY|os.O_CREATE, 0600)
		check(err)
		writer := bufio.NewWriter(m)
		writer.WriteString("TOKEN_LENGHT=4\n")
		writer.WriteString("LINE_MAX=3\n")
		writer.WriteString("LINE_TOL=1\n")
		writer.Flush()
		m.Close()
	}
	memoryCfg, err := os.OpenFile(cfgName, os.O_RDONLY, 0600)
	check(err)
	defer memoryCfg.Close()

	tokenLen = -1
	lineMax = -1
	lineTol = -1
	cfg_scanner := bufio.NewScanner(memoryCfg)
	for cfg_scanner.Scan() {
		line := strings.ToUpper(cfg_scanner.Text())
		if strings.Contains(line, "TOKEN_LENGHT") {
			tokenLen = OptionValue(line)
		}
		if strings.Contains(line, "LINE_MAX") {
			lineMax = OptionValue(line)
		}
		if strings.Contains(line, "LINE_TOL") {
			lineTol = OptionValue(line)
		}
	}
	if err := cfg_scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if tokenLen <= 0 {
		log.Fatal(fmt.Sprintf("TOKEN_LENGHT can't be less or equals to 0, got %d", tokenLen))
	}
	if lineMax <= 0 {
		log.Fatal(fmt.Sprintf("LINE_MAX can't be less or equals to 0, got %d", lineMax))
	}
	if lineTol < 0 {
		log.Fatal(fmt.Sprintf("LINE_TOL can't be less than 0, got %d", lineTol))
	}

	return tokenLen, lineMax, lineTol
}

func OptionValue(line string) (value int) {
	sLine := strings.Split(line, "=")
	param := sLine[0]
	sVal := sLine[1]
	val, err := strconv.Atoi(sVal)
	check(err)
	if val <= 0 {
		fmt.Printf("Value %s should be greater than 0, you used %d\n", param, val)
		os.Exit(1)
	}
	return val
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
