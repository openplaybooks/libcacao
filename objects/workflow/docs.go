// Copyright 2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package workflow implements the CACAO 3.0 workflow step objects.
//
// Workflows contain a series of steps that are stored in a dictionary (see the
// workflow property in section 6.1), where the key is the step ID and the value
// is a workflow step object. These workflow steps along with the associated
// agents and targets form the building blocks for playbooks and are used to
// control the commands that need to be processed. Workflows steps are processed
// either sequentially or in parallel. In addition to simple processing, workflow
// steps can also contain conditional logic, looping logic, and/or temporal
// processing to control the execution of the playbook.
//
// Conditional processing means executing steps or commands after some sort of
// condition is met. Loop processing means executing steps or commands for each
// element in a list of data. Temporal processing means executing steps or
// commands after some period of time has passed.
//
// This section defines the various workflow steps and how they may be used to
// define a playbook.
package workflow
