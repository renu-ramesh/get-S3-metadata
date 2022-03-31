package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"aws-lambda-in-go-lang/models"

	"github.com/alecthomas/jsonschema"
)

var topLevelModels = map[string]interface{}{
	"InputPayLoadParams": models.InputPayLoadParams{},
	"RequestParams":      models.RequestParams{},
}

//
// GenerateSchema
// This method will generat json schema for validation
func GenerateSchema() {

	var serializedSchema []byte
	var err error
	for k, v := range topLevelModels {
		fmt.Printf("generating json schema for %s\n", k)
		schema := jsonschema.Reflect(v)
		schema.Title = k

		serializedSchema, err = json.MarshalIndent(schema, "", "\t")
		if err != nil {
			panic(err)
		}

		func() {
			fileName := fmt.Sprintf("./schema/json/%s.json", k)
			f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			fmt.Println("writing ", fileName)
			_, err = f.Write(serializedSchema)
			if err != nil {
				panic(err)
			}
		}()
	}

}
