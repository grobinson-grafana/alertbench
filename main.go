package main

import (
	"log"
)

func main() {
	parseFlags()

	example, err := ExampleFromFile(file)
	if err != nil {
		log.Fatalf("Failed to open example: %s", err)
	}

	// make sure that Grafana is working and username/password are correct
	if err := ping(); err != nil {
		log.Fatalf("Failed to ping: %s", err)
	}

	if delete {
		log.Println("Deleting provisioned rules")
		deleteProvisionedRules()
	} else {
		log.Printf("Provisioning %d rules, at most %d rules per group", rules, rulesPerGroup)
		provisionRules(example)
	}
}
