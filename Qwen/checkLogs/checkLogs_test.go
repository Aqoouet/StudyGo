package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testData = []TestCase{}
var testFiles = []string{}

func init() {
	testData = ParseTests("examples.txt", true)
	for i := 1; i <= 39; i++ {
		testFiles = append(testFiles, fmt.Sprintf("testData/log%d.txt", i))
	}
}

func TestGetConsole(t *testing.T) {
	for _, tt := range testData {
		t.Run(tt.Name+" | "+tt.Descr, func(t *testing.T) {
			ans := getConsole("testData/" + tt.Name)
			if !reflect.DeepEqual(tt.Console, ans) {
				t.Errorf("want %s, get %s", tt.Console, ans)
			}
		})
	}
}

func TestGetCSV(t *testing.T) {
	for _, tt := range testData {
		t.Run(tt.Name+" | "+tt.Descr, func(t *testing.T) {
			ans := getCSV("testData/" + tt.Name)
			if !reflect.DeepEqual(tt.CSV, ans) {
				t.Errorf("want %s, get %s", tt.CSV, ans)
			}
		})
	}
}

func TestGetJSON(t *testing.T) {
	for _, tt := range testData {
		t.Run(tt.Name+" | "+tt.Descr, func(t *testing.T) {
			ans := getJSON("testData/" + tt.Name)
			if !reflect.DeepEqual(tt.JSON, ans) {
				t.Errorf("want %s, get %s", tt.JSON, ans)
			}
		})
	}
}
