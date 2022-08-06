package lenny

import (
	"math/rand"
	"time"
)

var predefined = map[string]*Lenny{
	"Lenny": {
		// ( ͡° ͜ʖ ͡°)
		leftEye:  " ͡°",
		rightEye: "͡°",
		leftEnd:  "(",
		rightEnd: ")",
		mouth:    " ͜ʖ ",
		mood:     Happy,
		name:     "Lenny",
	},
	"Raging-Mad": {
		// ᕕ(ᗒ෴ᗕ)ᕗ
		leftEye:  "ᗒ",
		rightEye: "ᗕ",
		leftEnd:  "ᕕ(",
		rightEnd: ")ᕗ",
		mouth:    "෴",
		mood:     Angry,
		name:     "Raging-Mad",
	},
	"Gasp": {
		// ∑(; °Д°)
		leftEye:  "°",
		rightEye: "°",
		leftEnd:  "∑(; ",
		rightEnd: ")",
		mouth:    "Д",
		mood:     Surprised,
		name:     "Gasp",
	},
	"Defeated": {
		// (◞д◟)
		leftEye:  "◞",
		rightEye: "◟",
		leftEnd:  "(",
		rightEnd: ")",
		mouth:    "д",
		mood:     Sad,
		name:     "Defeated",
	},
	"Juicer": {
		// ᕦ⊙෴⊙ᕤ
		leftEye:  "⊙",
		rightEye: "⊙",
		leftEnd:  "ᕦ",
		rightEnd: "ᕤ",
		mouth:    "෴",
		mood:     Strange,
		name:     "Juicer",
	},
}

func LoadPredefinedByName(name string) *Lenny {
	return predefined[name]
}

func LoadRandomPredefined() *Lenny {
	var randSlice []*Lenny
	for _, l := range predefined {
		randSlice = append(randSlice, l)
	}

	rand.Seed(time.Now().Unix())
	randPredefinedLenny := randSlice[rand.Intn(len(randSlice))]

	return randPredefinedLenny
}

func LoadRandomPredefinedByMood(mood Mood) *Lenny {
	var moodSlice []*Lenny
	for _, l := range predefined {
		if l.mood == mood {
			moodSlice = append(moodSlice, l)
		}
	}

	rand.Seed(time.Now().Unix())
	randLennyByMood := moodSlice[rand.Intn(len(moodSlice))]

	return randLennyByMood
}
