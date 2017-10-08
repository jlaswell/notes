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

package notes

import (
	"fmt"
	"os"
	"os/user"

	"github.com/jlaswell/conerror"
)

func GetNotesDir() (conerror.ConError, string) {
	u, err := user.Current()
	if err != nil {
		cerr := conerror.NewFromError(err)
		cerr.Set("reason", "user")
		return cerr, ""
	}
	notesDir := fmt.Sprintf("%s/.notes/", u.HomeDir)
	if _, err = os.Stat(notesDir); err != nil {
		cerr := conerror.NewFromError(err)
		cerr.Set("reason", "notesdir")
		return cerr, ""
	}
	return nil, notesDir
}
