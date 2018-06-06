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
	"strconv"

	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

var editUserCmd = &cobra.Command{
	Use:        "user",
	Aliases:    []string{"u"},
	SuggestFor: []string{"users"},
	Short:      "Modify a user",
	Example: `# modify a user bio
gitlabctl edit user --user=john.smith --bio="frontend devloper"`,
	Args:          cobra.ExactArgs(0),
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runEditUser(cmd)
	},
}

func init() {
	editCmd.AddCommand(editUserCmd)
	addEditUserFlags(editUserCmd)
}

func runEditUser(cmd *cobra.Command) error {
	opts, err := assignEditUserOptions(cmd)
	if err != nil {
		return err
	}
	user := getFlagString(cmd, "user")
	uid, err := strconv.Atoi(user)
	// if user is not a number,
	// search for the username's user id and assign it to uid
	if err != nil {
		u, err := getUserByUsername(user)
		if err != nil {
			return err
		}
		uid = u.ID
	}
	editedUser, err := editUser(uid, opts)
	if err != nil {
		return err
	}
	printUsersOut(cmd, editedUser)
	return nil
}

func editUser(uid int, opt *gitlab.ModifyUserOptions) (*gitlab.User, error) {
	git, err := newGitlabClient()
	if err != nil {
		return nil, err
	}
	user, _, err := git.Users.ModifyUser(uid, opt, nil)
	if err != nil {
		return nil, err
	}
	return user, nil
}