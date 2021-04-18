package main

import (
	"fmt"
	"os/exec"
	//"github.com/faiface/beep"
)

// TODO : use a sound library ffs
func click_sound() {
	var errA error
	mpv_cmd := fmt.Sprintf("mpv --really-quiet audio/click.wav")
	cmd_run, errA = exec.Command("bash", "-c", mpv_cmd).Output()
	if errA != nil {
		fmt.Printf("Error playing sound\n")
	}
}

func dice_sound() {
	var errA error
	mpv_cmd := fmt.Sprintf("mpv --really-quiet --volume=80 audio/dice.mp3")
	cmd_run, errA = exec.Command("bash", "-c", mpv_cmd).Output()
	if errA != nil {
		fmt.Printf("Error playing sound\n")
	}
}
