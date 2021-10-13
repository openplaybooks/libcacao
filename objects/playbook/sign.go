// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/gowebpki/jcs"
	"github.com/openplaybooks/libcacao/objects/signature"
)

// Sign - This method will sign a playbook object
func (p *Playbook) Sign(method string, key *rsa.PrivateKey, sig *signature.Signature) error {
	var err error

	// Save any existing signatures and then zero out the property for signing
	savedSignatures := p.Signatures
	p.Signatures = nil

	// Convert playbook object to a JSON byte[]
	playbookJSON, err := p.Encode()
	if err != nil {
		return err
	}

	// Create JCS version of playbook
	jcsPlaybook, err := jcs.Transform(playbookJSON)
	if err != nil {
		return err
	}

	// SHA256 encode JCS version. The Sum256 functions returns a [32]byte
	hashPlaybook := sha256.Sum256(jcsPlaybook)

	// Create base64URL.encoded version and remove base64 padding characters "="
	// per RFC 7515 section 2 - Base64url Encoding
	// Save the b64 hash to the signature object
	b64Playbook := base64.RawURLEncoding.EncodeToString(hashPlaybook[:])
	sig.SHA256 = b64Playbook

	// ------------------------------------------------------------
	// Digitally sign the signature object
	// ------------------------------------------------------------
	sigJSON, err := sig.Encode()
	if err != nil {
		return err
	}

	// Create JCS version of signature object
	jcsSig, err := jcs.Transform(sigJSON)
	if err != nil {
		return err
	}

	b64Sig := base64.RawURLEncoding.EncodeToString([]byte(jcsSig))

	var signingMethod jwt.SigningMethod
	if method == "RS256" {
		// This will set the signingMethod to *jwt.SigningMethodRSA
		signingMethod = jwt.SigningMethodRS256
	} else {
		return errors.New("no valid signing method was given")
	}

	sigData, err := signingMethod.Sign(b64Sig, key)
	if err != nil {
		panic(err)
	}

	// Add signature to signature object
	sig.Value = sigData
	// Add original signatures back to the playbook
	p.Signatures = append(p.Signatures, savedSignatures...)
	// Add the new signature
	p.Signatures = append(p.Signatures, *sig)

	return nil
}
