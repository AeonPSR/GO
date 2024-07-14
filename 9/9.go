package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	largeur, hauteur int
}

type Cercle struct {
	rayon int
}

type Forme interface {
	aire() int
}

func (c Cercle) aire() int {
	return int(math.Pi * float64(c.rayon) * float64(c.rayon)) //Cast into float to be able to use math.Pi, casted back into int afteward)
}

func (r Rectangle) aire() int {
	return r.largeur * r.hauteur
}
//In GO apparently funct can have the same name if we specify the type going along with it.
//So calling "aire" for a cercle or a rectangle will just use the right funct depending on the type.
//Weird.

func main() {
	cercle := Cercle{rayon: 10}
	rectangle := Rectangle{largeur: 10, hauteur: 10}

	formes := []Forme{cercle, rectangle} //Slice (?) qui les regroupent

	for _, forme := range formes {
		fmt.Printf("L'aire de la forme est: %d\n", forme.aire()) //Affiche les valeurs dans le même ordre qu'ils sont ordonnées dans la slice.
	}
}
