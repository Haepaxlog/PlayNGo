package main

import(
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"flag"
	"github.com/Haepaxlog/playngo/ui"
)


type Button_State int

const(
	UP Button_State = iota
	HOVER
	DOWN
)


type Button struct{
	Rect	sdl.Rect
	State	Button_State
	Pressed bool
}

type Input struct{
	Rect	sdl.Rect
	State	Button_State
	Text	string
	Pressed bool
	Display sdl.Rect
}



const(
	INITIAL_WINDOW_WIDTH = 1920
	INITIAL_WINDOW_HEIGHT = 1080
	FONT_SIZE = 22
	PIXEL_FONT_SIZE = FONT_SIZE * (72/69)
	FONT_PATH = "./fonts/Hack-Regular.ttf"
)

var(
	err error
	SongLoaded string
	font *ttf.Font
	SongDisplay *sdl.Surface
	FileOpenerDisplay *sdl.Surface
	viewport_size sdl.Rect
	surface *sdl.Surface
	window *sdl.Window
	renderer *sdl.Renderer
	LoadButton *ui.Button
	mouseData ui.Mouse
	FileOpenerInput *ui.Input
	pressedKeys []uint8
	FileOpenerData *ui.InputData
)

func initUI(){
	LoadButton = ui.CreateButton(renderer, sdl.Rect{0, 100, 1000, 200})

	FileOpenerInput = ui.CreateInput(renderer, sdl.Rect{50, 350, 800, 150})
}

func updateRendering(){

	updateText()

	renderer.Clear()

	viewport_size = renderer.GetViewport()

	SongDisplayRect := sdl.Rect{viewport_size.W/2 - (SongDisplay.W + viewport_size.W/10),
		viewport_size.H - (SongDisplay.H + 50) , SongDisplay.W , SongDisplay.H}

	SongDisplayTexture, _ := renderer.CreateTextureFromSurface(SongDisplay)

	//Background
	renderer.SetDrawColor(0x33,0x2c,0x2c,0xFF)
	renderer.FillRect(&viewport_size)
	//loadButton
	LoadButton.Render(renderer, sdl.Color{0xFF,0x00,0x00,0xFF})
	//FileOpener
	FileOpenerInput.Render(renderer, sdl.Color{0xFF,0x00,0x00,0xFF}, sdl.Color{0x00,0x00,0xFF,0xFF}, FileOpenerDisplay)

	/*TextRendering*/
	//SongDisplay
	renderer.Copy(SongDisplayTexture, nil, &SongDisplayRect)

	renderer.Present()


	/*	loadButton.Rect = sdl.Rect{0, viewport_size.H/10, viewport_size.W/2, viewport_size.H/5}
	surface.FillRect(&loadButton.Rect, 0xffff0000)

	fileOpenerInput.Rect = sdl.Rect{viewport_size.W/50, viewport_size.H/3, viewport_size.W/3, viewport_size.H/10}
	surface.FillRect(&fileOpenerInput.Rect, 0xffff0000)

	fileOpenerInput.Display = sdl.Rect{int32(float32(fileOpenerInput.Rect.X) + float32(fileOpenerInput.Rect.X)*1.75), fileOpenerInput.Rect.Y,
		int32(float32(fileOpenerInput.Rect.W)*0.75), fileOpenerInput.Rect.H}
	surface.FillRect(&fileOpenerInput.Display, 0x3458eb)
	*/
}

func updateText(){

	if SongDisplay, err = font.RenderUTF8Blended("<"+SongLoaded+">", sdl.Color{R: 255, G: 0, B: 0, A: 255}); err != nil {
		return
	}


	if FileOpenerDisplay, err = font.RenderUTF8Blended(" " +FileOpenerData.Text, sdl.Color{R: 0, G: 255, B: 0, A: 255}); err != nil {
		return
	}

	println(FileOpenerDisplay.W, FileOpenerDisplay.H)

}


func mouseClick(){

	println(mouseData.X, mouseData.Y)
	LoadButton.CheckState(&mouseData)

	if LoadButton.Pressed == true {
		println("loading music") //Insert dir open function
	}
	LoadButton.Pressed = false


	FileOpenerData.CheckState(FileOpenerInput, &mouseData, pressedKeys)
}

func TrimLastElement(input string) string{
	length:= len(input)
	if (length - 1) >= 0 {
	return input[:(length - 1)]
	} else {
		return ""
	}
}

func checkWritable(input string, rectWidth int32) bool {
	length := len(input)
	charPosX := int32(PIXEL_FONT_SIZE * float32(length))
	if charPosX < rectWidth {
		return true
	}
	return false
}



func main() {
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
	defer surface.Free()


	if font, err = ttf.OpenFont(FONT_PATH, FONT_SIZE); err != nil {
		panic(err)
	}
	defer font.Close()

	FileOpenerData = ui.InitInputData()

	running := true
	for running {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.TextInputEvent:
				println("activated")
				if checkWritable(FileOpenerData.Text,FileOpenerInput.Rect.W) {
				FileOpenerData.Text += t.GetText()
				}
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_RETURN {
					println("return pressed")
				}
				if t.Keysym.Sym == sdl.K_BACKSPACE && t.State == sdl.PRESSED {
					FileOpenerData.Text = TrimLastElement(FileOpenerData.Text)
				}

			}

		}

		initUI()

		updateRendering()

		pressedKeys = sdl.GetKeyboardState()
		mouseData.X, mouseData.Y, mouseData.State = sdl.GetMouseState()
		mouseClick()
		println(LoadButton.State)
		println(FileOpenerData.State)

		if FileOpenerData.Pressed == true {
			sdl.StartTextInput()
		} else {
			sdl.StopTextInput()
		}

		//Reduces CPU consumption by decreasing the number of rendering cycles (FPS)
		sdl.Delay(16)

	}

}
