// Copyright 2019-2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"fmt"

	"github.com/openplaybooks/libcacao/objects/playbook"
)

func main() {
	assignmentcommand(prep(playbook.New()))
	manualcommand(prep(playbook.New()))
	bashcommand(prep(playbook.New()))
	sshcommand(prep(playbook.New()))
	powershell1(prep(playbook.New()))
	powershell2(prep(playbook.New()))
	httpcommand1(prep(playbook.New()))
	httpcommand2(prep(playbook.New()))
	sigma(prep(playbook.New()))
	yara(prep(playbook.New()))
	kestrel(prep(playbook.New()))
	elastic(prep(playbook.New()))
}

func prep(p *playbook.Playbook) *playbook.Playbook {
	p.Created = ""
	p.Modified = ""
	return p
}

func assignmentcommand(p *playbook.Playbook) {

	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.2.1"
	step1.Description = "Example of an action step with an assignment command."

	// Create command
	cmd1, _ := step1.NewAssignmentCommand()
	cmd1.Description = "Copy various variables to new variables."
	cmd1.AddVariableAssignment("__new_variable_1__", "__existing_variable__")
	cmd1.AddVariableAssignment("__new_variable_2__", "__http--481b0ab0-df88-409e-ba58-7d77d7ffa4af.body__", "to_list", ",")

	header("assignment")
	encode(p)
}

func manualcommand(p *playbook.Playbook) {

	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.3.1"
	step1.Description = "Example of an action step with a manual command for an infected machine."

	// Create command
	cmd1, _ := step1.NewManualCommand()
	cmd1.Command = "Gather information about the infected machine."

	// Create four questions
	cmd1.AddQuestion("What OS is it running?", "string")
	cmd1.AddQuestion("What is the OS version?", "string")
	r3, _ := cmd1.AddQuestion("What is the IP address in ddd.ddd.ddd.ddd format?", "ipv4-addr")
	cmd1.AddQuestion("What is the MAC address in xx:xx:xx:xx:xx:xx format?", "mac-addr")

	// Create command
	cmd2, _ := step1.NewManualCommand()
	cmd2.Command = "Quarantine the machine to vlan __vlan_quarantine_number__:value."

	// Create one question
	cmd2.AddQuestion("What is the new IP address in ddd.ddd.ddd.ddd format?", "ipv4-addr")

	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	resid := r3.ID

	header("manual")
	encode(p)
	fmt.Printf("__%s.response__:value\n\n", resid)
}

func bashcommand(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.4.1"
	step1.Description = "Example of an action step with a bash command."
	cmd1, _ := step1.NewBashCommand()
	cmd1.Description = "View failed login attempts."
	cmd1.Command = "cat /var/log/auth.log | grep -i 'failed password'"

	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("bash")
	encode(p)
	fmt.Printf("__%s.stdout__:value\n\n", cmdid)
}

func sshcommand(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.5.1"
	step1.Description = "Example of an action step with an ssh command."
	cmd1, _ := step1.NewSSHCommand()
	cmd1.Description = "View failed login attempts."
	cmd1.Command = "cat /var/log/auth.log | grep -i 'failed password'"

	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("ssh")
	encode(p)
	fmt.Printf("__%s.stdout__:value\n\n", cmdid)
}

func powershell1(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.6.1"
	step1.Description = "Example of an action step with a powershell command."
	cmd1, _ := step1.NewPowerShellCommand()
	cmd1.Description = "Stop process"
	cmd1.Command = "Get-Process -ID 20496 | Stop-Process"

	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("powershell 1")
	encode(p)
	fmt.Printf("__%s.stdout__:value\n\n", cmdid)
}

func powershell2(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.6.2"
	step1.Description = "Example of an action step with a powershell command."
	cmd1, _ := step1.NewPowerShellCommand()
	cmd1.Description = "Disable Windows Defender"
	cmd1.CommandB64 = "IyAgIERlc2NyaXB0aW9uOgojIFRoaXMgc2NyaXB0IGRpc2FibGVzIFdpbmRvd3MgRGVmZW5kZXIuIFJ1biBpdCBvbmNlICh3aWxsIHRocm93IGVycm9ycyksIHRoZW4KIyByZWJvb3QsIHJ1biBpdCBhZ2FpbiAodGhpcyB0aW1lIG5vIGVycm9ycyBzaG91bGQgb2NjdXIpIGZvbGxvd2VkIGJ5IGFub3RoZXIKIyByZWJvb3QuCgpJbXBvcnQtTW9kdWxlIC1EaXNhYmxlTmFtZUNoZWNraW5nICRQU1NjcmlwdFJvb3RcLi5cbGliXE5ldy1Gb2xkZXJGb3JjZWQucHNtMQpJbXBvcnQtTW9kdWxlIC1EaXNhYmxlTmFtZUNoZWNraW5nICRQU1NjcmlwdFJvb3RcLi5cbGliXHRha2Utb3duLnBzbTEKCldyaXRlLU91dHB1dCAiRWxldmF0aW5nIHByaXZpbGVkZ2VzIGZvciB0aGlzIHByb2Nlc3MiCmRvIHt9IHVudGlsIChFbGV2YXRlLVByaXZpbGVnZXMgU2VUYWtlT3duZXJzaGlwUHJpdmlsZWdlKQoKJHRhc2tzID0gQCgKICAgICJcTWljcm9zb2Z0XFdpbmRvd3NcV2luZG93cyBEZWZlbmRlclxXaW5kb3dzIERlZmVuZGVyIENhY2hlIE1haW50ZW5hbmNlIgogICAgIlxNaWNyb3NvZnRcV2luZG93c1xXaW5kb3dzIERlZmVuZGVyXFdpbmRvd3MgRGVmZW5kZXIgQ2xlYW51cCIKICAgICJcTWljcm9zb2Z0XFdpbmRvd3NcV2luZG93cyBEZWZlbmRlclxXaW5kb3dzIERlZmVuZGVyIFNjaGVkdWxlZCBTY2FuIgogICAgIlxNaWNyb3NvZnRcV2luZG93c1xXaW5kb3dzIERlZmVuZGVyXFdpbmRvd3MgRGVmZW5kZXIgVmVyaWZpY2F0aW9uIgopCgpmb3JlYWNoICgkdGFzayBpbiAkdGFza3MpIHsKICAgICRwYXJ0cyA9ICR0YXNrLnNwbGl0KCdcJykKICAgICRuYW1lID0gJHBhcnRzWy0xXQogICAgJHBhdGggPSAkcGFydHNbMC4uKCRwYXJ0cy5sZW5ndGgtMildIC1qb2luICdcJwoKICAgIFdyaXRlLU91dHB1dCAiVHJ5aW5nIHRvIGRpc2FibGUgc2NoZWR1bGVkIHRhc2sgJG5hbWUiCiAgICBEaXNhYmxlLVNjaGVkdWxlZFRhc2sgLVRhc2tOYW1lICIkbmFtZSIgLVRhc2tQYXRoICIkcGF0aCIKfQoKV3JpdGUtT3V0cHV0ICJEaXNhYmxpbmcgV2luZG93cyBEZWZlbmRlciB2aWEgR3JvdXAgUG9saWNpZXMiCk5ldy1Gb2xkZXJGb3JjZWQgLVBhdGggIkhLTE06XFNPRlRXQVJFXFdvdzY0MzJOb2RlXFBvbGljaWVzXE1pY3Jvc29mdFxXaW5kb3dzIERlZmVuZGVyIgpTZXQtSXRlbVByb3BlcnR5IC1QYXRoICJIS0xNOlxTT0ZUV0FSRVxXb3c2NDMyTm9kZVxQb2xpY2llc1xNaWNyb3NvZnRcV2luZG93cyBEZWZlbmRlciIgIkRpc2FibGVBbnRpU3B5d2FyZSIgMQpTZXQtSXRlbVByb3BlcnR5IC1QYXRoICJIS0xNOlxTT0ZUV0FSRVxXb3c2NDMyTm9kZVxQb2xpY2llc1xNaWNyb3NvZnRcV2luZG93cyBEZWZlbmRlciIgIkRpc2FibGVSb3V0aW5lbHlUYWtpbmdBY3Rpb24iIDEKTmV3LUZvbGRlckZvcmNlZCAtUGF0aCAiSEtMTTpcU09GVFdBUkVcV293NjQzMk5vZGVcUG9saWNpZXNcTWljcm9zb2Z0XFdpbmRvd3MgRGVmZW5kZXJcUmVhbC1UaW1lIFByb3RlY3Rpb24iClNldC1JdGVtUHJvcGVydHkgLVBhdGggIkhLTE06XFNPRlRXQVJFXFdvdzY0MzJOb2RlXFBvbGljaWVzXE1pY3Jvc29mdFxXaW5kb3dzIERlZmVuZGVyXFJlYWwtVGltZSBQcm90ZWN0aW9uIiAiRGlzYWJsZVJlYWx0aW1lTW9uaXRvcmluZyIgMQoKV3JpdGUtT3V0cHV0ICJEaXNhYmxpbmcgV2luZG93cyBEZWZlbmRlciBTZXJ2aWNlcyIKVGFrZW93bi1SZWdpc3RyeSgiSEtFWV9MT0NBTF9NQUNISU5FXFNZU1RFTVxDdXJyZW50Q29udHJvbFNldFxTZXJ2aWNlc1xXaW5EZWZlbmQiKQpTZXQtSXRlbVByb3BlcnR5IC1QYXRoICJIS0xNOlxTWVNURU1cQ3VycmVudENvbnRyb2xTZXRcU2VydmljZXNcV2luRGVmZW5kIiAiU3RhcnQiIDQKU2V0LUl0ZW1Qcm9wZXJ0eSAtUGF0aCAiSEtMTTpcU1lTVEVNXEN1cnJlbnRDb250cm9sU2V0XFNlcnZpY2VzXFdpbkRlZmVuZCIgIkF1dG9ydW5zRGlzYWJsZWQiIDMKU2V0LUl0ZW1Qcm9wZXJ0eSAtUGF0aCAiSEtMTTpcU1lTVEVNXEN1cnJlbnRDb250cm9sU2V0XFNlcnZpY2VzXFdkTmlzU3ZjIiAiU3RhcnQiIDQKU2V0LUl0ZW1Qcm9wZXJ0eSAtUGF0aCAiSEtMTTpcU1lTVEVNXEN1cnJlbnRDb250cm9sU2V0XFNlcnZpY2VzXFdkTmlzU3ZjIiAiQXV0b3J1bnNEaXNhYmxlZCIgMwpTZXQtSXRlbVByb3BlcnR5IC1QYXRoICJIS0xNOlxTWVNURU1cQ3VycmVudENvbnRyb2xTZXRcU2VydmljZXNcU2Vuc2UiICJTdGFydCIgNApTZXQtSXRlbVByb3BlcnR5IC1QYXRoICJIS0xNOlxTWVNURU1cQ3VycmVudENvbnRyb2xTZXRcU2VydmljZXNcU2Vuc2UiICJBdXRvcnVuc0Rpc2FibGVkIiAzCgpXcml0ZS1PdXRwdXQgIlJlbW92aW5nIFdpbmRvd3MgRGVmZW5kZXIgY29udGV4dCBtZW51IGl0ZW0iClNldC1JdGVtICJIS0xNOlxTT0ZUV0FSRVxDbGFzc2VzXENMU0lEXHswOUE0Nzg2MC0xMUIwLTREQTUtQUZBNS0yNkQ4NjE5OEE3ODB9XElucHJvY1NlcnZlcjMyIiAiIgoKV3JpdGUtT3V0cHV0ICJSZW1vdmluZyBXaW5kb3dzIERlZmVuZGVyIEdVSSAvIHRyYXkgZnJvbSBhdXRvcnVuIgpSZW1vdmUtSXRlbVByb3BlcnR5ICJIS0xNOlxTT0ZUV0FSRVxNaWNyb3NvZnRcV2luZG93c1xDdXJyZW50VmVyc2lvblxSdW4iICJXaW5kb3dzRGVmZW5kZXIiIC1lYSAw"
	r, _ := cmd1.NewExternalReference()
	r.Name = "Disable Windows Defender"
	r.Description = "This script disables Windows Defender. Run it once (will throw errors), then reboot, run it again (this time no errors should occur) followed by another reboot."
	r.URL = "https://github.com/W4RH4WK/Debloat-Windows-10/blob/master/scripts/disable-windows-defender.ps1"
	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("powershell 2")
	encode(p)
	fmt.Printf("__%s.stdout__:value\n\n", cmdid)
}

func httpcommand1(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.7.1"
	step1.Description = "Example of an action step with an http GET command."
	cmd1, _ := step1.NewHTTPCommand()
	cmd1.Description = "Get current data for an ID"
	cmd1.Command = "GET /v1/getData?id=__some_data_id__:value HTTP/1.1"
	cmd1.AddHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) Chrome/109.0.0.0 Safari/537.36")

	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("http 1")
	encode(p)
	fmt.Printf("__%s.body__:value\n\n", cmdid)
}

func httpcommand2(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.7.2"
	step1.Description = "Example of an action step with an http POST command."
	cmd1, _ := step1.NewHTTPCommand()
	cmd1.Description = "Post data to endpoint"
	cmd1.Command = "POST /api1/newObjects/ HTTP/1.1"
	cmd1.AddHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) Chrome/109.0.0.0 Safari/537.36")
	cmd1.ContentB64 = "ewogICJ0eXBlIjogImhvdXNlIiwKICAidmFsdWUiOiAic29tZSBkYXRhIiwKICAuLi4KfQ=="

	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("http 2")
	encode(p)
	fmt.Printf("__%s.status__:value\n\n", cmdid)
}

func sigma(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.8.1"
	step1.Description = "Example of an action step with a sigma command."
	cmd1, _ := step1.NewSigmaCommand()
	cmd1.Description = "Detects a suspicious DLL load"
	cmd1.CommandB64 = "dGl0bGU6IFBvdGVudGlhbCBDT0xEU1RFRUwgUGVyc2lzdGVuY2UgU2VydmljZSBETEwgTG9hZAppZDogMWQ3YTU3ZGEtMDJlMC00ZjdmLTkyYjEtYzdiNDg2Y2NmZWQ1CnN0YXR1czogZXhwZXJpbWVudGFsCmRlc2NyaXB0aW9uOiB8CiAgICBEZXRlY3RzIGEgc3VzcGljaW91cyBETEwgbG9hZCBieSBhbiAic3ZjaG9zdCIgcHJvY2VzcyBiYXNlZCBvbiBsb2NhdGlvbiBhbmQgbmFtZSB0aGF0IG1pZ2h0IGJlIHJlbGF0ZWQgdG8gQ29sZFN0ZWVsIFJBVC4gVGhpcyBETEwgbG9jYXRpb24gYW5kIG5hbWUgaGFzIGJlZW4gc2VlbiB1c2VkIGJ5IENvbGRTdGVlbCBhcyB0aGUgc2VydmljZSBETEwgZm9yIGl0cyBwZXJzaXN0ZW5jZSBtZWNoYW5pc20KcmVmZXJlbmNlczoKICAgIC0gaHR0cHM6Ly93d3cubmNzYy5nb3YudWsvc3RhdGljLWFzc2V0cy9kb2N1bWVudHMvbWFsd2FyZS1hbmFseXNpcy1yZXBvcnRzL2NvbGQtc3RlZWwvTkNTQy1NQVItQ29sZC1TdGVlbC5wZGYKYXV0aG9yOiBOYXNyZWRkaW5lIEJlbmNoZXJjaGFsaSAoTmV4dHJvbiBTeXN0ZW1zKQpkYXRlOiAyMDIzLzA1LzAyCnRhZ3M6CiAgICAtIGF0dGFjay5wZXJzaXN0ZW5jZQogICAgLSBhdHRhY2suZGVmZW5zZV9ldmFzaW9uCiAgICAtIGRldGVjdGlvbi5lbWVyZ2luZ190aHJlYXRzCmxvZ3NvdXJjZToKICAgIHByb2R1Y3Q6IHdpbmRvd3MKICAgIGNhdGVnb3J5OiBpbWFnZV9sb2FkCmRldGVjdGlvbjoKICAgIHNlbGVjdGlvbjoKICAgICAgICBJbWFnZXxlbmRzd2l0aDogJ1xzdmNob3N0LmV4ZScKICAgICAgICBJbWFnZUxvYWRlZHxlbmRzd2l0aDogJ1xBcHBEYXRhXFJvYW1pbmdcbmV3ZGV2LmRsbCcKICAgIGNvbmRpdGlvbjogc2VsZWN0aW9uCmZhbHNlcG9zaXRpdmVzOgogICAgLSBVbmxpa2VseQpsZXZlbDogaGlnaA=="
	r, _ := cmd1.NewExternalReference()
	r.Name = "SigmaHQ Coldsteel Malware"
	r.Description = "Detects a suspicious DLL load by an \"svchost\" process based on location and name that might be related to ColdSteel RAT. This DLL location and name has been seen used by ColdSteel as the service DLL for its persistence mechanism."
	r.URL = "https://github.com/SigmaHQ/sigma/blob/master/rules-emerging-threats/2023/Malware/COLDSTEEL/image_load_malware_coldsteel_persistence_service_dll.yml"
	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("sigma")
	encode(p)
	fmt.Printf("__%s.stdout__:value\n\n", cmdid)
}

func yara(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.9.1"
	step1.Description = "Example of an action step with a yara command."
	cmd1, _ := step1.NewYaraCommand()
	cmd1.Description = "Detects files containing Badwinmail"
	cmd1.CommandB64 = "cnVsZSBleHBsb2l0X09mZmljZV9CYWR3aW5tYWlsIHsKICBtZXRhOgogICAgYXV0aG9yID0gIkRhdmlkIENhbm5pbmdzIgogICAgZGVzY3JpcHRpb24gPSAiU3BlY2lmaWMgcnVsZSB0byBkZXRlY3QgZmlsZXMgY29udGFpbmluZyBTV0Ygb2JqZWN0cywgZS5nLiBCYWR3aW5tYWlsIgogICAgcmVmID0gImh0dHBzOi8vY2Fuc2Vjd2VzdC5jb20vc2xpZGVzLzIwMTYvQ1NXMjAxNl9MaS1YdV9CYWRXaW5tYWlsX2FuZF9FbWFpbFNlY3VyaXR5T3V0bG9va19maW5hbC5wZGYiCiAgICAKICBzdHJpbmdzOgogICAgJGhlYWRlcl90bmVmID0geyA3OCA5RiAzRSAyMiB9CiAgICAkaGVhZGVyX2RvY2YgPSB7IEQwIENGIDExIEUwIH0KCiAgICAvLyBSZWR1Y2UgRlBzIG9uIG90aGVyIERPQ0YgZG9jdW1lbnRzIGJ5IHJlcXVpcmluZyBPdXRsb29rIHNwZWNpZmljIHByb3BlcnRpZXMKICAgIC8vIENvdWxkIGJlIGltcHJvdmVkIGJ5IHRha2luZyBmdXJ0aGVyIGl0ZW1zIGZyb20gTVMtT1hNU0cgc3BlY3MuCiAgICAkbXNnX3JlY2lwID0gIl9fcmVjaXBfdmVyc2lvbjEuMCIgd2lkZQogICAgJG1zZ19hdHRhY2ggPSAiX19hdHRhY2hfdmVyc2lvbjEuMCIgd2lkZQogICAgJG1zZ19wcm9wcyA9ICJfX3Byb3BlcnRpZXNfdmVyc2lvbiIgd2lkZQogICAgCiAgICAvLyBUT0RPOiBJcyB0aGVyZSBhbnkgcmVxdWlyZW1lbnQgdG8gc2lnbmF0dXJlIFJGQzgyMiBlbWFpbHM/CiAgICAKICAgIC8vIFNXRiBjbGFzcyBpZGVudGlmaWVycywgYXMgZW1iZWRkZWQgaW4gdGhlIGRvY3VtZW50CiAgICAkZW1iZWRkZWRfY2xzaWRfaGV4ID0gIkQyN0NEQjZFLUFFNkQtMTFjZi05NkI4LTQ0NDU1MzU0MDAwMCIgbm9jYXNlIHdpZGUgYXNjaWkKICAgICRlbWJlZGRlZF9jbGFzcyA9ICJvYmpjbGFzcyBTaG9ja3dhdmVGbGFzaC4iCiAgICAkZW1iZWRkZWRfY2xzaWQgPSB7IDZlIGRiIDdjIGQyIDZkIGFlIGNmIDExIDk2IGI4IDQ0IDQ1IDUzIDU0IDAwIDAwIH0KCiAgY29uZGl0aW9uOgogICAgKAogICAgICRoZWFkZXJfdG5lZiBhdCAwIG9yIAogICAgICgKICAgICAgJGhlYWRlcl9kb2NmIGF0IDAgYW5kCiAgICAgIDIgb2YgKCRtc2dfKikKICAgICApCiAgICApIGFuZCAKICAgIDEgb2YgKCRlbWJlZGRlZCopCn0="
	r, _ := cmd1.NewExternalReference()
	r.Name = "NCC Group Badwinmail Exploit"
	r.Description = "Specific rule to detect files containing SWF objects, e.g. Badwinmail."
	r.URL = "https://github.com/nccgroup/Cyber-Defence/blob/master/Signatures/yara/badwinmail.yara"
	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("yara")
	encode(p)
	fmt.Printf("__%s.stdout__:value\n\n", cmdid)
}

func kestrel(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.10.1"
	step1.Description = "Example of an action step with a kestrel command."
	cmd1, _ := step1.NewKestrelCommand()
	cmd1.Description = "Get network traffic and sort them by their destination port"
	cmd1.CommandB64 = "IyBnZXQgbmV0d29yayB0cmFmZmljIGFuZCBzb3J0IHRoZW0gYnkgdGhlaXIgZGVzdGluYXRpb24gcG9ydApudCA9IEdFVCBuZXR3b3JrLXRyYWZmaWMgRlJPTSBzdGl4c2hpZnRlcjovL2lkc1ggV0hFUkUgZHN0X3JlZl92YWx1ZSA9ICcxLjIuMy40JwpudHggPSBTT1JUIG50IEJZIGRzdF9wb3J0IEFTQwoKIyBkaXNwbGF5IGFsbCBkZXN0aW5hdGlvbiBwb3J0IGFuZCBub3cgaXQgaXMgZWFzeSB0byBjaGVjayBpbXBvcnRhbnQgcG9ydHMKRElTUCBudHggQVRUUiBkc3RfcG9ydA=="

	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("kestrel")
	encode(p)
	fmt.Printf("__%s.stdout__:value\n\n", cmdid)
}

func elastic(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.11.1"
	step1.Description = "Example of an action step with a elastic command."
	cmd1, _ := step1.NewElasticCommand()
	cmd1.Description = "Investigating remote system discovery commands"
	cmd1.CommandB64 = "cHJvY2VzcyB3aGVyZSBob3N0Lm9zLnR5cGUgPT0gIndpbmRvd3MiIGFuZCBldmVudC50eXBlID09ICJzdGFydCIgYW5kCiAgKChwcm9jZXNzLm5hbWUgOiAibmJ0c3RhdC5leGUiIGFuZCBwcm9jZXNzLmFyZ3MgOiAoIi1uIiwgIi1zIikpIG9yCiAgKHByb2Nlc3MubmFtZSA6ICJhcnAuZXhlIiBhbmQgcHJvY2Vzcy5hcmdzIDogIi1hIikgb3IKICAocHJvY2Vzcy5uYW1lIDogIm5sdGVzdC5leGUiIGFuZCBwcm9jZXNzLmFyZ3MgOiAoIi9kY2xpc3QiLCAiL2RzZ2V0ZGMiKSkgb3IKICAocHJvY2Vzcy5uYW1lIDogIm5zbG9va3VwLmV4ZSIgYW5kIHByb2Nlc3MuYXJncyA6ICIqX2xkYXAuX3RjcC5kYy4qIikgb3IKICAocHJvY2Vzcy5uYW1lOiAoImRzcXVlcnkuZXhlIiwgImRzZ2V0LmV4ZSIpIGFuZCBwcm9jZXNzLmFyZ3M6ICJzdWJuZXQiKSBvcgogICgoKChwcm9jZXNzLm5hbWUgOiAibmV0LmV4ZSIgb3IgcHJvY2Vzcy5wZS5vcmlnaW5hbF9maWxlX25hbWUgPT0gIm5ldC5leGUiKSBvcgogICAgKChwcm9jZXNzLm5hbWUgOiAibmV0MS5leGUiIG9yIHByb2Nlc3MucGUub3JpZ2luYWxfZmlsZV9uYW1lID09ICJuZXQxLmV4ZSIpIGFuZCBub3QKICAgICAgIHByb2Nlc3MucGFyZW50Lm5hbWUgOiAibmV0LmV4ZSIpKSBhbmQKICAgICAgIHByb2Nlc3MuYXJncyA6ICJncm91cCIgYW5kIHByb2Nlc3MuYXJncyA6ICIvZG9tYWluIiBhbmQgbm90IHByb2Nlc3MuYXJncyA6ICIvYWRkIikpKSBhbmQKICBub3QKICAoCiAgICAoCiAgICAgIHByb2Nlc3MubmFtZSA6ICJhcnAuZXhlIiBhbmQKICAgICAgcHJvY2Vzcy5wYXJlbnQuZXhlY3V0YWJsZSA6ICgKICAgICAgICAiPzpcXFByb2dyYW1EYXRhXFxDZW50cmFTdGFnZVxcQUVNQWdlbnRcXEFFTUFnZW50LmV4ZSIsCiAgICAgICAgIj86XFxQcm9ncmFtIEZpbGVzICh4ODYpXFxDaXRyaXhcXFdvcmtzcGFjZSBFbnZpcm9ubWVudCBNYW5hZ2VtZW50IEFnZW50XFxDaXRyaXguV2VtLkFnZW50LlNlcnZpY2UuZXhlIiwKICAgICAgICAiPzpcXFByb2dyYW0gRmlsZXMgKHg4NilcXExhbnN3ZWVwZXJcXFNlcnZpY2VcXExhbnN3ZWVwZXJTZXJ2aWNlLmV4ZSIKICAgICAgKQogICAgKQogICkK"
	r, _ := cmd1.NewExternalReference()
	r.Name = "Elastic"
	r.Description = "Remote system discovery commands."
	r.URL = "https://www.elastic.co/docs/reference/security/prebuilt-rules/rules_building_block/discovery_remote_system_discovery_commands_windows#investigating-remote-system-discovery-commands"

	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("elastic")
	encode(p)
	fmt.Printf("__%s.stdout__:value\n\n", cmdid)
}

// ----------------------------------------------------------------------
// Define Supporting Functions and Methods
// ----------------------------------------------------------------------

func header(t string) {
	fmt.Println("\n// ----------------------------------------")
	fmt.Printf("// %s command example", t)
	fmt.Println("\n// ----------------------------------------")
}

func encode(p *playbook.Playbook) {
	// Encode
	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
