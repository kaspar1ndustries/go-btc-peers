package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/btcsuite/btcd/wire"
)

type Network string

const (
	NetworkMainnet Network = "mainnet"
	NetworkTestnet Network = "testnet"
)

type Config struct {
	Network          Network
	NodesDB          string
	GoodNodesDB      string
	NodesPort        uint16
	NodeTimeout      time.Duration
	PingInterval     time.Duration
	PingTimeout      time.Duration
	ListenInterval   time.Duration
	ConnectionsLimit int
	LogsDir          string
	LogsFilename     string

	DnsAddress string
	DnsTimeout time.Duration
	DnsSeeds   []string

	// Wire
	Pver uint32

	// var btcnet = wire.MainNet
	Btcnet wire.BitcoinNet
}

func New() *Config {
	cfg := &Config{
		// var dnsAddress = "1.1.1.1:53" // cloudflare dns, 2x slower
		// google dns
		// DnsAddress:     "8.8.8.8:53",
		// cloudflare dns
		DnsAddress: "1.1.1.1:53",
		// quad dns
		// DnsAddress:     "9.9.9.9:53",

		Pver:           wire.ProtocolVersion, // 70016
		NodeTimeout:    5 * time.Second,
		PingInterval:   1 * time.Minute,
		PingTimeout:    15 * time.Second,
		ListenInterval: 1 * time.Second,
		LogsDir:        "logs",
		LogsFilename:   fmt.Sprintf("logs_%s.log", time.Now().Format("2006-01-02_15-04-05")),
		// Pver: 70013,
	}
	if os.Getenv("DEBUG") == "1" {
		cfg.ConnectionsLimit = 30
	} else {
		cfg.ConnectionsLimit = 50
	}
	// override connections limit
	if os.Getenv("CONN") != "" {
		conn, err := strconv.Atoi(os.Getenv("CONN"))
		if err != nil {
			log.Fatalf("error converting CONN env variable to int: %v", err)
		}
		cfg.ConnectionsLimit = conn
	}
	if os.Getenv("TESTNET") == "1" {
		cfg.Network = NetworkTestnet
		cfg.Btcnet = wire.TestNet3
		cfg.DnsTimeout = 10 * time.Second
		cfg.NodesDB = "data/nodes_testnet.json"
		cfg.GoodNodesDB = "data/nodes_good_testnet.json"
		cfg.NodesPort = 18333
		cfg.DnsSeeds = []string{
			"testnet-seed.bitcoin.jonasschnelli.ch",
			"seed.tbtc.petertodd.org",
			"seed.testnet.bitcoin.sprovoost.nl",
			"testnet-seed.bluematt.me",
		}
	} else {
		cfg.Network = NetworkMainnet
		cfg.Btcnet = wire.MainNet

		cfg.DnsTimeout = 5 * time.Second
		cfg.NodesDB = "data/nodes_mainnet.json"
		cfg.GoodNodesDB = "data/nodes_good_mainnet.json"
		cfg.NodesPort = 8333
		cfg.DnsSeeds = []string{
			"dnsseed.emzy.de",
			"dnsseed.bluematt.me",
			"dnsseed.bitcoin.dashjr.org",
			"seed.bitcoin.sipa.be",
			"seed.bitcoinstats.com",
			"seed.bitcoin.jonasschnelli.ch",
			"seed.btc.petertodd.org",
			"seed.bitcoin.sprovoost.nl",
			"seed.bitcoin.wiz.biz",
			"seed.bitnodes.io",
		}
	}
	return cfg
}
