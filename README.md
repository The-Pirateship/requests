# Requests - A simple package to manage API calls

Hey yall's! I'm new to the go ecosystem and really loved the language and wanted something fun as a starter project.
Since I come from the python world, here is a quick and simple abstraction for http requests for API similar to the requests library in python!
Any feedback, messages or improvements are welcome, I hope you find it useful!

# Installation

To install the package, use `go get`:

```bash
go get github.com/The-Pirateship/requests
```

# Usage

Using requests in go is very simple and similar to the requests package in python

1) First import the package

```go
import "github.com/The-Pirateship/requests"
```

2) Define your url and some payload (for post requests)
```go

url := "https://example.com/endpoint"

// A struct or map can be used here (anything that can be marshalled into json)
payload := map[string]interface{}{
  "email": "example.com",
  "data":  "dogs rule",
}
```

3) Make a request and process the output!

```go
response, err := requests.Post(url, payload)

if err != nil {
  // handle errors caused by parsing the payload, or if were unable to set the request here
}

if response.StatusCode != http.StatusOK {
  // handle non-200 status codes here!
}

// all good!
data := response.Body	// this is a []bytes type and can be unmarshalled into a struct or any format you need it in!

```

Heres examples for Get and Delete requests too!

## Get
```go
package main

import (
    "fmt"
    "github.com/The-Pirateship/requests"
)

func main() {
    url := "https://jsonplaceholder.typicode.com/posts/1"

    // Making the GET request
    response, err := requests.Get(url)
    if err != nil {
        fmt.Printf("Error occurred: %v\n", err)
        return
    }

    if response.StatusCode != http.StatusOK {
        fmt.Printf("Non-200 status code: %d\n", response.StatusCode)
        return
    }

    // Converting the response body to a string
    fmt.Printf("Response: %s\n", string(response.Body))
}
```

## Delete
```go
package main

import (
    "fmt"
    "github.com/The-Pirateship/requests"
)

func main() {
    url := "https://jsonplaceholder.typicode.com/posts/1"

    // Making the DELETE request
    response, err := requests.Delete(url)
    if err != nil {
        fmt.Printf("Error occurred: %v\n", err)
        return
    }

    if response.StatusCode != http.StatusOK { 
        fmt.Printf("Non-200 status code: %d\n", response.StatusCode)
        fmt.Printf("Response Body: %s\n", string(response.Body))
        return
    }

    fmt.Println("Delete request successful!")
}
```

I hope you find this project and I really learned a lot about go making it and in the past week learning the langauge, future plans surround extending this library to suppourt sessions and headers in an easy and convinent way!
Any feedback and advice is welcome, thanks!

contact: _team@thepirateship.net_

~ The PirateShip :)
