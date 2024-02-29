package flags

import "fmt"

var (
	BaseFlag      = "geodata"
	InjectionFlag = GetFlag("SQL_injections_are_B4D")
	HardCodedFlag = GetFlag("H4rdc0ded_S3cr3t")
	TraversalFlag = GetFlag("all_your_base_are_belong_to_us")
	CmonsterFlag  = GetFlag("cookie_conster_stea1s_all_your_c00kies")
	XssFlag       = GetFlag("XSS_is_just_the_beginning")
	SecurityFlag  = GetFlag("s3cur1ty_1s_just_4n_1llus10n")
)

func GetFlag(flag string) string {
	return fmt.Sprintf("%s{%s}", BaseFlag, flag)
}
