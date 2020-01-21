package input

import "github.com/hajimehoshi/ebiten"

type Layout string

const (
	RU Layout = `ru`
	EN Layout = `en`
)

var (
	LayoutNumbers = map[ebiten.Key]string{
		ebiten.Key0: "0",
		ebiten.Key1: "1",
		ebiten.Key2: "2",
		ebiten.Key3: "3",
		ebiten.Key4: "4",
		ebiten.Key5: "5",
		ebiten.Key6: "6",
		ebiten.Key7: "7",
		ebiten.Key8: "8",
		ebiten.Key9: "9",
	}

	LayoutUppercaseNumbersEN = map[ebiten.Key]string{
		ebiten.Key0: ")",
		ebiten.Key1: "!",
		ebiten.Key2: "@",
		ebiten.Key3: "#",
		ebiten.Key4: "$",
		ebiten.Key5: "%",
		ebiten.Key6: "^",
		ebiten.Key7: "&",
		ebiten.Key8: "*",
		ebiten.Key9: "(",
	}

	LayoutUppercaseNumbersRU = map[ebiten.Key]string{
		ebiten.Key0: ")",
		ebiten.Key1: "!",
		ebiten.Key2: "",
		ebiten.Key3: "№",
		ebiten.Key4: ";",
		ebiten.Key5: "%",
		ebiten.Key6: ":",
		ebiten.Key7: "?",
		ebiten.Key8: "*",
		ebiten.Key9: "(",
	}

	LayoutKPNumbers = map[ebiten.Key]string{
		ebiten.KeyKP0: "0",
		ebiten.KeyKP1: "1",
		ebiten.KeyKP2: "2",
		ebiten.KeyKP3: "3",
		ebiten.KeyKP4: "4",
		ebiten.KeyKP5: "5",
		ebiten.KeyKP6: "6",
		ebiten.KeyKP7: "7",
		ebiten.KeyKP8: "8",
		ebiten.KeyKP9: "9",
	}

	LayoutRU = map[ebiten.Key]string{
		ebiten.KeyF:            "а",
		ebiten.KeyPause:        "б",
		ebiten.KeyD:            "в",
		ebiten.KeyU:            "г",
		ebiten.KeyL:            "д",
		ebiten.KeyT:            "е",
		ebiten.KeyGraveAccent:  "ё",
		ebiten.KeySemicolon:    "ж",
		ebiten.KeyP:            "з",
		ebiten.KeyB:            "и",
		ebiten.KeyR:            "к",
		ebiten.KeyK:            "л",
		ebiten.KeyV:            "м",
		ebiten.KeyY:            "н",
		ebiten.KeyJ:            "о",
		ebiten.KeyG:            "п",
		ebiten.KeyH:            "р",
		ebiten.KeyC:            "с",
		ebiten.KeyN:            "т",
		ebiten.KeyE:            "у",
		ebiten.KeyA:            "ф",
		ebiten.KeyLeftBracket:  "х",
		ebiten.KeyW:            "ц",
		ebiten.KeyX:            "ч",
		ebiten.KeyI:            "ш",
		ebiten.KeyO:            "щ",
		ebiten.KeyRightBracket: "ъ",
		ebiten.KeyS:            "ы",
		ebiten.KeyM:            "ь",
		ebiten.KeyApostrophe:   "э",
		ebiten.KeyPeriod:       "ю",
		ebiten.KeyZ:            "я",
		//todo font for this
	}

	LayoutUppercaseRU = map[ebiten.Key]string{
		ebiten.KeyF:            "А",
		ebiten.KeyPause:        "Б",
		ebiten.KeyD:            "В",
		ebiten.KeyU:            "Г",
		ebiten.KeyL:            "Д",
		ebiten.KeyT:            "Е",
		ebiten.KeyGraveAccent:  "Ё",
		ebiten.KeySemicolon:    "Ж",
		ebiten.KeyP:            "З",
		ebiten.KeyB:            "И",
		ebiten.KeyR:            "К",
		ebiten.KeyK:            "Л",
		ebiten.KeyV:            "М",
		ebiten.KeyY:            "Н",
		ebiten.KeyJ:            "О",
		ebiten.KeyG:            "П",
		ebiten.KeyH:            "Р",
		ebiten.KeyC:            "С",
		ebiten.KeyN:            "Т",
		ebiten.KeyE:            "У",
		ebiten.KeyA:            "Ф",
		ebiten.KeyLeftBracket:  "Х",
		ebiten.KeyW:            "Ц",
		ebiten.KeyX:            "Ч",
		ebiten.KeyI:            "Ш",
		ebiten.KeyO:            "Щ",
		ebiten.KeyRightBracket: "Ъ",
		ebiten.KeyS:            "Ы",
		ebiten.KeyM:            "Ь",
		ebiten.KeyApostrophe:   "Э",
		ebiten.KeyPeriod:       "Ю",
		ebiten.KeyZ:            "Я",
	}

	LayoutSpecialRU = map[ebiten.Key]string{
		//ebiten.KeyRightBracket: "[",
		//ebiten.KeyLeftBracket:  "]",
		ebiten.KeySlash: ".",
		//ebiten.KeySemicolon:    ";",
		//ebiten.KeyApostrophe:   "'",
		//ebiten.KeyPause:        ",",
		//ebiten.KeyPeriod:       ".",
	}

	LayoutUpperSpecialRU = map[ebiten.Key]string{
		//ebiten.KeyRightBracket: "{",
		//ebiten.KeyLeftBracket:  "}",
		ebiten.KeySlash: ",",
		//ebiten.KeySemicolon:    ":",
		//ebiten.KeyApostrophe:   `"`,
		//ebiten.KeyPause:        "<",
		//ebiten.KeyPeriod:       ">",
	}

	LayoutEN = map[ebiten.Key]string{
		ebiten.KeyA: "a",
		ebiten.KeyB: "b",
		ebiten.KeyC: "c",
		ebiten.KeyD: "d",
		ebiten.KeyE: "e",
		ebiten.KeyF: "f",
		ebiten.KeyG: "g",
		ebiten.KeyH: "h",
		ebiten.KeyI: "i",
		ebiten.KeyJ: "j",
		ebiten.KeyK: "k",
		ebiten.KeyL: "l",
		ebiten.KeyM: "m",
		ebiten.KeyN: "n",
		ebiten.KeyO: "o",
		ebiten.KeyP: "p",
		ebiten.KeyQ: "q",
		ebiten.KeyR: "r",
		ebiten.KeyS: "s",
		ebiten.KeyT: "t",
		ebiten.KeyU: "u",
		ebiten.KeyV: "v",
		ebiten.KeyW: "w",
		ebiten.KeyX: "x",
		ebiten.KeyY: "y",
		ebiten.KeyZ: "z",
	}

	LayoutUppercaseEN = map[ebiten.Key]string{
		ebiten.KeyA: "A",
		ebiten.KeyB: "B",
		ebiten.KeyC: "C",
		ebiten.KeyD: "D",
		ebiten.KeyE: "E",
		ebiten.KeyF: "F",
		ebiten.KeyG: "G",
		ebiten.KeyH: "H",
		ebiten.KeyI: "I",
		ebiten.KeyJ: "J",
		ebiten.KeyK: "K",
		ebiten.KeyL: "L",
		ebiten.KeyM: "M",
		ebiten.KeyN: "N",
		ebiten.KeyO: "O",
		ebiten.KeyP: "P",
		ebiten.KeyQ: "Q",
		ebiten.KeyR: "R",
		ebiten.KeyS: "S",
		ebiten.KeyT: "T",
		ebiten.KeyU: "U",
		ebiten.KeyV: "V",
		ebiten.KeyW: "W",
		ebiten.KeyX: "X",
		ebiten.KeyY: "Y",
		ebiten.KeyZ: "Z",
	}

	LayoutSpecialEN = map[ebiten.Key]string{
		ebiten.KeyRightBracket: "[",
		ebiten.KeyLeftBracket:  "]",
		ebiten.KeySlash:        "/",
		ebiten.KeySemicolon:    ";",
		ebiten.KeyApostrophe:   "'",
		ebiten.KeyPause:        ",",
		ebiten.KeyPeriod:       ".",
	}

	LayoutUpperSpecialEN = map[ebiten.Key]string{
		ebiten.KeyRightBracket: "{",
		ebiten.KeyLeftBracket:  "}",
		ebiten.KeySlash:        "?",
		ebiten.KeySemicolon:    ":",
		ebiten.KeyApostrophe:   `"`,
		ebiten.KeyPause:        "<",
		ebiten.KeyPeriod:       ">",
	}
)
