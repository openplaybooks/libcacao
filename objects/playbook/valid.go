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

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// Valid - This method will verify that the object is correct. It will return a
// boolean, an integer that tracks the number of problems found, and a slice of
// strings that contain the detailed results, whether good or bad.
func (p *Playbook) Valid(debug bool) (bool, int, []string) {
	r := new(results)
	r.debug = debug

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

func requiredButMissing(r *results, propertyName string) {
	str := fmt.Sprintf("-- the %s property is required but missing", propertyName)
	logProblem(r, str)
}

func requiredAndFound(r *results, propertyName string) {
	str := fmt.Sprintf("++ the %s property is required and is found", propertyName)
	logValid(r, str)
}

func logProblem(r *results, msg string) {
	r.problemsFound++
	r.resultDetails = append(r.resultDetails, msg)
}

func logValid(r *results, msg string) {
	if r.debug {
		r.resultDetails = append(r.resultDetails, msg)
	}
}

// Each of these methods will check a specific property. It is done this way
// to reduce the complexity of the main valid() function. This way all of the
// checks for each property are self contained in their own function.

func (p *Playbook) checkObjectType(r *results) {
	if p.ObjectType == "" {
		requiredButMissing(r, "type")
	} else {
		requiredAndFound(r, "type")

		if p.ObjectType != "playbook" && p.ObjectType != "playbook-template" {
			logProblem(r, "-- the type property does not contain a value of playbook or playbook-template")
		} else {
			str := fmt.Sprintf("++ the type property contains a valid type value of \"%s\"", p.ObjectType)
			logValid(r, str)
		}
	}
}

func (p *Playbook) checkSpecVersion(r *results) {
	if p.SpecVersion == "" {
		requiredButMissing(r, "spec_version")
	} else {
		requiredAndFound(r, "spec_version")

		if p.SpecVersion != "1.0" {
			logProblem(r, "-- the spec_version property does not contain a value of 1.0")
		} else {
			str := fmt.Sprintf("++ the spec_version property contains a valid spec_version value of \"%s\"", p.SpecVersion)
			logValid(r, str)
		}
	}
}

func (p *Playbook) checkID(r *results) {
	if p.ID == "" {
		requiredButMissing(r, "id")
	} else {
		requiredAndFound(r, "id")

		if valid := objects.IsIDValid(p.ID); valid == false {
			logProblem(r, "-- the id property does not contain a valid identifier")
		} else {
			str := fmt.Sprintf("++ the id property contains a valid identifier value of \"%s\"", p.ID)
			logValid(r, str)
		}
	}
}

func (p *Playbook) checkName(r *results) {
	if p.Name == "" {
		requiredButMissing(r, "name")
	} else {
		requiredAndFound(r, "name")
	}
}

func (p *Playbook) checkPlaybookTypes(r *results) {
	if len(p.PlaybookTypes) == 0 {
		requiredButMissing(r, "playbook_types")
	} else {
		requiredAndFound(r, "playbook_types")

		ptvocab := p.GetPlaybookTypesVocab()
		for i := 0; i < len(p.PlaybookTypes); i++ {
			value := p.PlaybookTypes[i]
			if _, found := ptvocab[value]; found {
				str := fmt.Sprintf("++ the playbook_types property contains a valid playbook_types value of \"%s\"", value)
				logValid(r, str)
			} else {
				str := fmt.Sprintf("-- the playbook_types property contains a value of \"%s\" that is not in the vocabulary", value)
				logProblem(r, str)
			}
		}
	}
}

func (p *Playbook) checkCreatedBy(r *results) {
	if p.CreatedBy == "" {
		requiredButMissing(r, "created_by")
	} else {
		requiredAndFound(r, "created_by")

		if valid := objects.IsIDValid(p.CreatedBy); valid == false {
			logProblem(r, "-- the created_by property does not contain a valid identifier")
		} else {
			str := fmt.Sprintf("++ the created_by property contains a valid identifier value of \"%s\"", p.CreatedBy)
			logValid(r, str)
		}
	}
}

func (p *Playbook) checkCreated(r *results) {
	if p.Created == "" {
		requiredButMissing(r, "created")
	} else {
		requiredAndFound(r, "created")

		if valid := objects.IsTimestampValid(p.Created); valid == false {
			logProblem(r, "-- the created property does not contain a valid timestamp")
		} else {
			str := fmt.Sprintf("++ the created property contains a valid timestamp value of \"%s\"", p.Created)
			logValid(r, str)
		}
	}
}

func (p *Playbook) checkModified(r *results) {
	if p.Modified == "" {
		requiredButMissing(r, "modified")
	} else {
		requiredAndFound(r, "modified")

		if valid := objects.IsTimestampValid(p.Modified); valid == false {
			logProblem(r, "-- the modified property does not contain a valid timestamp")
		} else {
			str := fmt.Sprintf("++ the modified property contains a valid timestamp value of \"%s\"", p.Modified)
			logValid(r, str)
		}

		// Make sure the modified timestampe is equal to or greater than created
		if p.Created != "" {
			created, _ := time.Parse(time.RFC3339, p.Created)
			modified, _ := time.Parse(time.RFC3339, p.Modified)
			if modified.After(created) || modified.Equal(created) {
				logValid(r, "++ the modified timestamp is later than or equal to the created timestamp")
			} else {
				logProblem(r, "-- the modified timestamp is not later than or eqaul to the created timestamp")
			}
		}
	}
}

func (p *Playbook) checkValidFrom(r *results) {
	if p.ValidFrom != "" {
		if valid := objects.IsTimestampValid(p.ValidFrom); valid == false {
			logProblem(r, "-- the valid_from property does not contain a valid timestamp")
		} else {
			logValid(r, "++ the valid_from property contains a valid timestamp")
		}
	}
}

func (p *Playbook) checkValidUntil(r *results) {
	if p.ValidUntil != "" {
		if valid := objects.IsTimestampValid(p.ValidUntil); valid == false {
			logProblem(r, "-- the valid_until property does not contain a valid timestamp")
		} else {
			logValid(r, "++ the valid_until property contains a valid timestamp")
		}

		// If there is a valid_until timestamp, then lets check to see if there is also a valid_from and if so
		// is the valid_until later than the valid_from
		if p.ValidFrom != "" {
			validFrom, _ := time.Parse(time.RFC3339, p.ValidFrom)
			validUntil, _ := time.Parse(time.RFC3339, p.ValidUntil)
			if validUntil.After(validFrom) {
				logValid(r, "++ the valid_until timestamp is later than the valid_from timestamp")
			} else {
				logProblem(r, "-- the valid_until timestamp is not later than the valid_from timestamp")
			}
		}
	}
}

func (p *Playbook) checkDerivedFrom(r *results) {
	if len(p.DerivedFrom) != 0 {

		for i := 0; i < len(p.DerivedFrom); i++ {
			value := p.DerivedFrom[i]

			if valid := objects.IsIDValid(value); valid == false {
				logProblem(r, "-- the derived_from property does not contain a valid identifier")
			} else {
				str := fmt.Sprintf("++ the derived_from property contains a valid identifier value of \"%s\"", value)
				logValid(r, str)
			}
		}
	}
}

func (p *Playbook) checkPriority(r *results) {
	if p.Priority < 0 {
		logProblem(r, "-- the priority property does not contain a valid value, it is less than zero")
	} else if p.Priority > 100 {
		logProblem(r, "-- the priority property does not contain a valid value, it is greater than 100")
	} else if p.Priority >= 0 && p.Priority <= 100 {
		logValid(r, "++ the priority property contains a valid timestamp")
	}
}

func (p *Playbook) checkSeverity(r *results) {
	if p.Severity < 0 {
		logProblem(r, "-- the severity property does not contain a valid value, it is less than zero")
	} else if p.Severity > 100 {
		logProblem(r, "-- the severity property does not contain a valid value, it is greater than 100")
	} else if p.Severity >= 0 && p.Severity <= 100 {
		logValid(r, "++ the severity property contains a valid timestamp")
	}
}

func (p *Playbook) checkImpact(r *results) {
	if p.Impact < 0 {
		logProblem(r, "-- the impact property does not contain a valid value, it is less than zero")
	} else if p.Impact > 100 {
		logProblem(r, "-- the impact property does not contain a valid value, it is greater than 100")
	} else if p.Impact >= 0 && p.Impact <= 100 {
		logValid(r, "++ the impact property contains a valid timestamp")
	}
}

func (p *Playbook) checkExternalReferences(r *results) {
	if len(p.ExternalReferences) > 0 {
		for i := range p.ExternalReferences {
			if p.ExternalReferences[i].Name == "" {
				logProblem(r, "-- the name property in an external reference is required but missing")
			} else {
				logValid(r, "++ the name property in an external reference is required and is present")
			}
		}
	}
}
