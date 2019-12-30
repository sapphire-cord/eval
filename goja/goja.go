package eval

import (
	"github.com/dop251/goja"
	"github.com/sapphire-cord/eval/utils"
	"github.com/sapphire-cord/sapphire"
)

func Eval(ctx *sapphire.CommandContext) {
	vm := goja.New()
	vm.Set("ctx", ctx)
	vm.Set("bot", ctx.Bot)
	vm.Set("session", ctx.Session)

	value, err := vm.RunString(utils.StripCodeBlock(ctx.JoinedArgs()))

	if err != nil {
		ctx.CodeBlock("js", "%s", err)
		return
	}

	ctx.CodeBlock("js", "%s", value)
}

func Init(bot *sapphire.Bot, name, category string, aliases []string) {
	bot.AddCommand(sapphire.NewCommand(name, category, Eval).
		AddAliases(aliases...).
		SetDescription("Evaluates arbitrary JavaScript").
		SetUsage("<code:string...>").
		SetOwnerOnly(true))
}
