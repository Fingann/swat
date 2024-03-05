package flags

import "fmt"

var (
	BaseFlag       = "geodata"
	InjectionFlag  = GetFlag("SQL_injections_are_B4D")
	HardCodedFlag  = GetFlag("H4rdc0ded_S3cr3t")
	TraversalFlag  = GetFlag("all_your_base_are_belong_to_us")
	CmonsterFlag   = GetFlag("cookie_conster_stea1s_all_your_c00kies")
	XssFlag        = GetFlag("XSS_is_just_the_beginning")
	SecurityFlag   = GetFlag("s3cur1ty_1s_just_4n_1llus10n")
	ValidationFlag = GetFlag("V4l1d4t10n_1s_4_m1r4ge")

	FlagTaskNameMap = map[string]string{
		InjectionFlag:  "Injection",
		XssFlag:        "XSS",
		TraversalFlag:  "Path Traversal",
		HardCodedFlag:  "Hardcoded",
		ValidationFlag: "Validation",
		CmonsterFlag:   "Cmonster",
		SecurityFlag:   "Security",
	}
)

func CheckFlag(flag string) (string, bool) {
	flag, ok := FlagTaskNameMap[flag]
	return flag, ok
}

func GetFlag(flag string) string {
	return fmt.Sprintf("%s{%s}", BaseFlag, flag)
}
