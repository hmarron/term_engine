package term_engine

import (
	"os"
	"os/signal"
	"time"

	"github.com/hmarron/term_engine/term"

	"github.com/mattn/go-tty"
)

type Engine struct {
	drawables []Drawable
	keyFuncs  map[rune]func()
	tick      time.Duration
}

func NewEngine(tick time.Duration) *Engine {
	return &Engine{tick: tick}
}

func (e *Engine) Start() {
	// Prepare the terminal
	term.HideCursor()

	// listen for exit event
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			e.cleanup()
		}
	}()

	// listen for key presses
	go e.listenForKeyPress()

	// run the main loop
	for {
		e.Draw()
		time.Sleep(e.tick)
	}
}

func (e *Engine) Draw() {
	term.Clear()
	for _, d := range e.drawables {
		x, y, str := d.Draw()
		term.MoveCursor(x, y)
		term.Draw(str)
	}
	term.Render()
}

// Drawable should return x, y, string
type Drawable interface {
	Draw() (int, int, string)
}

func (e *Engine) AddDrawable(d Drawable) {
	e.drawables = append(e.drawables, d)
}

func (e *Engine) cleanup() {
	term.Clear()
	term.ShowCursor()

	term.MoveCursor(0, 0)
	term.Render()
	os.Exit(0)
}

func (e *Engine) SetKeyFunctions(keyFuncs map[rune]func()) {
	e.keyFuncs = keyFuncs
}

func (e *Engine) listenForKeyPress() {
	tty, err := tty.Open()
	if err != nil {
		panic(err)
	}
	defer tty.Close()

	for {
		char, err := tty.ReadRune()
		if err != nil {
			panic(err)
		}

		f, ok := e.keyFuncs[char]
		if ok {
			f()
		}
	}
}

func (e *Engine) GetMaxPos() (int, int) {
	mX, mY := term.GetSize()
	return mX, mY
}
