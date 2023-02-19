// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/gowebpki/jcs"
	"github.com/openplaybooks/libcacao/objects/playbook"
	"github.com/openplaybooks/libcacao/objects/signature"
)

func main() {
	// 	publicKeyData := `
	// -----BEGIN PUBLIC KEY-----
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAptKZyFPStvmOlb0WihOBhlHUr6wFDHC+
	// tW7hJAudfTQ5mHZQpB8PoMz07udZA+dG8dhUIPkmXlp1TgREeYTHdhxhuf0y/GhbpZv5JPYHx3wa
	// tO+HWO2qYkjRMEcrWhPMdaVkS/Xe/liaMcow4jYoWaFm8VobeYsyVD2bWWdyl4joTEETm1Z47Rnn
	// fR15kVhVudVrDzEFmM4nXV/6dmIg184RJE4httwBFxR8qZCQCwTiJmsoyJxfUR0Gs4ePKc5sB0NT
	// kmFZc5klQSitd67RJn2ldhbqE7EpDl4XlIt+UyLJm1guCBltia8Agke7dXuhpB7hQ6LJwY4Ejzth
	// kJ8IPwIDAQAB
	// -----END PUBLIC KEY-----
	// `

	privateKeyData := []byte(`
-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCm0pnIU9K2+Y6VvRaKE4GGUdSv
rAUMcL61buEkC519NDmYdlCkHw+gzPTu51kD50bx2FQg+SZeWnVOBER5hMd2HGG5/TL8aFulm/kk
9gfHfBq074dY7apiSNEwRytaE8x1pWRL9d7+WJoxyjDiNihZoWbxWht5izJUPZtZZ3KXiOhMQROb
VnjtGed9HXmRWFW51WsPMQWYziddX/p2YiDXzhEkTiG23AEXFHypkJALBOImayjInF9RHQazh48p
zmwHQ1OSYVlzmSVBKK13rtEmfaV2FuoTsSkOXheUi35TIsmbWC4IGW2JrwCCR7t1e6GkHuFDosnB
jgSPO2GQnwg/AgMBAAECggEAKT6KTNAEmb5rdTPxvaOC832J0wD5opDBZcQLH8lLX6go0Tv3Rgxz
5bKmn+ZMyL1GegadDiXrSYqd0/MUJuMgGWB8/OnP0D3Q4soEOBIn7DcPt0o9MUxZQsF0DraZzkR0
2WVRvcIFJucrAEJYAaWYJkjUVbmMb2ltwQwWO21rFHGbpE73nsfr/oAWsZEvKsQZoYm4fh5jVI5+
wKyRnKaN1uqAcNgj75cdywCHBVwgEefEgOPM77CDMH0+JumSirQiBfR35+HWRwHwpm09wI6Aqtvg
y5bzxvLDDRgrhX4LCPtUHGrUXNJHRKYiHQX6P6bIVuBrHV6VFpyS+5weu0w6kQKBgQDQo4QeLtO7
S3KH8UL3lX4lhH1K7/Q99uBHmvLXdiDkHjLbBbh0JfrHgHtnK9bvJ2GvVcwhI9fTiO1p1o5RM5jb
iVUSCS91sLcTPFv8X83sExBZnrvlSlb/va+4yW+Lzvr6ZiDlZYsVRNvNAHUTojHRCOH2P4eX1+ql
5P4FMdfvSQKBgQDMsQ4LBpxjD9KdDzJzw9a0xbL47QdCeZBqNUy6MvwLE0+KsF+prvoigNZCaTcJ
2FfoPxpE3/o0A/byCTuDkfddrd/hcAO0gd1R9CYJDXJfnIbZfheUmHW7ShbXyqhpqQKVjzH+jnLq
VjbGD6tz3dN+AwNgULD/vvwXM2TWpu9TRwKBgGkPPdMZD2NLzaNouKkFbR0lRxY6GEovi6Zi/w/C
GzPjhQZHLifGjC5zozBDohqRQR5SXNT/QInzdGGMOePn0HwT/nNzjqN71eRoy4UdFQtgWiZWyRTf
x0lGUjsBrBrBoh3+2WfKJywRnYDwTwQQ83boOyiNuxCaGD1rPwKMo8iJAoGAPIePE4uc615edbts
u/cJouNjjWDqaKnyHrYsPlOdXNkVCHonj9ICffmDYpgignLLbA5dAkkJgCA8Ak7gnoOnlrg4ID4z
mklc3UNJjBvB2qw65E35QyPijMPYBXAUZUppTTjPG+ub59ge0msH1Hegdv8FHJJABSDBA0tbYm5z
DzkCgYA9/0KtWKFMhF3v01L54AXF5b15RroBhZAfzI1U0wPO4J6Tz+1KqmtrwHTBPI36nzITIhlM
hcoTsMRMgnv0NHzxlcQQmAy3foFBFOyHXql3hPtWbEViB5jQs4cP5ts1oivVhrEtrrE51TG4V/Ef
fD1JKiHl7MECYEMyBz31PsRCuw==
-----END PRIVATE KEY-----
`)

	// ---------------------------------------------------------------------
	// This is minimal demo playbook used at the start of this example
	// ---------------------------------------------------------------------
	// 	playbookData := `{
	//   "type": "playbook",
	//   "spec_version": "2.0",
	//   "id": "playbook--a0777575-5c4c-4710-9f01-15776103837f",
	//   "name": "Playbook 1",
	//   "created_by": "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1",
	//   "created": "2022-05-18T11:31:31.319Z",
	//   "modified": "2022-05-18T11:31:31.319Z",
	//   "signatures": [
	//     {
	//       "type": "jss",
	//       "id": "jss--af4b4bf3-677a-411d-887a-1f6fa5090c05",
	//       "created_by": "identity--be59c641-b2d5-4930-94fc-6fd583524fc6",
	//       "created": "2023-02-02T14:31:31.319Z",
	//       "modified": "2023-02-02T14:31:31.319Z",
	//       "signee": "Existing Example Company",
	//       "valid_from": "2023-02-02T14:31:31.319Z",
	//       "valid_until": "2023-06-18T11:31:31.319Z",
	//       "related_to": "playbook--a0777575-5c4c-4710-9f01-15776103837f",
	//       "related_version": "2022-05-18T11:31:31.319Z",
	//       "hash_algorithm": "sha-256",
	//       "algorithm": "RS256",
	//       "public_key": "some public key",
	//       "value": "some signature"
	//     }
	//   ]
	// }`

	// ---------------------------------------------------------------------
	// Parse PEM encoded private key
	// ---------------------------------------------------------------------
	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		panic("failed to parse PEM block containing the private key")
	}
	parseResult, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded private key: " + err.Error())
	}
	privateKey := parseResult.(*rsa.PrivateKey)

	publicKeyDer, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	// ---------------------------------------------------------------------
	// Demo playbook object
	// ---------------------------------------------------------------------
	p := playbook.New()
	// Manually overwrite the values so I get consistent results each time
	p.SpecVersion = "2.0"
	p.ID = "playbook--a0777575-5c4c-4710-9f01-15776103837f"
	p.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	p.Created = "2022-05-18T11:31:31.319Z"
	p.Modified = p.Created
	p.Name = "Playbook 1"

	// Create dummy signature so we can show that you need to remove it
	var sigExisting signature.Signature
	sigExisting.ObjectType = "jss"
	sigExisting.ID = "jss--af4b4bf3-677a-411d-887a-1f6fa5090c05"
	sigExisting.CreatedBy = "identity--be59c641-b2d5-4930-94fc-6fd583524fc6"
	sigExisting.Created = "2023-02-02T14:31:31.319Z"
	sigExisting.Modified = sigExisting.Created
	sigExisting.Signee = "Existing Example Company"
	sigExisting.ValidFrom = sigExisting.Created
	sigExisting.ValidUntil = "2023-06-18T11:31:31.319Z"
	sigExisting.RelatedTo = p.ID
	sigExisting.RelatedVersion = p.Modified
	sigExisting.HashAlgorithm = "sha-256"
	sigExisting.Algorithm = "RS256"
	sigExisting.PublicKey = "some public key"
	sigExisting.Value = "some signature"
	p.Signatures = append(p.Signatures, sigExisting)

	// ---------------------------------------------------------------------
	// Demo new signature object used in step 3
	// ---------------------------------------------------------------------
	var s signature.Signature
	s.ObjectType = "jss"
	s.ID = "jss--af892292-c4b4-47eb-9be6-4897ff4b9388"
	s.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	s.Created = "2023-01-10T17:39:31.319Z"
	s.Modified = s.Created
	s.Signee = "ACME Cyber Company"
	s.ValidFrom = s.Created
	s.ValidUntil = "2023-06-10T17:39:31.319Z"
	s.RelatedTo = p.ID
	s.RelatedVersion = p.Modified
	s.HashAlgorithm = "sha-256"
	s.Algorithm = "RS256"
	s.PublicKey = base64.RawStdEncoding.EncodeToString(publicKeyDer)

	// ---------------------------------------------------------------------
	// Step 1
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 1: Create or parse a JSON playbook object to sign")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	// Convert playbook object to a JSON byte[]
	rawPlaybookData, _ := p.Encode()
	fmt.Println(string(rawPlaybookData))

	// ---------------------------------------------------------------------
	// Step 2
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 2: Temporarily remove any existing signature objects contained in the playbook's signatures property")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	savedSignatures := p.Signatures
	p.Signatures = nil
	jsonPlaybookData, _ := p.Encode()
	fmt.Println(string(jsonPlaybookData))

	// ---------------------------------------------------------------------
	// Step 3
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 3: Create and add signature to playbook")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	p.Signatures = append(p.Signatures, s)
	jsonWithSigData, _ := p.Encode()
	fmt.Println(string(jsonWithSigData))

	// ---------------------------------------------------------------------
	// Step 4
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 4: Create a JCS [RFC8785] canonical version of the playbook from step 3")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	jcsPlayBookData, _ := jcs.Transform(jsonWithSigData)
	fmt.Println(string(jcsPlayBookData))

	// ---------------------------------------------------------------------
	// Step 5
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 5: Create a hash (SHA256 in hex) of the canonical version of playbook from step 4")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	hashhexjcsPlaybookData := sha256.Sum256(jcsPlayBookData)
	// fmt.Println(fmt.Sprintf("%x", hashjcsPlaybookData[:]))
	hashjcsPlaybookData := hex.EncodeToString(hashhexjcsPlaybookData[:])
	fmt.Println(hashjcsPlaybookData)

	// // ---------------------------------------------------------------------
	// // Step 1.4
	// // ---------------------------------------------------------------------
	// fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	// fmt.Println("Step 1.4: Create base64URL.encoded version of the SHA256 hash from step 1.3 and remove any padding")
	// fmt.Println("------------------------------------------------------------------------------------------------------------")
	// // Remove base64 padding characters "=" per RFC 7515 section 2 - Base64url Encoding
	// b64hashjcsPlaybookData := base64.RawURLEncoding.EncodeToString(hashjcsPlaybookData[:])
	// fmt.Println(b64hashjcsPlaybookData)

	// ---------------------------------------------------------------------
	// Step 6
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 6: Sign the hash from step 5 using the algorithm defined in the signature object and base64URL.encode it (RS256)")
	fmt.Println("------------------------------------------------------------------------------------------------------------")

	method := jwt.SigningMethodRS256
	sigData, err := method.Sign(string(hashjcsPlaybookData), privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Signature: ", sigData)

	// ---------------------------------------------------------------------
	// Step 7
	// ---------------------------------------------------------------------
	fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	fmt.Println("Step 7: Add the new digital signature from step 6 to the signature value property (with existing signatures, if any)")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	s.Value = sigData
	// Clear out the signature for a second
	p.Signatures = nil
	// Add original signatures back to the playbook
	p.Signatures = append(p.Signatures, savedSignatures...)
	// Add the new signature
	p.Signatures = append(p.Signatures, s)
	finaldata, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(finaldata))

	// // ---------------------------------------------------------------------
	// // Step 2.2
	// // ---------------------------------------------------------------------
	// fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	// fmt.Println("Step 2.2: Create base64URL.encoded version of the JCS signature from step 2.1")
	// fmt.Println("------------------------------------------------------------------------------------------------------------")
	// b64jcsSigData := base64.RawURLEncoding.EncodeToString([]byte(jcsSigData))
	// fmt.Println(b64jcsSigData)

}
