// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package signature

import (
	"errors"

	"github.com/openplaybooks/libcacao/objects"
)

// ----------------------------------------------------------------------
// Signature Type
// ----------------------------------------------------------------------

// SetNewID - This method takes in a string value representing an object type and
// creates a new ID based on the specification format and update the id property
// for the object.
func (s *Signature) SetNewID(objType string) error {

	if valid := objects.IsTypeValid(objType); valid == false {
		return errors.New("the object type is not valid for a CACAO id")
	}

	s.ID, _ = objects.CreateID(objType)
	return nil
}

// GetCurrentSpecVersion - This method returns the current specification version
// that this library is using.
func (s *Signature) GetCurrentSpecVersion() string {
	return objects.GetCurrentSpecVersion()
}

// GetCurrentTime - This method takes in a value of either milli or micro and
// returns the current time in RFC 3339 format
func (s *Signature) GetCurrentTime(precision string) string {
	return objects.GetCurrentTime(precision)
}

// SetModified - This method takes in a timestamp in either time.Time or string
// format and updates the modified property with it. The value is stored as a
// string, so if the value is in time.Time format, it will be converted to the
// correct timestamp format.
func (s *Signature) SetModified(t interface{}) error {
	ts, _ := objects.TimeToString(t, "milli")
	s.Modified = ts
	return nil
}
