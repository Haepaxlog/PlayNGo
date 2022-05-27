package main

import(
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"flag"
	"fmt"
)

const(
	INITIAL_WINDOW_WIDTH = 1920
	INITIAL_WINDOW_HEIGHT = 1080
	FONT_SIZE = 32
	FONT_PATH = "./fonts/Hack-Regular.ttf"
)

var(
	err error
	SongLoaded string
	font *ttf.Font
	text *sdl.Surface
	viewport_size sdl.Rect
	surface *sdl.Surface
	window *sdl.Window
	renderer *sdl.Renderer
	rect sdl.Rect
)

func update(){

	surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}

	viewport_size = renderer.GetViewport()

	//Background Color
	surface.FillRect(nil, 0x332c2c)



	if err = text.Blit(nil, surface, &sdl.Rect{X: viewport_size.W - (text.W + 200),
		Y: 200 - (text.H), W: 0, H: 0}); err != nil {
			panic(err)
	}


	rect = sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()


}



func main() {

	//Parse for audio file inputs
	InputFile := flag.String("f","Empty","Put filename of mp3")
	flag.Parse()
	SongLoaded = *InputFile


	if err = ttf.Init(); err != nil {
		panic(err)
	}
	defer ttf.Quit()


	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()


	window, err = sdl.CreateWindow("Play 'N Go", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	//Required by WMs
	window.SetResizable(true)

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	width, height, err := renderer.GetOutputSize()
	if err != nil {
		panic(err)
	}
	fmt.Println(width, height)

	viewport_size = renderer.GetViewport()
	err = window.SetFullscreen(0)



	surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}

	//Background Color
	surface.FillRect(nil, 0x332c2c)

	if font, err = ttf.OpenFont(FONT_PATH, FONT_SIZE); err != nil {
		panic(err)
	}
	defer font.Close()

	// Create a red text with the font
	if text, err = font.RenderUTF8Blended(SongLoaded, sdl.Color{R: 255, G: 0, B: 0, A: 255}); err != nil {
		panic(err)
	}
	defer text.Free()

	if err = text.Blit(nil, surface, &sdl.Rect{X: viewport_size.W - (text.W + 200),
			Y: 200 - (text.H), W: 0, H: 0}); err != nil {
				panic(err)
	}


	rect = sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break

			}

		 }
		update()
		}

	}
