# Sapphire Eval
This repository contains a few backends for an eval command for your sapphire bots.

## What is eval?
Eval is a command that executes dynamic code on the bot's host, it's very dangerous if given access to anyone untrusted but for the owner's personal use it's very powerful and lets you test some stuff quickly.

If you are coming from discord.js you sure heard of it, in Go however the language is compiled and doesn't have runtime execution of code so we need to plug-in an embeddable scripting language and pass values in between to allow dynamic behaviour, this repository does just that, it includes some backends for the most popular scripting engines to plug in an eval command to your bots.

## Install
Install your preferred backend and call some initialization in your bot, the available backends currently are:
- **goja** - JavaScript VM in Go, supports ES5 only.
- **otto** - JavaScript VM in Go, supports ES5 only, lighter than Goja with no dependencies but slower.
- **anko** - A Go like interpreter, this is the closest thing you can get to a Go eval.
- **gopher-lua** - Lua VM in Go.

```sh
$ go get github.com/sapphire-cord/eval/<your-backend>
```
Replace `<your-backend>` with one of the above and import it in your entry file and initialize it, for example
```go
// import
import (
  "github.com/sapphire-cord/eval/<your-backend>"
)

// wherever you made a bot instance also initialize eval before logging in.
eval.Init(bot, "eval", "Owner", []string{"ev"})
```
Notes:
- The first argument is the bot instance, the second is the category, the third is the command name to use, the fourth is the aliases.
- Again replace `<your-backend>` with the backend you previously installed.
- Each backend exposes the package with the name `eval` which can conflict if you want to use multiple backends, prefer to use go import aliases, **see below**

### Using multiple backends
Sometimes you want to have more eval backends, we made it easy by allowing you to specify the name, category and aliases when registering so it doesn't conflict, now only the packages conflict but you can use an import alias:
```go
import (
  one "github.com/sapphire-cord/eval/backend1"
  two "github.com/sapphire-cord/eval/backend2"
)

// When initializing.
one.Init(bot, "eval1", "Owner", []string{"ev1"})
two.Init(bot, "eval2", "Owner", []string{"ev2"})
```

## Using the command
After initialized call the command with your prefix + whatever name you choosed at initialization and pass it a script of the choosen backend language, additionally these extra variables are exposed globally in the context of eval to make it more useful:
- `ctx` The command context.
- `bot` Alias to ctx.Bot
- `session` Alias to ctx.Session

Which lets you do almost anything you can do in Go, only dynamically from Discord!

## More
Feel free to suggest or contribute new backends.

## License
[MIT License](LICENSE)
