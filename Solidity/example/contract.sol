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

    function encode() public view returns(bytes memory result) {
        uint x = 10;
        address addr = 0x7A58c0Be72BE218B41C608b7Fe7C5bB630736C71;
        string name = "0xAA";
        uint[2] array = [5, 6];

        //result = abi.encode(x, addr, name, array);
        //result = abi.encodePacked(x, addr, name, array);
        //result = abi.encodeWithSignatur("foo(uint256, address, string, uint256[2])", x, addr, name, array);
        result = abi.encodeWithSelector(bytes4(keccak256("foo(uint256,address,string,uint256[2])")), x, addr, name, array);
    }

    func decode(bytes memory data) public pure returns(uint dex, address deaddr, string memory dename, uint[2] memory dearray) {
        (dex, deaddr, dename, dearray) = abi.decode(data, (uint, address, string, uint[2]));
    }
}

contract OnlyEven {
    constructor(uint a) {
        require(a != 0, "not equal zero");
        assert(a != 1);
    }

    function onlyEven(uint256 b) external pure returns(bool ok) {
        require(b % 2 == 0, "Ops, reverting...");
        ok = true;
    }
}

contract TryCatch {
    event SuccessEvent();

    event CatchEvent(string msg);
    event CatchByte(byte data);

    OnlyEven even;

    constructor() {
        even = new OnlyEven(2);
    }

    function exec(uint amount) external returns (bool ok) {
        try even.onlyEven(amount) returns(bool _ok) {
            emit SuccessEvent();
            return _ok;
        } catch Error(string memory reason) {
            emit CatchEvent(reason);
        }
    }

    function execNew(uint amount) external returns (bool ok) {
        try new OnlyEven(amount) returns(OnlyEven _even) {
            emit SuccessEvent();
            _ok = _even.onlyEven(amount);
        } catch Error(string memory reason) {
            emit CatchEvent(reason);
        } catch (bytes memory reason) {
            emit CatchByte(reason);
        }
    }
}