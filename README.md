# elefind-go
Elefind is allowing testers to find any html elements in cnx books with css selector in very short time.

# How to use
Update `./books` with `.xhtml` files.
Update `config.go`
Run `go run main.go config.go` or build with `go build main.go config.go`
Web server will start listening at port 3000
`/` -> `{status: "active"}`
`/books` -> json list of books from `config.go`
`/element?bookName=Prelagebra&element=[data-type="newline"]`

You can use `https://github.com/katalysteducation/elefind` for front-end

# Available selectors
All css selectors are available +

Find text inside element:
`element:hasText(text to search)`

Find elements inside other elements:
`table:has(img, p)`

IMPORTANT
Please consider that Elefind is searching elements inside `[data-type="page"] and [data-type="composite-page"]` so you will not be able to find those attributes.
