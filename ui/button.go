package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Button_State int

const(
	UP	Button_State = iota
	HOVER
	DOWN
)

type Button struct{
	Rect	sdl.Rect
	State	Button_State
	Pressed	bool
}


type Mouse struct{
	X int32
	Y int32
	State uint32
}

const(
	SOURCE_WINDOW_WIDTH = 1920
	SOURCE_WINDOW_HEIGHT = 1080
	)

func CreateButton(renderer *sdl.Renderer, rect sdl.Rect) (*Button){
	viewport := renderer.GetViewport()

	rect.X = int32(float32(rect.X) * float32(viewport.W)/float32(SOURCE_WINDOW_WIDTH))
	rect.Y = int32(float32(rect.Y) * float32(viewport.H)/float32(SOURCE_WINDOW_HEIGHT))
	rect.W = int32(float32(rect.W) * float32(viewport.W)/float32(SOURCE_WINDOW_WIDTH))
	rect.H = int32(float32(rect.H) * float32(viewport.H)/float32(SOURCE_WINDOW_HEIGHT))
	button := Button{rect, UP, false}

	return &button
}

func (button *Button) Render(renderer *sdl.Renderer, color sdl.Color){
	renderer.SetDrawColor(color.R,color.G,color.B,color.A)
	renderer.FillRect(&button.Rect)
}

func (button *Button) CheckState(mouseData *Mouse) {
	button.State = UP
	if button.Rect.HasIntersection(&sdl.Rect{mouseData.X, mouseData.Y, 1, 1}){
		button.State = HOVER
		if mouseData.State == 1{
			button.State = DOWN
			button.Pressed = true
		}
	}
}
