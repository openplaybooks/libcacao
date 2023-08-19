// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package signature

import (
	"errors"
	"time"

	"github.com/openplaybooks/libcacao/objects"
)

// ----------------------------------------------------------------------
// Public Signature Type Methods
// ----------------------------------------------------------------------

// SetNewID - This method takes in a string value representing an object type
// and creates a new ID based on the specification format and updates the id
// property for the object.
func (s *Signature) SetNewID(objType string) error {

	if objType != "jss" {
		return errors.New("the object type is not valid for a CACAO signature id")
	}

	s.ID, _ = objects.CreateID(objType)
	return nil
}

// // GetCurrentSpecVersion - This method returns the current specification version
// // that this library is using.
// func (s *Signature) GetCurrentSpecVersion() string {
// 	return objects.GetCurrentSpecVersion()
// }

// GetCurrentTime - This method takes in a value of either milli or micro and
// returns the current time in RFC 3339 format
func (s *Signature) GetCurrentTime(precision string) string {
	return objects.GetCurrentTime(precision)
}

// SetCreated - This method takes in a timestamp in either time.Time or string
// format and updates the created property with it. The value is stored as a
// string, so if the value is in time.Time format, it will be converted to the
// correct timestamp format.
func (s *Signature) SetCreated(t interface{}) error {
	ts, _ := objects.TimeToString(t, "milli")

	s.Created = ts
	return nil
}

// SetModified - This method takes in a timestamp in either time.Time or string
// format and updates the modified property with it. The value is stored as a
// string, so if the value is in time.Time format, it will be converted to the
// correct timestamp format.
func (s *Signature) SetModified(t interface{}) error {
	ts, _ := objects.TimeToString(t, "milli")

	// Make sure the modified timestampe is equal to or greater than created
	if s.Created == "" {
		return errors.New("the created property is null, but must be populated")
	}

	created, _ := time.Parse(time.RFC3339, s.Created)
	modified, _ := time.Parse(time.RFC3339, ts)
	if modified.Before(created) {
		return errors.New("the modified timestamp is invalid, it is before the created timestamp")
	}

	s.Modified = ts
	return nil
}
