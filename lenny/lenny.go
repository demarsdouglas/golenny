package lenny

type Mood int64

const (
    Happy Mood = iota
    Sad
    Angry
    Strange
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
    }
    return "unknown"
}

type Identifier struct {
    name string
    mood Mood
}

type Lenny struct {
    leftEye     string
    rightEye    string
    leftEnd     string
    rightEnd    string
    mouth       string
    identifiers []Identifier
}
