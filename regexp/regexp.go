/*
Supplements the regexp package from the standard and makes named capturing groups easily accessible.

Usage:
    func main() {
      pattern := `/(?P<country>[^/]+)/(?P<city>[^/]+)`
      regex := MustCompile(pattern)
      candidate := "/Germany/Dresden"
      match := regex.Match(candidate)
      
      // accessing capturing group by name
      fmt.Println("country: ", match.NamedGroups["country"])
      
      // accessing capturing group by index 
      fmt.Println("city: ", match.Groups[2])
    }
*/
package regexp

import "regexp"

// Custom type for extending regexp.Regexp.
type RichRegexp struct {
	regexp.Regexp
}

// Like regexp.Compile from the standard library, but returns *RichRegexp.
func Compile(expr string) (*RichRegexp, error) {
	re, error := regexp.Compile(expr)
	if error != nil {
		return nil, error
	}
	return &RichRegexp{*re}, error
}

// Like regexp.MustCompile from the standard library, but returns *RichRegexp.
func MustCompile(expr string) *RichRegexp {
	return &RichRegexp{*regexp.MustCompile(expr)}
}

// Like regexp.CompilePOSIX from the standard library, but returns *RichRegexp.
func CompilePOSIX(expr string) (*RichRegexp, error) {
	re, error := regexp.CompilePOSIX(expr)
	return &RichRegexp{*re}, error
}

// Like regexp.MustCompilePOSIX from the standard library, but returns *RichRegexp.
func MustCompilePOSIX(expr string) *RichRegexp {
	return &RichRegexp{*regexp.MustCompilePOSIX(expr)}
}

// Matches against the candidate and returns a Match struct or nil if the candidate doesn't match.
func (re *RichRegexp) Match(candidate string) *Match {
	matches := re.FindStringSubmatch(candidate)
	if matches == nil {
		return nil
	}
	return re.newMatch(matches)
}

func (re *RichRegexp) newMatch(matches []string) *Match {
	groupMap := make(map[string]string)
	for i, name := range re.SubexpNames() {
		// ignore the whole match and unnamed groups
		if i == 0 || name == "" {
			continue
		}
		groupMap[name] = matches[i]
	}
	return &Match{matches, groupMap}
}

/* 
'Groups' contains all the strings captured by the capturing groups, the first element is the whole match. 

'NamedGroups' contains only the strings captured by named capturing groups, the names of the capturing groups are used as keys.
*/
type Match struct {
	Groups      []string
	NamedGroups map[string]string
}
