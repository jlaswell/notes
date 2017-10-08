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
	"os"
	"os/exec"

	"github.com/jlaswell/conerror"
	notes "github.com/jlaswell/notes/src"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a note.",
	Long:  `Delete an existing note.`,
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

		switch len(args) {
		case 0:
			log.Fatal("notes needs the path to a note to delete")
		case 1:
			break
		default:
			log.Fatal("notes can only delete one note at a time")
		}

		// touch the path to create the file if not there
		path := fmt.Sprintf("%s%s", notesDir, args[0])
		if _, err := os.Stat(path); err != nil {
			log.Fatal("notes couldn't find this note, or there was an issue when resolving it's path")
		}

		if confirm("Are you sure you'd like to remove this note?") {
			if _, err := exec.Command("rm", path).Output(); err != nil {
				log.Fatal(err)
			}

		}
	},
}

func init() {
	RootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
