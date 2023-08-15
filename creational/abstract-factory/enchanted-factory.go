package abstract

type (
	EnchantedMazeFactory struct{}
	EnchantedRoom        struct{}
	EnchantedDoor        struct{}
)

func (e EnchantedMazeFactory) makeMaze() *IMaze {
	panic("not implemented") // TODO: Implement
}

func (e EnchantedMazeFactory) makeDoor(r1 *IRoom, r2 *IRoom) *IDoor {
	panic("not implemented") // TODO: Implement
}

func (e EnchantedMazeFactory) makeWall() *IWall {
	panic("not implemented") // TODO: Implement
}

func (e EnchantedMazeFactory) makeRoom(n int) *IRoom {
	panic("not implemented") // TODO: Implement
}
