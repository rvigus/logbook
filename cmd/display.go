package cmd

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "Displays the logbook",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		display()
	},
}

func init() {
	rootCmd.AddCommand(displayCmd)
}

func display() {

	var logbook Logbook
	var err error

	logbook, err = ReadLogbook(Filename)
	if err != nil {
		fmt.Println("Error reading logbook: ", err)
		return
	}

	sortRecordsAsc(logbook)
	displayTable(logbook)
}

func sortRecordsAsc(logbook Logbook) {
	sort.Slice(logbook.Records, func(i, j int) bool {
		t1, _ := time.Parse(time.RFC3339, logbook.Records[i].Timestamp)
		t2, _ := time.Parse(time.RFC3339, logbook.Records[j].Timestamp)
		return t1.Before(t2)
	})
}

func displayTable(logbook Logbook) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Activity", "Impact", "Category", "Date", "Time"})

	for i, record := range logbook.Records {
		t.AppendRow(table.Row{
			i, record.Activity, record.Impact, record.Category, record.timestampToDate(), record.timestampToTime(),
		})
	}
	t.SetStyle(table.StyleRounded)
	t.Render()
}
