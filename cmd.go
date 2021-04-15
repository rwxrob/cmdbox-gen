package cmd

import (
	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.New("gen")
	x.Summary = `generates KEG content, mainly index.html and README.md`
	x.Usage = ``
	x.Version = `v0.0.1`

	x.Description = `
		The *gen* command generates static content --- primarily index.html
		and README.md --- from KEG content. It is usually used when composed
		into the *kn* CmdBox monolith.`

	x.Method = func(args []string) (err error) {
		// TODO
		return cmdbox.ErrorUnimplemented()
	}
}
