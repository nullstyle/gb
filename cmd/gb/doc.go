package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/constabulary/gb"
	"github.com/constabulary/gb/cmd"
)

func init() {
	registerCommand(&cmd.Command{
		Name:      "doc",
		UsageLine: `doc <pkg> <sym>[.<method>]`,
		Short:     "show documentation for a package or symbol",
		Long: `
Doc shows documentation for a package or symbol.

See 'go help doc'.
`,
		Run: func(ctx *gb.Context, args []string) error {
			env := cmd.MergeEnv(os.Environ(), map[string]string{
				"GOPATH": fmt.Sprintf("%s:%s", ctx.Projectdir(), filepath.Join(ctx.Projectdir(), "vendor")),
			})
			if len(args) == 0 {
				args = append(args, ".")
			}
			args = append([]string{filepath.Join(ctx.Context.GOROOT, "bin", "godoc")}, args...)

			cmd := exec.Cmd{
				Path: args[0],
				Args: args,
				Env:  env,

				Stdin:  os.Stdin,
				Stdout: os.Stdout,
				Stderr: os.Stderr,
			}
			return cmd.Run()
		},
		SkipParseArgs: true,
	})
}
