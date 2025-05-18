package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testData = []TestCase{}
var testFiles = []string{}

func init() {
	testData = parseTests("examples.txt")
	for i := 1; i <= 39; i++ {
		testFiles = append(testFiles, fmt.Sprintf("log%d.txt", i))
	}
}

func TestGetConsole(t *testing.T) {
	for _, tt := range testData {
		t.Run(tt.Name, func(t *testing.T) {
			for _, file := range testFiles {
				ans := getConsole(file)
				if !reflect.DeepEqual(tt.Console, ans) {
					t.Errorf("want %s, get %s", tt.Console, ans)
				}
			}

		})
	}
}

func TestGetCSV(t *testing.T) {
	for _, tt := range testData {
		t.Run(tt.Name, func(t *testing.T) {
			for _, file := range testFiles {
				ans := getCSV(file)
				if !reflect.DeepEqual(tt.CSV, ans) {
					t.Errorf("want %s, get %s", tt.CSV, ans)
				}
			}

		})
	}
}

func TestGetJSON(t *testing.T) {
	for _, tt := range testData {
		t.Run(tt.Name, func(t *testing.T) {
			for _, file := range testFiles {
				ans := getJSON(file)
				if !reflect.DeepEqual(tt.JSON, ans) {
					t.Errorf("want %s, get %s", tt.JSON, ans)
				}
			}

		})
	}
}
