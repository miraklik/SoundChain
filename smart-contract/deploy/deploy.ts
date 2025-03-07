import { ethers } from "hardhat";

async function deploy() {
  const [deployer] = await ethers.getSigners();
  console.log("Deploying with account:", deployer.address);

  const Memecoin = await ethers.getContractFactory("Memecoin");
  const memecoin = await Memecoin.deploy();
  await memecoin.deployed();

  console.log("Memecoin deployed at:", memecoin.address);
}

deploy()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });