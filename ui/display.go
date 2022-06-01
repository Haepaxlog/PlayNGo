package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Display struct{
	Rect	sdl.Rect
	Text	string
}


func CreateDisplay(renderer *sdl.Renderer, rect sdl.Rect, text string) (*Display){
	viewport := renderer.GetViewport()

	rect.X = int32(float32(rect.X) * float32(viewport.W)/float32(SOURCE_WINDOW_WIDTH))
	rect.Y = int32(float32(rect.Y) * float32(viewport.H)/float32(SOURCE_WINDOW_HEIGHT))
	rect.W = int32(float32(rect.W) * float32(viewport.W)/float32(SOURCE_WINDOW_WIDTH))
	rect.H = int32(float32(rect.H) * float32(viewport.H)/float32(SOURCE_WINDOW_HEIGHT))
	display := Display{rect, text}

	return &display
}

func (display *Display) Render(renderer *sdl.Renderer, Display *sdl.Surface) (error){
	DisplayTexture, err := renderer.CreateTextureFromSurface(Display)
	if err != nil {
		return err
	}

	renderer.Copy(DisplayTexture, nil, &sdl.Rect{display.Rect.X, display.Rect.Y, Display.W, Display.H})
	return nil
}
