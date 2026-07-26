# General Client concepts
- Client sends an HTTP request to a Server, and the Server sends a Response
- HTTP is an internet protocol.
# Go Specific Things to learn regarding clients

## Example code
### Typical structure
```golang
import (
    "io"
    "net/http"
)

func getHttpCall() {
    res, err := http.Get("https://website.website.com/path/to/specific/part")
    if err != nil {
        // handle error
        // or return a more specific
    }
    defer res.Body.Close()
    
    data, err := io.ReadAll(res.Body)
    if err != nil {
        // handle error
        // or throw more specific error
    }

    return data, err
}
    
```