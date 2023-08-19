// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/gowebpki/jcs"
	"github.com/openplaybooks/libcacao/objects"
	"github.com/openplaybooks/libcacao/objects/signature"
)

// Sign - This method will sign a playbook object.
// It takes in a signing method like "RS256", a key, and a CACAO signature object.
func (p *Playbook) Sign(method string, key interface{}, sig *signature.Signature) error {
	var err error

	if !objects.IsVocabValueValid(method, objects.GetSigningMethodsVocab()) {
		return fmt.Errorf("the signing method %s is not valid", method)
	}

	// Step 2: Save any existing signatures and then zero out the property for signing
	savedSignatures := p.Signatures
	p.Signatures = nil

	// Step 3: Add new signature object
	p.Signatures = append(p.Signatures, *sig)
	// Convert playbook object to a JSON byte[] so we can run it through JCS
	pbData, err := p.Encode()
	if err != nil {
		return err
	}

	// Step 4: Create JCS version of playbook
	jcsData, err := jcs.Transform(pbData)
	if err != nil {
		return err
	}

	// Step 5: SHA256 encode JCS version. The Sum256 functions returns a [32]byte
	hashhex := sha256.Sum256(jcsData)
	hash := hex.EncodeToString(hashhex[:])

	// Step 6: Digitally sign the hash
	var signingMethod jwt.SigningMethod
	// If value is an RSA method the signingMethod will be *jwt.SigningMethodRSA
	if method == "RS256" {
		signingMethod = jwt.SigningMethodRS256
	} else if method == "RS384" {
		signingMethod = jwt.SigningMethodRS384
	} else if method == "RS512" {
		signingMethod = jwt.SigningMethodRS512
	} else if method == "ES256" {
		signingMethod = jwt.SigningMethodES256
	} else if method == "ES384" {
		signingMethod = jwt.SigningMethodES384
	} else if method == "ES512" {
		signingMethod = jwt.SigningMethodES512
	} else {
		return errors.New("no valid signing method was given")
	}
	sigData, err := signingMethod.Sign(hash, key)
	if err != nil {
		panic(err)
	}

	// Step 8
	// Add signature to signature object
	sig.Value = sigData
	// Add original signatures back to the playbook
	p.Signatures = nil
	p.Signatures = append(p.Signatures, savedSignatures...)
	// Add the new signature
	p.Signatures = append(p.Signatures, *sig)

	return nil
}
