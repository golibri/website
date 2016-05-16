[![Build Status](https://travis-ci.org/golibri/website.svg?branch=master)](https://travis-ci.org/golibri/website)
[![Code Climate](https://codeclimate.com/github/golibri/website/badges/gpa.svg)](https://codeclimate.com/github/golibri/website)
[![GoDoc](https://godoc.org/github.com/golibri/website?status.svg)](https://godoc.org/github.com/golibri/website)
[![Built with Spacemacs](https://cdn.rawgit.com/syl20bnr/spacemacs/442d025779da2f62fc86c2082703697714db6514/assets/spacemacs-badge.svg)](http://github.com/syl20bnr/spacemacs)

# golibri/website
Get Metadata from HTML

# Requirements
`go get -u github.com/golibri/fetch`

# installation
`go get -u github.com/golibri/website`

# usage
````go
import "github.com/golibri/website"

func main() {
    ws := website.FromURL("http://example.com/whatever")
    // OR:
    ws := website.Parse("website-html-string")
    // ws is a Website object, see below
}
````

# data fields
A **Website** has the following data fields:

````go
type Website struct {
    URL         string
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
