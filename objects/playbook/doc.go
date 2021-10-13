// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package playbook implements the CACAO 1.1 playbook object. CACAO playbooks
// are made up of six parts; playbook metadata, the workflow logic, a list of
// targets, a list of extensions, a list of data markings, and a list of
// signatures. Playbooks MAY refer to other playbooks in the workflow, similar
// to how programs refer to function calls or modules that comprise the
// program.
package playbook
