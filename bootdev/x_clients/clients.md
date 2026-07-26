# General Client concepts

### HTTP

- Client sends an HTTP request to a Server, and the Server sends a Response
- HTTP is an internet protocol.

### JSON

- JSON gets sent as a stream of bytes. These can be converted to a string, or however the receiving machine would like to handle it.

### DNS

- DNS handles the human readable equivalent of the specific Domain, IP, and port of a website

### HTTP Headers

- Additional case-insensitive information send with the request / response.

### HTTP Methods

- GET
- POST
- PUT
- DELETE

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

### Decoding JSON via json.Decoder

- Note that this pattern is slightly different from the above

```golang
import (
    "encoding/json"
)

// struct for JSON
// Must be expored (first leter capitalized)
type Issue struct {
    Id          string  `json:"id"`
    Title       string  `json:"title"`
    Estimate    int     `json:"estimate"`
}

func getIssues(url string) ([]Issue, error) {
    res, err := http.get(url)
    if err != nil {
        // handle error
        // or return more specific error
    }
    defer res.Body.Close()

    var issues []Issue
    decoder := json.NewDecoder(res.Body)
    if err := decoder.Decode(&issues); err != nil {
        // handle error
        // or return more specific error
    }

    return issues, nil
}
```

### Decoding JSON via json.Unmarshal

- Typically reach for `json.Unmarshal` if you already have JSON data in memory. For HTTP calls generally stick with `json.Decoder`
- Note in the example code that information was first read through `io.ReadAll(res.Body)`.
- The reverse of `json.Unmarshal` can be accomplished using `json.Marshal`.

```golang
...
// I believe this first puts it into memory
data, err := io.ReadAll(res.Body)
if err != nil {
    // handle error
    // or return more specific error
}

var issues []Issue
if err := json.Unmarshal(data, &issues); err != nil {
    // handle error
    // or return more specific error
}
...
```

### Use "net/url" to parse a URL

- You can [parse a URL](https://pkg.go.dev/net/url#Parse) string into a URL struct that can be further examined for specific parts of the URL

```golang
func main() {
	// Parse + String preserve the original encoding.
	u, err := url.Parse("https://example.com/foo%2fbar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Path)
	fmt.Println(u.RawPath)
	fmt.Println(u.String())
}
```

### Use "net/url" to manage headers

- Can use the `.Header.Set()` and `.Header.Get()` on a `http.Request` / `http.Response` [link](https://pkg.go.dev/net/http#Header).

```golang
import (
    "net/http"
)

func getResponseContentType() {
    res, err := http.Get("https://website.website.com/path/to/specific/part")
    if err != nil {
      // handle error
      // or return a more specific
    }
    defer res.Body.Close()

    res.Header.Get("Content-Type")
}

func createNewReqWithHeader() {
  req, err := http.NewRequest("GET", "https://website.website.com/path/to/specific/part", nil)
  if err != nil {
    // handle error
    // or return a more specific
  }

  req.Header.Set("super-amazing-key", "QWERT")

  client := http.Client{}
  res, err := client.Do(req)
  if err != nil {
    // handle error
    // or return a more specific
  }
}
```

### Creating a POST request

```golang

import (
  "bytes"
  "encoding/json"
)

type BodyStruct struct {
  Id string `json:"id"`
  Username string `json:"username`
}

func createPostRequest(url string, body BodyStruct) {
  jsonData, err := json.Marsal(body)
  if err != nil {
    // handle error
  }

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
  if err != nil {
    // handle error
  }
}
```

### Retrieve status code from an `http.Response`

```golang
res, err := http.Get("https://website.website.com/path/to/specific/part")
if err != nil {
  // handle error
}
defer res.Body.Close()
statusCode := res.StatusCode
```
