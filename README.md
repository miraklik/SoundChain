# 🎵 SoundChain — Web3 Music Platform

**SoundChain** is a decentralized music application that empowers artists to create, promote, and monetize their own crypto tokens. Built on Web3 technologies, SoundChain integrates smart contracts, a Go backend, and a user-friendly frontend to revolutionize the music industry.

---

## 🚀 Features

- 🎤 Artists can mint and manage their own crypto tokens
- 🎧 Users can stream and purchase music using tokens
- 🔗 Seamless integration with blockchain networks
- 🛠️ Backend services developed in Go
- 📜 Smart contracts written in Solidity
- 🌐 Web interface for user interaction

---

## 🛠️ Tech Stack

- **Backend:** Go (Golang)
- **Smart Contracts:** Solidity
- **Frontend:** In development
- **Blockchain Integration:** Web3.js or Ethers.js
- **Containerization:** Docker

---

## 📁 Project Structure

```
SoundChain/
├── backend/ # Go backend services
├── smart-contract/ # Solidity smart contracts
├── frontend/ #Frontend
├── .github/ # GitHub workflows
├── README.md # Project documentation
└── Dockerfile # Docker configuration
```

## 📦 Installation & Setup

### 1. Clone the repository

```bash
git clone https://github.com/miraklik/SoundChain.git
cd SoundChain
```

### 2. Set up the backend

```bash 
cd backend
go mod tidy
go run main.go
```

### 3. Deploy smart contracts

```bash
cd smart-contract
npm install
npx hardhat compile
npx hardhat run scripts/deploy.js --network mainnet
```
### 4. Run with Docker (optional)

```bash
docker build -t soundchain .
docker run -p 8080:8080 soundchain
```
