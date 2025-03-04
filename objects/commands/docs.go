// Copyright 2019-2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package commands implements the CACAO 3.0 command objects.
//
// The CACAO command object (command-data) contains detailed information about
// the commands that are to be executed automatically or manually as part of an
// action step (see section 7.6). Each command listed in an action step may be
// of a different command type, however, all commands listed in a single step
// MUST be processed or executed by all of the agents defined in that step.
//
// Commands can make use of variables just like other parts of the playbook.
package commands
