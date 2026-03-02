const fs = require('fs');
require('./wasm_exec.js');

const buf = fs.readFileSync('main.wasm');
const go = new Go();

(async function() {
    const inst = await WebAssembly.instantiate(buf, go.importObject);
    go.run(inst.instance);
})()
