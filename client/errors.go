package main

import "errors"

// ErrGameStatePending _
var ErrGameStatePending = errors.New("Game state pending")

// ErrTurnNotOver _
var ErrTurnNotOver = errors.New("Turn not over")
