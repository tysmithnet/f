const process = require("process");
const readline = require("readline");
var blessed = require('blessed')
  , contrib = require('blessed-contrib')
  , screen = blessed.screen()
  , grid = new contrib.grid({rows: 12, cols: 12, screen: screen})
  , map = grid.set(0, 0, 4, 4, contrib.map, {label: 'World Map'})
  , box = grid.set(4, 4, 4, 4, blessed.box, {content: 'My Box'})

screen.render()

const rl = readline.createInterface({
    input: process.stdin,
});

rl.on("line", (line) => {
    box.setContent(line);
    screen.render();
});