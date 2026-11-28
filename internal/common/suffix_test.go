package common

import "testing"

func TestSuffixMatcherAdd(t *testing.T) {
	sm := SuffixMatcher{""}

	sm.Add('*')
	if sm.string != "*" {
		t.Errorf(`should be "*" [was: "%v"]`, sm)
	}
}

func TestSuffixMatcherMatchesEndingSuffix(t *testing.T) {
	sm := SuffixMatcher{""}
	sm.Add('/')

	if !sm.Matches("src/") {
		t.Errorf(`Matches("src/") should be true`)
	}
}

func TestSuffixMatcherIgnoresLeadingSuffix(t *testing.T) {
	sm := SuffixMatcher{""}
	sm.Add('/')

	if sm.Matches("/src") {
		t.Errorf(`Matches("/src") should be false`)
	}
}
