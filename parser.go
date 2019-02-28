package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type entry struct {
	Score uint64 `json:"score"`
	ID    string `json:"id"`
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
	entries = sortEntries(entries)
	exitCode, err = outputHighestScores(entries, n)
	if err != nil {
		log.Println(err)
		return exitCode
	}
	return 0
}

func outputHighestScores(entries []entry, n int) (int, error) {
	highestScores := make([]entry, n)
	copy(highestScores, entries[0:n])
	output, err := json.MarshalIndent(highestScores, "", "  ")
	if err != nil {
		return 1, err
	}
	fmt.Fprintf(os.Stdout, "%s\n", output)
	return 0, nil
}

func sortEntries(entries []entry) []entry {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Score > entries[j].Score
	})
	return entries
}

func processFile(fname string, n int) ([]entry, int, error) {
	var entries []entry
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
				entries = append(entries, entry)
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

func processRecord(record string) (entry, error) {
	var dictionary map[string]interface{}
	var entry entry
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

// - check for  whitespace only lines.
// - validate Json
// - check for line breaks.
// - check for ID field.
