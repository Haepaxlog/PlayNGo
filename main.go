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
	fileOpenerDisplay *sdl.Surface
	viewport_size sdl.Rect
	surface *sdl.Surface
	window *sdl.Window
	renderer *sdl.Renderer
	loadButton Button
	mouseData Mouse
	fileOpenerInput Input
	pressedKeys []uint8
)

func updateRendering(){

	surface, err = window.GetSurface()
	if err != nil {
		return

	}

	viewport_size = renderer.GetViewport()

	//Background Color
	surface.FillRect(nil, 0x332c2c)


	loadButton.Rect = sdl.Rect{0, viewport_size.H/10, viewport_size.W/2, viewport_size.H/5}
	surface.FillRect(&loadButton.Rect, 0xffff0000)

	fileOpenerInput.Rect = sdl.Rect{viewport_size.W/50, viewport_size.H/3, viewport_size.W/3, viewport_size.H/10}
	surface.FillRect(&fileOpenerInput.Rect, 0xffff0000)

	fileOpenerInput.Display = sdl.Rect{int32(float32(fileOpenerInput.Rect.X) + float32(fileOpenerInput.Rect.X)*1.75), fileOpenerInput.Rect.Y,
		int32(float32(fileOpenerInput.Rect.W)*0.75), fileOpenerInput.Rect.H}
	surface.FillRect(&fileOpenerInput.Display, 0x3458eb)

}

func updateText(){

	//SongDisplay
	if SongDisplay, err = font.RenderUTF8Blended("<"+SongLoaded+">", sdl.Color{R: 255, G: 0, B: 0, A: 255}); err != nil {
		return
	}


	if err = SongDisplay.Blit(nil, surface, &sdl.Rect{X: viewport_size.W/2 - (SongDisplay.W + viewport_size.W/10),
		Y: viewport_size.H - (SongDisplay.H + 50), W: 0, H: 0}); err != nil {
		return
	}

	//fileOpenerDisplay
	if fileOpenerDisplay, err = font.RenderUTF8Blended(" " +fileOpenerInput.Text, sdl.Color{R: 0, G: 255, B: 0, A: 255}); err != nil {
		return
	}



	if err = fileOpenerDisplay.Blit(nil,surface, &sdl.Rect{fileOpenerInput.Display.X, int32(float32(fileOpenerInput.Display.Y) * 1.125),
		fileOpenerInput.Display.W, fileOpenerInput.Display.H }); err != nil {
		return
	}


}


func mouseClick(){

	//loadButton interaction
	println(mouseData.PosX, mouseData.PosY)

	if loadButton.Rect.HasIntersection(&sdl.Rect{X: mouseData.PosX, Y: mouseData.PosY, W: 1, H: 1}){
		loadButton.State = HOVER
		if mouseData.State == 1 {
			loadButton.State = DOWN
			loadButton.Pressed = true
		}
	} else {
		loadButton.State = UP
	}

	if loadButton.Pressed == true {
		println("loading music") //Insert dir open function
	}

	loadButton.Pressed = false

	//fileOpenerInput interaction
	if fileOpenerInput.Pressed == true {
		sdl.StartTextInput()
	} else {
		sdl.StopTextInput()
	}

	if fileOpenerInput.Rect.HasIntersection(&sdl.Rect{X: mouseData.PosX, Y: mouseData.PosY, W: 1, H: 1}){
		fileOpenerInput.State = HOVER
		if mouseData.State == 1 {
			fileOpenerInput.State = DOWN
			fileOpenerInput.Pressed = true
		}
	} else {
		fileOpenerInput.State = UP
	}

	if fileOpenerInput.State == UP && mouseData.State == 1{
		fileOpenerInput.Pressed = false
	}

	if pressedKeys[sdl.SCANCODE_ESCAPE] != 0 {
		fileOpenerInput.State = UP
		fileOpenerInput.Pressed = false
	}
}

func TrimLastElement(input string) string{
	length:= len(input)
	if (length - 1) >= 0 {
	return input[:(length - 1)]
	} else {
		return ""
	}
}

func checkWritable(input string) bool {
	length := len(input)
	charPosX := 0
	charPosX += int(PIXEL_FONT_SIZE * float32(length))
	if charPosX < int(fileOpenerInput.Rect.W) {
		return true
	}
	return false
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


	//	if SongDisplay, err = font.RenderUTF8Blended("<"+SongLoaded+">", sdl.Color{R: 255, G: 0, B: 0, A: 255}); err != nil {
	//	panic(err)
	//}
	//defer SongDisplay.Free()


	fileOpenerInput.Text = "aaa"

	//if fileOpenerDisplay, err = font.RenderUTF8Blended(fileOpenerInput.Text, sdl.Color{R: 0, G: 255, B: 0, A: 255}); err != nil {
	//	panic(err)
	//	}
	//defer fileOpenerDisplay.Free()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.TextInputEvent:
				input := t.GetText()
				if checkWritable(fileOpenerInput.Text) {
				fileOpenerInput.Text += input
				}
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_RETURN {
					println("return pressed")
				}
				if t.Keysym.Sym == sdl.K_BACKSPACE && t.State == sdl.PRESSED {
					fileOpenerInput.Text = TrimLastElement(fileOpenerInput.Text)
				}

			}

		}

		//Reduces CPU consumption by decreasing the number of rendering cycles
		sdl.Delay(16)

		pressedKeys = sdl.GetKeyboardState()

		updateRendering()
		updateText()




		mouseData.PosX, mouseData.PosY, mouseData.State = sdl.GetMouseState()
		mouseClick()
		println(loadButton.State)
		println(fileOpenerInput.State)


		if fileOpenerInput.Pressed == true {
			sdl.StartTextInput()
		} else {
			sdl.StopTextInput()
		}


		window.UpdateSurface()
	}

}
