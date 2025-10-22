package main

import (
	"log"
	"os"
	"fmt"
	"github.com/jieliu2000/anyi"
	"github.com/jieliu2000/anyi/llm/openai"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	
	fmt.Println(os.Getenv("OPENAI_API_KEY"))
	// Create client
	config := openai.DefaultConfig("gpt-4")
	config.APIKey = os.Getenv("OPENAI_API_KEY")
	client, err := anyi.NewClient("gpt4", config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create a two-step workflow
	step1, _ := anyi.NewLLMStepWithTemplate(
		"Generate a short story about {{.Text}}",
		"You are a creative fiction writer.",
		client,
	)
	step1.Name = "story_generation"

	step2, _ := anyi.NewLLMStepWithTemplate(
		"Create an engaging title for the following story:\n\n{{.Text}}",
		"You are an editor skilled at creating titles.",
		client,
	)
	step2.Name = "title_creation"

	// Create and register the flow
	myFlow, _ := anyi.NewFlow("story_flow", client, *step1, *step2)
	anyi.RegisterFlow("story_flow", myFlow)

	// Run the workflow
	result, _ := myFlow.RunWithInput("a detective in future London")

	log.Printf("Title: %s", result.Text)
}