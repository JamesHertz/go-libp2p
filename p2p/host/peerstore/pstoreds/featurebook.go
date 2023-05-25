package pstoreds

import (
	"github.com/libp2p/go-libp2p/core/peer"
)

var featureKey = "features"

type dsFeatureBook struct{}

func NewFeatureBook() *dsFeatureBook {
	return nil
}

func (fb *dsFeatureBook) Features(pid peer.ID) peer.Features {
	// features, err := fb.meta.Get(pid, featureKey)
	panic("called GetFeatures of dsFeatureBook (not implemented)")
}

func (fb *dsFeatureBook) SetFeatures(pid peer.ID, features ...peer.Feature) {
	// features, err := fb.meta.Get(pid, featureKey)
	panic("called SetFeatures of dsFeatureBook (not implemented)")
}

func (fb *dsFeatureBook) HasFeature(pid peer.ID, feature peer.Feature) bool {
	panic("called HasFeature of dsFeatureBook (not implemented)")
}
