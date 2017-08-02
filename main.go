package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"time"
)

var Out *os.File
var In *os.File

var player Character

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	Out = os.Stdout
	In = os.Stdin
}

func setPlayerInfo() Character {
	player = *new(Character)
	Output("cyan", "Hello user, what is your name?")
	name := UserInputln()
	player.Name = name
	player.Speed = 1 + rand.Intn(100)
	player.Alive = true
	player.Weap = 1
	player.CurrentLocation = "Room"
	player.Health = 100

	Output("cyan", "Good luck my friend!")
	return player
}

func main() {
	player = setPlayerInfo()
	// setUpDB(player)
	player.Play()
}

func Outputf(c string, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	Output(c, s)
}

func Output(c string, args ...interface{}) {
	s := fmt.Sprint(args...)
	col := color.WhiteString
	switch c {
	case "green":
		col = color.GreenString
	case "red":
		col = color.RedString
	case "blue":
		col = color.BlueString
	case "yellow":
		col = color.YellowString
	case "magenta":
		col = color.MagentaString
	case "cyan":
		col = color.CyanString
	}
	fmt.Fprintln(Out, col(s))
}

func UserInput(i *int) {
	fmt.Fscan(In, i)
}

func UserInputln() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n >>> ")
	text, _ := reader.ReadString('\n')
	return text
}
