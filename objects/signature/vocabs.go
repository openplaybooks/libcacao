// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package signature

// GetSigningMethodsVocab - This will return a slice of officially supported
// signing methods
func GetSigningMethodsVocab() []string {
	return []string{
		"RS256",
		"RS384",
		"RS512",
		"ES256",
		"ES384",
		"ES512",
		"PS256",
		"PS384",
		"PS512",
		"Ed25519",
		"Ed448",
		"XMSS-SHA2_10_256",
		"XMSS-SHA2_16_256",
		"XMSS-SHA2_20_256",
		"LMS_SHA256_M32_H5",
		"LMS_SHA256_M32_H10",
		"LMS_SHA256_M32_H15",
		"LMS_SHA256_M32_H20",
		"LMS_SHA256_M32_H25",
	}
}
