// Copyright © 2018 github.com/devopsctl authors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	gitlab "github.com/xanzy/go-gitlab"
)

func TestDeleteGroup(t *testing.T) {
	setBasicAuthEnvs()
	tt := []struct {
		name   string
		args   []string
		expect testResult
	}{
		{
			name:   "delete an existing group",
			args:   []string{"GroupTBD"},
			expect: pass,
		},
		{
			name:   "deleting a non existent group should fail",
			args:   []string{"GroupUnknown"},
			expect: fail,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// SETUP:
			// create the test group to be deleted
			if tc.expect == pass {
				if _, err := newGroup(&gitlab.CreateGroupOptions{
					Path: gitlab.String(tc.args[0]),
					Name: gitlab.String(tc.args[0]),
				}); err != nil {
					tInfo("Test data setup failure for delete group")
					os.Exit(1)
				}
				tInfo("Test group to be deleted is created")
			}

			execT := execTestCmdFlags{
				t:    t,
				cmd:  deleteGroupCmd,
				args: tc.args,
			}
			stdout, execResult := execT.executeCommand()
			fmt.Println(stdout)
			require.Equal(t, tc.expect, execResult, stdout)
		})
	}
}
