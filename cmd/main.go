package main

import (
	// "log"
	// "github.com/jieliu2000/anyi"
	// "github.com/joho/godotenv"
	// "context"
	"fmt"
	// "log"

	"github.com/shaaibu7/AgentX/executors"
	"log"
	"github.com/jieliu2000/anyi"
	"github.com/jieliu2000/anyi/flow"
)



func init() {
	stepWrapper := executors.NewGetBlockStepWrapper()

	anyi.RegisterExecutor("get_current_block_number", stepWrapper)


}


func main() {
	getGasPriceContext := &executors.GetGasPriceExecutor{}

	// context := anyi.NewFlowContext("What is the current block number", "no need to remember anything..")
	context := flow.FlowContext{}

	result, err := getGasPriceContext.Execute(&context)

	if err != nil {
		log.Fatalf("Execution failed: %v", err)
	}

	fmt.Printf("Result: %s\n", result.Text)
	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatalf("Environment variables cannot be loaded: %v", err)
	// }


	// err = anyi.ConfigFromFile("./config.yaml")

	// if err != nil {
	// 	log.Fatalf("Failed to load config: %v", err)
	// }

	// flow, err := anyi.GetFlow("agentic_flow")
	// if err != nil {
	// 	log.Fatalf("Failed to get flow: %v", err)
	// }

	// result, err := flow.RunWithInput("Fetch the current block number of the somnia blockchain")

	// if err != nil {
	// 	log.Fatalf("Flow execution failed: %v", err)
	// }


	// log.Printf("Result: %s", result.Text)
}