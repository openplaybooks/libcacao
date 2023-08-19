// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package markings implements the CACAO 2.0 data marking definition objects.
//
// CACAO data marking definition objects contain detailed information about a
// specific data marking. Data markings typically represent handling or sharing
// requirements and are applied via the markings property in a playbook.
//
// Data marking objects MUST NOT be versioned because it would allow for
// indirect changes to the markings on a playbook. For example, if a statement
// marking definition is changed from "Reuse Allowed" to "Reuse Prohibited",
// all playbooks marked with that statement marking definition would
// effectively have an updated marking without being updated themselves.
// Instead, in this example a new statement marking definition with the new
// text should be created and the marked objects updated to point to the new
// data marking object.
//
// Playbooks may be marked with multiple marking statements. In other words, the
// same playbook can be marked with both a statement saying "Copyright 2020"
// and a statement saying, "Terms of use are ..." and both statements apply.
// This specification does not define rules for how multiple markings applied
// to the same object should be interpreted.
package markings
