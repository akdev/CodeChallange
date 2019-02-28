package main

import (
	"log"
	"os"
	"testing"
)

const (
	HAPPYPATH_FILE  = "./data/score_recs.data"
	EMPTY_FILE      = "./data/score_recs_empty.data"
	EMPTYLINES_FILE = "./data/score_recs_emptyLines.data"
	BADJSON_FILE    = "./data/score_recs_badJson.data"
	LINEBREAKS_FILE = "./data/score_recs_lineBreaks.data"
	NOID_FILE       = "./data/score_recs_noId.data"
	BAD_FILE_NAME   = "./data/NoSuchFile"
	N_ENTRIES       = 5
)

func TestHappyPath(*testing.T) {
	if HighestNScores(HAPPYPATH_FILE, N_ENTRIES) != 0 {
		log.Fatal("Expected 0 exit code")
	}
}

func TestEmptyLines(*testing.T) {
	if HighestNScores(EMPTYLINES_FILE, N_ENTRIES) != 0 {
		log.Fatal("Expected 0 exit code")
	}
}
func TestEmptyFile(*testing.T) {
	if HighestNScores(EMPTY_FILE, N_ENTRIES) != 2 {
		log.Fatal("Expected 2 exit code")
	}
}

func TestBadJSON(*testing.T) {
	if HighestNScores(BADJSON_FILE, N_ENTRIES) != 2 {
		log.Fatal("Expected 2 exit code")
	}
}

func TestLineBreaks(*testing.T) {
	if HighestNScores(LINEBREAKS_FILE, N_ENTRIES) != 2 {
		log.Fatal("Expected 2 exit code")
	}
}

func TestNoIds(*testing.T) {
	if HighestNScores(NOID_FILE, N_ENTRIES) != 2 {
		log.Fatal("Expected 2 exit code")
	}
}

func TestBadFileName(*testing.T) {
	if HighestNScores(BAD_FILE_NAME, N_ENTRIES) != 1 {
		log.Fatal("Expected 1 exit code")
	}
}
func TestMain(m *testing.M) {
	// log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}
