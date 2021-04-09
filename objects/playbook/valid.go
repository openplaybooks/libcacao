// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

import (
	"time"

	"github.com/openplaybooks/libcacao/objects"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// Valid - This method will verify that the object is correct. It will return a
// boolean, an integer that tracks the number of problems found, and a slice of
// strings that contain the detailed results, whether good or bad.
func (p *Playbook) Valid() (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Type
	if p.ObjectType == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the type property is required but missing")
	}
	resultDetails = append(resultDetails, "++ the type property is required and is present")

	// Spec Version
	if p.SpecVersion == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the spec_version property is required but missing")
	}
	resultDetails = append(resultDetails, "++ the spec_version property is required and is present")

	// ID
	if p.ID == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the id property is required but missing")
	}
	resultDetails = append(resultDetails, "++ the id property is required and is present")

	// Name
	if p.Name == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the name property is required but missing")
	}
	resultDetails = append(resultDetails, "++ the name property is required and is present")

	// Description
	// No requirements

	// Playbook Types
	// No requirements

	// Created By
	// No requirements

	// Created
	if p.Created == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the created property is required but missing")
	}
	resultDetails = append(resultDetails, "++ the created property is required and is present")

	// Modified
	if p.Modified == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the modified property is required but missing")
	}
	resultDetails = append(resultDetails, "++ the modified property is required and is present")

	// Revoked
	// No requirements

	// Valid From
	if p.ValidFrom != "" {
		if valid := objects.IsTimestampValid(p.ValidFrom); valid == false {
			problemsFound++
			resultDetails = append(resultDetails, "-- the valid_from property does not contain a valid timestamp")
		}
		resultDetails = append(resultDetails, "++ the valid_from property contains a valid timestamp")
	}

	// Valid Until
	if p.ValidUntil != "" {
		if valid := objects.IsTimestampValid(p.ValidUntil); valid == false {
			problemsFound++
			resultDetails = append(resultDetails, "-- the valid_until property does not contain a valid timestamp")
		}
		resultDetails = append(resultDetails, "++ the valid_until property contains a valid timestamp")

		// If there is a valid_until timestamp, then lets check to see if there is also a valid_from and if so
		// is the valid_until later than the valid_from
		if p.ValidFrom != "" {
			validFrom, _ := time.Parse(time.RFC3339, p.ValidFrom)
			validUntil, _ := time.Parse(time.RFC3339, p.ValidUntil)
			if yes := validUntil.After(validFrom); yes != true {
				problemsFound++
				resultDetails = append(resultDetails, "-- the valid_until timestamp is not later than the valid_from timestamp")
			}
			resultDetails = append(resultDetails, "++ the valid_until timestamp is later than the valid_from timestamp")
		}
	}

	// Derived From
	// No requirements

	// Priority
	if p.Priority < 0 {
		problemsFound++
		resultDetails = append(resultDetails, "-- the priority property does not contain a valid value, it is less than zero")
	} else if p.Priority > 100 {
		problemsFound++
		resultDetails = append(resultDetails, "-- the priority property does not contain a valid value, it is greater than 100")
	} else if p.Priority >= 0 && p.Priority <= 100 {
		resultDetails = append(resultDetails, "++ the priority property contains a valid timestamp")
	}

	// Severity
	if p.Severity < 0 {
		problemsFound++
		resultDetails = append(resultDetails, "-- the severity property does not contain a valid value, it is less than zero")
	} else if p.Severity > 100 {
		problemsFound++
		resultDetails = append(resultDetails, "-- the severity property does not contain a valid value, it is greater than 100")
	} else if p.Severity >= 0 && p.Severity <= 100 {
		resultDetails = append(resultDetails, "++ the severity property contains a valid timestamp")
	}

	// Impact
	if p.Impact < 0 {
		problemsFound++
		resultDetails = append(resultDetails, "-- the impact property does not contain a valid value, it is less than zero")
	} else if p.Impact > 100 {
		problemsFound++
		resultDetails = append(resultDetails, "-- the impact property does not contain a valid value, it is greater than 100")
	} else if p.Impact >= 0 && p.Impact <= 100 {
		resultDetails = append(resultDetails, "++ the impact property contains a valid timestamp")
	}

	// Labels
	// No requirements

	// External References
	if len(p.ExternalReferences) > 0 {
		for i, _ := range p.ExternalReferences {
			if p.ExternalReferences[i].Name == "" {
				problemsFound++
				resultDetails = append(resultDetails, "-- the name property in an external reference is required but missing")
			}
			resultDetails = append(resultDetails, "++ the name property in an external reference is required and is present")
		}
	}

	// Features
	// No requirements

	// Finished Checks

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}
	return true, problemsFound, resultDetails
}
