# urlquery

```go
package main

import (
    "fmt"
    "net/url"

    "github.com/yhat/urlquery"
)

type EndpointOpts struct {
    Verbose  bool   `url:"v"`
    Username string `url"username"`
}

func main() {
    opts := EndpointOpts{}
    urlAddr, _ := url.Parse("http://example.com/?v=1&username=foo")
    urlquery.Unmarshal(urlAddr.Query(), &opts)
    fmt.Println(opts)
}
```

```bash
$ go run example.go
{true foo}
```

Shamelessly stolen from [fsouza's](https://github.com/fsouza) Docker client.
