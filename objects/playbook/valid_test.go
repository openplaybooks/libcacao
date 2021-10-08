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
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("1.1 checkObjectType returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check invalid value
	setup(r)
	p.ObjectType = "foo"
	p.checkObjectType(r)
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("1.2 checkObjectType returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check correct value
	setup(r)
	p.ObjectType = "playbook"
	p.checkObjectType(r)
	if r.problemsFound != 0 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("1.3 checkObjectType returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
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
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("2.1 checkSpecVersion returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check invalid value
	setup(r)
	p.SpecVersion = "0.9"
	p.checkSpecVersion(r)
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("2.2 checkSpecVersion returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check correct value
	setup(r)
	p.SpecVersion = "1.0"
	p.checkSpecVersion(r)
	if r.problemsFound != 0 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("2.3 checkSpecVersion returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// TestCheckID - This will check the id
func TestCheckID(t *testing.T) {
	p := new(Playbook)
	r := new(results)
	r.debug = true

	// Check when property is missing
	setup(r)
	p.checkID(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("3.1 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check invalid value
	setup(r)
	p.ID = "playbook--uuid1"
	p.checkID(r)
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("3.2 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check invalid value
	setup(r)
	p.ID = "foo--60cfe320-f6b4-4523-8558-14a042223797"
	p.checkID(r)
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("3.3 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check correct value
	setup(r)
	p.ID = "playbook--60cfe320-f6b4-4523-8558-14a042223797"
	p.checkID(r)
	if r.problemsFound != 0 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("3.4 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// TestCheckName - This will check the id
func TestCheckName(t *testing.T) {
	p := new(Playbook)
	r := new(results)
	r.debug = true

	// Check when property is missing
	setup(r)
	p.checkName(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("4.1 checkName returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}
