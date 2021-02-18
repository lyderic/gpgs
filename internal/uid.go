package internal

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bitfield/script"
)

type Keyinfo struct {
	Grip   string
	Status string
}

func (k Keyinfo) String() string {
	return fmt.Sprintf("%s [%s]", k.Grip, k.Status)
}

func GetKeyinfos() (keyinfos []Keyinfo) {
	listing := script.Exec(`gpg-connect-agent "keyinfo --list" /bye`)
	listing.EachLine(func(line string, output *strings.Builder) {
		if strings.HasPrefix(line, "S KEYINFO") {
			var keyinfo Keyinfo
			bits := strings.Fields(line)
			keyinfo.Grip = bits[2]
			keyinfo.Status = bits[6]
			keyinfos = append(keyinfos, keyinfo)
		}
	})
	return
}

func GetGrips(uid string) (grips []string) {
	command := "gpg --list-secret-keys --with-keygrip --with-colons " + uid
	output := script.Exec(command).MatchRegexp(regexp.MustCompile("^grp"))
	output.EachLine(func(line string, out *strings.Builder) {
		grips = append(grips, strings.Split(line, ":")[9])
	})
	return
}

func GetUids() (uids []string) {
	pipe := script.Exec("gpg --list-keys --with-colons").
		MatchRegexp(regexp.MustCompile("^uid"))
	pipe.EachLine(func(line string, out *strings.Builder) {
		fullUid := strings.Split(line, ":")[9]
		uids = append(uids, strings.Fields(fullUid)[0])
	})
	return
}
