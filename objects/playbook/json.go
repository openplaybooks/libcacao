// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

import "encoding/json"

// Decode - This function is a simple wrapper for decoding JSON data. It will
// decode a slice of bytes into an actual struct and return a pointer to that
// object along with any errors.
func Decode(data []byte) (*Playbook, error) {
	var p Playbook

	if err := json.Unmarshal(data, &p); err != nil {
		return nil, err
	}

	return &p, nil
}

// // UnmarshalJSON - This method will over write the default UnmarshalJSON method
// // to enable custom processing that we may need to do.
// func (p *Playbook) UnmarshalJSON(b []byte) error {

// 	type alias Playbook
// 	temp := &struct {
// 		*alias
// 	}{
// 		alias: (*alias)(p),
// 	}
// 	if err := json.Unmarshal(b, &temp); err != nil {
// 		return err
// 	}

// 	return nil
// }

// Encode - This method is a simple wrapper for encoding an object into JSON
func (p *Playbook) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return nil, err
	}

	// Any needed preprocessing would be done here
	return data, nil
}

// EncodeToString - This method is a simple wrapper for encoding an object into
// JSON
func (p *Playbook) EncodeToString() (string, error) {
	data, err := p.Encode()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
