package eval

import (
	"github.com/sapphire-cord/eval/utils"
	"github.com/sapphire-cord/sapphire"
	"github.com/starlight-go/starlight/convert"
	"go.starlark.net/starlark"
	"strings"
)

// Evaluates Arbitrary Python
// Usage: <code:string...>
// Aliases: py, python, sl
func Eval(ctx *sapphire.CommandContext) {
	// Hacky way but because strings.Split messes up indents.
	// For other eval backends it wasn't needed due to them not being whitespace sensitive
	// But this one is, so we have to get it exactly as the user entered it.
	code := utils.StripCodeBlock(strings.Trim(ctx.Message.Content[len(ctx.Prefix)+len(ctx.InvokedName):], " "))

	stdout := []string{}

	thread := &starlark.Thread{Name: "eval", Print: func(_ *starlark.Thread, msg string) {
		stdout = append(stdout, msg)
	}}

	dict, err := convert.MakeStringDict(map[string]interface{}{
		"ctx":     ctx,
		"bot":     ctx.Bot,
		"session": ctx.Session,
	})

	// Should not happen.
	if err != nil {
		ctx.Error(err)
		return
	}

	_, err = starlark.ExecFile(thread, "eval", code, dict)

	if err != nil {
		if trace, ok := err.(*starlark.EvalError); ok {
			ctx.CodeBlock("py", trace.Backtrace())
		} else {
			ctx.CodeBlock("py", err.Error())
		}

		return
	}

	if len(stdout) < 1 {
		ctx.Reply("No output returned.")
		return
	}

	ctx.CodeBlock("py", "%+v", strings.Join(stdout, "\n"))
}

func Init(bot *sapphire.Bot, name, category string, aliases []string) {
	bot.AddCommand(sapphire.NewCommand(name, category, Eval).
		AddAliases(aliases...).
		SetUsage("<code:string...>").
		SetDescription("Evaluates Arbitrary starlark").
		SetOwnerOnly(true))
}
