package main

import (
	"fmt"
	"runtime"

	graphics "github.com/teh-cmc/go-sfml/CSFML/include/SFML/Graphics"
	window "github.com/teh-cmc/go-sfml/CSFML/include/SFML/Window"
)

const title = "some window title"

func main() {
	runtime.LockOSThread()

	vm := window.NewSfVideoMode()
	vm.SetWidth(800)
	vm.SetHeight(600)
	cs := window.NewSfContextSettings()
	w := window.SfWindow_create(vm, title, 32, cs)

	e := window.NewSfEvent()

	for window.SfWindow_isOpen(w) > 0 {
		for window.SfWindow_pollEvent(w, e) > 0 {
			fmt.Println(e)
		}
		graphics.SfRenderWindow_clear(w, graphics.GetSfBlack())
	}
}
