package main

type Character struct {
	Name    string
	Health  int
	Evasion int
	Alive   bool
	Speed   int
	Weap    int
	Npc     bool
	Items   []int

	Welcome         string
	CurrentLocation string
}

func (p *Character) Equip(w int) {
	p.Weap = w
}

func (p *Character) Attack() int {
	return Weaps[p.Weap].Fire()
}

type Players []Character

func (slice Players) Len() int {
	return len(slice)
}

func (slice Players) Less(i, j int) bool {
	return slice[i].Speed > slice[j].Speed //Sort descending
	//return slice[i].Speed < slice[j].Speed;		//Sort ascending
}

func (slice Players) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

var numberOfPlays = 0

func (p *Character) Play() {
	if numberOfPlays == 0 {
		numberOfPlays++
		Output("magenta", "********************************************************")
		Output("magenta", "Hello there friend!\n I'm so happy you came to play with me!\n   Tip #1: type `go PLACE` to move around!  Tip #2: `help` is always there for you!\n     Have fun!")
	}
	numberOfPlays++

	Output(p.Welcome)
	for {
		Output("cyan", LocationMap[p.CurrentLocation].Description)
		Output("white", "\n********************************************************")
		p.ProcessEvents(LocationMap[p.CurrentLocation].Events)
		if p.Health <= 0 {
			Output("white", "You are dead, GAME OVER YO!!!")
			return
		}
		Output("magenta", "\nHealth:", p.Health)
		if p.Health < 40 {
			Output("white", "Uhh, you're getting low on health levels there. Might wanna step up your game a bit. ;)")
		}
		if len(LocationMap[p.CurrentLocation].Items) > 0 {
			Output("cyan", "Oh looky! You can see...")
			for _, itm := range LocationMap[p.CurrentLocation].Items {
				Outputf("cyan", "\t%s", Items[itm].Name)
			}
		}
		Output("yellow", "You can `go` to these places:")
		for _, loc := range LocationMap[p.CurrentLocation].Transitions {
			Outputf("yellow", "\t%s", loc)
		}
		cmd := UserInputln()
		ProcessCommands(p, cmd)
	}
}

func (p *Character) ProcessEvents(events []string) {
	for _, evtName := range events {
		p.Health += evts[evtName].ProcessEvent(p)
	}
}
