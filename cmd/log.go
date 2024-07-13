package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log()
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}

func log() {

	activity, impact, category, timestamp := readInput()

	entry := Record{
		Timestamp: timestamp,
		Activity:  activity,
		Impact:    impact,
		Category:  category,
	}

	var err error
	var file *os.File
	var logbook Logbook

	// Create if not exists
	if _, err = os.Stat(Filename); os.IsNotExist(err) {
		file, err = os.Create("logbook.json")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		file.Close()
	}

	logbook, err = ReadLogbook(Filename)
	if err != nil {
		fmt.Println("Error reading logbook: ", err)
		return
	}

	logbook.Records = append(logbook.Records, entry)

	updatedData, err := json.MarshalIndent(logbook, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	if err = os.WriteFile("logbook.json", updatedData, 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Entry logged successfully.")
}

func readInput() (string, string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Task: ")
	activity, _ := reader.ReadString('\n')
	activity = strings.Replace(activity, "\n", "", -1)

	fmt.Print("Impact: ")
	impact, _ := reader.ReadString('\n')
	impact = strings.Replace(impact, "\n", "", -1)

	fmt.Print("Category: ")
	category, _ := reader.ReadString('\n')
	category = strings.Replace(category, "\n", "", -1)

	timestamp := time.Now().Format(time.RFC3339)

	return activity, impact, category, timestamp
}
