package sounds

import (
	"embed"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//go:embed res/*
var contents embed.FS

var sounds = make(map[string]*beep.Buffer)

func LoadSound(name, path string) beep.Format {
	f, err := contents.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	sounds[name] = buffer

	return format
}

func InitSounds() {

	format := LoadSound("gunag", "res/gunag.mp3")

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

}

func PlaySound(name string) {
	var buffer = sounds[name]
	sfx := buffer.Streamer(0, buffer.Len())
	speaker.Play(sfx)
}
