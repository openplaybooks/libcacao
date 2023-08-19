// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package signature

import "encoding/json"

// Encode - This method is a simple wrapper for encoding an object into JSON
func (s *Signature) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return nil, err
	}

	// Any needed preprocessing would be done here
	return data, nil
}

// EncodeToString - This method is a simple wrapper for encoding an object into
// JSON
func (s *Signature) EncodeToString() (string, error) {
	data, err := s.Encode()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
