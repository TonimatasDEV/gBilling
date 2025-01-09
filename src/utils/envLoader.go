package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadEnvFile() {
	file, err := os.Open(".env")

	if err != nil {
		log.Fatal("Error opening the env fil:", err)
	}

	defer func(file *os.File) {
		err := file.Close()

		if err != nil {
			log.Println("Error closing env file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		err := os.Setenv(key, value)

		if err != nil {
			log.Println("Error setting env variable:", err)
		}
	}

	if scanner.Err() != nil {
		log.Fatal("Error scanning the env file: ", scanner.Err())
	}

	log.Println("Loaded env file successfully.")
}
