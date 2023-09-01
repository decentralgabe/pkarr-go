package pkarr_go

import (
	"strconv"
	"strings"

	"github.com/nictuku/dht"
	"github.com/sirupsen/logrus"
)

type DHT struct {
	d *dht.DHT
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

	return &DHT{d: d}, nil
}

func (d DHT) Start() error {
	return d.d.Start()
}

func (d DHT) Stop() {
	d.d.Stop()
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
