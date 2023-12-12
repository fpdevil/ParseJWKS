package utils

import (
	"fmt"
	"io"
)

// The Processor function takes the Result channel populated by
// the Fetcher function and processes the same and populates the
// channel Data with data complying to the Data struct
func Processor(done <-chan interface{}, results <-chan Result) <-chan Data {
	processed := make(chan Data)
	go func() {
		defer close(processed)
		var errcount int
		for res := range results {
			var data Data
			if res.Error != nil {
				fmt.Printf("error %v\n", res.Error.Error())
				errcount++
				data.Count = errcount
				data.Error = res.Error
			} else {
				data.Url = res.Response.Request.URL.Host
				data.Status = res.Response.Status

				res, err := io.ReadAll(res.Response.Body)
				if err != nil {
					data.Error = err
				}
				data.Response = res
			}

			select {
			case <-done:
				return
			case processed <- data:
				// nothing here
			}
		}
	}()

	return processed
}
