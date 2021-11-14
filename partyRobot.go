package main

import (
	"fmt"
	"strconv"
)

func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

func AssignTable(name string, table int, neighbour string, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.",
		name, table, direction, distance, neighbour)
}