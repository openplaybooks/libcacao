// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

// ----------------------------------------------------------------------
// Vocabulary Functions
// ----------------------------------------------------------------------

// GetObjectTypes - This function will return a map of the object types.
func GetObjectTypes() map[string]bool {
	objectTypes := map[string]bool{
		"playbook":             true,
		"playbook-template":    true,
		"signature":            true,
		"step":                 true,
		"target":               true,
		"extension-definition": true,
		"marking-definition":   true,
		"identity":             true,
	}
	return objectTypes
}

// IsTypeValid - This function will take in a string representing an object type
// and return true or false if it is an officially supported object.
func IsTypeValid(s string) bool {
	objectTypes := GetObjectTypes()

	if _, found := objectTypes[s]; found == true {
		return true
	}
	return false
}

// GetPlaybookTypesVocab - This function will return a map of the playbook types
func GetPlaybookTypesVocab() map[string]bool {
	playbookTypesVocab := map[string]bool{
		"notification":  true,
		"detection":     true,
		"investigation": true,
		"prevention":    true,
		"mitigation":    true,
		"remediation":   true,
		"attack":        true,
	}
	return playbookTypesVocab
}

// IsPlaybookTypeValid - This function will take in a string representing an
// playbook type and return true or false if it is an officially supported type.
func IsPlaybookTypeValid(s string) bool {
	playbookTypes := GetPlaybookTypesVocab()

	if _, found := playbookTypes[s]; found == true {
		return true
	}
	return false
}

// GetIndustrySectorVocab - This funtion will return a list of industry sectors
func GetIndustrySectorVocab() map[string]bool {
	industrySectorVocab := map[string]bool{
		"aerospace":                             true,
		"aviation":                              true,
		"agriculture":                           true,
		"automotive":                            true,
		"biotechnology":                         true,
		"chemical":                              true,
		"commercial":                            true,
		"consulting":                            true,
		"construction":                          true,
		"cosmetics":                             true,
		"critical-infrastructure":               true,
		"dams":                                  true,
		"defense":                               true,
		"education":                             true,
		"emergency-services":                    true,
		"energy":                                true,
		"non-renewable-energy":                  true,
		"renewable-energy":                      true,
		"media":                                 true,
		"financial":                             true,
		"food":                                  true,
		"gambling":                              true,
		"government":                            true,
		"local-government":                      true,
		"national-government":                   true,
		"regional-government":                   true,
		"public-services":                       true,
		"healthcare":                            true,
		"information-communications-technology": true,
		"electronics-hardware":                  true,
		"software":                              true,
		"telecommunications":                    true,
		"legal-services":                        true,
		"lodging":                               true,
		"manufacturing":                         true,
		"maritime":                              true,
		"metals":                                true,
		"mining":                                true,
		"non-profit":                            true,
		"humanitarian-aid":                      true,
		"human-rights":                          true,
		"nuclear":                               true,
		"petroleum":                             true,
		"pharmaceuticals":                       true,
		"research":                              true,
		"transportation":                        true,
		"logistics-shipping":                    true,
		"utilities":                             true,
		"video-game":                            true,
		"water":                                 true,
	}
	return industrySectorVocab
}

// IsIndustrySectorValid - This function will take in a string representing an
// industry sector and return true or false if it is an officially supported.
func IsIndustrySectorValid(s string) bool {
	industrySectors := GetIndustrySectorVocab()

	if _, found := industrySectors[s]; found == true {
		return true
	}
	return false
}

// GetSigningMethod - This function will return a map of the valid signing methods
func GetSigningMethod() map[string]bool {
	signingMethodsVocab := map[string]bool{
		"RS256": true,
		"RS384": true,
		"RS512": true,
		"ES256": true,
		"ES384": true,
		"ES512": true,
	}
	return signingMethodsVocab
}

// IsSigningMethodValid - This function will take in a string representing a
// signing method and return true or false if it is an officially supported type.
func IsSigningMethodValid(s string) bool {
	signingMethods := GetSigningMethod()

	if _, found := signingMethods[s]; found == true {
		return true
	}
	return false
}
