package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var playerScore, computerScore int = 0, 0
var playerChoice, computerChoice string
var winningScore int

func randInt(max int, min int) int {
	random_number := rand.Intn(max-min+1) + min
	return random_number
}

func getPlayerChoice() string {
	for {
		fmt.Println("\nPick (R)ock, (P)aper, (S)cissors: ")
		fmt.Scanln(&playerChoice)
		strings.ToLower(playerChoice)

		switch playerChoice {
		case "r":
			break
		case "p":
			break
		case "s":
			break
		default:
			continue
		}

		return playerChoice
	}
}

func getComputerChoice() string {
	randomNumber := randInt(100, 1)

	if randomNumber <= 33 {
		computerChoice = "r"
	} else if randomNumber <= 66 {
		computerChoice = "p"
	} else {
		computerChoice = "s"
	}

	return computerChoice
}

func compareChoices() {
	if strings.Compare(playerChoice, computerChoice) == 0 {
		fmt.Printf("Tie!\nYou both picked %s!\n", playerChoice)
	} else {
		fmt.Printf("You picked %s, Computer picked %s\n", playerChoice, computerChoice)

		if computerChoice == "r" && playerChoice == "p" {
			playerScore++
		} else if computerChoice == "p" && playerChoice == "r" {
			computerScore++
		}

		if computerChoice == "p" && playerChoice == "s" {
			playerScore++
		} else if computerChoice == "s" && playerChoice == "p" {
			computerScore++
		}

		if computerChoice == "s" && playerChoice == "r" {
			playerScore++
		} else if computerChoice == "r" && playerChoice == "s" {
			computerScore++
		}
	}
}

func determineWinner() bool {
	if computerScore == winningScore || playerScore == winningScore {
		fmt.Printf("Final Score:\n\nPlayer: %d\nComputer: %d\n\n", playerScore, computerScore)
		if computerScore > playerScore {
			fmt.Println("Computer Wins!")
		} else if playerScore > computerScore {
			fmt.Println("Player Wins!")
		}
		return false
	} else {
		fmt.Printf("Score:\n\nPlayer: %d\nComputer: %d\n\n", playerScore, computerScore)
		return true
	}
}

func game() {
	fmt.Println("What is the winning score you want to have?")
	fmt.Scanln(&winningScore)

	for (playerScore < winningScore) && (computerScore < winningScore) {
		playerChoice = getPlayerChoice()
		computerChoice = getComputerChoice()
		compareChoices()
		determineWinner()
	}
}

func main() {
	print("Welcome to Rock, Paper, Scissors!\n\n")
	game()
}
