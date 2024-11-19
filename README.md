# Make HTTP requests from inside WASM in Rye-lang

This devcontainer is configured to provide you a latest stable version of Go toolset.

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/wasm-outbound-http-examples/rye)

As the Rye [embedding API](https://github.com/refaktor/rye-website/blob/1bea85013b9873168952683ef4472dc9bf836f18/content/cookbook/deployment/console.md?plain=1#L42) 
is not released at the time of this writing ([some proof](https://www.reddit.com/r/golang/comments/1adtt0f/gos_concurrency_and_more_go_in_a_dynamic_language/),
look at *I also tested Rye as an embedded language in a Go app* part),
use the following oversimplified instruction to make HTTP request from the Rye WASM Playground 
([its source code](https://github.com/refaktor/rye-website/tree/1bea85013b9873168952683ef4472dc9bf836f18/content/ryeshell)):

1. Navigate your browser to official Ryelang website: [https://ryelang.org/](https://ryelang.org/).

2. Click darkgrey `Rye console` button on the top-right of your screen to open the Rye Playground (about 23M to download).

3. When the prompt and the blinking cursor appeared, type the following snippet into Rye Playground console:

```clojure
go does { get https://httpbin.org/get |print }
```

4. Press ENTER and see the resulting output in Rye Playground console (and browser developer console too).

<sub>Created for (wannabe-awesome) [list](https://github.com/vasilev/HTTP-request-from-inside-WASM)</sub>
