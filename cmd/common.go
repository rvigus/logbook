package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const Filename = "logbook.json"

type Logbook struct {
	Records []Record `json:"records"`
}

type Record struct {
	Activity  string `json:"activity"`
	Impact    string `json:"impact"`
	Category  string `json:"category"`
	Timestamp string `json:"timestamp"`
}

func (r *Record) timestampToDate() string {
	t, err := time.Parse(time.RFC3339, r.Timestamp)
	if err != nil {
		fmt.Println("Error parsing timestamp:", err)
	}
	date := t.Format("2006-01-02")
	return date
}

func (r *Record) timestampToTime() string {
	t, err := time.Parse(time.RFC3339, r.Timestamp)
	if err != nil {
		fmt.Println("Error parsing timestamp:", err)
	}
	timeStr := t.Format("15:04:05")
	return timeStr
}

func ReadLogbook(filename string) (Logbook, error) {
	emptyLogbook := Logbook{Records: []Record{}}

	bytes, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return emptyLogbook, err
	}

	var logbook Logbook
	if len(bytes) > 0 {
		if err = json.Unmarshal(bytes, &logbook); err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return emptyLogbook, err
		}
	}
	return logbook, nil
}
