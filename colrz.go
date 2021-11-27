package colrz

import "fmt"

const (
	fgBack = 30 + iota
	fgRed
	fgGreen
	fgYellow
	fgBlue
	fgMagenta
	fgCyan
	fgWhite
	_
	fgReset
)

const (
	FgBack    = "\033[30m"
	FgRed     = "\033[31m"
	FgGreen   = "\033[32m"
	FgYellow  = "\033[33m"
	FgBlue    = "\033[34m"
	FgMagenta = "\033[35m"
	FgCyan    = "\033[36m"
	FgWhite   = "\033[37m"

	FgReset = "\033[39m"
)

const (
	bgBack = 40 + iota
	bgRed
	bgGreen
	bgYellow
	bgBlue
	bgMagenta
	bgCyan
	bgWhite
	_
	bgReset
)

const (
	BgBack    = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"

	BgReset = "\033[49m"
)

func Back(s string) string {
	return setColor(fgBack, s) + FgReset
}
func Red(s string) string {
	return setColor(fgRed, s) + FgReset
}
func Green(s string) string {
	return setColor(fgGreen, s) + FgReset
}
func Yellow(s string) string {
	return setColor(fgYellow, s) + FgReset
}
func Blue(s string) string {
	return setColor(fgBlue, s) + FgReset
}
func Magenta(s string) string {
	return setColor(fgMagenta, s) + FgReset
}
func Cyan(s string) string {
	return setColor(fgCyan, s) + FgReset
}
func White(s string) string {
	return setColor(fgWhite, s) + FgReset
}

func setColor(c int, s string) string {
	s2 := fmt.Sprintf("\033[%dm%s\033[%dm", c, s, c)
	return s2
}

func BackBg(s string) string {
	return setColorBg(bgBack, s) + BgReset
}
func RedBg(s string) string {
	return setColorBg(bgRed, s) + BgReset
}
func GreenBg(s string) string {
	return setColorBg(bgGreen, s) + BgReset
}
func YellowBg(s string) string {
	return setColorBg(bgYellow, s) + BgReset
}
func BlueBg(s string) string {
	return setColorBg(bgBlue, s) + BgReset
}
func MagentaBg(s string) string {
	return setColorBg(bgMagenta, s) + BgReset
}
func CyanBg(s string) string {
	return setColorBg(bgCyan, s) + BgReset
}
func WhiteBg(s string) string {
	return setColorBg(bgWhite, s) + BgReset
}

func setColorBg(c int, s string) string {
	s2 := fmt.Sprintf("\033[%dm%s\033[%dm", c, s, c)
	return s2
}
