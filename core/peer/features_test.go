package peer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScores(t *testing.T) {
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

func TestFeaturesSet(t *testing.T) {

	fts := Features{"a", "b", "c", "d"}
	fs := NewFeatureSet(fts...)

	require.True(t, equals(fs.Features(), fts))

	for _, ft := range fts {
		require.True(t, fs.HasFeatures(ft))
	}

	ftsII := Features{"d", "e", "f"}

	fs.SetFeatures(ftsII...)
	require.True(t, equals(fs.Features(), ftsII))

	require.True(t, fs.HasFeatures(ftsII...))
	require.True(t, fs.Features().HasFeatures(ftsII...))
}

// TODO: deduplicate this (also used in peerstore_suit.go)
func equals(f1 Features, f2 Features) bool {
	if f1.Size() != f2.Size() {
		return false
	}
	helper := make(map[Feature]struct{}, f1.Size())
	for _, f := range f1 {
		helper[f] = struct{}{}
	}
	for _, f := range f2 {
		_, ok := helper[f]
		if !ok {
			return false
		}
	}
	return true
}
