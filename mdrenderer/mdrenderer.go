package mdrenderer

const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	italic = "\033[3m"
	blue   = "\033[1;34m"
	green  = "\033[32m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	yellow = "\033[33m"
)

func Print(text string) {
    print(bold + cyan + text)
}
