// Copyright 2023 Bret Jordan, All rights reserved.
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
	r.debug = true
}

// TestCheckObjectType - This will test the object type property
func TestCheckObjectType(t *testing.T) {
	p := new(Playbook)
	r := new(results)

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
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("1.2 checkObjectType returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("1.3 checkObjectType returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check correct value
	setup(r)
	p.ObjectType = "playbook"
	p.checkObjectType(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("1.4 checkObjectType returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 0 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("1.5 checkObjectType returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// TestCheckSpecVersion - This will check the spec_version property
func TestCheckSpecVersion(t *testing.T) {
	p := new(Playbook)
	r := new(results)

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
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("2.2 checkSpecVersion returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("2.3 checkSpecVersion returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check correct value
	setup(r)
	p.SpecVersion = "1.0"
	p.checkSpecVersion(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("2.4 checkSpecVersion returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 0 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("2.5 checkSpecVersion returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// TestCheckID - This will check the id property
func TestCheckID(t *testing.T) {
	p := new(Playbook)
	r := new(results)

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
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("3.2 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("3.3 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check invalid value
	setup(r)
	p.ID = "foo--60cfe320-f6b4-4523-8558-14a042223797"
	p.checkID(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("3.4 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("3.5 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check correct value
	setup(r)
	p.ID = "playbook--60cfe320-f6b4-4523-8558-14a042223797"
	p.checkID(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("3.6 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 0 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("3.7 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// TestCheckName - This will check the name property
func TestCheckName(t *testing.T) {
	p := new(Playbook)
	r := new(results)

	// Check when property is missing
	setup(r)
	p.checkName(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("4.1 checkName returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// Nothing to do for Description (property #5)

// TestCheckPlaybookTypes - This will check the playbook_types property
func TestCheckPlaybookTypes(t *testing.T) {
	p := new(Playbook)
	r := new(results)

	// Check when property is missing
	setup(r)
	p.checkPlaybookTypes(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("6.1 checkPlaybookTypes returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check invalid value
	setup(r)
	p.PlaybookTypes = nil
	p.PlaybookTypes = append(p.PlaybookTypes, "test")
	p.checkPlaybookTypes(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("6.2 checkPlaybookTypes returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("6.3 checkPlaybookTypes returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check correct value
	setup(r)
	p.PlaybookTypes = nil
	p.PlaybookTypes = append(p.PlaybookTypes, "notification")
	p.checkPlaybookTypes(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("6.4 checkPlaybookTypes returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 0 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("6.5 checkPlaybookTypes returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// TestCheckCreatedBy - This will check the created_by property
func TestCheckCreatedBy(t *testing.T) {
	p := new(Playbook)
	r := new(results)

	// Check when property is missing
	setup(r)
	p.checkCreatedBy(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("7.1 checkCreatedBy returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check invalid value
	setup(r)
	p.CreatedBy = "identity--uuid1"
	p.checkCreatedBy(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("7.2 checkCreatedBy returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("7.3 checkCreatedBy returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	// No other checks will be made since it will fail the first part

	// Check invalid value
	setup(r)
	p.CreatedBy = "foo--60cfe320-f6b4-4523-8558-14a042223797"
	p.checkCreatedBy(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("7.4 checkCreatedBy returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("7.5 checkCreatedBy returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	// No other checks will be made since it will fail the first part

	// Check invalid value
	setup(r)
	p.CreatedBy = "step--60cfe320-f6b4-4523-8558-14a042223797"
	p.checkCreatedBy(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("7.6 checkCreatedBy returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("7.7 checkCreatedBy returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check valid value
	setup(r)
	p.CreatedBy = "identity--60cfe320-f6b4-4523-8558-14a042223797"
	p.checkCreatedBy(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("7.8 checkCreatedBy returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 0 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("7.9 checkCreatedBy returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// TestCheckCreated - This will check the created property
func TestCheckCreated(t *testing.T) {
	p := new(Playbook)
	r := new(results)

	// Check when property is missing
	setup(r)
	p.checkCreated(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("8.1 checkCreated returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check invalid value
	setup(r)
	p.Created = "Some data 2021"
	p.checkCreated(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("8.2 checkCreated returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("8.3 checkCreated returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check correct value
	setup(r)
	p.Created = "2021-02-02T12:12:12.123Z"
	p.checkCreated(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("8.4 checkCreated returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 0 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("8.5 checkCreated returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// TestCheckModified - This will check the modified property
func TestCheckModified(t *testing.T) {
	p := new(Playbook)
	r := new(results)

	// Check when property is missing
	setup(r)
	p.checkModified(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("9.1 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check invalid value
	setup(r)
	p.Modified = "Some data 2021"
	p.checkModified(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("9.2 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("9.3 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check correct value
	setup(r)
	p.Modified = "2021-02-02T12:12:12.123Z"
	p.checkModified(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("9.4 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 0 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("9.5 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check if before Created value
	setup(r)
	p.Created = "2021-02-03T12:12:12.123Z"
	p.Modified = "2021-02-02T12:12:12.123Z"
	p.checkModified(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("9.6 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("9.7 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[2][0:2] != "--" {
		t.Errorf("9.8 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check if same as Created value
	setup(r)
	p.Created = "2021-02-02T12:12:12.123Z"
	p.Modified = "2021-02-02T12:12:12.123Z"
	p.checkModified(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("9.9 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 0 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("9.10 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 0 || r.resultDetails[2][0:2] != "++" {
		t.Errorf("9.11 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check if after Created value
	setup(r)
	p.Created = "2021-02-01T12:12:12.123Z"
	p.Modified = "2021-02-02T12:12:12.123Z"
	p.checkModified(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("9.12 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 0 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("9.13 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 0 || r.resultDetails[2][0:2] != "++" {
		t.Errorf("9.14 checkModified returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// Nothing to do for Revoked (property #10)

// TestCheckValidFrom - This will check the valid_from property
func TestCheckValidFrom(t *testing.T) {
	p := new(Playbook)
	r := new(results)

	// Check invalid value
	setup(r)
	p.ValidFrom = "Some data 2021"
	p.checkValidFrom(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("11.1 checkValidFrom returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check correct value
	setup(r)
	p.ValidFrom = "2021-02-02T12:12:12.123Z"
	p.checkValidFrom(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("11.2 checkValidFrom returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// TestCheckValidUntil - This will check the valid_from property
func TestCheckValidUntil(t *testing.T) {
	p := new(Playbook)
	r := new(results)

	// Check invalid value
	setup(r)
	p.ValidUntil = "Some data 2021"
	p.checkValidUntil(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("12.1 checkValidUntil returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check correct value
	setup(r)
	p.ValidUntil = "2021-02-02T12:12:12.123Z"
	p.checkValidUntil(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("12.2 checkValidUntil returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check if before Valid From value
	setup(r)
	p.ValidFrom = "2021-02-03T12:12:12.123Z"
	p.ValidUntil = "2021-02-02T12:12:12.123Z"
	p.checkValidUntil(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("12.3 checkValidUntil returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("12.4 checkValidUntil returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check if same as Valid From value
	setup(r)
	p.ValidFrom = "2021-02-02T12:12:12.123Z"
	p.ValidUntil = "2021-02-02T12:12:12.123Z"
	p.checkValidUntil(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("12.5 checkValidUntil returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 1 || r.resultDetails[1][0:2] != "--" {
		t.Errorf("12.6 checkValidUntil returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check if after Valid From value
	setup(r)
	p.ValidFrom = "2021-02-01T12:12:12.123Z"
	p.ValidUntil = "2021-02-02T12:12:12.123Z"
	p.checkValidUntil(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("12.7 checkValidUntil returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	if r.problemsFound != 0 || r.resultDetails[1][0:2] != "++" {
		t.Errorf("12.8 checkValidUntil returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

}

// TestCheckDerivedFrom - This will check the id
func TestCheckDerivedFrom(t *testing.T) {
	p := new(Playbook)
	r := new(results)

	// Check invalid value
	setup(r)
	p.DerivedFrom = nil
	p.DerivedFrom = append(p.DerivedFrom, "playbook--uuid1")
	p.checkDerivedFrom(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("13.1 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	// No other checks will be made since it will fail the first part

	// Check invalid value
	setup(r)
	p.DerivedFrom = nil
	p.DerivedFrom = append(p.DerivedFrom, "foo--60cfe320-f6b4-4523-8558-14a042223797")
	p.checkDerivedFrom(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("13.2 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
	// No other checks will be made since it will fail the first part

	// Check invalid value
	setup(r)
	p.DerivedFrom = nil
	p.DerivedFrom = append(p.DerivedFrom, "step--60cfe320-f6b4-4523-8558-14a042223797")
	p.checkDerivedFrom(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("13.4 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check correct value
	setup(r)
	p.DerivedFrom = nil
	p.DerivedFrom = append(p.DerivedFrom, "playbook--60cfe320-f6b4-4523-8558-14a042223797")
	p.checkDerivedFrom(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("13.5 checkID returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// TestCheckPriority - This will check the priority property
func TestCheckPriority(t *testing.T) {
	p := new(Playbook)
	r := new(results)

	// Check when property is invalid
	setup(r)
	p.Priority = -1
	p.checkPriority(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("14.1 checkPriority returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check when property is invalid
	setup(r)
	p.Priority = 105
	p.checkPriority(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("14.2 checkPriority returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check when property is valid
	setup(r)
	p.Priority = 0
	p.checkPriority(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("14.3 checkPriority returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check when property is valid
	setup(r)
	p.Priority = 10
	p.checkPriority(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("14.4 checkPriority returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check when property is valid
	setup(r)
	p.Priority = 100
	p.checkPriority(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("14.5 checkPriority returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// TestCheckSeverity - This will check the priority property
func TestCheckSeverity(t *testing.T) {
	p := new(Playbook)
	r := new(results)

	// Check when property is invalid
	setup(r)
	p.Severity = -1
	p.checkSeverity(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("15.1 checkSeverity returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check when property is invalid
	setup(r)
	p.Severity = 105
	p.checkSeverity(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("15.2 checkSeverity returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check when property is valid
	setup(r)
	p.Severity = 0
	p.checkSeverity(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("15.3 checkSeverity returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check when property is valid
	setup(r)
	p.Severity = 10
	p.checkSeverity(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("15.4 checkSeverity returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check when property is valid
	setup(r)
	p.Severity = 100
	p.checkSeverity(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("15.5 checkSeverity returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// TestCheckImpact - This will check the priority property
func TestCheckImpact(t *testing.T) {
	p := new(Playbook)
	r := new(results)

	// Check when property is invalid
	setup(r)
	p.Impact = -1
	p.checkImpact(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("16.1 checkImpact returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check when property is invalid
	setup(r)
	p.Impact = 105
	p.checkImpact(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("16.2 checkImpact returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check when property is valid
	setup(r)
	p.Impact = 0
	p.checkImpact(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("16.3 checkImpact returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check when property is valid
	setup(r)
	p.Impact = 10
	p.checkImpact(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("16.4 checkImpact returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check when property is valid
	setup(r)
	p.Impact = 100
	p.checkImpact(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("16.5 checkImpact returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}

// TestCheckIndustrySectors - This will check the industry_sectors property
func TestCheckIndustrySectors(t *testing.T) {
	p := new(Playbook)
	r := new(results)

	// Check invalid value
	setup(r)
	p.IndustrySectors = nil
	p.IndustrySectors = append(p.IndustrySectors, "test")
	p.checkIndustrySectors(r)
	if r.problemsFound != 1 || r.resultDetails[0][0:2] != "--" {
		t.Errorf("17.2 checkIndustrySectors returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}

	// Check correct value
	setup(r)
	p.IndustrySectors = nil
	p.IndustrySectors = append(p.IndustrySectors, "energy")
	p.checkIndustrySectors(r)
	if r.problemsFound != 0 || r.resultDetails[0][0:2] != "++" {
		t.Errorf("17.4 checkIndustrySectors returned errors %d and results %s which is invalid", r.problemsFound, r.resultDetails)
	}
}
