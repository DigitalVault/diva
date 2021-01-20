package cli

import (
  "fmt"
  "runtime/debug"
)

func Version() string {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		panic("couldn't read build info")
	}

	fmt.Printf("%s version %s\n", bi.Path, bi.Main.Version)

	for _, d := range bi.Deps {
		fmt.Printf("\tbuilt with %s version %s\n", d.Path, d.Version)
	}
  return "v0.2"
}
