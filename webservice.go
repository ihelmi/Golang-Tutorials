package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// DateInterval struct for JSON
type DateInterval struct {
	StartDate string   `json:"start_date"`
	Intervals []string `json:"intervals"`
	Results   []string `json:"results"`
}

func dateUpdate(w http.ResponseWriter, r *http.Request) {
	// Declare a new Person struct.
	var dateinterval DateInterval

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&dateinterval)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the start date and intervals
	StartDate := dateinterval.StartDate
	Intervals := dateinterval.Intervals
	// Manipulate the start date for easy changing
	t, err := time.Parse(time.RFC3339, StartDate)
	if err != nil {
		log.Fatal(err)
	}
	// Create variables to save the calculated time
	var Hour, Min, Sec int
	var calculated []string
	// Iterate through each interval to unmerge and calculate the time unit
	for i := 0; i < len(Intervals); i++ {

		// Iterate through each interval and unmerge them
		UnmergedIntervals := strings.FieldsFunc(Intervals[i], Split)

		// Create variable to chech the time unit to calculate
		H := strings.Contains(Intervals[i], "h")
		M := strings.Contains(Intervals[i], "m")
		S := strings.Contains(Intervals[i], "s")

		for x := 0; x < len(UnmergedIntervals); x++ {
			if H {
				Hour, _ = strconv.Atoi(UnmergedIntervals[x])
				H = false
			} else if M {
				Min, _ = strconv.Atoi(UnmergedIntervals[x])
				M = false
			} else if S {
				Sec, _ = strconv.Atoi(UnmergedIntervals[x])
				S = false
			}
		}
		// Calculate the new date and time
		NewCalculatedDate := t.Add(time.Hour*time.Duration(Hour) +
			time.Minute*time.Duration(Min) +
			time.Second*time.Duration(Sec))
		// Convert the date-time formate
		FormatedTime := NewCalculatedDate.Format(time.RFC3339)
		// Reset the time variables
		Hour = 0
		Min = 0
		Sec = 0
		// Do something with the struct
		calculated = append(calculated, FormatedTime)
	}
	// Display the results as JSON
	Output := DateInterval{
		StartDate: StartDate,
		Intervals: Intervals,
		Results:   calculated,
	}
	json.NewEncoder(w).Encode(Output)
}

//Split Function to split each intervals into h,m,sassuimg that the pattern of each interval is h,m,s respectively
func Split(r rune) bool {
	return r == 'h' || r == 'm' || r == 's'
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/date/update", dateUpdate)
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
