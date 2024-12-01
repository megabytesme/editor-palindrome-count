# editor-palindrome-count
A web service in Go which provides the palindrome count in a provided string.

## Usage
First, build and run the service:

### Docker
- `docker build -t editor-palindrome-count .`
- `docker run -p 8080:8080 editor-palindrome-count`

### Directly
- `cd src/`
- `go mod tidy`
- `go run palindrome-count.go`

Then open `http://localhost:8080/count_palindromes?text=your_text_here` in your favourite browser.