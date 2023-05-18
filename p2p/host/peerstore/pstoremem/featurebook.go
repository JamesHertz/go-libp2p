package pstoremem

import (
	"sync"

	"github.com/libp2p/go-libp2p/core/peer"
)

type memoryFeatureBook struct {
	store map[peer.ID]peer.FeatureList
	lock sync.RWMutex
}

func NewFeatureBook() *memoryFeatureBook{
	return &memoryFeatureBook{
		store: make(map[peer.ID]peer.FeatureList),
		lock: sync.RWMutex{},
	} 
}

func (fb * memoryFeatureBook) SetFeatures(pid peer.ID, features ...peer.Feature) {
	fb.lock.Lock()
	aux := make(peer.FeatureList, len(features))
	copy(aux, features)
	fb.store[pid] = aux
	fb.lock.Unlock()
}

func (fb * memoryFeatureBook) GetFeatures(pid peer.ID) peer.FeatureList {
	fb.lock.RLock()
	defer fb.lock.RUnlock()
	features, ok := fb.store[pid]

	if !ok {
		return nil
	}

	res := make(peer.FeatureList, features.Size())
	copy(res, features)
	return res
}

func (fb * memoryFeatureBook) HasFeature(pid peer.ID, feature peer.Feature) bool {
	fb.lock.RLock()
	defer fb.lock.RUnlock()
	features, ok := fb.store[pid]
	if !ok {
		return false
	}

	for _, ft := range features {
		if ft == feature {
			return true
		}
	}

	return false
}

func (fb * memoryFeatureBook) RemovePeer(pid peer.ID) {
	fb.lock.Lock()
	delete(fb.store, pid)
	fb.lock.Unlock()
}

// func (fb * memoryFeatureBook) RemoveFeature(pid peer.ID, features ...peer.Feature){
// 	fb.lock.Lock()
// 	defer fb.lock.Unlock()
// 	pfeatures, ok := fb.store[pid]
// 	if ok {
// 		//newfl := make(peer.FeatureList, 0, pfeatures.Size())
// 	}
// }

