# elefind-go
Elefind is allowing testers to find any html elements in cnx books with css selector in very short time.

# IMPORTANT
Please consider that Elefind is searching elements inside `[data-type="page"] and [data-type="composite-page"]` so you will not be able to find those attributes.

# How to use
1. Update `./books` with `.xhtml` files.
2. Run `./scripts/remove-namespaces ./books/book-name.xhtml` - this is removing namespaces from elements. Ex. `<m:math>` is changed to `<math>`. If you will omit this step then there will be no way to find those elements.
3. Update `config.go`
4. Run `go run main.go config.go limit.go` or build with `go build main.go config.go limit.go` and start with `./main`
You can pass port as a argument: `./main 3001`

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
