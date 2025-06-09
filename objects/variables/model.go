// Copyright 2019-2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package variables

// VariableObject - This interface defines a variable object.
type VariableObject interface {
	GetCommon() CommonProperties
	ClearName()
}

// CommonProperties - Each variable definition contains some base properties
// that are common across all variables. These common properties are defined in
// the following table. The name property here is just to help make processing
// easier, it will be removed and added as a key in the dictionary when it is
// added to the playbook.
type CommonProperties struct {
	ObjectType  string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Constant    bool   `json:"constant,omitempty"`
	External    bool   `json:"external,omitempty"`
}

// ClearName - This method will clear the ID from the object
func (v *CommonProperties) ClearName() {
	v.Name = ""
}

// StringVariable - This type implmenets the CACAO 3.0 string variable and
// defines all of its properties.
type StringVariable struct {
	CommonProperties
	Value string `json:"value,omitempty"`
}

// GetCommon - Implement the VariableObject interface and return common properties
func (v *StringVariable) GetCommon() CommonProperties {
	return v.CommonProperties
}
