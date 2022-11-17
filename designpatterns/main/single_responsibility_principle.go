package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Journal struct {
	entries []string
}

var entryCount = 0

func (j *Journal) ToString() string {
	return strings.Join(j.entries, "\n")
}
func (j *Journal) AddEntry(s string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, s)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// implement
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) WriteToFile(j *Journal, filename string) {
	ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I am hungry")
	j.AddEntry("I am full")
	fmt.Println(j.ToString())

	p := Persistence{lineSeparator: "\n"}
	p.WriteToFile(&j, "myJournal.txt")
}
