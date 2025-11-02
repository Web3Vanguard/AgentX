package executors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/jieliu2000/anyi/flow"
)

type GetBlockNumberExecutor struct {}

func (e *GetBlockNumberExecutor) Execute(context *flow.FlowContext) (*flow.FlowContext, error) {
	err := godotenv.Load()

	rpc_url := os.Getenv("SOMNIA_RPC_URL")

	if err != nil {
		log.Printf("Could not read .env: %v", err)
	}

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber",
		"params":  []interface{}{},
		"id":      1,
	}

	data, err := json.Marshal(payload)

	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}

	resp, err := http.Post(rpc_url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error making request: ", err)
	}
	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	var result map[string]interface{}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	// fmt.Println("Raw JSON response", string(body), result["result"])

	if result["result"] != nil {
		blockHexNumber := result["result"].(string)

		var blockNumber int64 
		_, err := fmt.Sscanf(blockHexNumber, "0x%x", &blockNumber)

		if err != nil {
			log.Fatal("Error in number conversion to hex.")
		}
		fmt.Println("the block number is:  ", blockNumber)

		resultText := fmt.Sprintf("The current block number is %d", blockNumber)
		fmt.Println(resultText)

		newContext := flow.FlowContext{
			Text: resultText,
			Memory: context.Memory,
			Think: context.Think,
			ImageURLs: context.ImageURLs,
			Flow: context.Flow,
			Variables: context.Variables,
		}

		return &newContext, nil
	}

	return nil, err
}

