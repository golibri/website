[![GoDoc](https://godoc.org/github.com/golibri/website?status.svg)](https://godoc.org/github.com/golibri/website)
[![Built with Spacemacs](https://cdn.rawgit.com/syl20bnr/spacemacs/442d025779da2f62fc86c2082703697714db6514/assets/spacemacs-badge.svg)](http://github.com/syl20bnr/spacemacs)

# golibri/website
Get Metadata from HTML

# installation
`go get -u github.com/golibri/website`

# usage
````go
import "github.com/golibri/website"

func main() {
    //... get a HTML string from anywhere, for example with golibri/fetch
    ws := website.Parse("website-html-string")
    // ws is a Website object, see below
}
````

# data fields
A **Website** has the following data fields:

````go
type Website struct {
    Url         string
    Body        string
    Language    string
    Title       string
    Description string
    Image       string
    Favicon     string
    Feeds       []string
    Tags        []string
    }
````

# license
LGPLv3. (You can use it in commercial projects as you like, but improvements/bugfixes must flow back to this lib.)
