require('lodash')
exports.initializeGame = (req,res) => { 
    resetGame()
    return res.status(200)
}

exports.announceUpDown = (req, res) => {
    if(turn[0]!=="states")
        return res.status(400).text("not in the state to announce states")
    nextTurn()
    return res.status(200).json(currentCards)
}

exports.playCard = (req, res) => {
    if(turn[0]!=="play")
        return res.status(400).text("not in the state to play")
    nextTurn()
    let playerID = req.param.playerID
    let card = req.param.card
    if(playedCards[playerID])
        return 
    playedCards[playerID] = parseInt(card)
    return res.status(200)
}

exports.getResults = (req, res) => {
    if(turn[0]!=="results")
        return res.status(400).text("not in the state to announce results")
    nextTurn()
    return res.status(200).json(playedCards)
  
}

exports.distributeCard = (req, res) => {
    if(turn[0]!=="distribute")
        return res.status(400).text("not in the state to distribute new card")
    newCardEach()
    playedCards = {
        1:undefined,
        2:undefined
    }
    nextTurn()
    return res.status(200)
}