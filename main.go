package main

import(
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"flag"
)


type Button_State int

const(
	UP Button_State = iota
	HOVER
	DOWN
)



type Mouse struct{
	PosX int32
	PosY int32
	State uint32
}

type Button struct{
	//Texture sdl.Texture
	Rect	sdl.Rect
	State	Button_State
}

type Input struct{
	Rect	sdl.Rect
	State	Button_State
	Text	string
}



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
	SongDisplay *sdl.Surface
	viewport_size sdl.Rect
	surface *sdl.Surface
	window *sdl.Window
	renderer *sdl.Renderer
	loadRect sdl.Rect
	loadButton Button
	mouseData Mouse
	fileOpenerInput Input
)

func updateRendering(){

	surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}

	viewport_size = renderer.GetViewport()

	//Background Color
	surface.FillRect(nil, 0x332c2c)


	loadRect = sdl.Rect{0, viewport_size.H/10, viewport_size.W/2, viewport_size.H/5}
	surface.FillRect(&loadRect, 0xffff0000)

	fileOpenerInput.Rect = sdl.Rect{viewport_size.W/50, viewport_size.H/3, viewport_size.W/3, viewport_size.H/10}
	surface.FillRect(&fileOpenerInput.Rect, 0xffff0000)


}

func updateText(){

	if err = SongDisplay.Blit(nil, surface, &sdl.Rect{X: viewport_size.W/2 - (SongDisplay.W + viewport_size.W/10),
		Y: viewport_size.H - (SongDisplay.H + 50), W: 0, H: 0}); err != nil {
			return
	}

	if SongDisplay, err = font.RenderUTF8Blended("<"+SongLoaded+">", sdl.Color{R: 255, G: 0, B: 0, A: 255}); err != nil {
		return
	}


}


func mouseClick(){

	//loadButton interaction

	loadButton.Rect = loadRect
	println(mouseData.PosX, mouseData.PosY)


	if loadButton.Rect.HasIntersection(&sdl.Rect{X: mouseData.PosX, Y: mouseData.PosY, W: 1, H: 1}){
		loadButton.State = HOVER
		if mouseData.State == 1 {
			loadButton.State = DOWN
		}
	} else {
		loadButton.State = UP
	}

}



func main() {
	//Parse for audio file inputs
	InputFile := flag.String("f","Empty","Put path to MP3 as a flag for song autostart!")
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

	err = window.SetFullscreen(0)
	if err != nil {
		panic(err)
	}



	surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}


	//Load Hack-Regular Font
	if font, err = ttf.OpenFont(FONT_PATH, FONT_SIZE); err != nil {
		panic(err)
	}
	defer font.Close()


	if SongDisplay, err = font.RenderUTF8Blended("<"+SongLoaded+">", sdl.Color{R: 255, G: 0, B: 0, A: 255}); err != nil {
		panic(err)
	}
	defer SongDisplay.Free()


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

		//Reduces CPU consumption by decreasing the number of rendering cycles
		sdl.Delay(10)


		updateRendering()
		updateText()

		mouseData.PosX, mouseData.PosY, mouseData.State = sdl.GetMouseState()
		mouseClick()
		println(loadButton.State)


		window.UpdateSurface()
	}

}
