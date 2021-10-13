// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package signature

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// Signature - This type defines the digital signature for a playbook object
type Signature struct {
	ObjectType     string   `json:"type,omitempty"`
	SpecVersion    string   `json:"spec_version,omitempty"`
	ID             string   `json:"id,omitempty"`
	CreatedBy      string   `json:"created_by,omitempty"`
	Created        string   `json:"created,omitempty"`
	Modified       string   `json:"modified,omitempty"`
	Revoked        bool     `json:"revoked,omitempty"`
	Signee         string   `json:"signee,omitempty"`
	ValidFrom      string   `json:"valid_from,omitempty"`
	ValidUntil     string   `json:"valid_until,omitempty"`
	RelatedTo      string   `json:"related_to,omitempty"`
	RelatedVersion string   `json:"related_version,omitempty"`
	SHA256         string   `json:"sha256,omitempty"`
	Algorithm      string   `json:"algorithm,omitempty"`
	PublicKeys     []string `json:"public_keys,omitempty"`
	CertURL        string   `json:"cert_url,omitempty"`
	Thumbprint     string   `json:"thumbprint,omitempty"`
	Value          string   `json:"value,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new signature object and return it as a
// pointer. It will also initialize the object by setting all of the basic
// properties.
func New() *Signature {
	s := new(Signature)
	s.Init()
	return s
}

// Init - This method will initialize a new signature object with the correct
// defaults.
func (s *Signature) Init() {
	s.ObjectType = "signature"
	s.SpecVersion = s.GetCurrentSpecVersion()
	s.SetNewID(s.ObjectType)
	s.Created = s.GetCurrentTime("milli")
	s.Modified = s.Created
	s.Revoked = false
}
