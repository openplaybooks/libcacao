// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

// GetPlaybookTypesVocab - This will return a slice of officially supported
// playbook types
func GetPlaybookTypesVocab() []string {
	return []string{
		"attack",
		"detection",
		"engagement",
		"investigation",
		"mitigation",
		"notification",
		"prevention",
		"remediation",
	}
}
