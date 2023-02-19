// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

// ----------------------------------------------------------------------
// Vocabulary Functions
// ----------------------------------------------------------------------

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

// GetWorkflowStepTypesVocab - This function will return a map of the workflow step types
func GetWorkflowStepTypesVocab() map[string]bool {
	workflowStepTypesVocab := map[string]bool{
		"start":            true,
		"end":              true,
		"action":           true,
		"playbook-action":  true,
		"parallel":         true,
		"if-condition":     true,
		"while-condition":  true,
		"switch-condition": true,
	}
	return workflowStepTypesVocab
}

// IsWorkflowStepTypeValid - This function will take in a string representing a
// workflow step and return true or false if it is an officially supported
// step.
func IsWorkflowStepTypeValid(s string) bool {
	workflowStepTypes := GetWorkflowStepTypesVocab()

	if _, found := workflowStepTypes[s]; found == true {
		return true
	}
	return false
}

// GetCommandDataTypesVocab - This funtion will return a list of command data types
func GetCommandDataTypesVocab() map[string]bool {
	commandDataTypesVocab := map[string]bool{
		"manual":      true,
		"bash":        true,
		"http-api":    true,
		"ssh":         true,
		"caldera-cmd": true,
		"elastic":     true,
		"jupyter":     true,
		"kestrel":     true,
		"openc2-json": true,
		"sigma":       true,
		"yara":        true,
	}
	return commandDataTypesVocab
}

// IsCommandDataTypeValid - This function will take in a string representing a
// command type and return true or false if it is an officially supported
// command.
func IsCommandDataTypeValid(s string) bool {
	workflowStepTypes := GetCommandDataTypesVocab()

	if _, found := workflowStepTypes[s]; found == true {
		return true
	}
	return false
}

// GetIndustrySectorsVocab - This funtion will return a list of industry sectors
func GetIndustrySectorsVocab() map[string]bool {
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
	industrySectors := GetIndustrySectorsVocab()

	if _, found := industrySectors[s]; found == true {
		return true
	}
	return false
}

// GetSigningMethodsVocab - This function will return a map of the valid signing methods
func GetSigningMethodsVocab() map[string]bool {
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
	signingMethods := GetSigningMethodsVocab()

	if _, found := signingMethods[s]; found == true {
		return true
	}
	return false
}

// GetVariableTypesVocab - This function will return a map of the valid variable
// types.
func GetVariableTypesVocab() map[string]bool {
	variableTypesVocab := map[string]bool{
		"bool":        true,
		"dictionary":  true,
		"float":       true,
		"hexstring":   true,
		"integer":     true,
		"ipv4-addr":   true,
		"ipv6-addr":   true,
		"long":        true,
		"mac-addr":    true,
		"sha256-hash": true,
		"string":      true,
		"uri":         true,
		"uuid":        true,
	}
	return variableTypesVocab
}

// IsVariableTypeValid - This function will take in a string representing a
// variable type and return true or false if it is an officially supported type.
func IsVariableTypeValid(s string) bool {
	variableTypes := GetVariableTypesVocab()

	if _, found := variableTypes[s]; found == true {
		return true
	}
	return false
}
