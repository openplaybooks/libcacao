// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package workflow implements the CACAO 2.0 workflow step objects.
//
// Workflows contain a series of steps that are stored in a dictionary (see the
// workflow property in section 3.1), where the key is the step ID and the
// value is a workflow step. These workflow steps along with the associated
// commands form the building blocks for playbooks and are used to control the
// commands that need to be executed. Workflows steps are processed either
// sequentially, in parallel, or both depending on the type of steps required
// by the playbook. In addition to simple processing, workflow steps MAY also
// contain conditional and/or temporal operations to control the execution of
// the playbook.
//
// Conditional processing means executing steps or commands after some sort of
// condition is met. Temporal processing means executing steps or commands
// either during a certain time window or after some period of time has
// passed.
//
// This section defines the various workflow steps and how they may be used to
// define a playbook.
package workflow
