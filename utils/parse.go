package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
)

// Parser function is  the downstream function that  takes data from
// the channels populated with data  complying to Data and populates
// the output channel KeyData with JWKS parsed information complying
// with the KeyData struct
func Parser(done <-chan interface{}, processed <-chan Data, jwks JWKS) <-chan KeyData {
	keys := make(chan KeyData)
	// anonymous go routine
	go func() {
		defer close(keys)
		var blockmap map[string]pem.Block
		for data := range processed {
			var k KeyData
			if err := json.Unmarshal(data.Response, &jwks); err != nil {
				k.Blocks = nil
				k.Error = err
			}
			blockmap = make(map[string]pem.Block)
			for i := 0; i < len(jwks.Keys); i++ {
				// log.Printf("processing key %d", i)
				key := jwks.Keys[i]

				if key.Kty != "RSA" {
					k.Blocks = nil
					k.Error = fmt.Errorf("error due to invalid key type: %s", key.Kty)
				}

				// decode the base64 encoded bytes of the modulus value
				nb, err := base64.URLEncoding.DecodeString(key.N)
				if err != nil {
					k.Blocks = nil
					k.Error = err
				}

				// exponent value; the default exponent is usually 65537, so we will
				// just compare the base64 value for [1,0,1] or [0,1,0,1]
				var exp int
				if key.E == "AQAB" || key.E == "AAEAQ" {
					exp = 65537
				} else {
					k.Blocks = nil
					k.Error = fmt.Errorf("unable to decode the exponent E: %s", key.E)
				}

				publicKey := &rsa.PublicKey{
					N: new(big.Int).SetBytes(nb),
					E: exp,
				}

				der, err := x509.MarshalPKIXPublicKey(publicKey)
				if err != nil {
					k.Blocks = nil
					k.Error = err
				}

				var block pem.Block
				block.Type = "RSA PUBLIC KEY"
				block.Bytes = der

				blockmap[key.Kid] = block
				k.Blocks = blockmap
				k.Error = nil
			}

			select {
			case <-done:
				return
			case keys <- k:
				// nothing to be done...
			}
		}
	}()
	return keys
}

func ParseJWKS(jwks []byte, data JWKS) (map[string]pem.Block, error) {
	blockmap := make(map[string]pem.Block)
	if err := json.Unmarshal(jwks, &data); err != nil {
		log.Printf("error during json unmarshalling: %s", err.Error())
		return nil, err
	}

	for i := 0; i < len(data.Keys); i++ {
		key := data.Keys[i]
		if key.Kty != "RSA" {
			err := fmt.Errorf("error due to invalid key type: %s", key.Kty)
			return nil, err
		}

		// decode the base64 encoded bytes of the modulus value
		// log.Printf("decoding the Modulus N: %v", key.N)
		nb, err := base64.URLEncoding.DecodeString(key.N)
		if err != nil {
			log.Printf("error decoding the Modulus N: %s", err.Error())
			return nil, err
		}

		// exponent value; the default exponent is usually 65537, so we will
		// just compare the base64 value for [1,0,1] or [0,1,0,1]
		var exp int
		if key.E == "AQAB" || key.E == "AAEAAQ" {
			exp = 65537
		} else {
			err := fmt.Errorf("unable to decode exponent E: %s", key.E)
			return nil, err
		}

		publicKey := &rsa.PublicKey{
			N: new(big.Int).SetBytes(nb),
			E: exp,
		}

		der, err := x509.MarshalPKIXPublicKey(publicKey)
		if err != nil {
			log.Printf("unable to convert public key to PKIX/ASN.1 DER: %s", err.Error())
			return nil, err
		}

		var block pem.Block
		block.Type = "RSA PUBLIC KEY"
		block.Bytes = der

		blockmap[key.Kid] = block
		// ps := new(bytes.Buffer)
		// pem.Encode(ps, &block)
		// log.Printf("key for kid: %s\n%s\n", key.Kid, ps.String())
	}

	return blockmap, nil
}
