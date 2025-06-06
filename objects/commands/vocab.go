// Copyright 2019-2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package commands

// ----------------------------------------------------------------------
// Vocabulary Functions
// ----------------------------------------------------------------------

// GetCommandTypesVocab - This will return a list of officially supported
// command data types
func GetCommandTypesVocab() []string {
	return []string{
		"assignment",
		"manual",
		"bash",
		"ssh",
		"http",
		"powershell",
		"caldera-cmd",
		"elastic",
		"jupyter",
		"kestrel",
		"openc2-json",
		"sigma",
		"yara",
	}
}
