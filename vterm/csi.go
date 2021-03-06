package vterm

import (
	"log"

	"github.com/aaronjanse/3mux/render"
)

func (v *VTerm) handleEraseInDisplay(parameterCode string) {
	seq := parseSemicolonNumSeq(parameterCode, 0)
	switch seq[0] {
	case 0: // clear from Cursor to end of screen
		for i := v.Cursor.X; i < len(v.Screen[v.Cursor.Y]); i++ {
			v.Screen[v.Cursor.Y][i].Rune = ' '
			v.Screen[v.Cursor.Y][i].IsWide = false
			v.Screen[v.Cursor.Y][i].PrevWide = false
		}
		if v.Cursor.Y+1 < len(v.Screen) {
			for j := v.Cursor.Y; j < len(v.Screen); j++ {
				for i := 0; i < len(v.Screen[j]); i++ {
					v.Screen[j][i].Rune = ' '
					v.Screen[j][i].IsWide = false
					v.Screen[j][i].PrevWide = false
				}
			}
		}
		v.RedrawWindow()
	case 1: // clear from Cursor to beginning of screen
		for j := 0; j < v.Cursor.Y; j++ {
			for i := 0; i < len(v.Screen[j]); i++ {
				v.Screen[j][i].Rune = ' '
				v.Screen[j][i].IsWide = false
				v.Screen[j][i].PrevWide = false
			}
		}
		v.RedrawWindow()
	case 2: // clear entire screen (and move Cursor to top left?)
		for i := range v.Screen {
			for j := range v.Screen[i] {
				v.Screen[i][j].Rune = ' '
				v.Screen[i][j].IsWide = false
				v.Screen[i][j].PrevWide = false
			}
		}
		v.setCursorPos(0, 0)
		v.RedrawWindow()
	case 3: // clear entire screen and delete all lines saved in scrollback buffer
		v.Scrollback = [][]render.Char{}
		for i := range v.Screen {
			for j := range v.Screen[i] {
				v.Screen[i][j].Rune = ' '
				v.Screen[i][j].IsWide = false
				v.Screen[i][j].PrevWide = false
			}
		}
		v.setCursorPos(0, 0)
		v.RedrawWindow()
	default:
		log.Printf("Unrecognized erase in display directive: %v", seq[0])
	}
}

func (v *VTerm) handleEraseInLine(parameterCode string) {
	seq := parseSemicolonNumSeq(parameterCode, 0)
	switch seq[0] {
	case 0: // clear from Cursor to end of line
		for i := v.Cursor.X; i < len(v.Screen[v.Cursor.Y]); i++ {
			v.Screen[v.Cursor.Y][i].Rune = ' '
			v.Screen[v.Cursor.Y][i].IsWide = false
			v.Screen[v.Cursor.Y][i].PrevWide = false
		}
	case 1: // clear from Cursor to beginning of line
		for i := 0; i < v.Cursor.X; i++ {
			v.Screen[v.Cursor.Y][i].Rune = ' '
			v.Screen[v.Cursor.Y][i].IsWide = false
			v.Screen[v.Cursor.Y][i].PrevWide = false
		}
	case 2: // clear entire line; Cursor position remains the same
		for i := 0; i < len(v.Screen[v.Cursor.Y]); i++ {
			v.Screen[v.Cursor.Y][i].Rune = ' '
			v.Screen[v.Cursor.Y][i].IsWide = false
			v.Screen[v.Cursor.Y][i].PrevWide = false
		}
	default:
		log.Printf("Unrecognized erase in line directive: %v", seq[0])
	}
	v.RedrawWindow()
}
