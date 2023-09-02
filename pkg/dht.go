package pkg

import (
	"context"

	"github.com/anacrolix/dht/v2"
	"github.com/anacrolix/dht/v2/exts/getput"
	"github.com/anacrolix/torrent/bencode"
	"github.com/anacrolix/torrent/types/infohash"
	"github.com/sirupsen/logrus"

	"pkarr-go/internal"
)

type DHT struct {
	*dht.Server
}

type BootstrapPeer struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func NewDHT() (*DHT, error) {
	c := dht.NewDefaultServerConfig()
	c.StartingNodes = func() ([]dht.Addr, error) { return dht.ResolveHostPorts(getDefaultBootstrapPeers()) }
	s, err := dht.NewServer(c)
	if err != nil {
		logrus.WithError(err).Error("failed to create dht server")
		return nil, err
	}
	return &DHT{Server: s}, nil
}

func (d *DHT) Get(key string) (string, error) {
	z32Decoded, err := internal.Z32Decode(key)
	if err != nil {
		logrus.WithError(err).Error("failed to decode key")
		return "", err
	}
	res, t, err := getput.Get(context.Background(), infohash.HashBytes(z32Decoded), d.Server, nil, nil)
	if err != nil {
		logrus.WithError(err).Errorf("failed to get key<%s> from dht; tried %d nodes, got %d responses", key, t.NumAddrsTried, t.NumResponses)
		return "", err
	}
	var payload string
	if err = bencode.Unmarshal(res.V, &payload); err != nil {
		logrus.WithError(err).Error("failed to unmarshal payload value")
		return "", err
	}
	decoded, err := internal.Decode([]byte(payload))
	if err != nil {
		logrus.WithError(err).Error("failed to decode value from dht")
		return "", err
	}
	return string(decoded), nil
}

func (d *DHT) Put(key string) (string, error) {
	return "", nil
}

func getDefaultBootstrapPeers() []string {
	return []string{
		// "router.magnets.im:6881",
		// "router.bittorrent.com:6881",
		// "dht.transmissionbt.com:6881",
		// "router.utorrent.com:6881",
		"router.nuh.dev:6881",
	}
}
