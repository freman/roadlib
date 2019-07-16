package tire

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type parser struct {
	input string
	index int
	info  Info
	step  step
	lastw int
	lastp string
	err   error
}

type step int

const (
	stepType step = iota
	stepWidth
	stepSlash
	stepAspectRatio
	stepConstruction
	stepDiameter
)

func (p *parser) Parse() (Info, error) {
	var loopSanity = 0
	for {
		if p.index >= len(p.input) {
			return p.info, p.err
		}
		if loopSanity > 20 {
			return p.info, errors.New("unable to parse input string")
		}
		switch p.step {
		case stepType:
			switch strings.ToUpper(p.peek()) {
			case "P":
				p.info.Type = Passenger
				p.pop()
			case "LT":
				p.info.Type = LightTruck
				p.pop()
			}
			p.step++
		case stepWidth:
			p.info.Width, p.err = strconv.ParseFloat(p.peek(), 64)
			if p.err != nil {
				return p.info, fmt.Errorf("parsing width failed due to %v at %d", p.err, p.index)
			}
			p.pop()
			p.step++
		case stepSlash:
			switch p.peek() {
			case "/":
				p.pop()
				p.step++
			default:
				return p.info, fmt.Errorf("expected / at %d", p.index)
			}
		case stepAspectRatio:
			p.info.AspectRatio, p.err = strconv.ParseFloat(p.peek(), 64)
			if p.err != nil {
				return p.info, fmt.Errorf("parsing aspect ratio failed due to %v at %d", p.err, p.index)
			}
			p.pop()
			p.step++
		case stepConstruction:
			switch strings.ToUpper(p.peek()) {
			case "R":
				p.info.Construction = Radial
				p.pop()
				p.step++
			default:
				return p.info, fmt.Errorf("expected construction (R) at %d", p.index)
			}
		case stepDiameter:
			p.info.Rim, p.err = strconv.ParseFloat(p.peek(), 64)
			if p.err != nil {
				return p.info, fmt.Errorf("parsing diameter failed due to %v at %d", p.err, p.index)
			}
			p.info.Rim *= 25.4
			return p.info, p.err
		}
	}
}

func (p *parser) peek() string {
	p.lastp, p.lastw = p.peekWithLength()
	return p.lastp
}

func (p *parser) pop() string {
	if p.lastp == "" {
		p.lastp, p.lastw = p.peekWithLength()
	}
	peeked := p.lastp
	p.lastp = ""
	p.index += p.lastw
	p.popWhitespace()
	return peeked
}

func (p *parser) peekWithLength() (string, int) {
	for _, tok := range []string{"P", "LT", "/", "R"} {
		token := p.input[p.index:min(len(p.input), p.index+len(tok))]
		if strings.EqualFold(tok, token) {
			return tok, len(tok)
		}
	}
	return p.peekFloatWidth()
}

func (p *parser) peekFloatWidth() (string, int) {
	v := ""
	for i := p.index; i < len(p.input); i++ {
		if (p.input[i] < '0' || p.input[i] > '9') && p.input[i] != '.' {
			break
		}
		v += p.input[i : i+1]
	}
	return v, len(v)
}

func (p *parser) popWhitespace() {
	for ; p.index < len(p.input) && p.input[p.index] == ' '; p.index++ {
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
