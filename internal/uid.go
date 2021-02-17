package internal

import (
	"regexp"
	"strings"

	"github.com/bitfield/script"
)

func GetUids() (uids []string) {
	pipe := script.Exec("gpg --list-keys --with-colons").
		MatchRegexp(regexp.MustCompile("^uid"))
	pipe.EachLine(func(line string, out *strings.Builder) {
		fullUid := strings.Split(line, ":")[9]
		uids = append(uids, strings.Fields(fullUid)[0])
	})
	return
}
