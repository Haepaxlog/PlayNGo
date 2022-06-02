package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Listbox struct{
	Rect	sdl.Rect
	ClickData []string
	ClickDataRects	[]sdl.Rect
	Click	int
	UpButton Button
	DownButton Button
	Song	string
}


func InitListboxData() (*Listbox){
	listbox := Listbox{Rect: sdl.Rect{0,0,0,0}, ClickData: []string{""}, ClickDataRects: []sdl.Rect{},
		Click: 0, UpButton: Button{sdl.Rect{0,0,0,0}, UP, false}, DownButton: Button{sdl.Rect{0,0,0,0}, UP, false}, Song: ""}
	return &listbox
}

func CreateListbox(renderer *sdl.Renderer, listboxData *Listbox,rect sdl.Rect, Data []string) (*Listbox){
	viewport := renderer.GetViewport()

	rect.X = int32(float32(rect.X) * float32(viewport.W)/float32(SOURCE_WINDOW_WIDTH))
	rect.Y = int32(float32(rect.Y) * float32(viewport.H)/float32(SOURCE_WINDOW_HEIGHT))
	rect.W = int32(float32(rect.W) * float32(viewport.W)/float32(SOURCE_WINDOW_WIDTH))
	rect.H = int32(float32(rect.H) * float32(viewport.H)/float32(SOURCE_WINDOW_HEIGHT))

	upButtonRect := sdl.Rect{rect.X + rect.W + 50, rect.Y, rect.W/10, rect.H/10}
	downButtonRect := sdl.Rect{rect.X + rect.W + 50, rect.Y + rect.H, rect.W/10, rect.H/10}

	upButton := Button{upButtonRect, UP, false}
	downButton := Button{downButtonRect, UP, false}

	dataCount := rect.H / 100
	y := 0

	dataRects := make([]sdl.Rect, int(dataCount))
	for i := 0; i < int(dataCount); i++ {
		dataRects[i] = sdl.Rect{rect.X + 20, rect.Y + 20 + int32(y), rect.W - 40, 90}
		y += 100
	}

	listbox := Listbox{Rect: rect, ClickData: Data, ClickDataRects: dataRects , Click: listboxData.Click, UpButton: upButton, DownButton: downButton, Song: listboxData.Song}
	return &listbox
}

func (listbox *Listbox) Render(renderer *sdl.Renderer, listboxData *Listbox, color sdl.Color, textColor sdl.Color, textFieldColor sdl.Color, font *ttf.Font) (error){
	var err error

	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	renderer.FillRect(&listbox.Rect)
	renderer.FillRect(&listbox.UpButton.Rect)
	renderer.FillRect(&listbox.DownButton.Rect)

	for i := 0; i < len(listbox.ClickDataRects); i++ {
		renderer.SetDrawColor(textFieldColor.R, textFieldColor.G, textFieldColor.B, textFieldColor.A)
		if listbox.ClickData[(i + listboxData.Click)] == listboxData.Song {
			renderer.SetDrawColor(100, textFieldColor.G, textFieldColor.B, textFieldColor.A)
		}
		renderer.FillRect(&listbox.ClickDataRects[i])
	}

	displays := make([]*sdl.Surface, len(listbox.ClickDataRects))
	textures := make([]*sdl.Texture, len(listbox.ClickDataRects))
	for i := 0; i < len(listbox.ClickDataRects); i++ {
		displays[i], err = font.RenderUTF8Blended(listbox.ClickData[(i + listboxData.Click)], textColor)
		if err != nil {
			return err
		}
		textures[i], err = renderer.CreateTextureFromSurface(displays[i])
		if err != nil {
			return err
		}
		renderer.Copy(textures[i], nil, &listbox.ClickDataRects[i])
	}
	return nil
}

func (listbox *Listbox) CheckState(mouseData *Mouse, listboxData *Listbox) {
	if mouseData.State == 1 && mouseData.PressedState == UNCLICKED{
		mouseData.PressedState = CLICKED
	}

	if mouseData.State != 1{
		mouseData.PressedState = UNCLICKED
	}

	listbox.UpButton.State = UP
	if listbox.UpButton.Rect.HasIntersection(&sdl.Rect{X: mouseData.X, Y: mouseData.Y, W: 1, H: 1}) {
		listbox.UpButton.State = HOVER
	}

	if listbox.UpButton.State == HOVER && mouseData.PressedState == CLICKED {
		listbox.UpButton.State = DOWN
		listbox.UpButton.Pressed = true
	}

	listbox.DownButton.State = UP
	if listbox.DownButton.Rect.HasIntersection(&sdl.Rect{X: mouseData.X, Y: mouseData.Y, W: 1, H: 1}) {
		listbox.DownButton.State = HOVER
	}

	if listbox.DownButton.State == HOVER && mouseData.PressedState == CLICKED {
		listbox.DownButton.State = DOWN
		listbox.DownButton.Pressed = true
	}


	if mouseData.State == 1 {
		for i := 0; i < len(listbox.ClickDataRects); i++ {
			if listbox.ClickDataRects[i].HasIntersection(&sdl.Rect{X: mouseData.X, Y: mouseData.Y, W: 1, H: 1}){
				listboxData.Song = listbox.ClickData[(i + listboxData.Click)]
			}
		}
	}

	listbox.actOnState(listboxData, mouseData)
}

func (listbox *Listbox) actOnState(listboxData *Listbox, mouseData *Mouse){
	if listbox.DownButton.Pressed == true && mouseData.PressedState == CLICKED{
		if listboxData.Click + 1 <= len(listbox.ClickData) - 3 {
			listboxData.Click += 1
			listbox.DownButton.Pressed = false
			mouseData.PressedState = PENDING
		}
	}
	if listbox.UpButton.Pressed == true && mouseData.PressedState == CLICKED{
		if listboxData.Click - 1 >= 0{
			listboxData.Click -= 1
			listbox.UpButton.Pressed = false
			mouseData.PressedState = PENDING
		}
	}

	/*if mouseData.State == 4{
		if listboxData.Click - 1 >= 0{
			listboxData.Click -= 1
			listbox.UpButton.Pressed = false
			mouseData.PressedState = PENDING
		}
	}*/
}
