// Copyright 2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package steps

// GetWorkflowStepTypesVocab - This will return a slice of officially supported
// workflow step types
func GetWorkflowStepTypesVocab() []string {
	return []string{
		"start",
		"return",
		"end",
		"action",
		"playbook-action",
		"parallel",
		"foreach",
		"while",
		"if-then",
		"switch",
	}
}
