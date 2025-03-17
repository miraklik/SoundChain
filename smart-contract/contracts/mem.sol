// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "../.deps/npm/@uniswap/v2-core/contracts/interfaces/IUniswapV2Factory.sol";
import "../.deps/npm/@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

contract MemCoin is ERC20, Ownable, ReentrancyGuard, Pausable {
    error TransferFromZero();
    error TransferToZero();
    error ArithmeticError();
    error TokenAmountMustBePositive();
    error TaxTooHigh();
    error InsufficientFunds();

    uint256 public buyTax = 5;
    uint256 public sellTax = 10;
    uint256 public liquidityShare = 50;
    uint256 public burnShare = 20;
    uint256 public rewardsShare = 30;

    IUniswapV2Router02 public uniswaprouter;
    address public uniswapPair;
    address public constant DEAD_ADDRESS = 0x000000000000000000000000000000000000dEaD;

    mapping(address => bool) public isExcludedFromTax;

    event TaxCollected(uint256 amount);
    event LiquidityAdded(uint256 tokens, uint256 eth);
    event TokensMinted(address indexed to, uint256 amount);

    constructor(
        string memory _name, 
        string memory _symbol, 
        address routerAddress,
        address owner
    ) ERC20(_name, _symbol) Ownable(owner) {
        uniswaprouter = IUniswapV2Router02(routerAddress);
        uniswapPair = IUniswapV2Factory(uniswaprouter.factory()).createPair(address(this), uniswaprouter.WETH());

        isExcludedFromTax[owner] = true;
        isExcludedFromTax[address(this)] = true;
    }

    function transfer(address sender, address recipient, uint256 amount) external{
        require(sender != address(0), TransferFromZero());
        require(recipient != address(0), TransferToZero());

        if (isExcludedFromTax[sender] || isExcludedFromTax[recipient]) {
            super._transfer(sender, recipient, amount);
            return;
        }

        bool isBuy = recipient == uniswapPair;
        bool isSell = sender == uniswapPair;

        uint256 taxAmount;
        if(isBuy) {
            taxAmount = (amount * buyTax) / 100;
        }else if(isSell) {
            taxAmount = (amount * sellTax) / 100;
        }

        uint256 netAmount = amount - taxAmount;
        require(amount == netAmount + taxAmount, ArithmeticError());

        if (taxAmount > 0) {
            super._transfer(sender, address(this), taxAmount);
            emit TaxCollected(taxAmount);

            uint256 liquidityTokens = (taxAmount * liquidityShare) / 100;
            uint256 burnTokens = (taxAmount * burnShare) / 100;
            uint256 rewardsTokens = taxAmount - liquidityTokens - burnTokens;

            super._transfer(address(this), uniswapPair, liquidityTokens);
            
            super._burn(address(this), burnTokens);

            super._transfer(address(this), owner(), rewardsTokens);
        }

        super._transfer(sender, recipient, netAmount);
    }

    function _update(address sender, address recipient, uint256 amount) internal virtual override {
        super._update(sender, recipient, amount);
    }


    function addLiquidity(uint256 tokenAmount) external onlyOwner nonReentrant {
        require(tokenAmount > 0, TokenAmountMustBePositive());

        _approve(address(this), address(uniswaprouter), tokenAmount);

        uniswaprouter.addLiquidityETH{value: address(this).balance}(
            address(this),
            tokenAmount,
            0,
            0,
            owner(),
            block.timestamp
        );

        emit LiquidityAdded(tokenAmount, address(this).balance);
    }

    function withdraw(uint256 amount) external onlyOwner nonReentrant {
        require(msg.sender != address(0), TransferToZero());
        require(amount <= address(this).balance, InsufficientFunds());

        uint256 balance = address(this).balance;

        payable(owner()).transfer(balance);
    }

    function excludeFromTax(address account, bool excluded) external onlyOwner {
        isExcludedFromTax[account] = excluded;
    }

    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }

    function setTaxes(uint256 _buyTax, uint256 _sellTax) external onlyOwner {
        require(_buyTax <= 10 && _sellTax <= 20, TaxTooHigh());
        buyTax = _buyTax;
        sellTax = _sellTax;
    }

    function pause() external onlyOwner {
        _pause();   
    }

    function unpause() external onlyOwner {
        _unpause();
    }
}