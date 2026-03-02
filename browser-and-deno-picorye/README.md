# Use PicoRye to send HTTP(s) requests from inside WASM

## Instructions for this devcontainer

Tested with Go 1.26.0, Bun 1.3.10, Deno 2.7.1, PicoRye [commit 35eb4f5f](https://github.com/refaktor/picorye/tree/35eb4f5fb9e8f33bfebe0a342eeb12a443308c71).

### Preparation

1. Open this repo in devcontainer, e.g. using Github Codespaces.
   Type or copy/paste following commands to devcontainer's terminal.

### Building

1. `cd` into the folder of this example:

```sh
cd browser-and-deno-picorye
```

2. Since picorye is not published (has no version tags) at the time of this writing, need to clone its repo:

```sh
git clone --depth=1 https://github.com/refaktor/picorye.git
```

3. Copy the example and its helper files to the folder of cloned repo:

```sh
cp httpget.go *.js index.html picorye/
```

4. `cd` into the folder of cloned repo:

```sh
cd picorye
```

5. Compile the example:

```sh
GOOS=js GOARCH=wasm go build -o main.wasm httpget.go
```

6. Copy the glue JS from Golang distribution to example's folder (note using `/lib/wasm/` because of Go 1.24+):

```sh
cp $(go env GOROOT)/lib/wasm/wasm_exec.js ./
```

### Test with browser

1. Run simple HTTP server to temporarily publish project to Web:

```sh
~/.deno/bin/deno run --allow-net --allow-read jsr:@std/http/file-server
```

Codespace will show you "Open in Browser" button. Just click that button or
obtain web address from "Forwarded Ports" tab.

2. As `index.html` and a **10.5M**-sized `main.wasm` are loaded into browser, refer to browser developer console
   to see the results.

### Test with Node.js

Impossible yet due to https://github.com/golang/go/issues/59605.

### Test with Bun

1. Install Bun:

```sh
curl -fsSL https://bun.sh/install | bash
```

2. Run with Bun:

```sh
~/.bun/bin/bun bun.js
```

### Test with Deno

1. Run with Deno:

```sh
~/.deno/bin/deno run --allow-read --allow-net deno.js
```

### Finish

Perform your own experiments if desired.
