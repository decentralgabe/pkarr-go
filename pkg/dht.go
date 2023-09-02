package pkg

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nictuku/dht"
	"github.com/sirupsen/logrus"

	"pkarr-go/internal"
)

type DHT struct {
	*dht.DHT
}

type BootstrapPeer struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func NewDHT() (*DHT, error) {
	c := dht.NewConfig()
	c.DHTRouters = getDefaultBootstrapPeersString()
	d, err := dht.New(c)
	if err != nil {
		logrus.WithError(err).Error("failed to create dht")
		return nil, err
	}
	if err = d.Start(); err != nil {
		logrus.WithError(err).Error("failed to start dht")
		return nil, err
	}

	return &DHT{DHT: d}, nil
}

func (d DHT) Start() error {
	return d.DHT.Start()
}

func (d DHT) Stop() {
	d.DHT.Stop()
}

func (d DHT) Get(key string) (string, error) {
	key = strings.Replace(key, "pk", "", 1)
	target := internal.Hash([]byte(key))
	targetHex := internal.Hex(target)

	infoHash, err := dht.DecodeInfoHash(targetHex)
	if err != nil {
		fmt.Printf("DecodeInfoHash faiure: %v", err)
		return "", err
	}

	infoHashPeers := d.QueryNodes(string(infoHash))
	for ih, peers := range infoHashPeers {
		if len(peers) > 0 {
			for _, peer := range peers {
				fmt.Println(dht.DecodePeerAddress(peer))
			}

			if fmt.Sprintf("%x", ih) == targetHex {
				return ih.String(), nil
			}
		}
	}
	return "", nil
}

func (d DHT) QueryNodes(infoHash string) map[dht.InfoHash][]string {
	tick := time.Tick(time.Second)
	timer := time.NewTimer(8 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-tick:
			fmt.Println("tick")
			d.PeersRequest(infoHash, true)
		case infoHashPeers := <-d.PeersRequestResults:
			return infoHashPeers
		case <-timer.C:
			fmt.Printf("Could not find new peers: timed out")
			return nil
		}
	}
}

func getDefaultBootstrapPeersString() string {
	defaultPeers := []BootstrapPeer{
		{Host: "router.magnets.im", Port: 6881},
		{Host: "router.bittorrent.com", Port: 6881},
		{Host: "dht.transmissionbt.com", Port: 6881},
		{Host: "router.utorrent.com", Port: 6881},
		{Host: "router.nuh.dev", Port: 6881},
	}
	b := new(strings.Builder)
	for _, peer := range defaultPeers {
		b.WriteString(strings.Join([]string{peer.Host, ":", strconv.Itoa(peer.Port)}, ""))
	}
	return b.String()
}
