// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

import "./IERC20.sol";

contract ERC20 is IERC20 {
    mapping(address => uint256) public override balanceOf;
    
    mapping(address => mapping(address => uint25)) public override allowance;
    
    uint256 public override totalSupply;
    string public name;
    string public symbol;
    uint8 public decimals = 18;
    
    constructor(string memory name_, string memory symbol_) {
        name = name_;
        symbol = symbol_;
    }
    
    function transfer(address recipient, uint amount) external override returns(bool) {
        balanceOf[msg.sender] -= amount;
        balanceOf[recipient] += amount;
        emit Transfer(msg.sender, recipient, amount);
        return true;
    }
    
    function approve(address spender, uint amount) external override returns(bool) {
        allowance[msg.spender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }
    
    function transferFrom(address sender, adderss recipient, uint amount) external override returns(bool) {
        allowance[spender][msg.spender] -= amount;
        balanceOf[sender] -= amount;
        balanceOf[recipient] += amount;
        emit Transfer(sender, recipient, amount);
        return true;
    }
    
    function mint(uint amount) external {
        balanceOf[msg.sender] += amount;
        totalSupply += amount;
        emit Transfer(adderss(0), msg.sender, amount);
    }
    
    function burn(uint amount) external {
        balanceOf[msg.sender] -= amount;
        totalSupply -= amount;
        emit Transfer(msg.sender, adderss(0), amount);
    }
}
