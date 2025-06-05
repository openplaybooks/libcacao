// Copyright 2019-2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

// ----------------------------------------------------------------------
// Vocabulary Functions
// ----------------------------------------------------------------------

// GetOperatorsVocab - This will return a slice of officially supported operator
// types
func GetOperatorsVocab() []string {
	return []string{
		"to_int",
		"to_float",
		"to_str",
		"to_bool",
		"to_list",
	}
}

// GetDelimitersVocab - This will return a slice of officially supported
// delimiters
func GetDelimitersVocab() []string {
	return []string{
		":",
		",",
		"space",
		"tab",
		"lf",
		"cr",
		"eol",
	}
}

// GetIndustrySectorsVocab - This will return a slice of officially supported
// industry sectors
func GetIndustrySectorsVocab() []string {
	return []string{
		"aerospace",
		"aviation",
		"agriculture",
		"automotive",
		"biotechnology",
		"chemical",
		"commercial",
		"consulting",
		"construction",
		"cosmetics",
		"critical-infrastructure",
		"dams",
		"defense",
		"education",
		"emergency-services",
		"energy",
		"non-renewable-energy",
		"renewable-energy",
		"media",
		"financial",
		"food",
		"gambling",
		"government",
		"local-government",
		"national-government",
		"regional-government",
		"public-services",
		"healthcare",
		"information-communications-technology",
		"electronics-hardware",
		"software",
		"telecommunications",
		"legal-services",
		"lodging",
		"manufacturing",
		"maritime",
		"metals",
		"mining",
		"non-profit",
		"humanitarian-aid",
		"human-rights",
		"nuclear",
		"petroleum",
		"pharmaceuticals",
		"research",
		"transportation",
		"logistics-shipping",
		"utilities",
		"video-game",
		"water",
	}
}
