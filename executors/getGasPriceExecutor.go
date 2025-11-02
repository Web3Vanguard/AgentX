package executors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jieliu2000/anyi/flow"
	"github.com/joho/godotenv"
)


type GetGasPriceExecutor struct {}

type GetGasPriceStepWrapper struct{}

func (e *GetGasPriceExecutor) Execute(context *flow.FlowContext) (*flow.FlowContext, error) {
	err := godotenv.Load()

	rpc_url := os.Getenv("SOMNIA_RPC_URL")

	if err != nil {
		log.Printf("Could not read .env: %v", err)
	}

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_gasPrice",
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
		log.Fatal("Error parsing JSON: ", err)
	}

	
	if result["result"] != nil {
		gasPriceData := result["result"].(string)

		var gasPrice int64

		_, err := fmt.Sscanf(gasPriceData, "0x%x", &gasPrice)

		if err != nil {
			log.Fatal("Error in number conversion to from hex.")
		}


		resultText := fmt.Sprintf("The current gasprice is %d", gasPrice)

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

func (s *GetGasPriceStepWrapper) Init() error {
	return nil
}

func (s *GetGasPriceStepWrapper) Run(flowContext flow.FlowContext, Step *flow.Step) (*flow.FlowContext, error) {
	getGasPriceExecutor := GetGasPriceExecutor{};
	context := flow.NewFlowContext("get_current_gas_price", "somnia blockchain interaction");
	newFlowContext, err := getGasPriceExecutor.Execute(context)
	if err != nil {
		return nil, err
	}
	return newFlowContext, nil
}

func NewGetGasPriceStepWrapper() *GetGasPriceStepWrapper {
	return &GetGasPriceStepWrapper{}
}
