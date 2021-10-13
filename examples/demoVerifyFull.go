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
  "spec_version": "1.0",
  "id": "playbook--a0777575-5c4c-4710-9f01-15776103837f",
  "name": "Playbook 1",
  "created": "2021-01-25T20:31:31.319Z",
  "modified": "2021-01-25T20:31:31.319Z",
  "signatures": [
    {
      "type": "signature",
      "spec_version": "1.0",
      "id": "signature--af892292-c4b4-47eb-9be6-4897ff4b9388",
      "created_by": "identity--uuid2",
      "created": "2021-01-25T20:31:31.319516Z",
      "modified": "2021-01-25T20:31:31.319516Z",
      "signee": "ACME Cyber Company",
      "valid_from": "2021-01-25T20:31:31.319516Z",
      "valid_until": "2022-01-01T12:12:12.123456Z",
      "related_to": "playbook--a0777575-5c4c-4710-9f01-15776103837f",
      "related_version": "2021-01-25T20:31:31.319Z",
      "sha256": "hHuhBwKscfqvLC3y2FfZtHi3DNkzE0o8kE8eE6x50pM",
      "algorithm": "RS256",
      "public_keys": [
        "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAptKZyFPStvmOlb0WihOBhlHUr6wFDHC+tW7hJAudfTQ5mHZQpB8PoMz07udZA+dG8dhUIPkmXlp1TgREeYTHdhxhuf0y/GhbpZv5JPYHx3watO+HWO2qYkjRMEcrWhPMdaVkS/Xe/liaMcow4jYoWaFm8VobeYsyVD2bWWdyl4joTEETm1Z47RnnfR15kVhVudVrDzEFmM4nXV/6dmIg184RJE4httwBFxR8qZCQCwTiJmsoyJxfUR0Gs4ePKc5sB0NTkmFZc5klQSitd67RJn2ldhbqE7EpDl4XlIt+UyLJm1guCBltia8Agke7dXuhpB7hQ6LJwY4EjzthkJ8IPwIDAQAB"
      ],
      "value": "lfmqOpMlNcUb4coQ9n6RhFqKCLCocqTEdyb9S4t5F4INN9Q4pXPAUpd28hnVS-D3BgmPACq6dQgNY1nXnU-QqcChlVDGeliRTu5OLULrBCkQTZ8OcAhyUprXYP4vhzN81w-eSmQz9urEGe98o2RbhLbZCrEuBUqgvmPdsu5cUnJr9wdkMHwoToS-rbc_xuWHQAFzqi0YarCAfbPop0jDQxO8KNDFIoy98mjbL2FXv0Y4GQOSZaJNgZpxdSmgqpQfF5vxOEzQpwirvoUkjGydroJsim7XhAsQwiQwEuegl0GzawhIODVMVz2ZIW0jByUnCH2G21oa1mlA2sX5nciGKw"
    }
  ]
}
`)

	// ---------------------------------------------------------------------
	// Step 1.0
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 1.0: Receive a JSON playbook object to verify")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	p, _ := playbook.Decode(rawPlaybookData)
	fmt.Println(string(rawPlaybookData))

	// ---------------------------------------------------------------------
	// Step 1.1
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 1.1: Capture the signature object from step 1.0")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	s := p.Signatures[0]
	jsonSigDataWithSig, _ := s.Encode()
	fmt.Println(string(jsonSigDataWithSig))

	// ---------------------------------------------------------------------
	// Step 1.2
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 1.2: Parse the public key from step 1.1")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	// Parse the public key
	publicKeyDer, err := base64.RawStdEncoding.DecodeString(s.PublicKeys[0])
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
	// Step 1.3
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 1.3: Remove the digital signature from the signature object from step 1.1")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	sigReceived := s.Value
	// Zero out the signature value so we can recompute the signature for verification
	s.Value = ""
	jsonSigData, _ := s.Encode()
	fmt.Println(string(jsonSigData))
	fmt.Println("\nFull Signature: ", sigReceived)

	// ---------------------------------------------------------------------
	// Step 1.4
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 1.4: Create canonical version of signature object from step 1.3")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	jcsSigData, _ := jcs.Transform(jsonSigData)
	fmt.Println(string(jcsSigData))

	// ---------------------------------------------------------------------
	// Step 1.5
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 1.5: Create base64URL.encoded version of the JCS signature from step 1.4")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	b64jcsSigData := base64.RawURLEncoding.EncodeToString([]byte(jcsSigData))
	fmt.Println(b64jcsSigData)

	// ---------------------------------------------------------------------
	// Step 2.0
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 2.0: Verify the signature received in step 1.3 of the B64JCS data from step 1.5 using the public key form 1.2 and using the algorithm from the signature object RS256")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	method := jwt.SigningMethodRS256
	err = method.Verify(b64jcsSigData, sigReceived, publicKey)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Signature is valid")
	}

	// ---------------------------------------------------------------------
	// Step 3.0
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 3.0: Compute the hash of the playbook and make sure it matches the hash in the signed signature")
	fmt.Println("------------------------------------------------------------------------------------------------------------")

	// ---------------------------------------------------------------------
	// Step 3.1
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 3.1: Remove existing signatures before computing hash")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	//savedSignatures := p.Signatures
	p.Signatures = nil
	jsonPlaybookData, _ := p.Encode()
	fmt.Println(string(jsonPlaybookData))

	// ---------------------------------------------------------------------
	// Step 3.2
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 3.2: Create JCS canonical version of the playbook from step 3.1")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	jcsPlayBookData, _ := jcs.Transform(jsonPlaybookData)
	fmt.Println(string(jcsPlayBookData))

	// ---------------------------------------------------------------------
	// Step 3.3
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 3.3: Create SHA256 (in hex) of canonical version of playbook from step 3.2")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	hashjcsPlaybookData := sha256.Sum256(jcsPlayBookData)
	fmt.Println(fmt.Sprintf("%x", hashjcsPlaybookData[:]))

	// ---------------------------------------------------------------------
	// Step 3.4
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 3.4: Create base64URL.encoded version of the SHA256 hash from step 1.3 and remove any padding")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	// Remove base64 padding characters "=" per RFC 7515 section 2 - Base64url Encoding
	b64hashjcsPlaybookData := base64.RawURLEncoding.EncodeToString(hashjcsPlaybookData[:])
	fmt.Println(b64hashjcsPlaybookData)

	// ---------------------------------------------------------------------
	// Step 4.0
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 4.0: Compare computed hash with hash found in signed signature")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Computed Hash: ", b64hashjcsPlaybookData)
	fmt.Println("Signed Hash:   ", s.SHA256)
	if b64hashjcsPlaybookData == s.SHA256 {
		fmt.Println("Hashes match and content is valid")
	} else {
		fmt.Println("Hashes do not match and content is not valid")
	}

}
