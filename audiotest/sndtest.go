package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/generators"
)

func old(){
	var tmp string
	fmt.Print("music:")
	fmt.Scanln(&tmp)

	f, err := os.Open(tmp)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done

}

func test(){
	var tmp string
	speaker.Init(beep.SampleRate(48000), 4800)//format.SampleRate, format.SampleRate.N(time.Second/10))

	s, err := generators.SinTone(beep.SampleRate(48000), 1000)
	if err != nil {
		panic(err)
	}

	speaker.Play(s)
	fmt.Scanln(&tmp)
	speaker.Clear()
	s2, err := generators.SinTone(beep.SampleRate(48000), 500)
	if err != nil {
		panic(err)
	}
	speaker.Play(s2)
	fmt.Scanln(&tmp)
	speaker.Clear()
	s3, err := generators.SinTone(beep.SampleRate(48000), 300)
	if err != nil {
		panic(err)
	}
	speaker.Play(s3)
	fmt.Scanln(&tmp)
	speaker.Clear()
	//fmt.Print("music:")	
}

func main() {
	test()
}
