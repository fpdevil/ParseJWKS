package utils

import (
	"net"
	"net/http"
	"time"
)

const (
	timeout = time.Duration(2 * time.Second)
)

// dialTimeout function provides a custom handler for making
// remote http calls using a custom timeout value
func dialTimeout(network, address string) (net.Conn, error) {
	return net.DialTimeout(network, address, timeout)
}

// get function makes a http GET call to a remote url and returns
// the output and it is the responsibility of the calling function
// to handle any errors if returned by this
func get(url string) (*http.Response, error) {
	transport := http.Transport{
		Dial: dialTimeout,
	}

	client := http.Client{
		Transport: &transport,
	}

	return client.Get(url)
}

// The Fetcher function calls a series of urls provided as argument
// and compiles all the information into a channel that is made
// available for any downstream function to call and process
func Fetcher(done <-chan interface{}, urls ...string) <-chan Result {
	results := make(chan Result)
	go func() {
		defer close(results)
		for _, url := range urls {
			var result Result
			resp, err := get(url)
			result = Result{
				Response: resp,
				Error:    err,
			}

			select {
			case <-done:
				return
			case results <- result:
				// do nothing...
			}
		}
	}()

	return results
}
