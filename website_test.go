package website

import (
	"io/ioutil"
	"testing"
)

func TestParser(t *testing.T) {
	fileBytes, _ := ioutil.ReadFile("./data/venturebeat.html")
	ws := Parse(string(fileBytes))
	if ws.Title != "VentureBeat" {
		t.Errorf("Wrong Title: %v", ws.Title)
	}
}
