package generator

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func GenerateFile(plugin *protogen.Plugin, file *protogen.File) error {
	output_file_path := fmt.Sprintf("%s.store.go", file.GeneratedFilenamePrefix)
	output_file := plugin.NewGeneratedFile(output_file_path, ".")

	output_file.Write([]byte("package pb"))
	return nil
}
