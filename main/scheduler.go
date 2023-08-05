package main

import (
	"fmt"
	schedulerConfig "ns-scheduler/config"
	"time"
)

func main() {
	var config schedulerConfig.YamlConfig

	// Load config from yaml if exist
	config = schedulerConfig.SchedulerConfig()

	// Accessing the data from the YAML structure
	standardSchedule := config.Schedules["standard"]

	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	dayOfWeek := int(date.Weekday())

	for _, weekday := range standardSchedule.Weekdays {
		for days := range weekday.Days {
			if dayOfWeek == days {
				fmt.Printf("Days: %v\n", weekday.Days)
				fmt.Printf("Start: %s\n", weekday.Start)
				fmt.Printf("Stop: %s\n", weekday.Stop)
			}
		}
	}
}
