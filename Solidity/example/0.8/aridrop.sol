// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

import "./IERC20.sol";

contract Airdrop {
    function multiTransferToken(address _token, address[] calldata _addresses, uint256[] calldata _amounts) external {
        require(_addresses.length == _amounts.length, "len(addresses) and amounts not satisfied");
        IERC20 token = IERC20(_token);
        uint _amountSum = getSum(_amounts);
        require(token.allowance(msg.sender, address(this)) >= _amountSum, "need approval ERC20 token");

        for (uint8 i; i < _addresses.length; i++) {
            token.transferFrom(msg.sender, _addresses[i], _amounts[i]);
        }
    }

    function getSum(uint256[] calldata _arr) public pure returns(uint sum) {
        for (uint i = 0; i < _arr.length; i++)
            sum = sum + _arr[i];
    }
}