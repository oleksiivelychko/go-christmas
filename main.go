package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

const (
	ledgeBranchCount = 3
	treeHeight       = 6
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
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c

		fmt.Println("\nBye bye... Merry Christmas and Happy New Year!")
		os.Exit(1)
	}()

	for {
		buildTree()

		time.Sleep(1 * time.Second)

		clearTerminal()
	}
}

func clearTerminal() {
	clearExec, ok := clearMap[runtime.GOOS]
	if ok {
		clearExec()
	} else {
		panic("Platform is undefined!")
	}
}

func buildTree() {
	stem := ledgeBranchCount*treeHeight + ledgeBranchCount + 1
	limit := ledgeBranchCount * treeHeight * 2

	fmt.Println(buildLine(limit*2+ledgeBranchCount, randomSnow))
	fmt.Printf("%s⭐️%s\n", buildLine(limit, randomSnow), buildLine(limit+1, randomSnow))

	for i := 1; i < stem+1; i++ {
		printLine(i, limit)

		if 0 == i%3 {
			printLine(i-1, limit)
		}
	}

	for i := 0; i < ledgeBranchCount-1; i++ {
		fmt.Printf("%s🟤%s\n", buildLine(limit, randomSnow), buildLine(limit+2, randomSnow))
	}
}

func putLeft(count int, limit int) string {
	return buildLine(limit-count-1, randomSnow) + buildLine(count+1, randomGarland)
}

func putRight(count int, limit int) string {
	return buildLine(count+1, randomGarland) + buildLine(limit-count+2, randomSnow)
}

func printLine(count int, limit int) {
	fmt.Printf("%s🟤%s\n", putLeft(count, limit), putRight(count, limit))
}

func randomGarland() string {
	random := [25]string{"🔴", "🟡", "🔵", "🟣", "🟠"}
	for i := 5; i < 25; i++ {
		random[i] = "🍀"
	}

	return random[rand.Intn(len(random))]
}

func randomSnow() string {
	return [2]string{"❄️", " ️"}[rand.Intn(2)]
}

func buildLine(limit int, fn func() string) string {
	line := ""

	for i := 0; i < limit; i++ {
		line += fn()
	}

	return line
}
