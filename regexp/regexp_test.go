package regexp

import (
	"testing"
)

func TestCompile(t *testing.T) {
	regex, err := Compile(`[a-z]+`)
	if err != nil {
		t.Errorf("Cannot compile")
	}
	
	match := regex.Match("avc")
	
	if match == nil {
		t.Errorf("regex doesn't work")
	}
}

func TestReturnCapturingGroups(t *testing.T) {
	pattern := `(?P<first>\d+)abc(?P<second>\s{2})(?P<third>\w{2}(?P<fourth>[xyz]+))`
	regex := MustCompile(pattern)
	candidate := "12abc  vwyz"

	match := regex.Match(candidate)

	expectedCaptures := []string{candidate, "12", "  ", "vwyz", "yz"}
	for i, expactedCapture := range expectedCaptures {
		if match.Groups[i] != expactedCapture {
			t.Errorf("expected '%s', but was '%s'", expactedCapture, match.Groups[i])
		}
	}
	if len(match.NamedGroups) != 4 {
		t.Errorf("wrong group count")
	}
	assertGroupAndCaptureInMap("first", expectedCaptures[1], match, t)
	assertGroupAndCaptureInMap("second", expectedCaptures[2], match, t)
	assertGroupAndCaptureInMap("third", expectedCaptures[3], match, t)
	assertGroupAndCaptureInMap("fourth", expectedCaptures[4], match, t)
}

func TestReturnCapturingGroups2(t *testing.T) {
	pattern := `/(?P<country>[^/]+)/(?P<city>[^/]+)`
	regex := MustCompile(pattern)
	candidate := "/Germany/Dresden"

	match := regex.Match(candidate)

	if len(match.NamedGroups) != 2 {
		t.Errorf("wrong group count")
	}
	assertGroupAndCaptureInMap("country", "Germany", match, t)
	assertGroupAndCaptureInMap("city", "Dresden", match, t)
}

func assertGroupAndCaptureInMap(group, capture string, match *Match, t *testing.T) {
	if match.NamedGroups[group] != capture {
		t.Errorf("capturing group '%s' missing", capture)
	}
}

func TestCompilePOSIX(t *testing.T) {
	regex, err := CompilePOSIX(`[a-z]+`)
	if err != nil {
		t.Errorf("Cannot compile")
	}
	
	match := regex.Match("avc")
	
	if match == nil {
		t.Errorf("regex doesn't work")
	}
}

func TestMustCompilePOSIX(t *testing.T) {
	regex := MustCompilePOSIX(`[a-z]+`)
	
	match := regex.Match("avc")
	
	if match == nil {
		t.Errorf("regex doesn't work")
	}
}

