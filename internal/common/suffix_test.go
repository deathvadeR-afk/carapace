package common

import "testing"

func TestSuffixMatcherAdd(t *testing.T) {
	sm := SuffixMatcher{""}

	sm.Add('*')
	if sm.string != "*" {
		t.Errorf(`should be "*" [was: "%v"]`, sm)
	}
}

func TestSuffixMatcherMatchesEndingWithSuffix(t *testing.T) {
	sm := SuffixMatcher{""}
	sm.Add('/')

	if !sm.Matches("src/") {
		t.Error("candidate ending in the suffix character should match")
	}
}

func TestSuffixMatcherDoesNotMatchStartingWithSuffix(t *testing.T) {
	sm := SuffixMatcher{""}
	sm.Add('/')

	if sm.Matches("/src") {
		t.Error("candidate only starting with the suffix character should not match")
	}
}
