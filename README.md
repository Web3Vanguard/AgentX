# 🧠 AgentX — Autonomous AI Agent for Somnia Blockchain

AgentX is a cutting-edge framework designed to deploy **autonomous AI agents** that interact seamlessly with the **Somnia blockchain**. Whether you're managing on-chain assets, querying smart contract data, or building intelligent decentralized workflows, AgentX empowers developers and researchers to leverage AI-driven automation in decentralized environments.

With a modular architecture, cross-platform support, and built-in blockchain connectivity, AgentX makes it easier than ever to create smart, proactive agents that act on real-time blockchain events—without human intervention.

> ⚙️ **All agent configuration and deployment logic happens through a single `config.yaml` file**. Define your agent’s behavior, blockchain credentials, and operational logic in one place—no code changes needed.

---

## 🚀 Features

- 🤖 Deploy and manage autonomous agents
- 🧩 Configure agents using a simple `config.yaml` file
- 🔗 Native Somnia blockchain integration
- Deploy Token contracts without ever touching code
- 🛠 Extensible and modular architecture
- 🔐 Secure and customizable agent execution
- ⚡ Cross-platform support (Linux, macOS, Windows)

---

## 📦 Tech Stack

- **Language:** Go (Golang)
- **Blockchain:** Somnia
- **Config:** YAML-based declarative setup


---

## 🛠 Installation

```
curl -fsSL https://raw.githubusercontent.com/Web3Vanguard/AgentX/main/scripts/install.sh | bash

curl -fsSL https://raw.githubusercontent.com/Web3Vanguard/AgentX/main/scripts/install.sh | sudo bash
```


### Prerequisites

- Go 1.21+
- Git
- Somnia Wallet / Node access

---

## 🔧 Usage

AgentX makes it incredibly easy to deploy and interact with autonomous AI agents on the Somnia blockchain using configuration files—**no code changes required**.

### 1. Create a `.env` File

Define your environment variables, including API keys and blockchain credentials:

```bash
# .env
OPENAI_API_KEY=your-openai-api-key-here
SOMNIA_RPC_URL=https://rpc.ankr.com/somnia_testnet
PRIVATE_KEY=your-private-key-here
```

---

### 2. Create a `config.yaml` File

Define your AI agents, flows, and basic blockchain interaction with a minimal configuration:

```yaml
clients:
  - name: "gpt4"
    type: "openai"
    config:
      model: "gpt-4"
      apiKey: "$OPENAI_API_KEY"
      temperature: 0.7

flows:
  - name: "agentic_flow_deploy_erc20"
    clientName: "gpt4" # Default client for the flow
    steps:
      - name: "deploy_erc20_token"
        executor:
          type: "deploy_erc20_token"
          withconfig:
            template: "Deploy an ERC20 contract on the somnia network"
            systemMessage: "deploy an ERC20 token on the somnia blockchain."
        maxRetryTimes: 2
```

---

## 🧩 Understanding the `config.yaml` Structure

The `config.yaml` file is the heart of AgentX — it tells the system **what actions your AI agent should take and how to execute them**. It is composed of two main parts:

- **Flows**: Define the sequential logic or workflows the AI agent follows.
- **Executors**: Define how the agent interacts with external systems like the Somnia blockchain.


---

## ✅ Supported Flows

| Flow Name                      | Description                                               |
|--------------------------------|-----------------------------------------------------------|
| `agentic_flow`                 | Standard blockchain-aware AI interaction                 |
| `agentic_flow_get_block_number` | Fetch and process current block number                   |
| `agentic_flow_price`           | Fetch current gas price from the Somnia network          |
| `agentic_flow_chain_id`        | Retrieve chain ID of the connected network               |
| `agentic_flow_deploy_erc20`    | Deploy an ERC20 token to the Somnia blockchain           |
| `agentic_flow_deploy_nft`      | Deploy an NFT contract to the Somnia blockchain          |

---

## ⚙️ Supported Executors

| Executor Name             | Description                                   |
|---------------------------|-----------------------------------------------|
| `get_current_block_number` | Fetches the current block number from Somnia |
| `get_current_gas_price`    | Retrieves current gas price                  |
| `get_chain_id`             | Gets the network chain ID                    |
| `deploy_erc20_token`       | Deploys an ERC20 token contract              |
| `deploy_nft_token`         | Deploys an NFT contract                      |

---

## 🧩 Flexible for Both Developers and Non-Developers

AgentX is built to be accessible to everyone — whether you're a seasoned blockchain developer or a no-code enthusiast.

### 👩‍💼 For Non-Developers (No-Code YAML Configuration)

AgentX makes it easy for anyone to deploy and run autonomous blockchain agents — **no programming skills required**.

All you need to do is:

1. **Define agent behavior** in a readable `config.yaml` file  
2. **Store credentials** securely in a `.env` file  
3. **Run the AgentX binary** with your config and environment file passed as flags

#### ✅ Example Workflow

1. Create a `.env` file with your keys:

```bash
# .env
OPENAI_API_KEY=your-openai-api-key
SOMNIA_RPC_URL=https://rpc.ankr.com/somnia_testnet
PRIVATE_KEY=your-private-key
```

### 👨‍💻 For Developers (Golang Package)

Developers can use AgentX as a Go library to build advanced AI-enabled blockchain applications programmatically. This offers maximum flexibility and full integration into existing Go-based systems.

---

## 🛠️ Roadmap & Future Features

AgentX is just getting started. We're building toward a robust platform for autonomous blockchain agents with capabilities that extend beyond basic flows.

Here are some of the exciting features coming soon:

### 🔍 Smart Contract Event Monitoring  
Enable your agents to watch for on-chain events in real time and automatically trigger workflows based on conditions (e.g., token transfers, contract state changes).

### 🚨 Real-Time Alerting & Notifications  
Set up monitoring and alerts for on-chain activities, anomalies, or predefined triggers. Integrate easily with email, SMS, Discord, Telegram, or webhook services.

### 📝 Advanced Smart Contract Interactions  
Support for interacting with any deployed smart contract — from calling functions to parsing return data and handling transaction receipts.

### 🔄 Workflow Chaining & Orchestration  
Build multi-stage workflows that respond to blockchain events, execute logic, and trigger secondary agents or webhooks.

### 🔐 Secure Multi-Agent Execution  
Run coordinated workflows between multiple agents with scoped access, wallet segregation, and enhanced encryption.


### 📦 AgentX as a Golang SDK  
Expose AgentX as a full-fledged Golang package that developers can import and use directly within their own Go applications. This unlocks deep integration and programmable control over agents, flows, and blockchain interactions.


---


## Contracts Addresses and TxHash of contracts deployed with AgentX

### ERC20 Token Contract

- Contract Address: 0x06d1b66811ee28819C1744B9a25c44db2A1a388F
- TxHash: 0xa09fb2fb6c97579bae345736d152fd42bda35b10ab5fc69d41d7621d856bafc3

### NFT Token contract:

- Contract Address: 0x5532845ee301d635377606cC192954B0E5941fE5
- TxHash: 0xcaa847bd397312d51b1c5b6e5047bae687adee77bd71d7fcd5e12e786d5d55c8
