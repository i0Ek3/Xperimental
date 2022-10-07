// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract Return {
    //function returnUnnamedValues() public pure returns(uint256 _n, bool _b, uint256[3] memory _arr) {
    function returnUnnamedValues() public pure returns(uint256, bool, uint256[3] memory) {
        return (1, true, [uint256(1), 2, 5]);
    }

    function returnNamedValues() public pure returns(uint256 _n, bool _b, uint256[3] memory _arr) {
        _n = 1;
        _b = true;
        _arr = [uint256(1), 2, 3];
    }

    function readValues() public pure {
        uint256 _n;
        bool _b1, _b2;
        uint256[3] memory _arr;
        (_n, _b1, _arr) = returnNamedValues();

        (_, _b2, _arr) = returnNamedValues();
    }

    uint[] X = [1, 2, 3];

    function storageExample() public {
        uint[] storage x = X;
        x[0] = 100;
        // print(X) = [100, 2, 3], you can think this with Go slice
    }

    function MemoryExample() public view {
        uint[] memory x = X;
        x[0] = 100;
        // print(X) = [1, 2, 3]
    }

    function calldataExample(uint[] calldata _x) public pure returns(uint[] calldata) {
        //_x[0] = 0; // panic
        return(_x);
    }
}
