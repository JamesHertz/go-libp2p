package peer

type FeatureScorer interface {
	FeatureScore(ft Features) int
	HasFeatures(...Feature) bool
}

type FeatureSet interface {
	FeatureScorer
	Features() Features
	SetFeatures(... Feature)
	Size() int
}


// features set for future plans :)
const FT_POINT = 1

type BasicFeatureSet struct {
	fts map[Feature]struct{}
}

func NewFeatureSet(fts ...Feature) FeatureSet {
	fs := BasicFeatureSet{
		fts: make(map[Feature]struct{}, len(fts)),
	}
	fs.SetFeatures(fts...)
	return fs
}

func (fs BasicFeatureSet) HasFeatures(fts ...Feature) bool {
	for _, ft := range fts {
		_, ok := fs.fts[ft]
		if !ok {
			return false;
		}

	}
	return true
}

func (fs BasicFeatureSet) FeatureScore(fts Features) int {
	if &fts == &emptyFeatureList {
		return 0
	}

	score := 0
	for _, ft := range fts {
		if fs.HasFeatures(ft) {
			score += FT_POINT
		}
	}
	return score
}


func (fs BasicFeatureSet) SetFeatures(fts ...Feature){
	for f := range fs.fts {
		delete(fs.fts, f)
	}
	for _, ft := range fts {
		fs.fts[ft] = struct{}{}
	}
}

func (fs BasicFeatureSet) Features() Features {
	if fs.Size() == 0 {
		return emptyFeatureList
	}

	res := make(Features, 0, len(fs.fts))
	for ft := range fs.fts {
		res = append(res, ft)
	}
	return res
}

func (fs BasicFeatureSet) Size() int {
	return len(fs.fts)
}