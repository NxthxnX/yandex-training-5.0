package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Point struct {
	X, Y int
}

type Fragment struct {
	L, R int
}

type Line struct {
	Y, H  int
	Frags []Fragment
}

type Surrounded struct {
	Y, H, X, W int
}

type Element struct {
	Type   string
	Text   []rune
	W, H   int
	Dx, Dy int
}

type State struct {
	PageW, InitH, CharW int
	Lines               []Line
	CurLine, CurFrag    int
	CurX                int
	Surroundeds         []Surrounded
	RefX, RefY          int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	firstLine := strings.Fields(scanner.Text())

	w, _ := strconv.Atoi(firstLine[0])
	h, _ := strconv.Atoi(firstLine[1])
	c, _ := strconv.Atoi(firstLine[2])

	var content strings.Builder
	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}
	lines := strings.Split(content.String(), "\n")

	var paragraphs [][]string
	var curParagraph []string
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			if len(curParagraph) > 0 {
				paragraphs = append(paragraphs, curParagraph)
				curParagraph = nil
			}
		} else {
			curParagraph = append(curParagraph, line)
		}
	}
	if len(curParagraph) > 0 {
		paragraphs = append(paragraphs, curParagraph)
	}

	var results []Point
	paragraphStartY := 0

	for _, paragraphLines := range paragraphs {
		text := strings.Join(paragraphLines, " ")
		elems := tokenize(text)
		if len(elems) == 0 {
			continue
		}

		state := NewState(w, h, c, paragraphStartY)

		for _, el := range elems {
			var x, y int
			switch el.Type {
			case "word":
				width := len(el.Text) * c
				x, y = state.placeWordOrEmbedded(width, 0, false)
			case "embedded":
				x, y = state.placeWordOrEmbedded(el.W, el.H, true)
			case "surrounded":
				x, y = state.placeSurrounded(el.W, el.H)
			case "floating":
				x, y = state.placeFloating(el.W, el.Dx, el.Dy)
			}

			if el.Type != "word" {
				results = append(results, Point{X: x, Y: y})
			}
		}
		paragraphStartY = getParagraphBottom(state)
	}

	for _, point := range results {
		fmt.Println(point.X, point.Y)
	}
}

func tokenize(text string) []Element {
	textRunes := []rune(text)
	var elems []Element
	i := 0
	n := len(textRunes)
	for i < n {
		if unicode.IsSpace(textRunes[i]) {
			i++
			continue
		}
		if textRunes[i] == '(' {
			j := i + 1
			for textRunes[j] != ')' {
				j++
			}
			imgStr := textRunes[i : j+1]
			elems = append(elems, parseImage(string(imgStr)))
			i = j + 1
		} else {
			j := i
			for j < n && !unicode.IsSpace(textRunes[j]) && textRunes[j] != '(' {
				j++
			}
			word := textRunes[i:j]
			if len(word) > 0 {
				elems = append(elems, Element{Type: "word", Text: word})
			}
			i = j
		}
	}
	return elems
}

func parseImage(s string) Element {
	s = strings.TrimPrefix(s, "(image")
	s = strings.TrimSuffix(s, ")")
	s = strings.TrimSpace(s)
	parts := strings.Fields(s)
	e := Element{Type: "image"}
	for _, p := range parts {
		param := strings.Split(p, "=")
		switch param[0] {
		case "layout":
			e.Type = param[1]
		case "width":
			e.W, _ = strconv.Atoi(param[1])
		case "height":
			e.H, _ = strconv.Atoi(param[1])
		case "dx":
			e.Dx, _ = strconv.Atoi(param[1])
		case "dy":
			e.Dy, _ = strconv.Atoi(param[1])
		}
	}
	return e
}

func NewState(w, h, c, startY int) *State {
	return &State{
		PageW: w, InitH: h, CharW: c,
		Lines:   []Line{{Y: startY, H: h, Frags: []Fragment{{L: 0, R: w}}}},
		CurLine: 0, CurFrag: 0, CurX: 0,
		RefX: 0, RefY: startY,
	}
}

func (s *State) placeWordOrEmbedded(width, height int, isEmbedded bool) (int, int) {
	for {
		s.ensureLine(s.CurLine)
		line := &s.Lines[s.CurLine]
		if len(line.Frags) == 0 {
			s.CurLine++
			s.CurFrag = 0
			continue
		}

		for i := s.CurFrag; i < len(line.Frags); i++ {
			f := line.Frags[i]
			tryX := s.CurX
			space := 0
			if i == s.CurFrag {
				if tryX > f.L {
					space = s.CharW
				}
			} else {
				tryX = f.L
				space = 0
			}
			if tryX < f.L {
				tryX = f.L
			}

			if tryX+space+width <= f.R {
				xPos := tryX + space
				yPos := line.Y
				s.CurFrag = i
				s.CurX = xPos + width
				s.updateRef(xPos+width, yPos)
				if isEmbedded && height > line.H {
					s.expandLine(s.CurLine, height)
				}
				return xPos, yPos
			}
		}

		s.CurLine++
		s.CurFrag = 0
		s.ensureLine(s.CurLine)
		if len(s.Lines[s.CurLine].Frags) > 0 {
			s.CurX = s.Lines[s.CurLine].Frags[0].L
		} else {
			s.CurX = 0
		}
	}
}

func (s *State) placeSurrounded(width, height int) (int, int) {
	for {
		s.ensureLine(s.CurLine)
		line := &s.Lines[s.CurLine]
		if len(line.Frags) == 0 {
			s.CurLine++
			s.CurFrag = 0
			continue
		}

		for i := s.CurFrag; i < len(line.Frags); i++ {
			f := line.Frags[i]
			startX := s.CurX
			if i > s.CurFrag {
				startX = f.L
			}
			if startX < f.L {
				startX = f.L
			}

			if f.R-startX >= width {
				xPos := startX
				yPos := line.Y

				line.Frags = cutFrags(line.Frags, xPos, xPos+width)
				s.CurFrag = i
				s.CurX = xPos + width
				s.updateRef(xPos+width, yPos)
				s.Surroundeds = append(s.Surroundeds, Surrounded{Y: yPos, H: height, X: xPos, W: width})

				for j := s.CurLine + 1; j < len(s.Lines); j++ {
					if s.Lines[j].Y < yPos+height && s.Lines[j].Y+s.Lines[j].H > yPos {
						s.Lines[j].Frags = cutFrags(s.Lines[j].Frags, xPos, xPos+width)
					}
				}
				return xPos, yPos
			}
		}
		s.CurLine++
		s.CurFrag = 0
		s.ensureLine(s.CurLine)
		if len(s.Lines[s.CurLine].Frags) > 0 {
			s.CurX = s.Lines[s.CurLine].Frags[0].L
		} else {
			s.CurX = 0
		}
	}
}

func (s *State) placeFloating(width, dx, dy int) (int, int) {
	x := s.RefX + dx
	y := s.RefY + dy
	if x < 0 {
		x = 0
	}
	if x+width > s.PageW {
		x = s.PageW - width
	}
	s.updateRef(x+width, y)
	return x, y
}

func getParagraphBottom(s *State) int {
	bottom := s.Lines[len(s.Lines)-1].Y + s.Lines[len(s.Lines)-1].H
	for _, sr := range s.Surroundeds {
		if b := sr.Y + sr.H; b > bottom {
			bottom = b
		}
	}
	return bottom
}

func (s *State) ensureLine(idx int) {
	for len(s.Lines) <= idx {
		last := s.Lines[len(s.Lines)-1]
		newY := last.Y + last.H
		newFrags := []Fragment{{L: 0, R: s.PageW}}

		for _, sr := range s.Surroundeds {
			if sr.Y < newY+s.InitH && sr.Y+sr.H > newY {
				newFrags = cutFrags(newFrags, sr.X, sr.X+sr.W)
			}
		}
		s.Lines = append(s.Lines, Line{Y: newY, H: s.InitH, Frags: newFrags})
	}
}

func cutFrags(frags []Fragment, cutL, cutR int) []Fragment {
	var res []Fragment
	for _, f := range frags {
		if f.R <= cutL || f.L >= cutR {
			res = append(res, f)
		} else {
			if f.L < cutL {
				res = append(res, Fragment{f.L, cutL})
			}
			if f.R > cutR {
				res = append(res, Fragment{cutR, f.R})
			}
		}
	}
	return res
}

func (s *State) updateRef(rx, ry int) {
	s.RefX = rx
	s.RefY = ry
}

func (s *State) expandLine(idx, newH int) {
	diff := newH - s.Lines[idx].H
	if diff <= 0 {
		return
	}
	s.Lines[idx].H = newH
	for i := idx + 1; i < len(s.Lines); i++ {
		s.Lines[i].Y += diff
	}

	for i := range s.Surroundeds {
		if s.Surroundeds[i].Y > s.Lines[idx].Y {
			s.Surroundeds[i].Y += diff
		}
	}
}
