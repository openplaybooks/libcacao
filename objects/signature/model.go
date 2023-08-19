// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package signature

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// Signature - This type defines the digital signature for a playbook object
type Signature struct {
	ObjectType      string     `json:"type,omitempty"`
	ID              string     `json:"id,omitempty"`
	CreatedBy       string     `json:"created_by,omitempty"`
	Created         string     `json:"created,omitempty"`
	Modified        string     `json:"modified,omitempty"`
	Revoked         bool       `json:"revoked,omitempty"`
	Signee          string     `json:"signee,omitempty"`
	ValidFrom       string     `json:"valid_from,omitempty"`
	ValidUntil      string     `json:"valid_until,omitempty"`
	RelatedTo       string     `json:"related_to,omitempty"`
	RelatedVersion  string     `json:"related_version,omitempty"`
	HashAlgorithm   string     `json:"hash_algorithm,omitempty"`
	Algorithm       string     `json:"algorithm,omitempty"`
	PublicKey       string     `json:"public_key,omitempty"`
	PublicCertChain []string   `json:"public_cert_chain,omitempty"`
	CertURL         string     `json:"cert_url,omitempty"`
	Thumbprint      string     `json:"thumbprint,omitempty"`
	Value           string     `json:"value,omitempty"`
	Signature       *Signature `json:"signature,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new signature object and return it as a
// pointer. It will also initialize the object by setting all of the basic
// properties.
func New() *Signature {
	s := new(Signature)
	s.ObjectType = "jss"
	s.SetNewID(s.ObjectType)
	s.Created = s.GetCurrentTime("milli")
	s.Modified = s.Created
	s.Revoked = false
	return s
}
