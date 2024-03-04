package simplemath

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

//value based
func (sv SemanticVersion) Incrementmajor() {
	sv.major += 1
}

//pointer based
//pointer receivers
func (sv *SemanticVersion) IncrementMajor() {
	sv.major += 1
}

//pointer based
func (sv *SemanticVersion) IncrementMinor() {
	sv.major += 1
}

func (sv *SemanticVersion) IncrementPatch() {
	sv.major += 1
}
