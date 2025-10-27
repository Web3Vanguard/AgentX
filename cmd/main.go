package main

import (
	"log"
	"github.com/jieliu2000/anyi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Environment variables cannot be loaded: %v", err)
	}


	err = anyi.ConfigFromFile("./config.yaml")

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	flow, err := anyi.GetFlow("agentic_flow")
	if err != nil {
		log.Fatalf("Failed to get flow: %v", err)
	}

	result, err := flow.RunWithInput("Interact with the somnia blockchain especially with smart contract")

	if err != nil {
		log.Fatalf("Floe execution failed: %v", err)
	}


	log.Printf("Result: %s", result.Text)
}