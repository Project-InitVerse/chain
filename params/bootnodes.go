// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/Project-InitVerse/chain/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main BSC network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://210d9d51a35d29ba5ef4c4c9f7512f492417e258bf71144c6eb4bd47df08170e80873f0f5768b814c08877c3f1a32eb1d60c292f2cbe20be6826f41be5dc03e0@bootnode-a.genesis-mainnet.inichain.com:30301",
	"enode://8d7e62fee7e364f308dedfd0a3a71c68ffc41e3abdf7be2eb7e0ae05291e0f709db33c292fa39a0b7648fa4dc33573b7141df26b2d828482d6a5e6baca353c81@bootnode-b.genesis-mainnet.inichain.com:30301",
	"enode://ded15ac04a75a8c08be778437f2f8ac3a4f5bf636b05b2ee76b1a9e0ccc7b2034606c104e0b5e3b59f345fac9c27b70a91a0be3368e96b8f8ad0d0923b27bafd@bootnode-c.genesis-mainnet.inichain.com:30301",
	"enode://804b480e6716289b3fcdeea2922489b25ded5b49d6ad34707bb708c585bdfde4ea822e4c4d6abdcd06c51c9ef16506e2da3ec5c95c4c569fb4e80312fad1115a@bootnode-a.genesis-mainnet.inichain.com:30304",
	"enode://863e6776fe0553af089e3783d2ca57df48215d64567f81ffb874095e6fe190dc14b3416ef79eac4d543846b37413f679d7b9e0fc30c69dea4666af9d0dc38200@bootnode-b.genesis-mainnet.inichain.com:30304",
	"enode://73a2a84a866f1db24a9bfffd511f59751985090b861bb82fe26927fba4d89b3c66237fcc72ce2d3528025f761da2c223e4ec76eec6cb0bc0109067dcd8a63acc@bootnode-c.genesis-mainnet.inichain.com:30304",
}
var TestnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://5b609e290c4c2d408c36eff0f2e5cbb7bd43339c5dc27259b122449c901420fda33c83670c0a51d1f40adbe38a18ea66c097a7a0f30b619d44c050e6e7c34feb@bootnode-a.genesis-testnet.inichain.com:30301",
	"enode://b04e54ff785ddfbca7309f66283f035dcf5388959e7a92235e2b4a0d1f721da10c773c94eef60e09c69b8caab2cce402a6539eab3c2af44bb1bbb4908ad7b178@bootnode-b.genesis-testnet.inichain.com:30301",
	"enode://be431192987395f9f5455733fe806f71c4cf7ea8bc6612451433b409dd600b654921019d505887f9d9a5ac60d6612d37c7189c5117dc46a1c23f44ca634f7c3a@bootnode-c.genesis-testnet.inichain.com:30301",
	"enode://8d95d155aef0d05011090ac96c2456382a3f27ad73f0c2ee10fe13235053e769560d16055e28b9909af6e8c572a1c57647695c03e8f31a5a383a31f4775d080b@bootnode-a.genesis-testnet.inichain.com:30304",
	"enode://ff91e1682b09ac717fe73421a7971adb0d62f7f7ed978ce075f3b0014b4a760180f1eb4e6b6ec2e89c4708c53c93430d457e2bdae40231145a5718fcee07db43@bootnode-b.genesis-testnet.inichain.com:30304",
	"enode://f83f309d96c289bb32f191700de79fa49454a22afbfa65306eb99c690a8b4fa4f839884d346b9becd6d7883032ba54b00aaa89ac6be8c91a69f2725814886ba1@bootnode-b.genesis-testnet.inichain.com:30304",
}
var DevnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://5b609e290c4c2d408c36eff0f2e5cbb7bd43339c5dc27259b122449c901420fda33c83670c0a51d1f40adbe38a18ea66c097a7a0f30b619d44c050e6e7c34feb@devnet-bootnode0.inichain.com:30301",
	"enode://b04e54ff785ddfbca7309f66283f035dcf5388959e7a92235e2b4a0d1f721da10c773c94eef60e09c69b8caab2cce402a6539eab3c2af44bb1bbb4908ad7b178@devnet-bootnode1.inichain.com:30301",
	"enode://be431192987395f9f5455733fe806f71c4cf7ea8bc6612451433b409dd600b654921019d505887f9d9a5ac60d6612d37c7189c5117dc46a1c23f44ca634f7c3a@devnet-bootnode2.inichain.com:30301",
}

var V5Bootnodes = []string{
	// Teku team's bootnode
	"enr:-KG4QMOEswP62yzDjSwWS4YEjtTZ5PO6r65CPqYBkgTTkrpaedQ8uEUo1uMALtJIvb2w_WWEVmg5yt1UAuK1ftxUU7QDhGV0aDKQu6TalgMAAAD__________4JpZIJ2NIJpcIQEnfA2iXNlY3AyNTZrMaEDfol8oLr6XJ7FsdAYE7lpJhKMls4G_v6qQOGKJUWGb_uDdGNwgiMog3VkcIIjKA", // # 4.157.240.54 | azure-us-east-virginia
	"enr:-KG4QF4B5WrlFcRhUU6dZETwY5ZzAXnA0vGC__L1Kdw602nDZwXSTs5RFXFIFUnbQJmhNGVU6OIX7KVrCSTODsz1tK4DhGV0aDKQu6TalgMAAAD__________4JpZIJ2NIJpcIQExNYEiXNlY3AyNTZrMaECQmM9vp7KhaXhI-nqL_R0ovULLCFSFTa9CPPSdb1zPX6DdGNwgiMog3VkcIIjKA", // 4.196.214.4  | azure-au-east-sydney
	// Prylab team's bootnodes
	"enr:-Ku4QImhMc1z8yCiNJ1TyUxdcfNucje3BGwEHzodEZUan8PherEo4sF7pPHPSIB1NNuSg5fZy7qFsjmUKs2ea1Whi0EBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQOVphkDqal4QzPMksc5wnpuC3gvSC8AfbFOnZY_On34wIN1ZHCCIyg", // 18.223.219.100 | aws-us-east-2-ohio
	"enr:-Ku4QP2xDnEtUXIjzJ_DhlCRN9SN99RYQPJL92TMlSv7U5C1YnYLjwOQHgZIUXw6c-BvRg2Yc2QsZxxoS_pPRVe0yK8Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMeFF5GrS7UZpAH2Ly84aLK-TyvH-dRo0JM1i8yygH50YN1ZHCCJxA", // 18.223.219.100 | aws-us-east-2-ohio
	"enr:-Ku4QPp9z1W4tAO8Ber_NQierYaOStqhDqQdOPY3bB3jDgkjcbk6YrEnVYIiCBbTxuar3CzS528d2iE7TdJsrL-dEKoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMw5fqqkw2hHC4F5HZZDPsNmPdB1Gi8JPQK7pRc9XHh-oN1ZHCCKvg", // 18.223.219.100 | aws-us-east-2-ohio
	// Lighthouse team's bootnodes
	"enr:-Le4QPUXJS2BTORXxyx2Ia-9ae4YqA_JWX3ssj4E_J-3z1A-HmFGrU8BpvpqhNabayXeOZ2Nq_sbeDgtzMJpLLnXFgAChGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISsaa0Zg2lwNpAkAIkHAAAAAPA8kv_-awoTiXNlY3AyNTZrMaEDHAD2JKYevx89W0CcFJFiskdcEzkH_Wdv9iW42qLK79ODdWRwgiMohHVkcDaCI4I", // 172.105.173.25 | linode-au-sydney
	"enr:-Le4QLHZDSvkLfqgEo8IWGG96h6mxwe_PsggC20CL3neLBjfXLGAQFOPSltZ7oP6ol54OvaNqO02Rnvb8YmDR274uq8ChGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISLosQxg2lwNpAqAX4AAAAAAPA8kv_-ax65iXNlY3AyNTZrMaEDBJj7_dLFACaxBfaI8KZTh_SSJUjhyAyfshimvSqo22WDdWRwgiMohHVkcDaCI4I", // 139.162.196.49 | linode-uk-london
	"enr:-Le4QH6LQrusDbAHPjU_HcKOuMeXfdEB5NJyXgHWFadfHgiySqeDyusQMvfphdYWOzuSZO9Uq2AMRJR5O4ip7OvVma8BhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISLY9ncg2lwNpAkAh8AgQIBAAAAAAAAAAmXiXNlY3AyNTZrMaECDYCZTZEksF-kmgPholqgVt8IXr-8L7Nu7YrZ7HUpgxmDdWRwgiMohHVkcDaCI4I", // 139.99.217.220 | ovh-au-sydney
	"enr:-Le4QIqLuWybHNONr933Lk0dcMmAB5WgvGKRyDihy1wHDIVlNuuztX62W51voT4I8qD34GcTEOTmag1bcdZ_8aaT4NUBhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISLY04ng2lwNpAkAh8AgAIBAAAAAAAAAA-fiXNlY3AyNTZrMaEDscnRV6n1m-D9ID5UsURk0jsoKNXt1TIrj8uKOGW6iluDdWRwgiMohHVkcDaCI4I", // 139.99.78.39 | ovh-singapore
	// EF bootnodes
	"enr:-Ku4QHqVeJ8PPICcWk1vSn_XcSkjOkNiTg6Fmii5j6vUQgvzMc9L1goFnLKgXqBJspJjIsB91LTOleFmyWWrFVATGngBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhAMRHkWJc2VjcDI1NmsxoQKLVXFOhp2uX6jeT0DvvDpPcU8FWMjQdR4wMuORMhpX24N1ZHCCIyg", // 3.17.30.69 | aws-us-east-2-ohio
	"enr:-Ku4QG-2_Md3sZIAUebGYT6g0SMskIml77l6yR-M_JXc-UdNHCmHQeOiMLbylPejyJsdAPsTHJyjJB2sYGDLe0dn8uYBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhBLY-NyJc2VjcDI1NmsxoQORcM6e19T1T9gi7jxEZjk_sjVLGFscUNqAY9obgZaxbIN1ZHCCIyg", // 18.216.248.220 | aws-us-east-2-ohio
	"enr:-Ku4QPn5eVhcoF1opaFEvg1b6JNFD2rqVkHQ8HApOKK61OIcIXD127bKWgAtbwI7pnxx6cDyk_nI88TrZKQaGMZj0q0Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDayLMaJc2VjcDI1NmsxoQK2sBOLGcUb4AwuYzFuAVCaNHA-dy24UuEKkeFNgCVCsIN1ZHCCIyg", // 54.178.44.198 | aws-ap-northeast-1-tokyo
	"enr:-Ku4QEWzdnVtXc2Q0ZVigfCGggOVB2Vc1ZCPEc6j21NIFLODSJbvNaef1g4PxhPwl_3kax86YPheFUSLXPRs98vvYsoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDZBrP2Jc2VjcDI1NmsxoQM6jr8Rb1ktLEsVcKAPa08wCsKUmvoQ8khiOl_SLozf9IN1ZHCCIyg", // 54.65.172.253 | aws-ap-northeast-1-tokyo
	// Nimbus team's bootnodes
	"enr:-LK4QA8FfhaAjlb_BXsXxSfiysR7R52Nhi9JBt4F8SPssu8hdE1BXQQEtVDC3qStCW60LSO7hEsVHv5zm8_6Vnjhcn0Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhAN4aBKJc2VjcDI1NmsxoQJerDhsJ-KxZ8sHySMOCmTO6sHM3iCFQ6VMvLTe948MyYN0Y3CCI4yDdWRwgiOM", // 3.120.104.18 | aws-eu-central-1-frankfurt
	"enr:-LK4QKWrXTpV9T78hNG6s8AM6IO4XH9kFT91uZtFg1GcsJ6dKovDOr1jtAAFPnS2lvNltkOGA9k29BUN7lFh_sjuc9QBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhANAdd-Jc2VjcDI1NmsxoQLQa6ai7y9PMN5hpLe5HmiJSlYzMuzP7ZhwRiwHvqNXdoN0Y3CCI4yDdWRwgiOM", // 3.64.117.223 | aws-eu-central-1-frankfurt}
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case TestnetGenesisHash:
		net = "testnet"
	case DevnetGenesisHash:
		net = "devnet"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
