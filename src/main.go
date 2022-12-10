package main

import (
	"fmt"
	"os"

	"github.com/2mer/goplaylist/sounds"
	tea "github.com/charmbracelet/bubbletea"
)

// func main() {

// 	const soundsDirPath = "C:\\Users\\user\\Sounds"
// 	soundFilePath := filepath.Join(soundsDirPath, "Gunag.mp3")

// 	f, err := os.Open(soundFilePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	streamer, format, err := mp3.Decode(f)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer streamer.Close()

// 	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

// 	done := make(chan struct{})
// 	speaker.Play(
// 		beep.Seq(
// 			streamer,
// 			beep.Callback(func() {
// 				done <- struct{}{}
// 			}),
// 		),
// 	)

// 	<-done
// }

func main() {
	// init
	sounds.InitSounds()

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
