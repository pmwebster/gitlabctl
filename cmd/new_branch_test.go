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
	"testing"
)

func TestNewBranch(t *testing.T) {
	tt := []struct {
		name     string
		flagsMap map[string]string
		args     []string
		expect   testResult
	}{
		{
			name: "create a new branch",
			flagsMap: map[string]string{
				"project": "Group2/project12",
				"ref":     "master",
			},
			args:   []string{"release-xy"},
			expect: pass,
		},
		{
			name:   "flags must be required",
			args:   []string{"master"},
			expect: fail,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			execT := execTestCmdFlags{
				t:        t,
				cmd:      newBranchCmd,
				flagsMap: tc.flagsMap,
				args:     tc.args,
			}
			stdout, execResult := execT.executeCommand()
			fmt.Println(stdout)
			assertEqualResult(t, execResult, tc.expect, printFlagsTable(tc.flagsMap, stdout))
			if tc.expect == pass {
				assertOutContains(t, stdout, tc.args[0])
				if err := deleteBranch(tc.flagsMap["project"], tc.args[0]); err != nil {
					tInfo(err)
				}
			}
		})
	}
}
