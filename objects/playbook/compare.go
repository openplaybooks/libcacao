// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

import "fmt"

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// Compare - This method will compare two objects to make sure they are the
// same and will return a boolean, an integer that tracks the number of problems
// found, and a slice of strings that contain the detailed results, whether good
// or bad.
func (p *Playbook) Compare(p2 *Playbook) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Object Type
	if p.ObjectType != p2.ObjectType {
		problemsFound++
		str := fmt.Sprintf("-- the type property values do not match: %s | %s", p.ObjectType, p2.ObjectType)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the type property values match: %s | %s", p.ObjectType, p2.ObjectType)
		resultDetails = append(resultDetails, str)
	}

	// Spec Version
	if p.SpecVersion != p2.SpecVersion {
		problemsFound++
		str := fmt.Sprintf("-- the spec_version property values do not match: %s | %s", p.SpecVersion, p2.SpecVersion)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the spec_version property values match: %s | %s", p.SpecVersion, p2.SpecVersion)
		resultDetails = append(resultDetails, str)
	}

	// ID
	if p.ID != p2.ID {
		problemsFound++
		str := fmt.Sprintf("-- the id property values do not match: %s | %s", p.ID, p2.ID)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the id property values match: %s | %s", p.ID, p2.ID)
		resultDetails = append(resultDetails, str)
	}

	// Name
	if p.Name != p2.Name {
		problemsFound++
		str := fmt.Sprintf("-- the names property values do not match: %s | %s", p.Name, p2.Name)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the names property values match: %s | %s", p.Name, p2.Name)
		resultDetails = append(resultDetails, str)
	}

	// Description
	if p.Description != p2.Description {
		problemsFound++
		str := fmt.Sprintf("-- the description property values do not match: %s | %s", p.Description, p2.Description)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the description property values match: %s | %s", p.Description, p2.Description)
		resultDetails = append(resultDetails, str)
	}

	// Playbook Types
	if len(p.PlaybookTypes) != len(p2.PlaybookTypes) {
		problemsFound++
		str := fmt.Sprintf("-- the number of entries in the playbook_types property do not match: %d | %d", len(p.PlaybookTypes), len(p2.PlaybookTypes))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the number of entries in the playbook_types property match: %d | %d", len(p.PlaybookTypes), len(p2.PlaybookTypes))
		resultDetails = append(resultDetails, str)

		// If lengths are the same, then check each value
		for index := range p.PlaybookTypes {
			if p.PlaybookTypes[index] != p2.PlaybookTypes[index] {
				problemsFound++
				str := fmt.Sprintf("-- the playbook_type values do not match: %s | %s", p.PlaybookTypes[index], p2.PlaybookTypes[index])
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ the playbook_type values match: %s | %s", p.PlaybookTypes[index], p2.PlaybookTypes[index])
				resultDetails = append(resultDetails, str)
			}
		}
	}

	// CreatedBy
	if p.CreatedBy != p2.CreatedBy {
		problemsFound++
		str := fmt.Sprintf("-- the created by values do not match: %s | %s", p.CreatedBy, p2.CreatedBy)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the created by values match: %s | %s", p.CreatedBy, p2.CreatedBy)
		resultDetails = append(resultDetails, str)
	}

	// Created
	if p.Created != p2.Created {
		problemsFound++
		str := fmt.Sprintf("-- the created dates do not match: %s | %s", p.Created, p2.Created)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the created dates match: %s | %s", p.Created, p2.Created)
		resultDetails = append(resultDetails, str)
	}

	// Modified
	if p.Modified != p2.Modified {
		problemsFound++
		str := fmt.Sprintf("-- the modified dates do not match: %s | %s", p.Modified, p2.Modified)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the modified dates match: %s | %s", p.Modified, p2.Modified)
		resultDetails = append(resultDetails, str)
	}

	// Revoked
	if p.Revoked != p2.Revoked {
		problemsFound++
		str := fmt.Sprintf("-- the revoked values do not match: %t | %t", p.Revoked, p2.Revoked)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the revoked values match: %t | %t", p.Revoked, p2.Revoked)
		resultDetails = append(resultDetails, str)
	}

	// Valid From
	if p.ValidFrom != p2.ValidFrom {
		problemsFound++
		str := fmt.Sprintf("-- the valid_from values do not match: %t | %t", p.ValidFrom, p2.ValidFrom)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the valid_from values match: %t | %t", p.ValidFrom, p2.ValidFrom)
		resultDetails = append(resultDetails, str)
	}

	// Valid Until
	if p.ValidUntil != p2.ValidUntil {
		problemsFound++
		str := fmt.Sprintf("-- the valid_until values do not match: %t | %t", p.ValidUntil, p2.ValidUntil)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the valid_until values match: %t | %t", p.ValidUntil, p2.ValidUntil)
		resultDetails = append(resultDetails, str)
	}

	// Derived From
	if len(p.DerivedFrom) != len(p2.DerivedFrom) {
		problemsFound++
		str := fmt.Sprintf("-- the number of entries in the derived_from property do not match: %d | %d", len(p.DerivedFrom), len(p2.DerivedFrom))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the number of entries in the derived_from property match: %d | %d", len(p.DerivedFrom), len(p2.DerivedFrom))
		resultDetails = append(resultDetails, str)

		// If lengths are the same, then check each value
		for index := range p.DerivedFrom {
			if p.DerivedFrom[index] != p2.DerivedFrom[index] {
				problemsFound++
				str := fmt.Sprintf("-- the derived_from values do not match: %s | %s", p.DerivedFrom[index], p2.DerivedFrom[index])
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ the derived_from values match: %s | %s", p.DerivedFrom[index], p2.DerivedFrom[index])
				resultDetails = append(resultDetails, str)
			}
		}
	}

	// Priority
	if p.Priority != p2.Priority {
		problemsFound++
		str := fmt.Sprintf("-- the priority values do not match: %t | %t", p.Priority, p2.Priority)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the priority values match: %t | %t", p.Priority, p2.Priority)
		resultDetails = append(resultDetails, str)
	}

	// Severity
	if p.Severity != p2.Severity {
		problemsFound++
		str := fmt.Sprintf("-- the severity values do not match: %t | %t", p.Severity, p2.Severity)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the severity values match: %t | %t", p.Severity, p2.Severity)
		resultDetails = append(resultDetails, str)
	}

	// Impact
	if p.Impact != p2.Impact {
		problemsFound++
		str := fmt.Sprintf("-- the impact values do not match: %t | %t", p.Impact, p2.Impact)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the impact values match: %t | %t", p.Impact, p2.Impact)
		resultDetails = append(resultDetails, str)
	}

	// Labels
	if len(p.Labels) != len(p2.Labels) {
		problemsFound++
		str := fmt.Sprintf("-- the number of entries in the labels property do not match: %d | %d", len(p.Labels), len(p2.Labels))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the number of entries in the labels property match: %d | %d", len(p.Labels), len(p2.Labels))
		resultDetails = append(resultDetails, str)

		// If lengths are the same, then check each value
		for index := range p.Labels {
			if p.Labels[index] != p2.Labels[index] {
				problemsFound++
				str := fmt.Sprintf("-- the label values do not match: %s | %s", p.Labels[index], p2.Labels[index])
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ the label values match: %s | %s", p.Labels[index], p2.Labels[index])
				resultDetails = append(resultDetails, str)
			}
		}
	}

	// Check External References
	if len(p.ExternalReferences) != len(p2.ExternalReferences) {
		problemsFound++
		str := fmt.Sprintf("-- the number of entries in the external references property do not match: %d | %d", len(p.ExternalReferences), len(p2.ExternalReferences))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the number of entries in the external references property match: %d | %d", len(p.ExternalReferences), len(p2.ExternalReferences))
		resultDetails = append(resultDetails, str)
		for index := range p.ExternalReferences {

			// Check External Reference Name
			if p.ExternalReferences[index].Name != p2.ExternalReferences[index].Name {
				problemsFound++
				str := fmt.Sprintf("-- the name values do not match: %s | %s", p.ExternalReferences[index].Name, p2.ExternalReferences[index].Name)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ the name values match: %s | %s", p.ExternalReferences[index].Name, p2.ExternalReferences[index].Name)
				resultDetails = append(resultDetails, str)
			}

			// Check External Reference Descriptions
			if p.ExternalReferences[index].Description != p2.ExternalReferences[index].Description {
				problemsFound++
				str := fmt.Sprintf("-- the description values do not match: %s | %s", p.ExternalReferences[index].Description, p2.ExternalReferences[index].Description)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ the description values match: %s | %s", p.ExternalReferences[index].Description, p2.ExternalReferences[index].Description)
				resultDetails = append(resultDetails, str)
			}

			// Check External Reference Source
			if p.ExternalReferences[index].Source != p2.ExternalReferences[index].Source {
				problemsFound++
				str := fmt.Sprintf("-- the source values do not match: %s | %s", p.ExternalReferences[index].Source, p2.ExternalReferences[index].Source)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ the source values match: %s | %s", p.ExternalReferences[index].Source, p2.ExternalReferences[index].Source)
				resultDetails = append(resultDetails, str)
			}

			// Check External Reference URLs
			if p.ExternalReferences[index].URL != p2.ExternalReferences[index].URL {
				problemsFound++
				str := fmt.Sprintf("-- the url values do not match: %s | %s", p.ExternalReferences[index].URL, p2.ExternalReferences[index].URL)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ the url values match: %s | %s", p.ExternalReferences[index].URL, p2.ExternalReferences[index].URL)
				resultDetails = append(resultDetails, str)
			}

			// Check External Reference External IDs
			if p.ExternalReferences[index].ExternalID != p2.ExternalReferences[index].ExternalID {
				problemsFound++
				str := fmt.Sprintf("-- the external_id values do not match: %s | %s", p.ExternalReferences[index].ExternalID, p2.ExternalReferences[index].ExternalID)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ the external_id values match: %s | %s", p.ExternalReferences[index].ExternalID, p2.ExternalReferences[index].ExternalID)
				resultDetails = append(resultDetails, str)
			}

			// Check External Reference Reference IDs
			if p.ExternalReferences[index].ReferenceID != p2.ExternalReferences[index].ReferenceID {
				problemsFound++
				str := fmt.Sprintf("-- the reference_id values do not match: %s | %s", p.ExternalReferences[index].ReferenceID, p2.ExternalReferences[index].ReferenceID)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ the reference_id values match: %s | %s", p.ExternalReferences[index].ReferenceID, p2.ExternalReferences[index].ReferenceID)
				resultDetails = append(resultDetails, str)
			}
		}
	}

	// Features
	if p.Features.ParallelProcessing != p2.Features.ParallelProcessing {
		problemsFound++
		str := fmt.Sprintf("-- the features parallel-processing values do not match: %s | %s", p.Features.ParallelProcessing, p2.Features.ParallelProcessing)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the features parallel-processing values match: %s | %s", p.Features.ParallelProcessing, p2.Features.ParallelProcessing)
		resultDetails = append(resultDetails, str)
	}

	if p.Features.IfLogic != p2.Features.IfLogic {
		problemsFound++
		str := fmt.Sprintf("-- the features if-logic values do not match: %s | %s", p.Features.IfLogic, p2.Features.IfLogic)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the features if-logic values match: %s | %s", p.Features.IfLogic, p2.Features.IfLogic)
		resultDetails = append(resultDetails, str)
	}

	if p.Features.WhileLogic != p2.Features.WhileLogic {
		problemsFound++
		str := fmt.Sprintf("-- the features while-logic values do not match: %s | %s", p.Features.WhileLogic, p2.Features.WhileLogic)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the features while-logic values match: %s | %s", p.Features.WhileLogic, p2.Features.WhileLogic)
		resultDetails = append(resultDetails, str)
	}

	if p.Features.SwitchLogic != p2.Features.SwitchLogic {
		problemsFound++
		str := fmt.Sprintf("-- the features switch-logic values do not match: %s | %s", p.Features.SwitchLogic, p2.Features.SwitchLogic)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the features switch-logic values match: %s | %s", p.Features.SwitchLogic, p2.Features.SwitchLogic)
		resultDetails = append(resultDetails, str)
	}

	if p.Features.TemporalLogic != p2.Features.TemporalLogic {
		problemsFound++
		str := fmt.Sprintf("-- the features temporal-logic values do not match: %s | %s", p.Features.TemporalLogic, p2.Features.TemporalLogic)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the features temporal-logic values match: %s | %s", p.Features.TemporalLogic, p2.Features.TemporalLogic)
		resultDetails = append(resultDetails, str)
	}

	if p.Features.DataMarkings != p2.Features.DataMarkings {
		problemsFound++
		str := fmt.Sprintf("-- the features data-markings values do not match: %s | %s", p.Features.DataMarkings, p2.Features.DataMarkings)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the features data-markings values match: %s | %s", p.Features.DataMarkings, p2.Features.DataMarkings)
		resultDetails = append(resultDetails, str)
	}

	if p.Features.Extensions != p2.Features.Extensions {
		problemsFound++
		str := fmt.Sprintf("-- the features extensions values do not match: %s | %s", p.Features.Extensions, p2.Features.Extensions)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the features extensions values match: %s | %s", p.Features.Extensions, p2.Features.Extensions)
		resultDetails = append(resultDetails, str)
	}

	// TODO need finish fleshing this out

	// End
	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
