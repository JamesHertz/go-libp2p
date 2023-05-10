package peer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScores(t *testing.T){
	testCases := []struct{
		f1 FeatureList 
		f2 FeatureList 
		score int
	}{
		{
			f1: FeatureList{"a", "b", "c"},
			f2: FeatureList{"a", "b", "c"},
			score: 0b111,
		},
		{
			f1: FeatureList{"a", "b", "c"},
			f2: FeatureList{"a", "d", "c"},
			score: 0b101,
		},
		{
			f1: FeatureList{"a", "b", "c"},
			f2: FeatureList{"c", "a", "b"},
			score: 0b000,
		},
		{
			f1: FeatureList{"a", "b", "c"},
			f2: FeatureList{"c"},
			score: 0b001,
		},
		{
			f1: FeatureList{"a", "b", "c"},
			f2: FeatureList{"b", "c"},
			score: 0b011,
		},
		{
			f1: FeatureList{"a", "b", "c"},
			f2: FeatureList{"c", "b"},
			score: 0b000,
		},
	};

	for _, tcase := range testCases {
		features1 := tcase.f1
		features2 := tcase.f2
		require.Equal(t, features2.FeaturesScore(features1), features1.FeaturesScore(features2))
		require.Equal(t, tcase.score, features1.FeaturesScore(features2))
	}
}