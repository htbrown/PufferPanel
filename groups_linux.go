/*
 Copyright 2020 Padduck, LLC
  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at
  	http://www.apache.org/licenses/LICENSE-2.0
  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package pufferpanel

import (
	"fmt"
	"os/user"
)

func UserInGroup(groups ...string) bool {
	//add root as an allowed group
	groups = append(groups, "root")

	u, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	allowedIds := make([]string, 0)
	for _, v := range groups {
		rootGroup, err := user.LookupGroup(v)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			allowedIds = append(allowedIds, rootGroup.Gid)
		}
	}

	g, err := u.GroupIds()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	for _, v := range g {
		for _, t := range allowedIds {
			if v == t {
				return true
			}
		}
	}

	return false
}
