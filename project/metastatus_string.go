// Code generated by "stringer -type=MetaStatus"; DO NOT EDIT

package project

import "fmt"

const _MetaStatus_name = "StatusLoadingStatusUnloadingStatusUpdatingStatusActiveStatusFailedStatusUnknown"

var _MetaStatus_index = [...]uint8{0, 13, 28, 42, 54, 66, 79}

func (i MetaStatus) String() string {
	if i < 0 || i >= MetaStatus(len(_MetaStatus_index)-1) {
		return fmt.Sprintf("MetaStatus(%d)", i)
	}
	return _MetaStatus_name[_MetaStatus_index[i]:_MetaStatus_index[i+1]]
}