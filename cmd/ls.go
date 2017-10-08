// Copyright © 2017 John Laswell
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/jlaswell/conerror"
	notes "github.com/jlaswell/notes/src"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List current notes.",
	Long: `ls will provide listing of notes.

ls will provide a tree view of existing notes. You may
pass a path to limit the listing, simliar to passing a
pass to 'ls'.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			cerr     conerror.ConError
			notesDir string
		)
		if cerr, notesDir = notes.GetNotesDir(); cerr != nil {
			switch cerr.Get("reason") {
			case "user":
				log.Fatal("notes was unable to find a user")
			case "notesdir":
				fmt.Println("notes was unable to find a notes directory")
				if confirm("would you like to initialize notes?") {
					initializeNotes()
				}
			}
		}

		if len(args) > 0 {
			for _, innerDir := range args {
				out, err := exec.Command("tree", fmt.Sprintf("%s%s", notesDir, innerDir)).Output()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s\n", out)
			}
		} else {
			out, err := exec.Command("tree", notesDir).Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", out)
		}
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
