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
	"os/user"

	notes "github.com/jlaswell/notes/src"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize notes.",
	Long: `Initialize notes.

init will create a .notes directory in you home dir so that
you can start writing notes.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cerr, _ := notes.GetNotesDir(); cerr != nil {
			switch cerr.Get("reason") {
			case "user":
				log.Fatal("notes was unable to find a user")
			case "notesdir":
				initializeNotes()
				return
			}
		}
		fmt.Println("notes is already initialized")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initializeNotes() {
	u, err := user.Current()
	if err != nil {
		log.Fatal("error resolving current user")
	}
	notesDir := fmt.Sprintf("%s/.notes/", u.HomeDir)
	if _, err := exec.Command("mkdir", notesDir).Output(); err != nil {
		log.Fatal(err)
	}
}
