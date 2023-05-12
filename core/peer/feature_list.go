package peer

type Feature string
type FeatureList []Feature

var emptyFeatureList FeatureList = nil

func (list FeatureList) FeaturesScore(other FeatureList) int {
	if &list == &emptyFeatureList || &other == &emptyFeatureList {
		return 0
	}

	var smallest, greater FeatureList
	diff := other.Size() - list.Size()
	if diff >= 0 {
		smallest = list
		greater = other
	} else {
		smallest = other
		greater = list
		diff = -diff
	}

	score := 0
	for i := 0; i < smallest.Size(); i++ {
		bit := 0
		if smallest[i] == greater[i+diff] {
			bit = 1
		}
		score = score<<1 | bit
	}
	return score
}

func (list FeatureList) Size() int {
	if &list == &emptyFeatureList {
		return 0
	}
	return len(list)
}
