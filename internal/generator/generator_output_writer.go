package generator

import (
	"fmt"
	"io"
	"text/template"
)

func (output *Output) Write(writer io.Writer) (err error) {
	err = output.writeHeaderComments(writer)
	if err != nil {
		return err
	}

	err = output.writePackageName(writer)
	if err != nil {
		return err
	}

	err = output.writeImports(writer)
	if err != nil {
		return err
	}

	err = output.writeStoreKeys(writer)
	if err != nil {
		return err
	}

	err = output.writeStoreTypes(writer)
	if err != nil {
		return err
	}

	err = output.writeMessagesFuncs(writer)
	if err != nil {
		return err
	}

	return nil
}

func writeNewLine(writer io.Writer) error {
	_, err := writer.Write([]byte("\n"))
	return err
}

func (output *Output) writeHeaderComments(writer io.Writer) error {
	commentsTemplateString := `{{range $comment := .HeaderComments}}
// {{$comment}}{{end}}
`

	commentTemplate, err := template.New("comment-template").Parse(commentsTemplateString)
	if err != nil {
		return err
	}

	err = commentTemplate.Execute(writer, output)
	if err != nil {
		return err
	}

	return writeNewLine(writer)
}

func (output *Output) writePackageName(writer io.Writer) error {
	_, err := writer.Write([]byte(fmt.Sprintf("package %s\n", output.PackageName)))

	if err != nil {
		return err
	}

	return writeNewLine(writer)
}

func (output *Output) writeImports(writer io.Writer) error {
	importsTemplateString := `import ({{range $import := .Imports}}
	{{$import}}{{end}}
)
`

	importsTemplate, err := template.New("imports-template").Parse(importsTemplateString)
	if err != nil {
		return err
	}

	err = importsTemplate.Execute(writer, output)
	if err != nil {
		return err
	}

	return writeNewLine(writer)
}

func (output *Output) writeStoreTypes(writer io.Writer) error {
	_, err := writer.Write([]byte("type StoreSoul []byte\n"))
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(`type Store struct {
	db *bolt.DB	
}`))
	if err != nil {
		return err
	}

	return writeNewLine(writer)
}

func (output *Output) writeStoreKeys(writer io.Writer) error {
	keysTemplateString := `{{$package := .ProtoPackage}}const ({{range $key := .Messages}}
	{{$key.Name}}Key string = "{{if $package }}{{$package}}.{{end}}{{$key.Name}}"{{end}}
)
`

	keysTemplate, err := template.New("keys-template").Parse(keysTemplateString)
	if err != nil {
		return err
	}

	err = keysTemplate.Execute(writer, output)
	if err != nil {
		return err
	}

	return writeNewLine(writer)
}

func (output *Output) writeMessagesFuncs(writer io.Writer) error {
	funcsTemplateString := `func (store *Store) Set{{ .Name }}(soul StoreSoul, data *{{ .Name }}) error {
	return store.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte({{ .Name }}Key))
		if err != nil {
			return err
		}

		dataBytes, err := proto.Marshal(data)
		if err != nil {
			return err
		}

		{{ if .HasID }}
		soulBucket, err := bucket.CreateBucketIfNotExists(soul)
		if err != nil {
			return err
		}

		return soulBucket.Put([]byte(data.GetId()), dataBytes)
		{{ else }}
		return bucket.Put(soul, dataBytes)
	{{end}}})
}

func (store *Store) Get{{ .Name }}(soul SoulStore) (result {{ if .HasID }}[]{{ end }}*{{ .Name }}, err error) {
	err = store.db.View(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte({{ .Name }}Key))
		if err != nil {
			return err
		}

		{{ if .HasID }}
		soulBucket, err := bucket.CreateBucketIfNotExists(soul)
		if err != nil {
			return err
		}

		cursor := soulBucket.Cursor()

		for key, data := cursor.First(); key != nil; key, data = cursor.Next() {
			value := &{{ .Name }}{}

			err := proto.Unmarshal(data, value)
			if err != nil {
				return err
			}

			result = append(result, value)
		}

		return nil
		{{ else }}
		data := bucket.Get(soul)
		return proto.Unmarshal(data, result)
	{{ end }}})

	return result, err
}
`

	funcsTemplate, err := template.New("funcs-template").Parse(funcsTemplateString)
	if err != nil {
		return err
	}

	for _, message := range output.Messages {
		err = funcsTemplate.Execute(writer, message)
		if err != nil {
			return err
		}

		err = writeNewLine(writer)
		if err != nil {
			return err
		}
	}

	return nil
}
