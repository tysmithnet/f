const process = require("process");
const readline = require("readline");
const blessed = require('blessed');
const contrib = require('blessed-contrib');
const screen = blessed.screen({
    smartCSR: true,
});

const grid = new contrib.grid({ rows: 20, cols: 1, screen: screen });
const input = grid.set(0, 0, 1, 1, blessed.textbox, { label: "Input!", height: "100%", width: "100%", style: {fg: "green"} })
const box = grid.set(1, 0, 19, 1, blessed.box, { label: "Box!", height: "100%", width: "100%" })
const buffer = [];

input.focus();
screen.render();

const rl = readline.createInterface({
    input: process.stdin,
    terminal: false
});

rl.on("line", (line) => {
    buffer.push(line);
    box.setContent(line);
    screen.render();
});