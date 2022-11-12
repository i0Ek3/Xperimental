# Solidity Notes

## YSK

- `type permission varName = value;`
  - type:
    - int, uint256 = uint
    - array
      - T[K]: fixed length array
      - T[]: variable length array (dynamic array)
    - struct
    - mapping(KT => VT)
      - KT = key type, VT = value type
      - KT only can be the default basic type(not struct) in solidity
      - VT can be any type
    - bool
    - string
    - bytes1, bytes8, bytes32
    - enum
    - address
  - permission
    - pure
      - `pure` tells us that not only does the function not save any data to the blockchain, but it also doesn't read any data from the blockchain
    - view
      - `view` functions are only free when they're called externally
      - `view` tells us that by running the function, no data will be saved/changed
    - private
      - `private` means it's only callable from other functions inside the contract
    - public
      - `public` can be called anywhere, both internally and externally
    - internal
      - `internal` is like `private` but can also be called by contracts that inherit from this one
    - external
      - `external` can only be called outside the contract
  - varName: N/A
  - value
    - bool: true, false
    - address: 0xXXXXXXXXXX
  
- store position
  - storage
    - default store position on the chain
    - assign to local storage var will create a ref, modify this ref will modify the original var
  - memory
    - store in the memory instead of the chain
    - used for function parameter and temp var
    - assign storage var to memory var will create new copy, which means modify this new copy will not modify the original var, vice versa
    - but assign memory var to another memory var will create a ref, modify this ref will modify the original var
    - use the `memory` keyword with arrays to create a new array inside a function without needing to write anything to storage. The array will only exist until the end of the function call, and this is a lot cheaper gas-wise than updating an array in `storage`
  - calldata
    - same with memory but used for external function parameter and cannot modify immutable var
    - tell the smart contract which function I want to call and what the arguments are
  
- var scope
  - state
    - the data is stored in the variables on the chain, all in-contract functions can access it, and the gas consumption is high
  - local
  - global
  
- some keywords
  - constant/immutable
    - these two keywords can save gas
    - only number var can be declared as constant or immutable
    - string and bytes can only declared as constant
  
- control flow
  - if (condition) {} else {}
  - for (int i = 0; i < 10; i++) {}
  - while (i < n) {}
  - do {} while(i < n);
  - condition ? true_expr : false_expr;
  
- constructor
  - used to initialize parameters in the contract
  
- modifier
  - define: modifier var_name
  - commonly used methods to control smart contract permissions
  - like function but not a function, add it to end of function
  
- payable
  
  - a special type of function that can receive Ether
  
- **event**
  - event is an abstraction of logs on EVM, it has two features
    - response
    - cheaper: 2000 gas/event instead of 20000 gas on chain
  - use emit to release
    - emit kinda like output of defer

- **inheritance**
  - rules: same with C++
    - virtual
    - override
  - simple inheritance
  - multiple inheritance
  - modifier inheritance
  - constructor inheritance

- abstract contract
  - if there is at least one unimplemented function in a smart contract, that is, a function lacks the content in the body {}, the contract must be marked as abstract, otherwise the compilation will report an error

- interface
  - bytes4 selector for each function in the contract and interface id contains in that

- error(equals exception we always said)
  - `error` needs `revert` to use, just like defer and recover in Go, **recommend**
  - `require` just like exception in Python
  - `assert` same with assert in Go
    - `assert` is similar to `require`, where it will throw an error if false
    - the difference between `assert`and `require` is that `require` will refund the user the rest of their gas when a function fails, whereas `assert` will not
    - most of the time you want to use `require` in your code; `assert` is typically used when something has gone horribly wrong with the code (like a `uint` overflow)

- function overloading: same with C++
  - functions with the same name but different input parameter types
  - but modifier overloading not allowed

- library contract
  - invoke
    - `using A for B;`
    - call Strings.xxXXX() directly

  - common library
    - String: uint256 -> String
    - Address: checks if an address is a contract address
    - Create2: safely use Create2 EVM opcode
    - Arrays
  - always use SafeMath

- import
  - similar with JavaScript's import
  - import method
    - `import url;`
    - `import '@openzeppelin/xxx/xxx/xxx.sol';`
    - `import {} form './xxx.sol';`

- callback function
  - used to receive ETH and handle function calls that do not exist in the contract
  - type
    - receive()
      - only used to process receiving ETH
      - a contract has at most one receive() function
      - no need function keyword while declaration
      - can't have any parameters, can't return any value, must contain external and payable keyword
    - fallback()
      - triggered when a function that does not exist in the contract is called
      - can receiving ETH and delegate(proxy) contract
      - no need function keyword while declaration but must contain external, payable is optional

- send ETH
  - transfer()
    - receiveAddress.transfer(ETHAmount)
    - will consume 2300 gas, but the fallback() or receive() function of the counterparty contract cannot implement too complicated logic
    - if fails will revert

  - send(), not recommend
    - receiveAddress.send(ETHAmount)
    - will consume 2300 gas, but the fallback() or receive() function of the counterparty contract cannot implement too complicated logic
    - if fails cannot revert
    - return bool value

  - call(), recommend
    - receiveAddress.call{value: ETHAmount}("")
    - no gas limitation
    - if fails cannot revert
    - return (bool, data)

- call third_party deployed contract
  - pass contract address
  - pass contract var
  - create contract var
  - call contract and send ETH 

- call(): use call() to call contract
  - call the function after declaring the contract variable
  - call the target contract without knowing the source code and ABI
  - `targetContractAddress.call{value:amount, gas:gasAmount}(binaryValue);`
  - `binaryValue = abi.encodeWithSignature("函数签名", 逗号分隔的具体参数);`

- delegatecall()
  - similar with call() but have different context with call()
  - delegatecall() have security risks. When using it, ensure that the state variable storage structure of the current contract and the target contract is the same, and the target contract is safe, otherwise it will cause asset loss
  - `targetContractAddress.call{value:amount, gas:gasAmount}(binaryValue);`
  - `binaryValue = abi.encodeWithSignature("函数签名", 逗号分隔的具体参数);`

- create a new contract in a contract
  - method
    - create()
      - `Contract x = new Contract{value: _value}(params)`

    - creat2()
      - to make the contract address independent of future events
      - the contract address created with CREATE2 is determined by 4 parts
        - 0xFF
        - creator address
        - salt value
        - bytecode

      - `newAddr = hash("0xFF", creatorAddr, salt, bytecode)`
      - `Contract x = new Contract{salt: _salt, value: _value}(params)`
      - practical application scenarios
        - the exchange reserves the creation of a wallet contract address for new users
        - obtains a definite pair address

- delete contract

  - `selfdestruct(addr)`: delete the smart contract and transfer the remaining ETH of the contract to the specified address

- ABI Encode & Decode

  - standard for interaction with Ethereum smart contracts

  - contains five functions
    - abi.encode

    - abi.encodePacked

    - abi.encodeWithSignature

    - abi.encodeWithSelector

    - abi.decode

  - In Ethereum, data must be encoded in bytecode to interact with the smart contract

- hash in Solidity

  - application
    - generate a data unique identifier

    - encrypted signature

    - security encryption

  - functions
    - keccak256(), different with sha3()

- selector

  - call the target function

- ERC20

  - basically a token is just a contract that keeps track of who owns how much of that token, and some functions so those users can transfer their tokens to other addresses

  - a tokens standard in Ethereum, implement the basic logic of tokens transfer

  - IERC20 is the interface contract of the ERC20 token standard, which specifies the functions and events to be implemented by ERC20 token
    - IERC20 defines two events
      - Transfer: `event Transfer(address indexed from, address indexed to, uint256 value);`

      - Approval: `event Approval(address indexed owner, address indexed spender, uint256 value);`

    - IERC20 defines six functions
      - totalSupply(): `function totalSupply() external view returns (uint256);`

      - balanceOf(): `function balanceOf(address account) external view returns (uint256);` 

      - transfer(): `function transfer(address to, uint256 amount) external returns (bool);`

      - allowance(): `function allowance(address owner, address spender) external view returns (uint256);`

      - approve(): `function approve(address spender, uint256 amount) external returns (bool);`

      - transferFrom(): `function transferFrom(address from, address to, uint256 amount) external returns (bool);`

- token faucet

  - the token faucet is a website/app that allows users to receive tokens for free

- airdrop contracts

  - airdrop is a marketing strategy in the currency circle, the project side will be free tokens issued to specific user groups

- ERC721

  - used to abstract non-homogeneous objects
  - EIP: Ethereum Imporvement Proposals
  - ERC: Ethereum Request For Comment
  - ERC belongs EIP
  - ERC165: used to check if an item is a smart contract

- comments

  - single line comments `//`
  - multiple lines comments `/* */`
  - natspec comment `/// `

  ```solidity
  /// @title
  /// @author
  /// @notice
  
  /// @param
  /// @return
  /// @dev
  ```



## Credit

- https://cryptozombies.io
- https://remix.ethereum.org/
- https://github.com/AmazingAng/WTF-Solidity
- https://gnidan.github.io/abi-to-sol/
