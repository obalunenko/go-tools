// Code generated by "stringer -type=TargetType"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TargetTypeCaarlos0-0]
	_ = x[TargetTypeCleanenv-1]
}

const _TargetType_name = "TargetTypeCaarlos0TargetTypeCleanenv"

var _TargetType_index = [...]uint8{0, 18, 36}

func (i TargetType) String() string {
	if i < 0 || i >= TargetType(len(_TargetType_index)-1) {
		return "TargetType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TargetType_name[_TargetType_index[i]:_TargetType_index[i+1]]
}
