package newton

import _ "embed"

// contract codes for Mainnet upgrade
var (
	//go:embed mainnet/AddressListContract
	MainnetAddressListContract string
	//go:embed testnet/AddressListContract
	TestnetAddressListContract string
)
