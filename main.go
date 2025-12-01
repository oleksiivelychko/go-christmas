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
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-ch

		fmt.Printf("\nBye bye... Have yourself a Merry Christmas and Happy New Year %d!\n", time.Now().Year()+1)
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
	cols, err := getTerminalCols()
	if err != nil {
		cols = 80 // fallback default
		log.Printf("Warning: Could not get terminal columns: %v. Using default: %d.", err, cols)
	}

	rows, err := getTerminalRows()
	if err != nil {
		rows = 24 // fallback default
		log.Printf("Warning: Could not get terminal rows: %v. Using default: %d.", err, rows)
	}

	// normalize to even numbers
	if remainder := cols % 2; remainder != 0 {
		cols -= remainder
	}

	if remainder := rows % 10; remainder != 0 {
		rows -= remainder
	}

	return cols, rows
}

func getTerminalCols() (int, error) {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.TrimSpace(string(output)))
}

func getTerminalRows() (int, error) {
	cmd := exec.Command("tput", "lines")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.TrimSpace(string(output)))
}
