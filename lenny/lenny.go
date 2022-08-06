package lenny

import "fmt"

type Mood int64

const (
	Happy Mood = iota
	Sad
	Angry
	Strange
	Surprised
)

func (m Mood) String() string {
	switch m {
	case Happy:
		return "happy"
	case Sad:
		return "sad"
	case Angry:
		return "angry"
	case Strange:
		return "strange"
	case Surprised:
		return "surprised"
	}
	return "unknown"
}

type Lenny struct {
	leftEye  string
	rightEye string
	leftEnd  string
	rightEnd string
	mouth    string
	mood     Mood
	name     string
}

func LoadDefault() *Lenny {
	return &Lenny{
		leftEye:  " ͡°",
		rightEye: "͡°",
		leftEnd:  "(",
		rightEnd: ")",
		mouth:    " ͜ʖ ",
		mood:     Happy,
		name:     "Lenny",
	}
}

func (l Lenny) Render() string {
	return fmt.Sprintf(
		"%s%s%s%s%s", l.leftEnd, l.leftEye, l.mouth, l.rightEye, l.rightEnd,
	)
}
