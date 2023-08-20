package main

func main() {
	player1 := Player{}
	player1.transport = &OnFoot{}
	player1.name = "player1"

	player1.travel()

	player2 := Player{}
	player2.transport = &Vehicle{}
	player2.name = "player2"

	player2.travel()

	npc1 := NPC{}
	npc1.transport = &OnFoot{}
	npc1.name = "npc1"
	npc1.travel()

	npc2 := NPC{}
	npc2.transport = &Vehicle{}
	npc2.name = "npc2"
	npc2.travel()
}
