# Solidity

## YSK

- `type permission varName = value;`
  - type:
    - int, uint, uint256
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
  - permission: public, payable public
  - varName: N/A
  - value:
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
  - calldata
    - same with memory but used for function parameter and cannot modify immutable var

- var scope
  - state
    - the data is stored in the variables on the chain, all in-contract functions can access it, and the gas consumption is high
  - local
  - global

- credit
  - https://remix.ethereum.org/
  - https://github.com/AmazingAng/WTF-Solidity
