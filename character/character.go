package character

const (
	orientedDown int = iota
	orientedLeft
	orientedRight
	orientedUp
)

// Character définit les caractéristiques du personnage.
// Pour le moment seules les coordonnées absolues de l'endroit
// où il se trouve sont exportées, mais vous pourrez
// ajouter des choses aux besoins lors de votre développement.
//
// Les champs sont :
//   - orientation : l'orientation du personnage (haut, bas, gauche, droite).
//   - animationStep : l'étape d'animation (-1 ou 1, représentant l'animation
//     d'un pas à gauche ou à droite).
//   - XInc, YInc : les incréments en X et Y à réaliser après la prochaine animation.
//   - Moving : l'information de si une animation est en cours ou pas.
//   - shift : la position actuelle en pixels du personnage relativement à ses
//     coordonnées absolues.
//   - animationFrameCount : le nombre d'appels à update (ou de 1/60 de seconde) qui
//     ont eu lieu depuis la dernière étape d'animation.
type Character struct {
	X, Y                int
	orientation         int
	animationStep       int
	XInc, YInc          int
	Moving              bool
	shift               int
	animationFrameCount int
}
