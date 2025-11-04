package main

import (
	// "log"
	// "github.com/jieliu2000/anyi"
	// "fmt"

	"github.com/joho/godotenv"
	// "context"
	"fmt"
	"log"
	"flag"
	// "log"

	"github.com/jieliu2000/anyi"
	"github.com/shaaibu7/AgentX/executors"
	// "github.com/jieliu2000/anyi/flow"
)



func init() {
	blockStepWrapper := executors.NewGetBlockStepWrapper()
	gasPriceStepWrapper := executors.NewGetGasPriceStepWrapper()
	chainIdStepWrapper := executors.NewGetChainIdStepWrapper()
	deployERC20TokenWrapper := executors.NewDeployERC20TokenStepWrapper()
	deployNFTTokenWrapper := executors.NewDeployNFTTokenStepWrapper()




	anyi.RegisterExecutor("get_current_block_number", blockStepWrapper)

	anyi.RegisterExecutor("get_current_gas_price", gasPriceStepWrapper)

	anyi.RegisterExecutor("get_chain_id", chainIdStepWrapper)

	anyi.RegisterExecutor("deploy_erc20_token", deployERC20TokenWrapper)

	anyi.RegisterExecutor("deploy_nft_token", deployNFTTokenWrapper)



}


func main() {
	configPath := flag.String("name", "./config.yaml", "config path string")

	flag.Parse()
	configFilePath := "./"+*configPath

	
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Environment variables cannot be loaded: %v", err)
	}


	err = anyi.ConfigFromFile(configFilePath)

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	flow, err := anyi.GetFlow("agentic_flow")
	if err != nil {
		log.Fatalf("Failed to get flow: %v", err)
	}

	result, err := flow.RunWithInput("Fetch the current block number of the somnia blockchain")

	if err != nil {
		log.Fatalf("Flow execution failed: %v", err)
	} 

	log.Printf("Result: %s", result.Text)

	fmt.Println("Still processing....")

	flow, err = anyi.GetFlow("agentic_flow_price")
	if err != nil {
		log.Fatalf("Failed to get flow: %v", err)
	}


	result, err = flow.RunWithInput("Fetch the current gas price of the somnia blockchain")

	if err != nil {
		log.Fatalf("Flow execution failed: %v", err)
	}


	log.Printf("Result: %s", result.Text)


	fmt.Println("Still processing....")

	flow, err = anyi.GetFlow("agentic_flow_chain_id")
	if err != nil {
		log.Fatalf("Failed to get flow: %v", err)
	}


	result, err = flow.RunWithInput("Fetch the chain id of the somnia blockchain")

	if err != nil {
		log.Fatalf("Flow execution failed: %v", err)
	}


	log.Printf("Result: %s", result.Text)


	fmt.Println("Still processing....")

	flow, err = anyi.GetFlow("agentic_flow_deploy_erc20")
	if err != nil {
		log.Fatalf("Failed to get flow: %v", err)
	}


	result, err = flow.RunWithInput("Deploy an ERC20 token on the somnia blockchain")

	if err != nil {
		log.Fatalf("Flow execution failed: %v", err)
	}


	log.Printf("Result: %s", result.Text)


	fmt.Println("Still processing....")

	flow, err = anyi.GetFlow("agentic_flow_deploy_nft")
	if err != nil {
		log.Fatalf("Failed to get flow: %v", err)
	}


	result, err = flow.RunWithInput("Deploy an NFT token on the somnia blockchain")

	if err != nil {
		log.Fatalf("Flow execution failed: %v", err)
	}


	log.Printf("Result: %s", result.Text)


}