package menu

type _colors struct {
	HEADER    string
	OKBLUE    string
	OKCYAN    string
	OKGREEN   string
	WARNING   string
	FAIL      string
	ENDC      string
	BOLD      string
	UNDERLINE string
}

var colors = &_colors{"\033[95m", "\033[94m", "\033[96m", "\033[92m", "\033[93m", "\033[91m", "\033[0m", "\033[1m", "\033[4m"}
