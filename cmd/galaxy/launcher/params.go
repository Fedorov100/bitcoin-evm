package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// mainnet
		"enode://2ff956608d772872155b89d51b9bf8f3262573132eb1c67c7599754f5405f720a000fd63e948da0c4d215e9f8fa2452c0767abdfe1946d532adbeff9767d5d2a@85.206.160.134:18887",
		"enode://1aed3eb16a1705a48ddd423f3f09a45f65fb626f9248ea0e1abf06ea2496026820a9bba5e2f1a57954e3a2d677d2186ffbea94f4c7baa931419617107e4d5881@85.206.160.135:18887",
		"enode://1bb9718da75b77ac8e3b1b8417435b9842eadce99e1b0d8069c79b4bd6c12149205aa1451eedbb04d3fa106a5a92154bd2a280fb271a29fda7a7cca37f8cddd2@185.25.48.207:18887",
		// testnet
		"enode://2ff956608d772872155b89d51b9bf8f3262573132eb1c67c7599754f5405f720a000fd63e948da0c4d215e9f8fa2452c0767abdfe1946d532adbeff9767d5d2a@31.220.57.80:18887",
		"enode://1aed3eb16a1705a48ddd423f3f09a45f65fb626f9248ea0e1abf06ea2496026820a9bba5e2f1a57954e3a2d677d2186ffbea94f4c7baa931419617107e4d5881@31.220.55.122:18887",
		"enode://1bb9718da75b77ac8e3b1b8417435b9842eadce99e1b0d8069c79b4bd6c12149205aa1451eedbb04d3fa106a5a92154bd2a280fb271a29fda7a7cca37f8cddd2@31.220.52.254:18887",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
