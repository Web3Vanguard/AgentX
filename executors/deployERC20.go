package executors

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"github.com/shaaibu7/AgentX/bindings" // Replace with your module path
)

func DeployERC20() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Environment variables cannot be loaded: %v", err)
	}

	rpc_url := os.Getenv("SOMNIA_RPC_URL")
	private_key := os.Getenv("PRIVATE_KEY")

	fmt.Println(rpc_url, private_key)

	// Connect to an Ethereum node
	client, err := ethclient.Dial(rpc_url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Load your private key
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
	auth.Value = big.NewInt(0)        // no ETH sent with the contract
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
	fmt.Println("The reciept status is:", receipt)
	if err != nil {
		log.Fatalf("Failed to get tx receipt: %v", err)
	}

	fmt.Println("Tx reciept: ", receipt)
}
