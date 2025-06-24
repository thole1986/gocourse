package main

import (
	"fmt"
	"time"
)

func main() {
	// Current local time

	fmt.Println(time.Now())

	// Specific time
	specificTime := time.Date(2024, time.July, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(specificTime)

	// Parse time
	// Convert time string to time object

	// Must be followed the format: (Mon Jan 2 15:04:05 MST 2006)
	parsedTime, _ := time.Parse("2006-01-02", "2024-05-01")
	parsedTime1, _ := time.Parse("06-01-02", "20-05-01") // yy-mm-dd
	parsedTime2, _ := time.Parse("06-1-2", "20-5-1")
	parsedTime3, _ := time.Parse("06-1-2 15-04", "20-5-1 18-03") // yy-mm-dd hh:ss
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)
	fmt.Println(parsedTime3)

	// Formatting time
	t := time.Now()
	// Convert time object to time string.
	fmt.Println("Formatted time: ", t.Format("Monday 06-01-02 15-04-05"))

	// Current time + 24 hours to get next datetime
	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	// Get the day name of next day
	fmt.Println(oneDayLater.Weekday())

	// Mean 12:16 -> Rounded to 12:00
	fmt.Println("Rounded Time:", t.Round(time.Hour))

	// loc, _ := time.LoadLocation("Asia/Kolkata") // Diff timezone
	// t = time.Date(2024, time.July, 8, 14, 16, 40, 00, time.UTC)

	// // Convert this to the specifc time zone.
	// tLocal := t.In(loc)

	// //  Perform rounding
	// roundedTime := t.Round(time.Hour)
	// roundedTimeLocal := roundedTime.In(loc)
	// fmt.Println("Original Time (UTC): ", t)
	// fmt.Println("Original Time (Local): ", tLocal)
	// fmt.Println("Rounded Time (UTC): ", roundedTime)
	// fmt.Println("Rounded Time (UTC): ", roundedTimeLocal)

	fmt.Println("Truncated Time:", t.Truncate(time.Hour))

	loc, _ := time.LoadLocation("America/New_York")

	// Convert time to location
	tInNY := time.Now().In(loc)

	fmt.Println("New York Time: ", tInNY)

	t1 := time.Date(2024, time.July, 4, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, time.July, 4, 18, 0, 0, 0, time.UTC)

	// Calculate duration between 2 time
	duration := t2.Sub(t1)
	fmt.Println("Duration:", duration)

	//  Compare times
	fmt.Println("t2 is after t1?", t2.After(t1))

}
