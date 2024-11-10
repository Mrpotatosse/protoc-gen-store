package main

import (
	"github.com/Mrpotatosse/protoc-gen-store/internal/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

// This controls the maxprocs environment variable in container runtimes.
// see https://martin.baillie.id/wrote/gotchas-in-the-go-network-packages-defaults/#bonus-gomaxprocs-containers-and-the-cfs

func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		for _, file := range plugin.Files {
			err := generator.GenerateFile(plugin, file)

			if err != nil {
				return err
			}
		}
		return nil
	})
}
