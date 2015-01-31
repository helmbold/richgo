Rich Go
=======

*Rich Go* is a library to enrich the [Go](http://golang.org/) (golang) standard library.

For now it contains only a more comfortable to use version of regexp.

Usage:

    import "github.com/helmbold/richgo"

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
