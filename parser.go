package main

import (
	"bufio"
	"container/heap"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	Score uint64 `json:"score"`
	ID    string `json:"id"`
	index int
}
type EntryHeap []*Entry

func (h EntryHeap) Len() int           { return len(h) }
func (h EntryHeap) Less(i, j int) bool { return h[i].Score > h[j].Score }
func (h EntryHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j

}
func (h *EntryHeap) Push(x interface{}) {
	n := len(*h)
	entry := x.(Entry)
	entry.index = n
	*h = append(*h, &entry)
	heap.Fix(h, n)
}

func (h *EntryHeap) Pop() interface{} {
	old := *h
	n := len(old)
	entry := old[n-1]
	entry.index = -1
	*h = old[0 : n-1]
	return *entry
}

func (h *EntryHeap) HighestScores(n int) []Entry {
	var highestScores []Entry
	for i := 0; i < n; i++ {
		highestScores = append(highestScores, heap.Pop(h).(Entry))
	}
	return highestScores
}
func main() {

	fname := flag.String("f", "", "Data file name")
	debug := flag.Bool("d", false, "Turn Debug on")
	n := flag.Int("n", 5, "The N highest scores")
	flag.Parse()
	if !*debug {
		log.SetOutput(ioutil.Discard)
	}
	if *fname == "" {
		log.Fatal("-f command argument is required")
	}
	if *n < 1 {
		log.Fatal("-n command argument is required with a positive integer")

	}
	os.Exit(HighestNScores(*fname, *n))
}

func outputHighestScores(highestScores []Entry) (int, error) {
	output, err := json.MarshalIndent(highestScores, "", "  ")
	if err != nil {
		return 1, err
	}
	fmt.Fprintf(os.Stdout, "%s\n", output)
	return 0, nil
}

func HighestNScores(fname string, n int) int {
	entries, exitCode, err := processFile(fname, n)
	if err != nil {
		log.Println(err)
		return exitCode
	}

	if n > len(entries) {
		// if the file is small we will return the sorted entries instead of failing
		n = len(entries)
	}
	highestScores := entries.HighestScores(n)
	exitCode, err = outputHighestScores(highestScores)
	if err != nil {
		log.Println(err)
		return exitCode
	}
	return 0
}

func processFile(fname string, n int) (EntryHeap, int, error) {
	var entries EntryHeap

	log.Println("Processing Data file:", fname)
	log.Println("n:", n)
	file, err := os.Open(fname)
	if err != nil {
		return entries, 1, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if entry, err := processRecord(scanner.Text()); err != nil {
			return entries, 2, err
		} else {
			if entry.Score != 0 {
				heap.Push(&entries, entry)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return entries, 1, err
	}
	if len(entries) == 0 {
		return entries, 2, errors.New("Empty input file")
	}
	return entries, 0, nil
}

func processRecord(record string) (Entry, error) {
	var dictionary map[string]interface{}
	var entry Entry
	if len(strings.TrimSpace(record)) == 0 {
		entry.Score = 0
		return entry, nil
	}
	tokens := strings.SplitN(record, ":", 2)
	if tokens == nil {
		return entry, errors.New("Encountered formating error for record: " + record)

	}
	score, err := strconv.ParseUint(tokens[0], 10, 64)
	if err != nil {
		return entry, err
	}
	entry.Score = score
	if err := json.Unmarshal([]byte(tokens[1]), &dictionary); err != nil {
		return entry, err
	}
	if dictionary["id"] == nil {
		return entry, errors.New("no id field in record")
	}
	id := dictionary["id"].(string)
	// TODO check if id exists
	entry.ID = id
	log.Println("score:", score, "id:", id)
	return entry, nil
}
