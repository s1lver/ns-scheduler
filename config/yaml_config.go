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
	Timezone string   `yaml:"timezone"`
	Weekdays []string `yaml:"weekdays"`
	Holidays []string `yaml:"holidays"`
}

func SchedulerConfig() YamlConfig {
	const schedulerConfigFilePath = "scheduler.yml"

	var config Config

	if _, err := os.Stat(schedulerConfigFilePath); err == nil {
		configData, err := os.ReadFile(schedulerConfigFilePath)
		if err != nil {
			fmt.Println("Ошибка чтения файла конфигурации YAML:", err)
			return YamlConfig{}
		}

		err = yaml.Unmarshal([]byte(configData), &config)
		if err != nil {
			fmt.Println("Ошибка декодирования файла конфигурации YAML:", err)
			return YamlConfig{}
		}
	}

	return config.Scheduler.Config
}
