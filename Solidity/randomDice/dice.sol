// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;


contract Dice {
    // default: 0
    uint balance;
    
    uint public reward;

    /*
    struct player {
        uint256 remain;
        address payable addr;
    }

    mapping(address => player[]) players;
    */

    modifier check(uint _input){
        require(_input == 100, "balance must be 100");
        _;
    }

    // add balance
    function addBalance(uint _balance) public check(_balance){
        balance += _balance;
    }

    // return balance
    function getBalance() public view returns(uint){
        return balance;
    }
    
    // generate [1, 6] randomly
    function getRandom() public view returns (uint) {
        return uint(keccak256(abi.encodePacked(this, block.timestamp))) % 6 + 1;
    }

    // reward
    function Start() public returns(string memory) {
        reward = getRandom();
        require(balance >= reward, "balance unenough!");
        if (reward <= 3) {
            return "Good job!";
        } else {
            return "Ops, you lost!";
        }
        
        /*// reward A reward and reward B reward-3
        if (reward <= 3) {
            players[msg.sender][0].remain += reward;
            if (msg.sender == players[msg.sender][0].addr) {
                players[msg.sender][0].addr.transfer(reward);
            } else {
                revert("wrong");
            }
            return "Good job!";
        } else {
            players[msg.sender][1].remain += reward-3;
            if (msg.sender == players[msg.sender][1].addr) {
                players[msg.sender][1].addr.transfer(reward-3);
            } else {
                revert("wrong");
            }
            return "Ops, you lost!";
        }*/
    }
}
