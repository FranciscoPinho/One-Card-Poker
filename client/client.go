package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/json-iterator/go"
)

// GameState _
type GameState map[int][2]int

// TurnResult _
type TurnResult map[int]int

// Client _
type Client struct {
	id   int
	addr string
}

// NewClient _
func NewClient(id int, addr string) *Client {
	return &Client{id: id, addr: "http://" + addr}
}

func (c *Client) init() error {
	// init/
	addr := fmt.Sprintf("%s/init", c.addr)

	fmt.Println("init addr", addr)
	_, err := http.Get(addr)

	return err
}

func (c *Client) playCard(card *Card) error {
	// playCard/:id/:card
	addr := fmt.Sprintf("%s/playCard/%d/%d", c.addr, c.id, card.val)

	_, err := http.Get(addr)

	return err
}

func (c *Client) fetchGameState() (GameState, error) {
	// announceStatus/
	addr := fmt.Sprintf("%s/announceStatus/", c.addr)

	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 400 {
		return nil, ErrGameStatePending
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var state GameState
	err = jsoniter.Unmarshal(data, &state)
	if err != nil {
		fmt.Println("data", string(data))
		return nil, err
	}

	return state, nil
}

func (c *Client) fetchTurnResult() (*TurnResult, error) {
	// results/
	addr := fmt.Sprintf("%s/results/", c.addr)

	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 400 {
		return nil, ErrTurnNotOver
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TurnResult
	err = jsoniter.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("data", string(data))
		return nil, err
	}

	return &result, nil
}

func (c *Client) drawNextTurn() error {
	time.Sleep(2 * time.Second)

	// distributeCard/
	addr := fmt.Sprintf("%s/distributeCard", c.addr)

	_, err := http.Get(addr)

	return err
}
