package website

import (
	"io/ioutil"
	"testing"
)

func TestParser(t *testing.T) {
	fileBytes, _ := ioutil.ReadFile("./data/venturebeat.html")
	ws, err := Parse(string(fileBytes))
	if err != nil {
		t.Errorf("Error while parsing: %v", err)
	}
	if ws.Title != "VentureBeat" {
		t.Errorf("Wrong Title: %v", ws.Title)
	}
}
