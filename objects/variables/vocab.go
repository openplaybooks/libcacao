// Copyright 2019-2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package variables

// ----------------------------------------------------------------------
// Vocabulary Functions
// ----------------------------------------------------------------------

// GetVariableTypesVocab - This will return a slice of officially supported
// variable types.
func GetVariableTypesVocab() []string {
	return []string{
		"bool",
		"dictionary",
		"float",
		"hexstring",
		"integer",
		"ipv4-addr",
		"ipv6-addr",
		"list",
		"long",
		"mac-addr",
		"hash",
		"md5-hash",
		"sha256-hash",
		"string",
		"uri",
		"uuid",
	}
}
