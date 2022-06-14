package main

import(
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"flag"
	"fmt"
	"os"
	"github.com/Haepaxlog/playngo/ui"
	"github.com/Haepaxlog/playngo/lib"
)

type Button_State int

const(
	UP Button_State = iota
	HOVER
	DOWN
)


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
	viewport sdl.Rect
	surface *sdl.Surface
	window *sdl.Window
	renderer *sdl.Renderer
	LoadButton *ui.Button
	mouseData *ui.Mouse
	FileOpenerInput *ui.Input
	SongPresenter	*ui.Display
	pressedKeys []uint8
	FileOpenerData *ui.InputData
	testListbox	*ui.Listbox
	testListboxData *ui.Listbox
)

func initUI(){
	LoadButton = ui.CreateButton(renderer, sdl.Rect{0, 100, 1000, 200})
	FileOpenerInput = ui.CreateInput(renderer, sdl.Rect{50, 350, 800, 150})
	SongPresenter = ui.CreateDisplay(renderer, sdl.Rect{800, 1000, 0, 0}, SongLoaded)
	testListbox = ui.CreateListbox(renderer, testListboxData, sdl.Rect{0, 600, 1000, 300}, []string{"Hello","Hi","Hey","Ahoi", "Morgen","Tag","Moin","Guten","Ge"})
}

func updateRendering(){
	if err = updateText(); err != nil {
		fmt.Fprintf(os.Stderr, "updateText: %q", err)
	}

	renderer.Clear()
	viewport = renderer.GetViewport()

	//SongDisplayRect := sdl.Rect{viewport.W/2 - (SongDisplay.W + viewport.W/10),
	//	viewport.H - (SongDisplay.H + 50) , SongDisplay.W , SongDisplay.H}

	//Background
	renderer.SetDrawColor(0x33,0x2c,0x2c,0xFF)
	renderer.FillRect(&viewport)
	//loadButton
	LoadButton.Render(renderer, sdl.Color{0xFF,0x00,0x00,0xFF})
	//FileOpener
	if err := FileOpenerInput.Render(renderer, sdl.Color{0xFF,0x00,0x00,0xFF}, sdl.Color{0x00,0x00,0xFF,0xFF}, FileOpenerDisplay); err != nil {
		fmt.Fprintf(os.Stderr, "FileOpenerInput.Render: %q", err)
	}
	//SongDisplay
	if err := SongPresenter.Render(renderer, SongDisplay); err != nil {
		fmt.Fprintf(os.Stderr, "SongPresenter.Render: &q", err)
	}

	if err := testListbox.Render(renderer, testListboxData, sdl.Color{0x00,0xFF,0x00,0xFF}, sdl.Color{0xFF,0x00,0x00,0xFF},
				sdl.Color{0x00,0x00,0xFF,0xFF}, font); err != nil {
		fmt.Fprintf(os.Stderr, "testListbox.Render: &q", err)
	}

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

func updateText() (error){
	if SongDisplay, err = font.RenderUTF8Blended("<"+SongLoaded+">", sdl.Color{R: 255, G: 0, B: 0, A: 255}); err != nil {
		return err
	}

	if FileOpenerDisplay, err = font.RenderUTF8Blended(" " +FileOpenerData.Text, sdl.Color{R: 0, G: 255, B: 0, A: 255}); err != nil {
		return err
	}

	println(FileOpenerDisplay.W, FileOpenerDisplay.H)
	return nil
}


func mouseClick(){
	println(mouseData.X, mouseData.Y)
	LoadButton.CheckState(mouseData)

	if LoadButton.Pressed == true {
		println("loading music") //Insert dir open function
	}
	LoadButton.Pressed = false

	FileOpenerData.CheckState(FileOpenerInput, mouseData, pressedKeys)
	testListbox.CheckState(mouseData, testListboxData)
	println("testListbox:",testListbox.DownButton.State)
	println("testListbox:",testListbox.DownButton.Pressed)
	println("testListbox:",testListbox.Click)
	println("SONG:",testListbox.Song)
}

func TrimLastElement(input string) string{
	length:= len(input)
	if (length - 1) >= 0 {
		return input[:(length - 1)]
	}
	return ""
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

	string, err := lib.GetAudioPlaylist(SongLoaded)
	if err != nil {
		fmt.Fprintf(os.Stderr, "lib.GetAudioPlaylist: %q", err)
	}
	fmt.Println(string)

	if err = ttf.Init(); err != nil {
		panic("ttf.Init")
	}
	defer ttf.Quit()

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic("sdl.Init")
	}
	defer sdl.Quit()

	window, err = sdl.CreateWindow("Play 'N Go", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
									INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic("sdl.CreateWindow")
	}
	defer window.Destroy()

	//Required by WMs
	window.SetResizable(true)

	if renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		panic("sdl.CreateRenderer")
	}
	defer renderer.Destroy()

	if err = window.SetFullscreen(0); err != nil{
		panic("window.SetFullscreen")
	}

	if surface, err = window.GetSurface(); err != nil {
		panic("window.GetSurface")
	}
	defer surface.Free()


	if font, err = ttf.OpenFont(FONT_PATH, FONT_SIZE); err != nil {
		panic("ttf.OpenFont")
	}
	defer font.Close()

	FileOpenerData = ui.InitInputData()
	testListboxData = ui.InitListboxData()

	mouseData = &ui.Mouse{0,0, 0, ui.UNCLICKED}

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
				if t.Keysym.Sym == sdl.K_BACKSPACE && t.State == sdl.PRESSED && FileOpenerData.Pressed == true {
					FileOpenerData.Text = TrimLastElement(FileOpenerData.Text)
				}
			}
		}

		initUI()
		updateRendering()

		pressedKeys = sdl.GetKeyboardState()
		mouseData.X, mouseData.Y, mouseData.State = sdl.GetMouseState()
		println("MOUSE:", mouseData.State)
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
