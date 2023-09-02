package pkg

import (
	"context"
	"fmt"

	"github.com/anacrolix/dht/v2"
	"github.com/anacrolix/dht/v2/exts/getput"
	"github.com/anacrolix/dht/v2/krpc"
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
	hashed := internal.Hash([]byte(key))
	hexed := internal.Hex(hashed)
	res, t, err := getput.Get(context.Background(), infohash.FromHexString(hexed), d.Server, nil, nil)
	if err != nil {
		logrus.WithError(err).Errorf("failed to get key<%s> from dht; tried %d nodes, got %d responses", key, t.NumAddrsTried, t.NumResponses)
		return "", err
	}
	var payload krpc.Bep46Payload
	if err = bencode.Unmarshal(res.V, &payload); err != nil {
		return "", fmt.Errorf("unmarshalling bep46 payload: %w", err)
	}
	s := payload.Ih.Bytes()
	decoded, err := internal.Decode(s)
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
