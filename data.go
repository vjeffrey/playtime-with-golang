package main

var Items = map[int]*Item{
	1: {Name: "Key"},
	2: {Name: "Chest", ItemForUse: 1, Contains: []int{3}},
	3: {Name: "Balloon"},
}

var Weaps = map[int]*Weapon{
	1: {Name: "Ice Cream Cone", minAtt: 5, maxAtt: 15},
	2: {Name: "Pie Throwing Machine", minAtt: 1, maxAtt: 15},
	3: {Name: "Silly String", minAtt: 3, maxAtt: 12},
}

var Enemies = map[int]*Character{
	0: {Name: "Koopa", Health: 50, Alive: true, Weap: 2, Npc: true},
	1: {Name: "Bowser", Health: 55, Alive: true, Weap: 3, Npc: true},
}

var evts = map[string]*Event{
	"zombieAttack":     {Type: "Combat", Chance: 20, Description: "A bunch of zombies jump in front of you and start dancing towards you.", Health: -50, Evt: "doctorTreatment"},
	"zombieBabyAttack": {Type: "Combat", Chance: 20, Description: "A crowd of zombie babies are thrown at you.  They have sharp nails.", Health: -50, Evt: "doctorTreatment"},
	"zombieBirdAttack": {Type: "Combat", Chance: 20, Description: "BIRDSSS!!!! Zombie birds start flying at you and pecking your head!", Health: -50, Evt: "doctorTreatment"},
	"doctorTreatment":  {Type: "Story", Chance: 10, Description: "The doctor rushes in and gives you a really cool zombie antidote.", Health: +30, Evt: ""},
	"unicorn":          {Type: "Story", Chance: 50, Description: "A unicorn strolls by, all shiny and pretty.", Health: +10, Evt: ""},
	"firefly":          {Type: "Story", Chance: 50, Description: "A firefly greets you with his adorable little lit up lower abdomen.", Health: +10, Evt: ""},
	"relaxing":         {Type: "Story", Chance: 100, Description: "Bob Marley really is the way to go.", Health: +30, Evt: ""},
}

var LocationMap = map[string]*Location{
	"Bench":  {Description: "Close your eyes and imagine you are sitting on a bench by the canal in Amsterdam, on a warm night, breathing in the fresh air, listening to the early morning animals awaken.", Transitions: []string{"Room", "Lift"}},
	"Coast":  {Description: "Sitting on the warm sand, you hear the waves crashing in the middle of the night, while music booms behind you on the coast of St. Louis, Senegal", Transitions: []string{"Bench", "Lounge", "Forest"}, Events: []string{"firefly"}},
	"Room":   {Description: "You are discussing the weirdness of life with The Man from Another Place in the Red Room.", Transitions: []string{"Bench"}, Events: []string{}, Items: []int{2}},
	"Lift":   {Description: "A Sailor Moon style awesome wind tunnel that takes you anywhere in the world.", Transitions: []string{"Bench", "Lounge", "Forest"}, Events: []string{"firefly"}},
	"Forest": {Description: "You are in a lush forest, breathing in the fresh air from the trees.", Transitions: []string{"Lift"}, Events: []string{"unicorn"}, Items: []int{1}},
	"Mall":   {Description: "You have entered a mall.  May the force be with you.", Transitions: []string{"Lift"}, Events: []string{"zombieAttack", "zombieBabyAttack", "zombieBirdAttack"}},
	"Lounge": {Description: "You are in the lounge, you feel very relaxed", Transitions: []string{"Lift"}, Events: []string{"relaxing"}},
}
