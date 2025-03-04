// Copyright 2019-2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package variables implements the CACAO 3.0 variables object.
//
// Variables can be defined and then used as the playbook is executed. Variables
// are stored in a dictionary where the key is the name of the variable and the
// value is a variable data type. Variables can represent stateful elements
// that may need to be captured to allow for the successful execution of the
// playbook. All playbook variables are mutable unless identified as a
// constant.
package variables
