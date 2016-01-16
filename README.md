# golibri/website
Get Metadata from HTML

# installation
`go get github.com/golibri/website`

# depenencies
`github.com/PuerkitoBio/goquery`
`github.com/endeveit/guesslanguage`
`

# usage
````go
ws := Parse("website-html-string")
````

# data fields
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
