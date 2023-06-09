package peer

type FeatureSet interface {
	Features() Features
	SetFeatures(... Feature)
	HasFeatures(...Feature) bool
	FeatureScore(ft Features) int
	MaxFeatureScore() int
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
	if len(fts) == fs.Size() && fs.HasFeatures(fts...){
		return // nothing changed so don't do nothing :)
	}

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

// ultimate need :)
func (fs BasicFeatureSet) MaxFeatureScore() int {
	return len(fs.fts) * FT_POINT
}

// it's here only for testint reasons
func SameFeatures(f1 Features, f2 Features) bool {
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