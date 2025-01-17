package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var clearMap map[string]func()

func init() {
	clearMap = make(map[string]func())

	unixCmd := func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}

	clearMap["darwin"] = unixCmd
	clearMap["linux"] = unixCmd
	clearMap["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}

	_, ok := clearMap[runtime.GOOS]
	if !ok {
		log.Fatalf("OS %s does not support!", runtime.GOOS)
	}
}

func main() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-ch

		fmt.Println("\nBye bye... Have yourself a Merry Christmas and Happy New Year!")
		os.Exit(0)
	}()

	tree := NewChristmasTree(getDimensions())
	clearTerminal, _ := clearMap[runtime.GOOS]

	for {
		tree.Draw()

		time.Sleep(1 * time.Second)

		clearTerminal()
	}
}

func getDimensions() (int, int) {
	tputCols := exec.Command("tput", "cols")
	tputCols.Stdin = os.Stdin
	tputColsOut, err := tputCols.Output()
	if err != nil {
		log.Fatal(err)
	}

	cols, err := strconv.Atoi(strings.TrimSuffix(string(tputColsOut), "\n"))
	if err != nil {
		log.Fatal(err)
	}

	tputRows := exec.Command("tput", "lines")
	tputRows.Stdin = os.Stdin
	tputRowsOut, err := tputRows.Output()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := strconv.Atoi(strings.TrimSuffix(string(tputRowsOut), "\n"))
	if err != nil {
		log.Fatal(err)
	}

	remainder := cols % 2
	if remainder != 0 {
		cols -= remainder
	}

	remainder = rows % 10
	if remainder != 0 {
		rows -= remainder
	}

	return cols, rows
}
