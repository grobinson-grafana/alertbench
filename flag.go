package main

import (
	"flag"
	"os"
)

var (
	delete        bool
	destURL       string
	file          string
	token         string
	username      string
	password      string
	offset        int
	rules         int
	rulesPerGroup int
)

func fromEnv(name, or_default string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return or_default
}

func parseFlags() {
	flag.BoolVar(&delete, "delete", false, "Delete all provisioned rules")
	flag.StringVar(&destURL, "url", "http://127.0.0.1:3000", "URL of the Grafana server")
	flag.StringVar(&file, "file", "example.json", "The file with the example rule")
	flag.StringVar(&token, "bearer-token", fromEnv("BEARER_TOKEN", ""), "The bearer token ($BEARER_TOKEN)")
	flag.StringVar(&username, "username", fromEnv("USERNAME", "admin"), "The username ($USERNAME)")
	flag.StringVar(&password, "password", fromEnv("PASSWORD", "password"), "The password ($PASSWORD)")
	flag.IntVar(&offset, "offset", 0, "The offset that rules should be provisioned from")
	flag.IntVar(&rules, "rules", 100, "The number of rules that should be provisioned")
	flag.IntVar(&rulesPerGroup, "rules-per-group", 10, "The maximum number of rules per rule group")
	flag.Parse()
}
