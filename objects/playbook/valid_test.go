// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

import (
	"testing"
)

func setup(r *results) {
	r.problemsFound = 0
	r.resultDetails = nil
}

// TestCheckObjectType - This will test the object type
func TestCheckObjectType(t *testing.T) {
	p := new(Playbook)
	r := new(results)
	r.debug = true

	// Check when property is missing
	setup(r)
	p.checkObjectType(r)
	if r.problemsFound != 1 && r.resultDetails[0][0:2] != "--" {
		t.Errorf("1.1 checkObjectType returned incorrect values when object missing")
	}

	// Check invalid value
	setup(r)
	p.checkObjectType(r)
	p.ObjectType = "foo"
	if r.problemsFound != 1 && r.resultDetails[0][0:2] != "--" {
		t.Errorf("1.2 checkObjectType returned incorrect values when value is invalid")
	}

	// Check correct value
	setup(r)
	p.checkObjectType(r)
	p.ObjectType = "playbook"
	if r.problemsFound != 0 && r.resultDetails[0][0:2] != "++" {
		t.Errorf("1.3 checkObjectType returned incorrect values when value is valid")
	}
}

// TestCheckSpecVersion - This will check the spec version
func TestCheckSpecVersion(t *testing.T) {
	p := new(Playbook)
	r := new(results)
	r.debug = true

	// Check when property is missing
	setup(r)
	p.checkSpecVersion(r)
	if r.problemsFound != 1 && r.resultDetails[0][0:2] != "--" {
		t.Errorf("2.1 checkSpecVersion returned incorrect values when object missing")
	}

	// Check invalid value
	setup(r)
	p.checkSpecVersion(r)
	p.SpecVersion = "0.9"
	if r.problemsFound != 1 && r.resultDetails[0][0:2] != "--" {
		t.Errorf("2.2 checkSpecVersion returned incorrect values when value is invalid")
	}

	// Check correct value
	setup(r)
	p.checkSpecVersion(r)
	p.SpecVersion = "1.0"
	if r.problemsFound != 0 && r.resultDetails[0][0:2] != "++" {
		t.Errorf("2.3 checkSpecVersion returned incorrect values when value is valid")
	}
}
