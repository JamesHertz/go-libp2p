package peer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScores(t *testing.T) {
	aux := NewFeatureSet();
	fmt.Println("aux.Features():", aux.Features())

	fs := NewFeatureSet("a", "b", "c", "d")

	require.Equal(t, fs.FeatureScore(Features(nil)), 0)

	testCases := []struct {
		fts   Features
		score int
	}{
		{
			fts:   Features{"a", "b"},
			score: 2 * FT_POINT,
		},
		{
			fts:   Features{"a", "b", "d"},
			score: 3 * FT_POINT,
		},
		{
			fts:   Features{"z", "h", "i"},
			score: 0,
		},
		{
			fts:   fs.Features(),
			score: FT_POINT * fs.Size(),
		},
	}

	for _, test := range testCases {
		require.Equal(t, test.score, fs.FeatureScore(test.fts))
	}
}
