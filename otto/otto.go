package eval

import (
  "github.com/sapphire-cord/sapphire"
  "github.com/sapphre-cord/eval/utils"
  "github.com/robertkrimen/otto"
)

func Eval(ctx *sapphire.CommandContext) {
  vm := otto.New()

  vm.Set("ctx", ctx)
  vm.Set("bot", ctx.Bot)
  vm.Set("session", ctx.Session)

  value, err := vm.Run(utils.StripCodeBlock(ctx.JoinedArgs()))

  if err != nil {
    ctx.CodeBlock("js", "%s", err)
    return
  }

  ctx.CodeBlock("js", "%s", value)
}

func Init(bot *sapphire.Bot, name, category string, aliases []string) {
  bot.AddCommand(sapphire.NewCommand(name, category, Eval).
    SetUsage("<code:string...>").
    SetDescription("Evaluates arbitrary JavaScript").
    AddAliases(aliases...).
    SetOwnerOnly(true))
}
