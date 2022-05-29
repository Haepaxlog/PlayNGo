package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)



func ListboxInitSurfaces(fileList []string, font *ttf.Font, color sdl.Color) ([]*sdl.Surface, error){
	var err error
	TextSurfaces := new([len(fileList)]*sdl.Surface)

	for i := 0; i < len(fileList); i++ {
		if SurfaceName[i], err = font.RenderUTF8Blended(fileList[i], color); err != nil {
			return err
		}
	}

	return TextSurfaces, err

}

func ListboxDraw(){

}


















	if SongDisplay, err = font.RenderUTF8Blended("<"+SongLoaded+">", sdl.Color{R: 255, G: 0, B: 0, A: 255}); err != nil {
		return
	}


	if err = SongDisplay.Blit(nil, surface, &sdl.Rect{X: viewport_size.W/2 - (SongDisplay.W + viewport_size.W/10),
		Y: viewport_size.H - (SongDisplay.H + 50), W: 0, H: 0}); err != nil {
			return
	}

}
