package peer

import (
	"unsafe"
)

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