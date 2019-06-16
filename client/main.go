package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
	"untracked/onecardpoker/internal/utils"

	"github.com/sirupsen/logrus"
)

var input *bufio.Reader

var addr = flag.String("addr", "127.0.0.1:3001", "Server address")
var id = flag.Int("id", 1, "Player id (1 or 2)")

func main() {
	utils.ClearTerminal()
	flag.Parse()
	logrus.SetLevel(logrus.DebugLevel)

	if *id != 1 && *id != 2 {
		panic("Invalid player id, must be 1 or 2")
	}

	input = bufio.NewReader(os.Stdin)

	c := NewClient(*id, *addr)

	logrus.Debug("001")
	initGameState(c)
	logrus.Debug("002")
	for {
		state := waitForGameState(c)
		logrus.Debug("003")

		myCards := [2]*Card{
			NewCard(state[c.id][0]),
			NewCard(state[c.id][1]),
		}

		var oid int
		if c.id == 1 {
			oid = 2
		} else {
			oid = 1
		}

		theirCards := [2]*Card{
			NewCard(state[oid][0]),
			NewCard(state[oid][1]),
		}

		logrus.Debug("004")
		printGameState(myCards, theirCards)
		logrus.Debug("005")

		toplay := askForPlay(myCards)
		logrus.Debug("006")

		err := c.playCard(toplay)
		if err != nil {
			panic(err)
		}
		logrus.Debug("007")

		results := waitForResults(c)

		logrus.Debug("008")

		myPlay := (*results)[c.id]
		theirPlay := (*results)[oid]

		winner := endTurn(c, oid, myPlay, theirPlay)
		logrus.Debug("009")

		fmt.Println("Your  play:", NewCard(myPlay).ToString())
		fmt.Println("Their play:", NewCard(theirPlay).ToString())

		if winner == 0 {
			fmt.Println("Draw!")
		} else if winner == c.id {
			fmt.Println("You win this turn!")
		} else {
			fmt.Println("You lost this turn!")
		}

		logrus.Debug("010")
	}
}

func initGameState(c *Client) {
	if c.id != 1 {
		return
	}

	for {
		err := c.init()
		if err == nil {
			logrus.Info("Game initialised")
			break
		}

		logrus.WithField("error", err)
		time.Sleep(2 * time.Second)
	}
}

func waitForGameState(c *Client) GameState {
	for {
		state, err := c.fetchGameState()

		if err == nil {
			return state
		}

		if err == ErrGameStatePending {
			time.Sleep(time.Second)
			continue
		}

		panic(err)
	}
}

func printGameState(myCards [2]*Card, theirCards [2]*Card) {
	fmt.Println("Opponent has:")

	if theirCards[0].isUpper() && theirCards[1].isUpper() {
		fmt.Println("UPPER UPPER")
	} else if theirCards[0].isLower() && theirCards[1].isLower() {
		fmt.Println("LOWER LOWER")
	} else {
		fmt.Println("LOWER UPPER")
	}

	fmt.Print("\n\n----------\n\n")

	fmt.Println("Your hand:")
	fmt.Println("1.", (*myCards[0]).ToString())
	fmt.Println("2.", (*myCards[1]).ToString())
}

func askForPlay(myCards [2]*Card) *Card {
	for {
		fmt.Print("\nPlay card: ")

		choice, err := input.ReadString('\n')
		if err != nil {
			panic(err)
		}

		choice = choice[:1]

		if choice != "1" && choice != "2" {
			fmt.Println("Choice must be either 1 or 2")
			continue
		}

		x, err := strconv.Atoi(choice)
		if err != nil {
			panic(err)
		}

		return myCards[x-1]
	}
}

func waitForResults(c *Client) *TurnResult {
	for {
		r, err := c.fetchTurnResult()
		if err == nil {
			return r
		}

		if err == ErrTurnNotOver {
			time.Sleep(time.Second)
			continue
		}

		panic(err)
	}
}

func endTurn(c *Client, oid, myPlay, theirPlay int) int {
	if myPlay < theirPlay {
		if myPlay == 2 && theirPlay == 14 {
			err := c.drawNextTurn()
			if err != nil {
				panic(err)
			}
			return c.id
		}

		return oid
	}

	if theirPlay < myPlay {
		if theirPlay == 2 && myPlay == 14 {
			return oid
		}

		err := c.drawNextTurn()
		if err != nil {
			panic(err)
		}
		return c.id
	}

	if c.id == 1 {
		err := c.drawNextTurn()
		if err != nil {
			panic(err)
		}
	}

	return 0
}
