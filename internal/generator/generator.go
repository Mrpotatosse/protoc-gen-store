package generator

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func GenerateFile(plugin *protogen.Plugin, file *protogen.File) error {
	outputFilePath := fmt.Sprintf("%s.store.go", file.GeneratedFilenamePrefix)
	outputFile := plugin.NewGeneratedFile(outputFilePath, ".")

	_, err := outputFile.Write([]byte("package pb"))
	return err
}
