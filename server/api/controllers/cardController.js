var _ = require('lodash')
var playedCards, currentCards, cardArray
var turn = ['states', 'states', 'play', 'play', 'results', 'results', 'distribute']
var nextTurn = () => {
    turn.push(turn.shift())
}
var resetGame = () => {
    cardArray = [
        2, 2, 2, 2,
        3, 3, 3, 3,
        4, 4, 4, 4,
        5, 5, 5, 5,
        6, 6, 6, 6,
        7, 7, 7, 7,
        8, 8, 8, 8,
        9, 9, 9, 9,
        10, 10, 10, 10,
        11, 11, 11, 11,
        12, 12, 12, 12,
        13, 13, 13, 13,
        14, 14, 14, 14
    ]
    turn = ['states', 'states', 'play', 'play', 'results', 'results', 'distribute']

    playedCards = {
        1: undefined,
        2: undefined
    }

    currentCards = {
        1: [],
        2: []
    }

    cardArray = _.shuffle(cardArray)
    currentCards['1'].push(cardArray.shift())
    currentCards['1'].push(cardArray.shift())
    currentCards['2'].push(cardArray.shift())
    currentCards['2'].push(cardArray.shift())

}

var newCardEach = () => {
    let played1 = playedCards['1']
    let played2 = playedCards['2']
    if (!played1 || !played2)
        return
    let playerWon = 0
    if (played1 < played2)
        if (played1 === 2 && played2 === 14)
            playerWon = 1
    else playerWon = 2
    if (played1 > played2)
        if (played2 === 2 && played1 === 14)
            playerWon = 2
    else playerWon = 1
    currentCards['1'].splice(currentCards['1'].indexOf(played1),1)
    currentCards['2'].splice(currentCards['2'].indexOf(played2),1)
    playedCards = {
        1: undefined,
        2: undefined
    }
    playerWon < 2 ? currentCards['1'].push(cardArray.shift()) : currentCards['2'].push(cardArray.shift())
}

setInterval(() => {
    //console.log(`Current State - ${turn[0]}`)
}, 2000)

exports.initializeGame = (req, res) => {
    resetGame()
    console.log("GAME INITIALIZED")
    return res.status(200).send()
}

exports.announceUpDown = (req, res) => {
    console.log("Received ask state")
    if (turn[0] !== "states")
        return res.status(400).send("not in the state to announce states")
    console.log("ASKED STATE SUCCESSFULLY")
    nextTurn()
    return res.status(200).json(currentCards)
}

exports.playCard = (req, res) => {
    if (turn[0] !== "play")
        return res.status(400).send("not in the state to play")
    console.log("PLAYED CARD SUCCESSFULLY")
    nextTurn()
    let playerID = req.params.playerID
    let card = req.params.card
    if (playedCards[playerID])
        return
    playedCards[playerID] = parseInt(card)
    return res.status(200).send()
}

exports.getResults = (req, res) => {
    if (turn[0] !== "results")
        return res.status(400).send("not in the state to announce results")
    console.log("RESULTS ANNOUNCED SUCCESSFULLY")
    nextTurn()
    return res.status(200).json(playedCards)
}

exports.distributeCard = (req, res) => {
    if (turn[0] !== "distribute")
        return res.status(400).send("not in the state to distribute new card")
    newCardEach()
    console.log("CARD DISTRIBUTED SUCCESSFULLY")
    nextTurn()
    return res.status(200).send()
}