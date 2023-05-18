package peer 

import "unsafe"

type Feature string
type FeatureList []Feature

var emptyFeatureList FeatureList = nil

func StringToFeatureList(fts []string) FeatureList {
	return *(*FeatureList)(unsafe.Pointer(&fts)) // https://stackoverflow.com/questions/29031353/conversion-of-a-slice-of-string-into-a-slice-of-custom-type
}

func (list FeatureList) FeaturesScore(other FeatureList) int {
	if &list == &emptyFeatureList || &other == &emptyFeatureList {
		return 0
	}

	smallest := list
	greater  := other
	diff     := other.Size() - list.Size()
	if diff < 0 {
		smallest = other
		greater  = list
		diff     = -diff
	}

	score := 0
	for i := 0; i < smallest.Size(); i++ {
		score <<= 1
		if smallest[i] == greater[i+diff] {
			score |= 1
		}
	}
	return score
}


func (list FeatureList) HasFeature(feature Feature) bool{
	for _, ft := range list {
		if ft == feature{
			return true
		}
	}
	return false
}


func (list FeatureList) Size() int {
	if &list == &emptyFeatureList {
		return 0
	}
	return len(list)
}

func (list FeatureList) StringArray() []string{
	return *(*[]string)(unsafe.Pointer(&list)) // https://stackoverflow.com/questions/29031353/conversion-of-a-slice-of-string-into-a-slice-of-custom-type
}
