package eval

import (
  "github.com/sapphire-cord/sapphire"
  "github.com/sapphire-cord/eval/utils"
  "github.com/mattn/anko/env"
  "github.com/mattn/anko/vm"
  "bytes"
  "fmt"
)

func Eval(ctx *sapphire.CommandContext) {
  e := env.NewEnv()
  stdout := &bytes.Buffer{}

  // define also returns an error but it really shouldn't happen in our case so we ignore it.

  e.Define("println", func(args ...interface{}) {
    stdout.WriteString(fmt.Sprintln(args...))
  })

  e.Define("ctx", ctx)
  e.Define("bot", ctx.Bot)
  e.Define("session", ctx.Session)

  value, err := vm.Execute(e, nil, utils.StripCodeBlock(ctx.JoinedArgs()))
  if err != nil {
    ctx.CodeBlock("go", "%v\n", err)
    return
  }

  ctx.CodeBlock("go", "%s%+v", stdout.String(), value)
}

func Init(bot *sapphire.Bot, name, category string, aliases []string) {
  bot.AddCommand(sapphire.NewCommand(name, category, Eval).
    SetUsage("<code:string...>").
    SetDescription("Evaluates arbitrary Anko").
    AddAliases(aliases...).
    SetOwnerOnly(true))
}
