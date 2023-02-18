// Copyright 2023 Bret Jordan, All rights reserved.
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
func (p *Playbook) Compare(p2 *Playbook, debug bool) (bool, int, []string) {
	var r *results = new(results)
	r.debug = debug

	// Object Type
	if p.ObjectType != p2.ObjectType {
		str := fmt.Sprintf("-- the type property values do not match: %s | %s", p.ObjectType, p2.ObjectType)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the type property values match: %s | %s", p.ObjectType, p2.ObjectType)
		logValid(r, str)
	}

	// Spec Version
	if p.SpecVersion != p2.SpecVersion {
		str := fmt.Sprintf("-- the spec_version property values do not match: %s | %s", p.SpecVersion, p2.SpecVersion)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the spec_version property values match: %s | %s", p.SpecVersion, p2.SpecVersion)
		logValid(r, str)
	}

	// ID
	if p.ID != p2.ID {
		str := fmt.Sprintf("-- the id property values do not match: %s | %s", p.ID, p2.ID)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the id property values match: %s | %s", p.ID, p2.ID)
		logValid(r, str)
	}

	// Name
	if p.Name != p2.Name {
		str := fmt.Sprintf("-- the names property values do not match: %s | %s", p.Name, p2.Name)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the names property values match: %s | %s", p.Name, p2.Name)
		logValid(r, str)
	}

	// Description
	if p.Description != p2.Description {
		str := fmt.Sprintf("-- the description property values do not match: %s | %s", p.Description, p2.Description)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the description property values match: %s | %s", p.Description, p2.Description)
		logValid(r, str)
	}

	// Playbook Types
	if len(p.PlaybookTypes) != len(p2.PlaybookTypes) {
		str := fmt.Sprintf("-- the number of entries in the playbook_types property do not match: %d | %d", len(p.PlaybookTypes), len(p2.PlaybookTypes))
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the number of entries in the playbook_types property match: %d | %d", len(p.PlaybookTypes), len(p2.PlaybookTypes))
		logValid(r, str)

		// If lengths are the same, then check each value
		for index := range p.PlaybookTypes {
			if p.PlaybookTypes[index] != p2.PlaybookTypes[index] {
				str := fmt.Sprintf("-- the playbook_type values do not match: %s | %s", p.PlaybookTypes[index], p2.PlaybookTypes[index])
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ the playbook_type values match: %s | %s", p.PlaybookTypes[index], p2.PlaybookTypes[index])
				logValid(r, str)
			}
		}
	}

	// CreatedBy
	if p.CreatedBy != p2.CreatedBy {
		str := fmt.Sprintf("-- the created by values do not match: %s | %s", p.CreatedBy, p2.CreatedBy)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the created by values match: %s | %s", p.CreatedBy, p2.CreatedBy)
		logValid(r, str)
	}

	// Created
	if p.Created != p2.Created {
		str := fmt.Sprintf("-- the created dates do not match: %s | %s", p.Created, p2.Created)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the created dates match: %s | %s", p.Created, p2.Created)
		logValid(r, str)
	}

	// Modified
	if p.Modified != p2.Modified {
		str := fmt.Sprintf("-- the modified dates do not match: %s | %s", p.Modified, p2.Modified)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the modified dates match: %s | %s", p.Modified, p2.Modified)
		logValid(r, str)
	}

	// Revoked
	if p.Revoked != p2.Revoked {
		str := fmt.Sprintf("-- the revoked values do not match: %t | %t", p.Revoked, p2.Revoked)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the revoked values match: %t | %t", p.Revoked, p2.Revoked)
		logValid(r, str)
	}

	// Valid From
	if p.ValidFrom != p2.ValidFrom {
		str := fmt.Sprintf("-- the valid_from values do not match: %s | %s", p.ValidFrom, p2.ValidFrom)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the valid_from values match: %s | %s", p.ValidFrom, p2.ValidFrom)
		logValid(r, str)
	}

	// Valid Until
	if p.ValidUntil != p2.ValidUntil {
		str := fmt.Sprintf("-- the valid_until values do not match: %s | %s", p.ValidUntil, p2.ValidUntil)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the valid_until values match: %s | %s", p.ValidUntil, p2.ValidUntil)
		logValid(r, str)
	}

	// Derived From
	if len(p.DerivedFrom) != len(p2.DerivedFrom) {
		str := fmt.Sprintf("-- the number of entries in the derived_from property do not match: %d | %d", len(p.DerivedFrom), len(p2.DerivedFrom))
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the number of entries in the derived_from property match: %d | %d", len(p.DerivedFrom), len(p2.DerivedFrom))
		logValid(r, str)

		// If lengths are the same, then check each value
		for index := range p.DerivedFrom {
			if p.DerivedFrom[index] != p2.DerivedFrom[index] {
				str := fmt.Sprintf("-- the derived_from values do not match: %s | %s", p.DerivedFrom[index], p2.DerivedFrom[index])
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ the derived_from values match: %s | %s", p.DerivedFrom[index], p2.DerivedFrom[index])
				logValid(r, str)
			}
		}
	}

	// Priority
	if p.Priority != p2.Priority {
		str := fmt.Sprintf("-- the priority values do not match: %d | %d", p.Priority, p2.Priority)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the priority values match: %d | %d", p.Priority, p2.Priority)
		logValid(r, str)
	}

	// Severity
	if p.Severity != p2.Severity {
		str := fmt.Sprintf("-- the severity values do not match: %d | %d", p.Severity, p2.Severity)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the severity values match: %d | %d", p.Severity, p2.Severity)
		logValid(r, str)
	}

	// Impact
	if p.Impact != p2.Impact {
		str := fmt.Sprintf("-- the impact values do not match: %d | %d", p.Impact, p2.Impact)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the impact values match: %d | %d", p.Impact, p2.Impact)
		logValid(r, str)
	}

	// Labels
	if len(p.Labels) != len(p2.Labels) {
		str := fmt.Sprintf("-- the number of entries in the labels property do not match: %d | %d", len(p.Labels), len(p2.Labels))
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the number of entries in the labels property match: %d | %d", len(p.Labels), len(p2.Labels))
		logValid(r, str)

		// If lengths are the same, then check each value
		for index := range p.Labels {
			if p.Labels[index] != p2.Labels[index] {
				str := fmt.Sprintf("-- the label values do not match: %s | %s", p.Labels[index], p2.Labels[index])
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ the label values match: %s | %s", p.Labels[index], p2.Labels[index])
				logValid(r, str)
			}
		}
	}

	// Check External References
	if len(p.ExternalReferences) != len(p2.ExternalReferences) {
		str := fmt.Sprintf("-- the number of entries in the external references property do not match: %d | %d", len(p.ExternalReferences), len(p2.ExternalReferences))
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the number of entries in the external references property match: %d | %d", len(p.ExternalReferences), len(p2.ExternalReferences))
		logValid(r, str)
		for index := range p.ExternalReferences {

			// Check External Reference Name
			if p.ExternalReferences[index].Name != p2.ExternalReferences[index].Name {
				str := fmt.Sprintf("-- the name values do not match: %s | %s", p.ExternalReferences[index].Name, p2.ExternalReferences[index].Name)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ the name values match: %s | %s", p.ExternalReferences[index].Name, p2.ExternalReferences[index].Name)
				logValid(r, str)
			}

			// Check External Reference Descriptions
			if p.ExternalReferences[index].Description != p2.ExternalReferences[index].Description {
				str := fmt.Sprintf("-- the description values do not match: %s | %s", p.ExternalReferences[index].Description, p2.ExternalReferences[index].Description)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ the description values match: %s | %s", p.ExternalReferences[index].Description, p2.ExternalReferences[index].Description)
				logValid(r, str)
			}

			// Check External Reference Source
			if p.ExternalReferences[index].Source != p2.ExternalReferences[index].Source {
				str := fmt.Sprintf("-- the source values do not match: %s | %s", p.ExternalReferences[index].Source, p2.ExternalReferences[index].Source)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ the source values match: %s | %s", p.ExternalReferences[index].Source, p2.ExternalReferences[index].Source)
				logValid(r, str)
			}

			// Check External Reference URLs
			if p.ExternalReferences[index].URL != p2.ExternalReferences[index].URL {
				str := fmt.Sprintf("-- the url values do not match: %s | %s", p.ExternalReferences[index].URL, p2.ExternalReferences[index].URL)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ the url values match: %s | %s", p.ExternalReferences[index].URL, p2.ExternalReferences[index].URL)
				logValid(r, str)
			}

			// Check External Reference External IDs
			if p.ExternalReferences[index].ExternalID != p2.ExternalReferences[index].ExternalID {
				str := fmt.Sprintf("-- the external_id values do not match: %s | %s", p.ExternalReferences[index].ExternalID, p2.ExternalReferences[index].ExternalID)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ the external_id values match: %s | %s", p.ExternalReferences[index].ExternalID, p2.ExternalReferences[index].ExternalID)
				logValid(r, str)
			}

			// Check External Reference Reference IDs
			if p.ExternalReferences[index].ReferenceID != p2.ExternalReferences[index].ReferenceID {
				str := fmt.Sprintf("-- the reference_id values do not match: %s | %s", p.ExternalReferences[index].ReferenceID, p2.ExternalReferences[index].ReferenceID)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ the reference_id values match: %s | %s", p.ExternalReferences[index].ReferenceID, p2.ExternalReferences[index].ReferenceID)
				logValid(r, str)
			}
		}
	}

	// Features
	// if p.Features.ParallelProcessing != p2.Features.ParallelProcessing {
	// 	str := fmt.Sprintf("-- the features parallel-processing values do not match: %s | %s", p.Features.ParallelProcessing, p2.Features.ParallelProcessing)
	// 	logProblem(r, str)
	// } else {
	// 	str := fmt.Sprintf("++ the features parallel-processing values match: %s | %s", p.Features.ParallelProcessing, p2.Features.ParallelProcessing)
	// 	logValid(r, str)
	// }

	// if p.Features.IfLogic != p2.Features.IfLogic {
	// 	str := fmt.Sprintf("-- the features if-logic values do not match: %s | %s", p.Features.IfLogic, p2.Features.IfLogic)
	// 	logProblem(r, str)
	// } else {
	// 	str := fmt.Sprintf("++ the features if-logic values match: %s | %s", p.Features.IfLogic, p2.Features.IfLogic)
	// 	logValid(r, str)
	// }

	// if p.Features.WhileLogic != p2.Features.WhileLogic {
	// 	str := fmt.Sprintf("-- the features while-logic values do not match: %s | %s", p.Features.WhileLogic, p2.Features.WhileLogic)
	// 	logProblem(r, str)
	// } else {
	// 	str := fmt.Sprintf("++ the features while-logic values match: %s | %s", p.Features.WhileLogic, p2.Features.WhileLogic)
	// 	logValid(r, str)
	// }

	// if p.Features.SwitchLogic != p2.Features.SwitchLogic {
	// 	str := fmt.Sprintf("-- the features switch-logic values do not match: %s | %s", p.Features.SwitchLogic, p2.Features.SwitchLogic)
	// 	logProblem(r, str)
	// } else {
	// 	str := fmt.Sprintf("++ the features switch-logic values match: %s | %s", p.Features.SwitchLogic, p2.Features.SwitchLogic)
	// 	logValid(r, str)
	// }

	// if p.Features.TemporalLogic != p2.Features.TemporalLogic {
	// 	str := fmt.Sprintf("-- the features temporal-logic values do not match: %s | %s", p.Features.TemporalLogic, p2.Features.TemporalLogic)
	// 	logProblem(r, str)
	// } else {
	// 	str := fmt.Sprintf("++ the features temporal-logic values match: %s | %s", p.Features.TemporalLogic, p2.Features.TemporalLogic)
	// 	logValid(r, str)
	// }

	// if p.Features.DataMarkings != p2.Features.DataMarkings {
	// 	str := fmt.Sprintf("-- the features data-markings values do not match: %s | %s", p.Features.DataMarkings, p2.Features.DataMarkings)
	// 	logProblem(r, str)
	// } else {
	// 	str := fmt.Sprintf("++ the features data-markings values match: %s | %s", p.Features.DataMarkings, p2.Features.DataMarkings)
	// 	logValid(r, str)
	// }

	// if p.Features.Extensions != p2.Features.Extensions {
	// 	str := fmt.Sprintf("-- the features extensions values do not match: %s | %s", p.Features.Extensions, p2.Features.Extensions)
	// 	logProblem(r, str)
	// } else {
	// 	str := fmt.Sprintf("++ the features extensions values match: %s | %s", p.Features.Extensions, p2.Features.Extensions)
	// 	logValid(r, str)
	// }

	// TODO need finish fleshing this out

	// End
	if r.problemsFound > 0 {
		return false, r.problemsFound, r.resultDetails
	}

	return true, 0, r.resultDetails
}
