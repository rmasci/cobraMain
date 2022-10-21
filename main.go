/*
* ============LICENSE_START=======================================================
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
* ============LICENSE_END=========================================================

*/
package main

import (
	"codecloud.web.att.com/users/rx7322/tsend/cmd"
	"os"
)

func main() {
	defaultCmd("push")
	cmd.Execute()
}

func defaultCmd(dcmd string) {
	commands:=[]string{"push", "web"}
	// if length of os.Args is 1 or less -- return user didn't supply argument
	if len(os.Args) > 1 {
		// if first arg is 'help' return
		if os.Args[1] == "--help" {
			return
		}
	} else {
		// os.Args is 1 or less
		return
	}

	for _,osargs := range os.Args {
		for _,cmd := range commands {
			if cmd == osargs {
				return
			}
		}
	}
	// At this point we inject a default
	os.Args = append(os.Args[:2], os.Args[1:]...)
	os.Args[1] = dcmd
}
