
var cors = require('cors');

var bodyParser = require('body-parser')
var express = require('express'),
app = express(),
port = process.env.PORT || 3001;
app.use(cors({origin: '*'}));
app.use(bodyParser.json());       // to support JSON-encoded bodies
app.use(bodyParser.urlencoded({     // to support URL-encoded bodies
  extended: true
})); 

var playedCards = {
  1:undefined,
  2:undefined
}
var currentCards = {
  1:[],
  2:[]
}
var cardArray = [
  2,2,2,2,
  3,3,3,3,
  4,4,4,4,
  5,5,5,5,
  6,6,6,6,
  7,7,7,7,
  8,8,8,8,
  9,9,9,9,
  10,10,10,10,
  11,11,11,11,
  12,12,12,12,
  13,13,13,13,
  14,14,14,14
]
var turn = ['states','states','play','play','results','results','distribute']

var nextTurn = () => {
  turn.push(turn.shift())
}
var resetGame = () => {
  cardArray = [
    2,2,2,2,
    3,3,3,3,
    4,4,4,4,
    5,5,5,5,
    6,6,6,6,
    7,7,7,7,
    8,8,8,8,
    9,9,9,9,
    10,10,10,10,
    11,11,11,11,
    12,12,12,12,
    13,13,13,13,
    14,14,14,14
  ]
  turn = ['states','states','play','play','results','results','distribute']

  playedCards = {
    1:undefined,
    2:undefined
  }

  currentCards = {
    1:[],
    2:[]
  }
  
  _.shuffle(cardArray) 
  currentCards['1'].push(cardArray.shift())
  currentCards['1'].push(cardArray.shift())
  currentCards['2'].push(cardArray.shift())
  currentCards['2'].push(cardArray.shift())
}

var newCardEach = () =>{
    let played1 = playedCards['1']
    let played2 = playedCards['2']
    if(!played1 || !played2)
      return
    let playerWon = 0
    if(played1<played2)
      if(played1===2 && played2===14)
        playerWon = 1
      else playerWon = 2
    if(played1>played2)
      if(played2===2 && played1===14)
        playerWon = 2
      else playerWon = 1

    playerWon<2 ? currentCards['1'].push(cardArray.shift()) : currentCards['2'].push(cardArray.shift())
}

var routes = require('./api/routes/Routes'); //importing route
routes(app); //register the route
app.listen(port);