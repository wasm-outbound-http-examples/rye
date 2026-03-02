import * as _ from './wasm_exec.js';
import { readAll } from 'jsr:@std/io@0.225.2';
const go = new Go();
const f = await Deno.open('main.wasm');
const buf = await readAll(f);
const inst = await WebAssembly.instantiate(buf, go.importObject);
go.run(inst.instance);
