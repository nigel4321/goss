package resource

import "github.com/aelsabbahy/goss/system"

type User struct {
	Username string   `json:"-"`
	Exists   bool     `json:"exists"`
	UID      string   `json:"uid,omitempty"`
	GID      string   `json:"gid,omitempty"`
	Groups   []string `json:"groups,omitempty"`
	Home     string   `json:"home,omitempty"`
}

func (u *User) ID() string      { return u.Username }
func (u *User) SetID(id string) { u.Username = id }

func (u *User) Validate(sys *system.System) []TestResult {
	sysuser := sys.NewUser(u.Username, sys)

	var results []TestResult

	results = append(results, ValidateValue(u.ID(), "exists", u.Exists, sysuser.Exists))
	if !u.Exists {
		return results
	}
	results = append(results, ValidateValue(u.ID(), "uid", u.UID, sysuser.UID))
	results = append(results, ValidateValue(u.ID(), "gid", u.GID, sysuser.GID))
	results = append(results, ValidateValue(u.ID(), "home", u.Home, sysuser.Home))
	results = append(results, ValidateValues(u.ID(), "groups", u.Groups, sysuser.Groups))

	return results
}

func NewUser(sysUser system.User) *User {
	username := sysUser.Username()
	exists, _ := sysUser.Exists()
	uid, _ := sysUser.UID()
	gid, _ := sysUser.GID()
	groups, _ := sysUser.Groups()
	home, _ := sysUser.Home()
	return &User{
		Username: username,
		Exists:   exists.(bool),
		UID:      uid.(string),
		GID:      gid.(string),
		Groups:   groups,
		Home:     home.(string),
	}
}
