// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "./mem.sol";

contract MemCoinFactory is Ownable, ReentrancyGuard, Pausable {
    event MemCoinCreated(address indexed tokenAddress, address indexed artist);

    address public routerAddress;
    address[] public allTokens;

    constructor(address _routerAddress) Ownable(msg.sender) {
        routerAddress = _routerAddress;
    }

    function createMemCoin(
        string memory name,
        string memory symbol,
        uint256 initialSupply
    ) external returns (address) {
        MemCoin token = new MemCoin(name, symbol, routerAddress, msg.sender);

        token.mint(msg.sender, initialSupply);

        allTokens.push(address(token));
        emit MemCoinCreated(address(token), msg.sender);

        return address(token);
    }

    function getAllTokens() external view returns (address[] memory) {
        return allTokens;
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }
}
