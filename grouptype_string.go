// generated by stringer -type=GroupType; DO NOT EDIT

package vk

import "fmt"

const _GroupType_name = "GroupOpenGroupClosedGroupPrivate"

var _GroupType_index = [...]uint8{0, 9, 20, 32}

func (i GroupType) String() string {
	if i < 0 || i+1 >= GroupType(len(_GroupType_index)) {
		return fmt.Sprintf("GroupType(%d)", i)
	}
	return _GroupType_name[_GroupType_index[i]:_GroupType_index[i+1]]
}
