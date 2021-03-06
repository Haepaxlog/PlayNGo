package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Input struct {
    Rect	sdl.Rect
	Display sdl.Rect
}

type InputData struct{
	Pressed bool
	Text	string
	State 	Button_State
}

func CreateInput(renderer *sdl.Renderer, rect sdl.Rect) (*Input){
	viewport := renderer.GetViewport()

	rect.X = int32(float32(rect.X) * float32(viewport.W)/float32(SOURCE_WINDOW_WIDTH))
	rect.Y = int32(float32(rect.Y) * float32(viewport.H)/float32(SOURCE_WINDOW_HEIGHT))
	rect.W = int32(float32(rect.W) * float32(viewport.W)/float32(SOURCE_WINDOW_WIDTH))
	rect.H = int32(float32(rect.H) * float32(viewport.H)/float32(SOURCE_WINDOW_HEIGHT))

	display := sdl.Rect{int32(float32(rect.X) + float32(rect.X)*1.75),
						rect.Y,
						int32(float32(rect.W)*0.75),
						rect.H}

	input := Input{rect, display}

	return &input
}

func (input *Input) Render(renderer *sdl.Renderer, colorRect sdl.Color, colorDisp sdl.Color, InputDisplay *sdl.Surface) (error){
	//Draw Background Rect
	renderer.SetDrawColor(colorRect.R, colorRect.G, colorRect.B, colorRect.A)
	renderer.FillRect(&input.Rect)

	//Draw Display Rect
	renderer.SetDrawColor(colorDisp.R, colorDisp.G, colorDisp.B, colorDisp.A)
	renderer.FillRect(&input.Display)

	//Draw Display Text
	DisplayTexture, err := renderer.CreateTextureFromSurface(InputDisplay)
	if err != nil {
		return err
	}
	renderer.Copy(DisplayTexture, nil, &sdl.Rect{input.Display.X, input.Display.Y, InputDisplay.W, input.Display.H})
	return nil
}

func (input *InputData) CheckState(inputRect *Input,mouseData *Mouse, pressedKeys []uint8) {
	input.State = UP
	if inputRect.Rect.HasIntersection(&sdl.Rect{X: mouseData.X, Y: mouseData.Y, W: 1, H: 1}){
		input.State = HOVER
	}

	if input.State == HOVER && mouseData.State == 1 {
		input.State = DOWN
		input.Pressed = true
	}

	if input.State == UP && mouseData.State == 1{
		input.Pressed = false
	}

	if pressedKeys[sdl.SCANCODE_ESCAPE] != 0 {
		input.State = UP
		input.Pressed = false
	}
}

func InitInputData() (*InputData){
	return &InputData{false, "", UP}
}
