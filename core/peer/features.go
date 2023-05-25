package peer

import "unsafe"

type Feature string
type Features []Feature

var emptyFeatureList Features = nil

func (list Features) HasFeature(feature Feature) bool {
	for _, ft := range list {
		if ft == feature {
			return true
		}
	}
	return false
}

func (list Features) Size() int {
	if &list == &emptyFeatureList {
		return 0
	}
	return len(list)
}

func (list Features) ToStrArray() []string {
	return *(*[]string)(unsafe.Pointer(&list)) // https://stackoverflow.com/questions/29031353/conversion-of-a-slice-of-string-into-a-slice-of-custom-type
}

func ToFeatures(fts []string) Features {
	return *(*Features)(unsafe.Pointer(&fts)) // https://stackoverflow.com/questions/29031353/conversion-of-a-slice-of-string-into-a-slice-of-custom-type
}

// features set for future plans :)
const FT_POINT = 1

type FeatureSet struct {
	fts map[Feature]struct{}
}

func NewFeatureSet(fts ...Feature) FeatureSet {
	fs := FeatureSet{
		fts: make(map[Feature]struct{}, len(fts)),
	}
	for _, ft := range fts {
		fs.fts[ft] = struct{}{}
	}
	return fs
}

func (fs FeatureSet) HasFeature(ft Feature) bool {
	_, ok := fs.fts[ft]
	return ok
}

func (fs FeatureSet) FeatureScore(fts Features) int {
	if &fts == &emptyFeatureList {
		return 0
	}

	score := 0
	for _, ft := range fts {
		if fs.HasFeature(ft) {
			score += FT_POINT
		}
	}
	return score
}

func (fs FeatureSet) Features() Features {
	res := make(Features, len(fs.fts))
	for ft := range fs.fts {
		res = append(res, ft)
	}
	return res
}

func (fs FeatureSet) Size() int {
	return len(fs.fts)
}
