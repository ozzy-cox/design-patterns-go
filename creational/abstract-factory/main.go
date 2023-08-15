package abstract

type Dir string

const (
	NORTH Dir = "NORTH"
	SOUTH Dir = "SOUTH"
	EAST  Dir = "EAST"
	WEST  Dir = "WEST"
)

type (
	IMapSite interface {
		Enter()
	}
	IMaze interface {
		addRoom(room *IRoom)
		IMapSite
	}
	IRoom interface {
		setSide(dir Dir, site IMapSite)
		getSide(dir Dir)
		IMapSite
	}
	IDoor interface {
		IMapSite
	}

	IWall interface {
		IMapSite
	}

	Room         struct{}
	Door         struct{}
	Wall         struct{}
	IMazeFactory interface {
		makeMaze() *IMaze
		makeDoor(r1 *IRoom, r2 *IRoom) *IDoor
		makeWall() *IWall
		makeRoom(n int) *IRoom
	}
)

func (*Room) Enter() {}
func (*Door) Enter() {}
func (*Wall) Enter() {}

func createMaze(factory IMazeFactory) *IMaze {
	amaze := factory.makeMaze()
	room1 := factory.makeRoom(1)
	room2 := factory.makeRoom(2)
	door := factory.makeDoor(room1, room2)

	(*amaze).addRoom(room1)
	(*amaze).addRoom(room2)

	(*room1).setSide(NORTH, *factory.makeWall())
	(*room1).setSide(EAST, *door)
	(*room1).setSide(SOUTH, *factory.makeWall())
	(*room1).setSide(WEST, *factory.makeWall())

	(*room2).setSide(NORTH, *factory.makeWall())
	(*room2).setSide(EAST, *factory.makeWall())
	(*room2).setSide(SOUTH, *factory.makeWall())
	(*room2).setSide(WEST, *door)

	return amaze
}

func main() {
}
