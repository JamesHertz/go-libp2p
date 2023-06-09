package pstoremem

import (
	"sync"

	"github.com/libp2p/go-libp2p/core/peer"
	ps "github.com/libp2p/go-libp2p/core/peerstore"
)

var _ ps.FeatureBook = (*memoryFeatureBook)(nil)


type memoryFeatureBook struct {
	store map[peer.ID]peer.FeatureSet
	lock  sync.RWMutex
}

func NewFeatureBook() *memoryFeatureBook {
	return &memoryFeatureBook{
		store: make(map[peer.ID]peer.FeatureSet),
		lock:  sync.RWMutex{},
	}
}

func (fb *memoryFeatureBook) SetFeatures(pid peer.ID, features ...peer.Feature) {
	fb.lock.Lock()
	defer fb.lock.Unlock()
	fs, ok := fb.store[pid]
	if !ok {
		fb.store[pid] = peer.NewFeatureSet(features...)
	} else {
		fs.SetFeatures(features...)
	}
}

func (fb *memoryFeatureBook) Features(pid peer.ID) peer.Features {
	fb.lock.RLock()
	defer fb.lock.RUnlock()

	fs, ok := fb.store[pid]

	if !ok {
		return nil
	}

	return fs.Features()
}

func (fb *memoryFeatureBook) HasFeatures(pid peer.ID, feature ...peer.Feature) bool {
	fb.lock.RLock()
	defer fb.lock.RUnlock()
	fs, ok := fb.store[pid]
	return ok && fs.HasFeatures(feature...)
}

func (fb *memoryFeatureBook) RemovePeer(pid peer.ID) {
	fb.lock.Lock()
	delete(fb.store, pid)
	fb.lock.Unlock()
}