package systemcontract

import (
	"math/big"
	"strings"

	"github.com/Project-InitVerse/chain/accounts/abi"
	"github.com/Project-InitVerse/chain/common"
	"github.com/Project-InitVerse/chain/params"
)

// ValidatorsInteractiveABI contains all methods to interactive with validator contracts.

const AddrListInteractiveABI = `
[
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "newAdmin",
          "type": "address"
        }
      ],
      "name": "AdminChanged",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "newAdmin",
          "type": "address"
        }
      ],
      "name": "AdminChanging",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "addr",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "enum AddressList.Direction",
          "name": "d",
          "type": "uint8"
        }
      ],
      "name": "BlackAddrAdded",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "addr",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "enum AddressList.Direction",
          "name": "d",
          "type": "uint8"
        }
      ],
      "name": "BlackAddrRemoved",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "addr",
          "type": "address"
        }
      ],
      "name": "DeveloperAdded",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "addr",
          "type": "address"
        }
      ],
      "name": "DeveloperRemoved",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bool",
          "name": "newState",
          "type": "bool"
        }
      ],
      "name": "EnableStateChanged",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "a",
          "type": "address"
        },
        {
          "internalType": "enum AddressList.Direction",
          "name": "d",
          "type": "uint8"
        }
      ],
      "name": "addBlacklist",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "addr",
          "type": "address"
        }
      ],
      "name": "addDeveloper",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "admin",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_stor",
          "type": "address"
        }
      ],
      "name": "checkValid",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "newAdmin",
          "type": "address"
        }
      ],
      "name": "commitChangeAdmin",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "confirmChangeAdmin",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "devVerifyEnabled",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "disableDevVerify",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "enableDevVerify",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getBlacksFrom",
      "outputs": [
        {
          "internalType": "address[]",
          "name": "",
          "type": "address[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getBlacksTo",
      "outputs": [
        {
          "internalType": "address[]",
          "name": "",
          "type": "address[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "initialize",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "initialized",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "a",
          "type": "address"
        }
      ],
      "name": "isBlackAddress",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        },
        {
          "internalType": "enum AddressList.Direction",
          "name": "",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "addr",
          "type": "address"
        }
      ],
      "name": "isDeveloper",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "pendingAdmin",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "a",
          "type": "address"
        },
        {
          "internalType": "enum AddressList.Direction",
          "name": "d",
          "type": "uint8"
        }
      ],
      "name": "removeBlacklist",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "addr",
          "type": "address"
        }
      ],
      "name": "removeDeveloper",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "_group",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "_value",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "_storage",
          "type": "uint256"
        }
      ],
      "name": "setValid",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`

const MinerDelegateABI = `[
        {
            "inputs": [],
            "stateMutability": "nonpayable",
            "type": "constructor"
        },
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "minter",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "delegate",
                    "type": "address"
                }
            ],
            "name": "SET_DELEGATE",
            "type": "event"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "",
                    "type": "address"
                }
            ],
            "name": "minter_delegate",
            "outputs": [
                {
                    "internalType": "address",
                    "name": "",
                    "type": "address"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "_delegate",
                    "type": "address"
                }
            ],
            "name": "set_delegate",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        }
    ]`

const DevMappingPosition = 2

var (
	/*
		FactoryAdminAddr        = common.HexToAddress("0x7a0BbA5EEbD9B84F46A39A9ffd488b8afB88979d")
		SysGovContractName      = "governance"
		AddressListContractName = "address_list"
		DposFactoryContractName = "dpos_factory"
		PunishV1ContractName    = "punish_v1"
		SysGovContractAddr      = common.HexToAddress("0x000000000000000000000000000000000000c000")
		AddressListContractAddr = common.HexToAddress("0x000000000000000000000000000000000000c001")
		DposFactoryContractAddr = common.HexToAddress("0x000000000000000000000000000000000000c002")
		PunishV1ContractAddr    = common.HexToAddress("0x000000000000000000000000000000000000c003")
		// SysGovToAddr is the To address for the system governance transaction, NOT contract address
		SysGovToAddr = common.HexToAddress("0x000000000000000000000000000000000000cccc")
	*/

	AddressListContractName          = "address_list"
	AddressListContractAddr          = common.HexToAddress("0x000000000000000000000000000000000000C001")
	MinerDelegateContractAddr        = common.HexToAddress("0x000000000000000000000000000000000000C009")
	ValidMethod                      = "checkValid"
	AddressListContractAdminAddr     = common.HexToAddress("0x2b9ac060e7d20cf91bbb6719178d957f9c441235")
	AddressListTestContractAdminAddr = common.HexToAddress("0x2b9ac060e7d20cf91bbb6719178d957f9c441235")
	AddressListDevContractAdminAddr  = common.HexToAddress("0x2b9ac060e7d20cf91bbb6719178d957f9c441235")

	abiMap map[string]abi.ABI
)

func init() {
	abiMap = make(map[string]abi.ABI, 0)

	tmpABI, _ := abi.JSON(strings.NewReader(AddrListInteractiveABI))
	abiMap[AddressListContractName] = tmpABI

}

func GetInteractiveABI() map[string]abi.ABI {

	return abiMap
}

func GetAddressListAdmin(chainId *big.Int) common.Address {
	if chainId.Cmp(params.MainnetChainConfig.ChainID) == 0 {
		return AddressListContractAdminAddr
	} else if chainId.Cmp(params.TestnetChainConfig.ChainID) == 0 {
		return AddressListTestContractAdminAddr
	} else if chainId.Cmp(params.DevnetChainConfig.ChainID) == 0 {
		return AddressListDevContractAdminAddr
	} else {
		return AddressListContractAdminAddr
	}
}
