package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"

	"pkarr-go/internal"
)

func TestGetDHT(t *testing.T) {
	d, err := NewDHT()
	require.NoError(t, err)

	got, err := d.Get("yj47pezutnpw9pyudeeai8cx8z8d6wg35genrkoqf9k3rmfzy58o")
	require.NoError(t, err)
	require.NotEmpty(t, got)
	println(got)
}

func TestPutDHT(t *testing.T) {
	d, err := NewDHT()
	require.NoError(t, err)

	records := [][]string{
		{"foo", "bar"},
	}
	pubKey, privKey, err := internal.GenerateKeypair()
	require.NoError(t, err)

	putReq, err := CreatePutRequest(pubKey, privKey, records)
	require.NoError(t, err)

	id, err := d.Put(pubKey, *putReq)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	println(id)

	got, err := d.Get(id)
	require.NoError(t, err)
	require.NotEmpty(t, got)
	println(got)
}
