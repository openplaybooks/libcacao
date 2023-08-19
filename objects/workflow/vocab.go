// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package workflow

// GetWorkflowStepTypesVocab - This will return a slice of officially supported
// workflow step types
func GetWorkflowStepTypesVocab() []string {
	return []string{
		"start",
		"end",
		"action",
		"playbook-action",
		"parallel",
		"if-condition",
		"while-condition",
		"switch-condition",
	}
}

// GetCommandDataTypesVocab - This will return a list of officially supported
// command data types
func GetCommandDataTypesVocab() []string {
	return []string{
		"manual",
		"bash",
		"http-api",
		"ssh",
		"caldera-cmd",
		"elastic",
		"jupyter",
		"kestrel",
		"openc2-json",
		"sigma",
		"yara",
	}
}
