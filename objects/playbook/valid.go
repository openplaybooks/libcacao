// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

import (
	"fmt"
	"time"

	"github.com/openplaybooks/libcacao/objects"
)

type results struct {
	problemsFound int
	resultDetails []string
}

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// Valid - This method will verify that the object is correct. It will return a
// boolean, an integer that tracks the number of problems found, and a slice of
// strings that contain the detailed results, whether good or bad.
func (p *Playbook) Valid() (bool, int, []string) {
	var r *results = new(results)

	// Check each property in the model
	p.checkObjectType(r)
	p.checkSpecVersion(r)
	p.checkID(r)
	p.checkName(r)
	// Description - No requirements
	p.checkPlaybookTypes(r)
	p.checkCreatedBy(r)
	p.checkCreated(r)
	p.checkModified(r)
	// Revoked - No requirements
	p.checkValidFrom(r)
	p.checkValidUntil(r)
	// Derived From - No requirements
	p.checkPriority(r)
	p.checkSeverity(r)
	p.checkImpact(r)
	// Labels - No requirements
	p.checkExternalReferences(r)
	// Features - No requirements

	// Finished Checks

	// Return real values not pointers
	if r.problemsFound > 0 {
		return false, r.problemsFound, r.resultDetails
	}
	return true, r.problemsFound, r.resultDetails
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

// Each of these methods will check a specific property. It is done this way
// to reduce the complexity of the main valid() function. This way all of the
// checks for each property are self contained in their own function.

func (p *Playbook) checkObjectType(r *results) {
	if p.ObjectType == "" {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the type property is required but missing")
	} else {
		r.resultDetails = append(r.resultDetails, "++ the type property is required and is present")
		if p.ObjectType != "playbook" && p.ObjectType != "playbook-template" {
			r.problemsFound++
			r.resultDetails = append(r.resultDetails, "-- the type property does not contain a value of playbook or playbook-template")
		} else {
			r.resultDetails = append(r.resultDetails, "++ the type property does contain a value of playbook or playbook-template")
		}
	}
}

func (p *Playbook) checkSpecVersion(r *results) {
	if p.SpecVersion == "" {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the spec_version property is required but missing")
	} else {
		r.resultDetails = append(r.resultDetails, "++ the spec_version property is required and is present")
	}
}

func (p *Playbook) checkID(r *results) {
	if p.ID == "" {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the id property is required but missing")
	} else {
		r.resultDetails = append(r.resultDetails, "++ the id property is required and is present")
		if valid := objects.IsIDValid(p.ID); valid == false {
			r.problemsFound++
			r.resultDetails = append(r.resultDetails, "-- the id property is not a valid timestamp")
		} else {
			r.resultDetails = append(r.resultDetails, "++ the id property is a valid timestamp")
		}
	}
}

func (p *Playbook) checkName(r *results) {
	if p.Name == "" {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the name property is required but missing")
	} else {
		r.resultDetails = append(r.resultDetails, "++ the name property is required and is present")
	}
}

func (p *Playbook) checkPlaybookTypes(r *results) {
	if len(p.PlaybookTypes) == 0 {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the playbook_types property is required but missing")
	} else {
		r.resultDetails = append(r.resultDetails, "++ the playbook_types property is required and is present")

		ptvocab := p.GetPlaybookTypesVocab()
		for i := 0; i < len(p.PlaybookTypes); i++ {
			value := p.PlaybookTypes[i]
			if _, found := ptvocab[value]; found {
				str := fmt.Sprintf("++ the playbook_types property contains a value of \"%s\" that is in the vocabulary", value)
				r.resultDetails = append(r.resultDetails, str)
			} else {
				r.problemsFound++
				str := fmt.Sprintf("-- the playbook_types property contains a value of \"%s\" that is not in the vocabulary", value)
				r.resultDetails = append(r.resultDetails, str)
			}
		}
	}
}

func (p *Playbook) checkCreatedBy(r *results) {
	if p.CreatedBy == "" {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the created_by property is required but missing")
	} else {
		r.resultDetails = append(r.resultDetails, "++ the created_by property is required and is present")
	}
}

func (p *Playbook) checkCreated(r *results) {
	if p.Created == "" {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the created property is required but missing")
	} else {
		r.resultDetails = append(r.resultDetails, "++ the created property is required and is present")
		if valid := objects.IsTimestampValid(p.Created); valid == false {
			r.problemsFound++
			r.resultDetails = append(r.resultDetails, "-- the created property does not contain a valid timestamp")
		} else {
			r.resultDetails = append(r.resultDetails, "++ the created property contains a valid timestamp")
		}
	}
}

func (p *Playbook) checkModified(r *results) {
	if p.Modified == "" {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the modified property is required but missing")
	} else {
		r.resultDetails = append(r.resultDetails, "++ the modified property is required and is present")
		if valid := objects.IsTimestampValid(p.Modified); valid == false {
			r.problemsFound++
			r.resultDetails = append(r.resultDetails, "-- the modified property does not contain a valid timestamp")
		} else {
			r.resultDetails = append(r.resultDetails, "++ the modified property contains a valid timestamp")
		}
	}
}

func (p *Playbook) checkValidFrom(r *results) {
	if p.ValidFrom != "" {
		if valid := objects.IsTimestampValid(p.ValidFrom); valid == false {
			r.problemsFound++
			r.resultDetails = append(r.resultDetails, "-- the valid_from property does not contain a valid timestamp")
		} else {
			r.resultDetails = append(r.resultDetails, "++ the valid_from property contains a valid timestamp")
		}
	}
}

func (p *Playbook) checkValidUntil(r *results) {
	if p.ValidUntil != "" {
		if valid := objects.IsTimestampValid(p.ValidUntil); valid == false {
			r.problemsFound++
			r.resultDetails = append(r.resultDetails, "-- the valid_until property does not contain a valid timestamp")
		} else {
			r.resultDetails = append(r.resultDetails, "++ the valid_until property contains a valid timestamp")
		}

		// If there is a valid_until timestamp, then lets check to see if there is also a valid_from and if so
		// is the valid_until later than the valid_from
		if p.ValidFrom != "" {
			validFrom, _ := time.Parse(time.RFC3339, p.ValidFrom)
			validUntil, _ := time.Parse(time.RFC3339, p.ValidUntil)
			if yes := validUntil.After(validFrom); yes != true {
				r.problemsFound++
				r.resultDetails = append(r.resultDetails, "-- the valid_until timestamp is not later than the valid_from timestamp")
			} else {
				r.resultDetails = append(r.resultDetails, "++ the valid_until timestamp is later than the valid_from timestamp")
			}
		}
	}
}

func (p *Playbook) checkPriority(r *results) {
	if p.Priority < 0 {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the priority property does not contain a valid value, it is less than zero")
	} else if p.Priority > 100 {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the priority property does not contain a valid value, it is greater than 100")
	} else if p.Priority >= 0 && p.Priority <= 100 {
		r.resultDetails = append(r.resultDetails, "++ the priority property contains a valid timestamp")
	}
}

func (p *Playbook) checkSeverity(r *results) {
	if p.Severity < 0 {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the severity property does not contain a valid value, it is less than zero")
	} else if p.Severity > 100 {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the severity property does not contain a valid value, it is greater than 100")
	} else if p.Severity >= 0 && p.Severity <= 100 {
		r.resultDetails = append(r.resultDetails, "++ the severity property contains a valid timestamp")
	}
}

func (p *Playbook) checkImpact(r *results) {
	if p.Impact < 0 {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the impact property does not contain a valid value, it is less than zero")
	} else if p.Impact > 100 {
		r.problemsFound++
		r.resultDetails = append(r.resultDetails, "-- the impact property does not contain a valid value, it is greater than 100")
	} else if p.Impact >= 0 && p.Impact <= 100 {
		r.resultDetails = append(r.resultDetails, "++ the impact property contains a valid timestamp")
	}
}

func (p *Playbook) checkExternalReferences(r *results) {
	if len(p.ExternalReferences) > 0 {
		for i := range p.ExternalReferences {
			if p.ExternalReferences[i].Name == "" {
				r.problemsFound++
				r.resultDetails = append(r.resultDetails, "-- the name property in an external reference is required but missing")
			} else {
				r.resultDetails = append(r.resultDetails, "++ the name property in an external reference is required and is present")
			}
		}
	}
}
