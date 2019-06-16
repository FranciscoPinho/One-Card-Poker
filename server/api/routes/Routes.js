'use strict';
module.exports = function(app) {
  var controller = require('../controllers/cardController.js')
  app.route('/init')
    .get(controller.initializeGame)
  app.route('/announceStatus')
    .get(controller.announceUpDown)
  app.route('/playCard/:playerID/:card')
    .get(controller.playCard)
  app.route('/results')
    .get(controller.getResults)
  app.route('/distributeCard')
    .get(controller.distributeCard)
};