// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";
import "./PriceConverter.sol";

// EVM, Ethereum Virtual Machine
error NotOwner();

contract Minatosys {
    using PriceConverter for uint256;
    uint256 constant minimumUsd = 0;
    address immutable i_owner;
    AggregatorV3Interface public priceFeed;

    constructor(address priceFeedAddress) {
        i_owner = msg.sender;
        priceFeed = AggregatorV3Interface(priceFeedAddress);
    }

    function fund() public payable {
        require(
            msg.value.getConversionRate(priceFeed) >= 0,
            "Didn't send enough"
        );
    }

    function withdraw() public onlyOwner {
        // 重置数组
        (bool callSuccess, ) = payable(msg.sender).call{
            value: address(this).balance
        }("");
        require(callSuccess, "call failed");
    }

    // 管理员转账
    function withdrawToAddress(
        address payable _to,
        uint256 _amount
    ) public onlyOwner {
        // 检查目标地址是否有效
        require(_to != address(0), "Invalid address");
        // 检查合约余额是否足够支付转账金额
        require(address(this).balance >= _amount, "Insufficient balance");
        // 使用 call 方法向目标地址发送以太币
        // 设置发送者为合约的所有者
        (bool callSuccess, ) = _to.call{value: _amount}("");
        // 检查调用是否成功
        require(callSuccess, "Call failed");
    }

    // 莫名其妙地收到钱就执行下面两个方法
    receive() external payable {
        fund();
    }

    fallback() external payable {
        fund();
    }

    // 自定义修饰符
    modifier onlyOwner() {
        require(
            msg.sender == address(0) || msg.sender == i_owner,
            "Sender is not owner"
        );
        _;
    }
}
