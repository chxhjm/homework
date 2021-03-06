// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"fmt"
)

// Injectors from main.go:

func InitMission(name string) Mission {
	player := NewPlayer(name)
	monster := NewMonster()
	mission := NewMission(player, monster)
	return mission
}

// main.go:

type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "kitty"}
}

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{Player: p, Monster: m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

func main() {
	monster := NewMonster()
	player := NewPlayer("dj")
	mission := NewMission(player, monster)

	mission.Start()
}
