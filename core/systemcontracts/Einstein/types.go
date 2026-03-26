package einstein

import _ "embed"

// contract codes for Mainnet upgrade
var (
	//go:embed mainnet/MinerDelegateContract
	MainnetDelegateContract string
	//go:embed testnet/MinerDelegateContract
	TestnetDelegateContract string
)
