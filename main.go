package main

import (
	"bytes"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fpdevil/parse-jwks/utils"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		// fmt.Fprintf(os.Stderr, "usage: go run %s <input jwks>\n", filepath.Base(args[0]))
		fmt.Fprintf(os.Stderr, "usage: go run %s <jwks url>\n", filepath.Base(args[0]))
		return
	}

	done := make(chan interface{})
	defer close(done)

	fetch := utils.Fetcher(done, args[1:]...)
	data := utils.Processor(done, fetch)

	// for v := range data {
	// 	fmt.Println(string(v.Response))
	// }

	var jwks utils.JWKS
	keys := utils.Parser(done, data, jwks)
	for v := range keys {
		if v.Error != nil {
			log.Printf("error parsing data: %v", v.Error)
			return
		}

		for key, value := range v.Blocks {
			fmt.Println()
			log.Printf("* Kid: %s", key)

			ps := new(bytes.Buffer)
			pem.Encode(ps, &value)
			log.Printf("* Public Key:\n%s", ps.String())
		}
	}
}
