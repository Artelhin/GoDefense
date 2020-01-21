package input

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

//Interface used for comfortable input handling from ebiten package.
//Can look for customizable slice of keys and mouse buttons and solve keys according to the layout you give to it.
type InputSolver interface {
	GetInput() []KeyAction
	SolveKey(key ebiten.Key, layout Layout, uppercase bool) string
}

type inputSolver struct {
	FromKeyboard []ebiten.Key
	FromMouse    []ebiten.MouseButton

	keyWasPressed    map[ebiten.Key]bool
	buttonWasPressed map[ebiten.MouseButton]bool
}

type KeyState int

const (
	Pressed KeyState = iota
	Holded
	Released
)

type KeyAction struct {
	IsKey       bool
	Key         ebiten.Key
	MouseButton ebiten.MouseButton
	State       KeyState
}

func (inp *inputSolver) GetInput() []KeyAction {
	actions := make([]KeyAction, 0)
	for _, key := range inp.FromKeyboard {
		if ebiten.IsKeyPressed(key) {
			if inp.keyWasPressed[key] {
				actions = append(actions, KeyAction{
					IsKey: true,
					Key:   key,
					State: Holded,
				})
			} else {
				actions = append(actions, KeyAction{
					IsKey: true,
					Key:   key,
					State: Pressed,
				})
				inp.keyWasPressed[key] = true
			}
		} else {
			if inp.keyWasPressed[key] {
				actions = append(actions, KeyAction{
					IsKey: true,
					Key:   key,
					State: Released,
				})
				inp.keyWasPressed[key] = false
			}
		}
	}
	for _, button := range inp.FromMouse {
		if ebiten.IsMouseButtonPressed(button) {
			if inp.buttonWasPressed[button] {
				actions = append(actions, KeyAction{
					IsKey:       false,
					MouseButton: button,
					State:       Holded,
				})
			} else {
				actions = append(actions, KeyAction{
					IsKey:       false,
					MouseButton: button,
					State:       Pressed,
				})
				inp.buttonWasPressed[button] = true
			}
		} else {
			if inp.buttonWasPressed[button] {
				actions = append(actions, KeyAction{
					IsKey:       false,
					MouseButton: button,
					State:       Released,
				})
				inp.buttonWasPressed[button] = false
			}
		}
	}
	return actions
}

//Returns the string, according to the passed key and layout.
//Uppercase supported
func (i *inputSolver) SolveKey(key ebiten.Key, layout Layout, uppercase bool) string {
	switch layout {
	case EN:
		if uppercase {
			if s, ok := LayoutUppercaseNumbersEN[key]; ok {
				return s
			}
			if s, ok := LayoutUppercaseEN[key]; ok {
				return s
			}
			if s, ok := LayoutUpperSpecialEN[key]; ok {
				return s
			}
		} else {
			if s, ok := LayoutNumbers[key]; ok {
				return s
			}
			if s, ok := LayoutEN[key]; ok {
				return s
			}
			if s, ok := LayoutSpecialEN[key]; ok {
				return s
			}
		}
		if s, ok := LayoutKPNumbers[key]; ok {
			return s
		}
		log.Printf("INPUTER ERROR: no binding for key '%s'", key)
		return ""
	case RU:
		if uppercase {
			if s, ok := LayoutUppercaseNumbersRU[key]; ok {
				return s
			}
			if s, ok := LayoutUppercaseRU[key]; ok {
				return s
			}
			if s, ok := LayoutUpperSpecialRU[key]; ok {
				return s
			}
		} else {
			if s, ok := LayoutNumbers[key]; ok {
				return s
			}
			if s, ok := LayoutRU[key]; ok {
				return s
			}
			if s, ok := LayoutSpecialRU[key]; ok {
				return s
			}
		}
		if s, ok := LayoutKPNumbers[key]; ok {
			return s
		}
		log.Printf("INPUTER ERROR: no binding for key '%s'", key)
		return ""
	default:
		log.Printf("INPUTER ERROR: this layout in not supported")
		return ""
	}
}

//Returns new input solver which will look only for given keys and mouse buttons
func NewInputSolver(keys []ebiten.Key, buttons []ebiten.MouseButton) *inputSolver {
	return &inputSolver{
		FromKeyboard:     keys,
		FromMouse:        buttons,
		keyWasPressed:    make(map[ebiten.Key]bool),
		buttonWasPressed: make(map[ebiten.MouseButton]bool),
	}
}
