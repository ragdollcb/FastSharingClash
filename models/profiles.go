package models

import (
	"fmt"
	"os"
)

func ReadProfiles() string {
	file, err := os.ReadFile("config/profiles.yaml")
	if err != nil {
		fmt.Print(err)
		return ""
	}
	return string(file)
}
