// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package markings

// ----------------------------------------------------------------------
// Define TLP Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common data marking properties
func (m *MarkingTLP) GetCommon() CommonProperties {
	return m.CommonProperties
}

// ----------------------------------------------------------------------
// Define Statement Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common data marking properties
func (m *MarkingStatement) GetCommon() CommonProperties {
	return m.CommonProperties
}

// ----------------------------------------------------------------------
// Define IEP Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common data marking properties
func (m *MarkingIEP) GetCommon() CommonProperties {
	return m.CommonProperties
}
