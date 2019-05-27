const process = require("process");
const readline = require("readline");
const term = require( 'terminal-kit' ).terminal;

const buffer = [];

const rl = readline.createInterface({
    input: process.stdin,
});

rl.on("line", (line) => {
    term.red(line);
});

