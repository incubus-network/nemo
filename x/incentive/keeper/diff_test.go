package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetDiff(t *testing.T) {
	tests := []struct {
		name     string
		setA     []string
		setB     []string
		expected []string
	}{
		{"empty", []string{}, []string{}, []string(nil)},
		{"diff equal sets", []string{"busd", "musd"}, []string{"busd", "musd"}, []string(nil)},
		{"diff set empty", []string{"bnb", "ufury", "musd"}, []string{}, []string{"bnb", "ufury", "musd"}},
		{"input set empty", []string{}, []string{"bnb", "ufury", "musd"}, []string(nil)},
		{"diff set with common elements", []string{"bnb", "btcb", "musd", "xrpb"}, []string{"bnb", "musd"}, []string{"btcb", "xrpb"}},
		{"diff set with all common elements", []string{"bnb", "musd"}, []string{"bnb", "btcb", "musd", "xrpb"}, []string(nil)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, setDifference(tt.setA, tt.setB))
		})
	}
}
