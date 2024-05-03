package interfaces

import "fmt"

type Player interface {
	KickBall() int
}

type FootballPlayer struct {
	speed   int
	stamina int
}

func (fp FootballPlayer) KickBall() int {
	return fp.speed * fp.stamina
}

type CR7Player struct {
	speed   int
	stamina int
}

func (cr7 CR7Player) KickBall() int {
	return cr7.speed * cr7.stamina * 2
}

func Exec() {
	players := generatePlayers()

	for _, player := range players {
		result(player.KickBall())
	}
}

func generatePlayers() []Player {
	players := []Player{
		FootballPlayer{speed: 5, stamina: 10},
		FootballPlayer{speed: 10, stamina: 5},
		FootballPlayer{speed: 3, stamina: 8},
		FootballPlayer{speed: 8, stamina: 3},
		CR7Player{speed: 10, stamina: 10},
	}

	return players
}

func result(power int) {
	fmt.Printf("Kicked with %d\n", power)
}
