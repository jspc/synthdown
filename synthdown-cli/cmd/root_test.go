package cmd

import (
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {
	os.MkdirAll("~/.config/synthdown", 0700)

	for _, fn := range []string{
		"",
		"testdata/config.yaml",
		"testdata/nonsuch",
	} {
		t.Run(fn, func(t *testing.T) {
			cfgFile = fn

			defer func() {
				err := recover()
				if err != nil {
					t.Fatal(err)
				}
			}()

			initConfig()
		})
	}
}
