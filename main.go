package main
import ("bufio"
		"strconv"
		"fmt"
		"os")


func main() {
	
	var size int = 3
	var sizeSquare int = size * size
	var grille[] string
	var player int
	var choice string
	var win bool = false
	// construction / initialisation de la grille
	
	grille = grilleCreate(sizeSquare)
	player = 1

	// tour de jeu
	for win == false {
		gameLoop(grille, size, player, choice)
		if player == 1 {
			player = 2
		} else {
			player = 1
		}
		// détection d'un gagnant
		win = isWin(grille, size)
	}
	winner(player)
}

func gameLoop(grille[] string, size int, player int, choice string) {
	var isValid bool
	var choiceInt int
	isValid = false
	for isValid == false {
		grilleDisplay(grille, size)
		choice = playerChoice(player)
		isValid, choiceInt = choiceValid(choice, grille)
		if isValid == true {
			if player == 1 {grille[choiceInt - 1] = "X"
			} else {
				grille[choiceInt - 1] = "O"
			}
		}
	}	
}

func grilleCreate(size int) ([]string) {
	var grille[]string
	grille = make([]string, size, size)
	var compteur = 1
	for i := 0; i < size; i++ {
		grille[i] = strconv.Itoa(compteur)
		compteur ++
	}
	return grille
}


func grilleDisplay(grille[]string, size int) {
	var compteur = 0
	for i := 0; i < size; i++ {
			fmt.Println(grille[compteur:compteur + size])
			compteur += size
		}
}

func playerChoice(player int) (string){
	var choice string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("C'est le tour de joueur ", player)
	fmt.Print("Dans quelle case veux tu jouer? ")
	
	scanner.Scan()
    choice = scanner.Text()
	return choice
}
func choiceValid(choice string, grille[] string) (bool, int) {
	// test si la case est libre
	var valid bool
	choiceInt, err := strconv.Atoi(choice)
	if err == nil {
		if (grille[choiceInt - 1] == "X") || (grille[choiceInt - 1] == "O") {
			fmt.Println("Case déjà utilisée, choisis en une autre")
			valid = false
		} else {valid = true}
	}
	
	return valid, choiceInt
}

func isWin(grille[] string, size int)(bool) {
	var win bool
	var i int
	win = false

	// diagonale 1
	if (grille[0] == grille[4]) && (grille[4] == grille[8]){
		win = true
	} // fin if
	// diagonale 2
	if (grille[2] == grille[4]) && (grille[4] == grille[6]){
		win = true
	} // fin if
// horizontal
	for i = 0; i < size; i++ {
		if (grille[i] == grille[i+1]) && (grille[i+1] == grille[i+2]){
			win = true
		}
	} // fin for

// vertical
	for i = 0; i < size; i++ {
		if (grille[i] == grille[i + 3]) && (grille[i + 3] == grille[i + 6]){
			win = true
		}
	} // fin for
	return win
} // fin isWin

func winner(player int) {
	if player == 1 {
		fmt.Println("Victoire de joueur 2")
	} else {
		fmt.Println("victoire de joueur 1")
	}
}