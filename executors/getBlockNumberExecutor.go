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
)

func GetBlockNumber() (int64, error){
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
		fmt.Sscanf(blockHexNumber, "0x%x", &blockNumber)

		fmt.Println("the block number is:  ", blockNumber)

		return blockNumber, nil
	}

	return 0, err

}

