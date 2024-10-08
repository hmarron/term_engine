package term

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

var screen = bufio.NewWriter(os.Stdout)

func HideCursor() {
	fmt.Fprint(screen, "\033[?25l")
}

func ShowCursor() {
	fmt.Fprint(screen, "\033[?25h")
}

func Clear() {
	fmt.Fprint(screen, "\033[2J")
}

func Draw(str string) {
	fmt.Fprint(screen, str)
}

func Render() {
	screen.Flush()
}

func GetSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}
	return width, height
}

func MoveCursor(x, y int) {
	fmt.Fprintf(screen, "\033[%d;%dH", y, x)
}
