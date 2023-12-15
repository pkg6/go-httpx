package regexpx

import "regexp"

func IsMatchString(reg, s string) bool {
	return regexp.MustCompile(reg).MatchString(s)
}
