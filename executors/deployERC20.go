package executors

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jieliu2000/anyi/flow"
	"github.com/joho/godotenv"
	"github.com/shaaibu7/AgentX/bindings"
	"log"
	"math/big"
	"os"
)

type DeployERC20TokenExecutor struct{}

type DeployERC20TokenStepWrapper struct{}

func (e *DeployERC20TokenExecutor) Execute(ctx *flow.FlowContext) (*flow.FlowContext, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Environment variables cannot be loaded: %v", err)
	}

	rpc_url := os.Getenv("SOMNIA_RPC_URL")
	private_key := os.Getenv("PRIVATE_KEY")


	// Connect to an Ethereum node
	client, err := ethclient.Dial(rpc_url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(private_key)
	if err != nil {
		log.Fatal(err)
	}

	// Get the account's public address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("Deploying contract from:", fromAddress.Hex())

	// Get the nonce (transaction count)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("nonce error:", err)
	}

	// Set gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("gas price error", err)
	}

	// Create the authenticated transactor
	chainID, err := client.NetworkID(context.Background())
	fmt.Println("The chain id: ", chainID)
	if err != nil {
		log.Fatal("network id error", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("keyed tx error", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // no ETH sent with the contract
	// auth.GasLimit = uint64(30000000000)   // Gas limit
	auth.GasPrice = gasPrice

	// Deploy contract
	address, tx, _, err := bindings.DeployBindings(auth, client)
	if err != nil {
		log.Fatal("deployment error ", err)
	}

	fmt.Println("Contract deployed to address:", address.Hex())
	fmt.Println("Transaction hash:", tx.Hash().Hex())

	// Wait for mining
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("Failed to get tx receipt: %v", err)
	}

	if receipt.Status == 1 {
		resultText := fmt.Sprintf("The deployed contract address is and the txHash is %s %s", address.Hex(), tx.Hash().Hex())

		newContext := flow.FlowContext{
			Text:      resultText,
			Memory:    ctx.Memory,
			Think:     ctx.Think,
			ImageURLs: ctx.ImageURLs,
			Flow:      ctx.Flow,
			Variables: ctx.Variables,
		}

		return &newContext, nil
	}

	return nil, err

}

func (s *DeployERC20TokenStepWrapper) Init() error {
	return nil
}

func (s *DeployERC20TokenStepWrapper) Run(flowContext flow.FlowContext, Step *flow.Step) (*flow.FlowContext, error) {
	deployERC20TokenExecutor := DeployERC20TokenExecutor{};
	context := flow.NewFlowContext("deploy_erc20_token", "somnia blockchain interaction");
	newFlowContext, err := deployERC20TokenExecutor.Execute(context)
	if err != nil {
		return nil, err
	}
	return newFlowContext, nil
}

func NewDeployERC20TokenStepWrapper() *DeployERC20TokenStepWrapper {
	return &DeployERC20TokenStepWrapper{}
}
