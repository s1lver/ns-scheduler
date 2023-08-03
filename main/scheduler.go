package main

import (
	"fmt"
	schedulerConfig "ns-scheduler/config"
)

func main() {
	var config schedulerConfig.YamlConfig

	// Load config from yaml if exist
	config = schedulerConfig.SchedulerConfig()

	// Accessing the data from the YAML structure
	standardSchedule := config.Schedules["standard"]

	for _, weekday := range standardSchedule.Weekdays {
		fmt.Printf("Days: %v\n", weekday.Days)
		fmt.Printf("Start: %s\n", weekday.Start)
		fmt.Printf("Stop: %s\n", weekday.Stop)
	}
}
