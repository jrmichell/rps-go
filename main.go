package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	ROCK     = "r"
	PAPER    = "p"
	SCISSORS = "s"
)

var playerScore, computerScore int = 0, 0
var playerChoice, computerChoice string
var winningScore int

func randInt(min int, max int) int {
	random_number := rand.Intn(max-min+1) + min
	return random_number
}

func getPlayerChoice() string {
	fmt.Println("\nPick (R)ock, (P)aper, (S)cissors: ")

	for {
		fmt.Scanln(&playerChoice)
		strings.ToLower(playerChoice)

		switch playerChoice {
		case ROCK:
			break
		case PAPER:
			break
		case SCISSORS:
			break
		default:
			fmt.Println("\nPlease pick a valid option - (R)ock, (P)aper, (S)cissors: ")
			continue
		}

		return playerChoice
	}
}

func getComputerChoice() string {
	randomNumber := randInt(1, 100)

	if randomNumber <= 33 {
		computerChoice = ROCK
	} else if randomNumber <= 66 {
		computerChoice = PAPER
	} else {
		computerChoice = SCISSORS
	}

	return computerChoice
}

func compareChoices() {
	if strings.Compare(playerChoice, computerChoice) == 0 {
		fmt.Printf("\nTie! You both picked %s!\n", playerChoice)
	} else {
		fmt.Printf("You picked %s, Computer picked %s\n", playerChoice, computerChoice)

		if computerChoice == ROCK && playerChoice == PAPER {
			playerScore++
		}
		if computerChoice == PAPER && playerChoice == ROCK {
			computerScore++
		}

		if computerChoice == PAPER && playerChoice == SCISSORS {
			playerScore++
		}
		if computerChoice == SCISSORS && playerChoice == PAPER {
			computerScore++
		}

		if computerChoice == SCISSORS && playerChoice == ROCK {
			playerScore++
		}
		if computerChoice == ROCK && playerChoice == SCISSORS {
			computerScore++
		}
	}
}

func determineWinner() bool {
	fmt.Printf("\nPlayer: %d\nComputer: %d\n\n", playerScore, computerScore)
	if computerScore == winningScore || playerScore == winningScore {
		if computerScore > playerScore {
			fmt.Println("Computer Wins!")
		} else if playerScore > computerScore {
			fmt.Println("Player Wins!")
		}
		return false
	} else {
		return true
	}
}

func game() {
	fmt.Println("What is the winning score you want to have?")
	fmt.Scanln(&winningScore)

	for (playerScore < winningScore) && (computerScore < winningScore) {
		getPlayerChoice()
		getComputerChoice()
		compareChoices()
		determineWinner()
	}
}

func main() {
	print("Welcome to Rock, Paper, Scissors!\n\n")
	game()
}
