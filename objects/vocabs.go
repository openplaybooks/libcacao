// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

// ----------------------------------------------------------------------
// Vocabulary Functions
// ----------------------------------------------------------------------

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
