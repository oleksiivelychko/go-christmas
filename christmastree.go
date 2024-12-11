package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type ChristmasTree struct {
	cols, rows int
}

func NewChristmasTree(cols, rows int) *ChristmasTree {
	return &ChristmasTree{cols: cols, rows: rows}
}

func (t *ChristmasTree) Draw() {
	fmt.Printf("%s⭐️%s\n", t.randomSnow(0), t.randomSnow(0))

	for i := 1; i <= t.rows; i++ {
		t.printLine(i)

		if 0 == i%3 {
			t.printLine(i - 1)
		}
	}
}

func (t *ChristmasTree) randomSnow(garlandCount int) string {
	cols := t.cols / 2
	if 0 != cols%10 {
		cols -= cols % 10
	}

	spaces := t.rows
	snowflakes := cols - garlandCount - spaces

	line := make([]string, snowflakes)
	for i := 0; i < snowflakes; i++ {
		line = append(line, "❄️")
	}

	for i := 0; i < spaces; i++ {
		line = append(line, "  ")
	}

	for i := range line {
		j := rand.Intn(i + 1)
		line[i], line[j] = line[j], line[i]
	}

	return strings.Join(line, "")
}

func (t *ChristmasTree) randomGarland(count int) string {
	garland := append(
		strings.Split(strings.Repeat("🟢", 20), ""),
		"🔴", "🟡", "🔵", "🟣", "🟠",
	)

	ledge := make([]string, count)
	for i := 0; i < count; i++ {
		ledge = append(ledge, garland[rand.Intn(len(garland))])
	}

	return strings.Join(ledge, "")
}

func (t *ChristmasTree) printLine(i int) {
	fmt.Printf("%s%s🟤%s%s\n",
		t.randomSnow(i),
		t.randomGarland(i),
		t.randomGarland(i),
		t.randomSnow(i),
	)
}
