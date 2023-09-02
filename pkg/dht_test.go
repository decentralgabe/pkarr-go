package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDHT(t *testing.T) {
	d, err := NewDHT()
	require.NoError(t, err)

	got, err := d.Get("yj47pezutnpw9pyudeeai8cx8z8d6wg35genrkoqf9k3rmfzy58o")
	require.NoError(t, err)
	require.NotEmpty(t, got)
	println(got)
}
