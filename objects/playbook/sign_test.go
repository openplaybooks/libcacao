// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log"
	"testing"

	"github.com/openplaybooks/libcacao/objects/signature"
)

// TestSign - This will test the Sign() method
func TestSign(t *testing.T) {

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
	p := New()
	// Manually overwrite the values so I get consistent results each time
	p.SpecVersion = "1.1"
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

	err = p.Sign("RS256", privateKey, &s)
	if err != nil {
		panic(err)
	}

	if p.Signatures[0].Value != "some signature" {
		t.Errorf("1.1 signature was not produced correctly")
	}

	correctSigValue := "QrVp0g3kJ8x6OssFnRCu8nkkK9l8gjhQZhbbvRRufAgyeHXJQymjlaZZ01lUnMIcRd22gYMPhsJ3EpsjAsFVj8DjO3BcNKzVZ_i2w4fH9O3hXKfAOSr0rX0eFlHdAPmdfCmNAOWMwubLP_J3k9duwrxWZf6EH3pn1bfi2nU6AGpfn3Ur8dn8G5qh4Hsso5FPmgf7LE8pcEHYU2IClkvRSu6fQfmZJp52jNU3uSZLR4K2PMHzzDAtGtzijzYLxm5MiHNwb-26_Rhb8Zr31aUtisY3gnYpRadsR6KjseMOtlOzJutbXpzKQF-0lMzSz79Q81lniQFmLn1xh1fuYFUskA"
	if p.Signatures[1].Value != correctSigValue {
		t.Errorf("1.2 signature was not produced correctly")
		t.Errorf("Expected: %s", correctSigValue)
		t.Errorf("Have: %s", p.Signatures[1].Value)
	}

	if p.Signatures[1].PublicKey != "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAptKZyFPStvmOlb0WihOBhlHUr6wFDHC+tW7hJAudfTQ5mHZQpB8PoMz07udZA+dG8dhUIPkmXlp1TgREeYTHdhxhuf0y/GhbpZv5JPYHx3watO+HWO2qYkjRMEcrWhPMdaVkS/Xe/liaMcow4jYoWaFm8VobeYsyVD2bWWdyl4joTEETm1Z47RnnfR15kVhVudVrDzEFmM4nXV/6dmIg184RJE4httwBFxR8qZCQCwTiJmsoyJxfUR0Gs4ePKc5sB0NTkmFZc5klQSitd67RJn2ldhbqE7EpDl4XlIt+UyLJm1guCBltia8Agke7dXuhpB7hQ6LJwY4EjzthkJ8IPwIDAQAB" {
		t.Errorf("1.3 public keys were not added to the signature correctly")
	}
}
