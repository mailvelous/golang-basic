package main

import "fmt"

type FootballPlayer interface {
	kick() string
	run() string
	tackle() string
}

type Defender struct {
	Name string
	Club string
}
type Midfielder struct {
	Name string
	Club string
}

func main() {
	name, club := getPlayerInput()

	player := &Midfielder{
		Name: name,
		Club: club,
	}

	var fp FootballPlayer = player
	fmt.Println(fp.kick())
	fmt.Println(fp.run())
	fmt.Println(fp.tackle())

	// printPlayerInfo(player)

}

	// func printPlayerInfo(p FootballPlayer) {
	// 	fmt.Println("A football player can", p.kick(), p.run(), p.tackle())
	// }

func (m *Midfielder) kick() string {
	return m.Name + " from " + m.Club + " " + "is kicking"
}

func (m *Midfielder) run() string {
	return m.Name + " from " + m.Club + " " + "is running"
}

func (m *Midfielder) tackle() string {
	return m.Name + " from " + m.Club + " " + "is tackling"
}

////////////////////////////

func (d *Defender) tackle() string {
	return d.Name + " from " + d.Club + " " + "is tackling"
}

func (d *Defender) run() string {
	return d.Name + " from " + d.Club + " " + "is running"
}

func (d *Defender) kicking() string {
	return d.Name + " from " + d.Club + " " + "is kicking"
}


func getPlayerInput() (string, string) {
	var name string
	var club string

	fmt.Println("Enter Name:")
	fmt.Scan(&name)

	fmt.Println("Enter Club:")
	fmt.Scan(&club)

	return name, club
}
