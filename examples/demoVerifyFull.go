// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/gowebpki/jcs"
	"github.com/openplaybooks/libcacao/objects/playbook"
)

func main() {
	// 	publicKeyData := []byte(`
	// -----BEGIN PUBLIC KEY-----
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAptKZyFPStvmOlb0WihOBhlHUr6wFDHC+
	// tW7hJAudfTQ5mHZQpB8PoMz07udZA+dG8dhUIPkmXlp1TgREeYTHdhxhuf0y/GhbpZv5JPYHx3wa
	// tO+HWO2qYkjRMEcrWhPMdaVkS/Xe/liaMcow4jYoWaFm8VobeYsyVD2bWWdyl4joTEETm1Z47Rnn
	// fR15kVhVudVrDzEFmM4nXV/6dmIg184RJE4httwBFxR8qZCQCwTiJmsoyJxfUR0Gs4ePKc5sB0NT
	// kmFZc5klQSitd67RJn2ldhbqE7EpDl4XlIt+UyLJm1guCBltia8Agke7dXuhpB7hQ6LJwY4Ejzth
	// kJ8IPwIDAQAB
	// -----END PUBLIC KEY-----
	// `)

	// 	privateKeyData := []byte(`
	// -----BEGIN PRIVATE KEY-----
	// MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCm0pnIU9K2+Y6VvRaKE4GGUdSv
	// rAUMcL61buEkC519NDmYdlCkHw+gzPTu51kD50bx2FQg+SZeWnVOBER5hMd2HGG5/TL8aFulm/kk
	// 9gfHfBq074dY7apiSNEwRytaE8x1pWRL9d7+WJoxyjDiNihZoWbxWht5izJUPZtZZ3KXiOhMQROb
	// VnjtGed9HXmRWFW51WsPMQWYziddX/p2YiDXzhEkTiG23AEXFHypkJALBOImayjInF9RHQazh48p
	// zmwHQ1OSYVlzmSVBKK13rtEmfaV2FuoTsSkOXheUi35TIsmbWC4IGW2JrwCCR7t1e6GkHuFDosnB
	// jgSPO2GQnwg/AgMBAAECggEAKT6KTNAEmb5rdTPxvaOC832J0wD5opDBZcQLH8lLX6go0Tv3Rgxz
	// 5bKmn+ZMyL1GegadDiXrSYqd0/MUJuMgGWB8/OnP0D3Q4soEOBIn7DcPt0o9MUxZQsF0DraZzkR0
	// 2WVRvcIFJucrAEJYAaWYJkjUVbmMb2ltwQwWO21rFHGbpE73nsfr/oAWsZEvKsQZoYm4fh5jVI5+
	// wKyRnKaN1uqAcNgj75cdywCHBVwgEefEgOPM77CDMH0+JumSirQiBfR35+HWRwHwpm09wI6Aqtvg
	// y5bzxvLDDRgrhX4LCPtUHGrUXNJHRKYiHQX6P6bIVuBrHV6VFpyS+5weu0w6kQKBgQDQo4QeLtO7
	// S3KH8UL3lX4lhH1K7/Q99uBHmvLXdiDkHjLbBbh0JfrHgHtnK9bvJ2GvVcwhI9fTiO1p1o5RM5jb
	// iVUSCS91sLcTPFv8X83sExBZnrvlSlb/va+4yW+Lzvr6ZiDlZYsVRNvNAHUTojHRCOH2P4eX1+ql
	// 5P4FMdfvSQKBgQDMsQ4LBpxjD9KdDzJzw9a0xbL47QdCeZBqNUy6MvwLE0+KsF+prvoigNZCaTcJ
	// 2FfoPxpE3/o0A/byCTuDkfddrd/hcAO0gd1R9CYJDXJfnIbZfheUmHW7ShbXyqhpqQKVjzH+jnLq
	// VjbGD6tz3dN+AwNgULD/vvwXM2TWpu9TRwKBgGkPPdMZD2NLzaNouKkFbR0lRxY6GEovi6Zi/w/C
	// GzPjhQZHLifGjC5zozBDohqRQR5SXNT/QInzdGGMOePn0HwT/nNzjqN71eRoy4UdFQtgWiZWyRTf
	// x0lGUjsBrBrBoh3+2WfKJywRnYDwTwQQ83boOyiNuxCaGD1rPwKMo8iJAoGAPIePE4uc615edbts
	// u/cJouNjjWDqaKnyHrYsPlOdXNkVCHonj9ICffmDYpgignLLbA5dAkkJgCA8Ak7gnoOnlrg4ID4z
	// mklc3UNJjBvB2qw65E35QyPijMPYBXAUZUppTTjPG+ub59ge0msH1Hegdv8FHJJABSDBA0tbYm5z
	// DzkCgYA9/0KtWKFMhF3v01L54AXF5b15RroBhZAfzI1U0wPO4J6Tz+1KqmtrwHTBPI36nzITIhlM
	// hcoTsMRMgnv0NHzxlcQQmAy3foFBFOyHXql3hPtWbEViB5jQs4cP5ts1oivVhrEtrrE51TG4V/Ef
	// fD1JKiHl7MECYEMyBz31PsRCuw==
	// -----END PRIVATE KEY-----
	// `)

	// ---------------------------------------------------------------------
	// This is minimal demo playbook used at the start of this example
	// ---------------------------------------------------------------------
	rawPlaybookData := []byte(`
{
  "type": "playbook",
  "spec_version": "2.0",
  "id": "playbook--a0777575-5c4c-4710-9f01-15776103837f",
  "name": "Playbook 1",
  "created_by": "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1",
  "created": "2022-05-18T11:31:31.319Z",
  "modified": "2022-05-18T11:31:31.319Z",
  "signatures": [
    {
      "type": "jss",
      "id": "jss--af892292-c4b4-47eb-9be6-4897ff4b9388",
      "created_by": "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1",
      "created": "2023-01-10T17:39:31.319Z",
      "modified": "2023-01-10T17:39:31.319Z",
      "signee": "ACME Cyber Company",
      "valid_from": "2023-01-10T17:39:31.319Z",
      "valid_until": "2023-06-10T17:39:31.319Z",
      "related_to": "playbook--a0777575-5c4c-4710-9f01-15776103837f",
      "related_version": "2022-05-18T11:31:31.319Z",
      "hash_algorithm": "sha-256",
      "algorithm": "RS256",
      "public_key": "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAptKZyFPStvmOlb0WihOBhlHUr6wFDHC+tW7hJAudfTQ5mHZQpB8PoMz07udZA+dG8dhUIPkmXlp1TgREeYTHdhxhuf0y/GhbpZv5JPYHx3watO+HWO2qYkjRMEcrWhPMdaVkS/Xe/liaMcow4jYoWaFm8VobeYsyVD2bWWdyl4joTEETm1Z47RnnfR15kVhVudVrDzEFmM4nXV/6dmIg184RJE4httwBFxR8qZCQCwTiJmsoyJxfUR0Gs4ePKc5sB0NTkmFZc5klQSitd67RJn2ldhbqE7EpDl4XlIt+UyLJm1guCBltia8Agke7dXuhpB7hQ6LJwY4EjzthkJ8IPwIDAQAB",
      "value": "hhezAh4ncNfLYZj4t85pR_cWBIGv3JoFSLY0a5kxzC53e5QLfBG1EQHC3CELnuL8zJLz9NGd9I_3XebpLrcehYbUPwv9wzvx0IwHXlZIPVTEmTvPJ6635rOuPMF6M42bRK2V7P8vEU3CWXme_u2rBKEEGAwSTx8u4yxZTpU5mQbBMZemuic6dzpGXoynhuFP8Y-fZXVIafWGBO0-9FWfjq8qJY0CjVuoUz-3Vnm1p9k_GiVCiBa0oPMXOEXCvlyyfPlGXpGFEXnkkSb_LfmbPEEdeSUlu9ak_aPF69l8hg-M7ceQJXOLNapmX6t6u5kQCdYFO4KWE14zodL45z6IEg"
    }
  ]
}
`)

	// ---------------------------------------------------------------------
	// Step 1
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 1: Receive and parse the JSON playbook object to verify")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	p, _ := playbook.Decode(rawPlaybookData)
	fmt.Println(string(rawPlaybookData))

	// ---------------------------------------------------------------------
	// Step 2
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 2: Capture and remove the digital signature from step 1")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	ds := p.Signatures[0].Value
	fmt.Println("DS: ", ds)
	p.Signatures[0].Value = ""
	// Since it was Decoded/Unmarshalled in Step 1 we need to encode/remarshal it here to print it out
	jsonSigDataWithoutSig, _ := p.Encode()
	fmt.Println(string(jsonSigDataWithoutSig))

	// ---------------------------------------------------------------------
	// Step 3
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 3: Parse or fetch the public key from step 2")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	// Parse the public key
	publicKeyDer, err := base64.RawStdEncoding.DecodeString(p.Signatures[0].PublicKey)
	if err != nil {
		panic("failed to decode public key: " + err.Error())
	}
	parsedKey, err := x509.ParsePKIXPublicKey(publicKeyDer)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	switch parsedKey := parsedKey.(type) {
	case *rsa.PublicKey:
		fmt.Println("public key is of type RSA:", parsedKey)
	case *dsa.PublicKey:
		fmt.Println("public key is of type DSA:", parsedKey)
	case *ecdsa.PublicKey:
		fmt.Println("public key is of type ECDSA:", parsedKey)
	case ed25519.PublicKey:
		fmt.Println("public key is of type Ed25519:", parsedKey)
	default:
		panic("unknown type of public key")
	}

	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		panic("Unable to parse RSA public key with rsa.PublicKey")
	}
	fmt.Println("Public RSA Key E: ", publicKey.E)
	fmt.Println("Public RSA Key N: ", publicKey.N)

	// ---------------------------------------------------------------------
	// Step 4
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 4: Create canonical version of signature object from step 2")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	jcsData, _ := jcs.Transform(jsonSigDataWithoutSig)
	fmt.Println(string(jcsData))

	// ---------------------------------------------------------------------
	// Step 5
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 5: Create a hash (SHA256 in hex) of the canonical version of playbook from step 4")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	temp1 := sha256.Sum256(jcsData)
	// fmt.Println(fmt.Sprintf("%x", hashjcsPlaybookData[:]))
	hashjcsPlaybookData := hex.EncodeToString(temp1[:])
	fmt.Println(hashjcsPlaybookData)

	// // ---------------------------------------------------------------------
	// // Step 1.5
	// // ---------------------------------------------------------------------
	// fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	// fmt.Println("Step 1.5: Create base64URL.encoded version of the JCS signature from step 1.4")
	// fmt.Println("------------------------------------------------------------------------------------------------------------")
	// b64jcsSigData := base64.RawURLEncoding.EncodeToString([]byte(jcsSigData))
	// fmt.Println(b64jcsSigData)

	// ---------------------------------------------------------------------
	// Step 6
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 6: Verify the signature received using the public key and algorithm (RS256) from the signature object")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	method := jwt.SigningMethodRS256
	err = method.Verify(hashjcsPlaybookData, ds, publicKey)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Signature is valid")
	}

}
