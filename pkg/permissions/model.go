package permissions

import "strings"

type Permissions string

const (
	PermissionsViewer    Permissions = "view"
	PermissionsSubmitter Permissions = "view,submit"
	PermissionsApprover  Permissions = "view,submit,approve"
)

func (p Permissions) ToString() string {
	return string(p)
}

func (p Permissions) ToArray() []string {
	return strings.Split(p.ToString(), ",")
}
