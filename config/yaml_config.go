package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Scheduler Scheduler `yaml:"scheduler"`
}

type Scheduler struct {
	Config YamlConfig `yaml:"config"`
}

type YamlConfig struct {
	Schedules map[string]Schedule `yaml:"schedules"`
}

type Schedule struct {
	Timezone string        `yaml:"timezone"`
	Weekdays []WeekdayInfo `yaml:"weekdays"`
	Holidays []HolidayInfo `yaml:"holidays"`
}

type WeekdayInfo struct {
	Days  []int  `yaml:"days"`
	Start string `yaml:"start"`
	Stop  string `yaml:"stop"`
}

type HolidayInfo struct {
	Start string `yaml:"start"`
	Stop  string `yaml:"stop"`
}

func SchedulerConfig() YamlConfig {
	const schedulerConfigFilePath = "scheduler.yml"

	var config Config

	if _, err := os.Stat(schedulerConfigFilePath); err == nil {
		configData, err := os.ReadFile(schedulerConfigFilePath)
		if err != nil {
			fmt.Println("Error reading YAML config file:", err)
			return YamlConfig{}
		}

		err = yaml.Unmarshal(configData, &config)
		if err != nil {
			fmt.Println("Error decoding YAML config file:", err)
			return YamlConfig{}
		}
	}

	return config.Scheduler.Config
}
