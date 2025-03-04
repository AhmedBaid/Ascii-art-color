package colors

// Définition des constantes ANSI pour les couleurs
const (
	RESET  = "\033[0m"
	BLACK  = "\033[30m"
	GRAY   = "\033[37m"
	RED    = "\033[31m"
	ORANGE = "\033[38;5;208m"
	CYAN   = "\033[36m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	PURPLE = "\033[35m"
	WHITE  = "\033[37m"
)

// Map globale des couleurs
var ColorMap = map[string]string{
	"black":  BLACK,
	"gray":   GRAY,
	"red":    RED,
	"orange": ORANGE,
	"cyan":   CYAN,
	"green":  GREEN,
	"yellow": YELLOW,
	"blue":   BLUE,
	"purple": PURPLE,
	"white":  WHITE,
	"reset":  RESET,
}
