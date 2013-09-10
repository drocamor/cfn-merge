package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type CloudFormationTemplate struct {
	AWSTemplateFormatVersion string
	Description              string
	Mappings                 map[string]interface{}
	Outputs                  map[string]interface{}
	Parameters               map[string]interface{}
	Resources                map[string]interface{}
}

func main() {

	template_list := os.Args[1:]

	if len(template_list) < 2 {
		fmt.Println("You must provide at least two template snippets to merge")
		os.Exit(1)
	}

	// Create an empty template with a default description and version
	outputTemplate := CloudFormationTemplate{
		AWSTemplateFormatVersion: "2010-09-09",
		Description:              "TODO: Add a relevant description",
		Mappings:                 make(map[string]interface{}),
		Outputs:                  make(map[string]interface{}),
		Parameters:               make(map[string]interface{}),
		Resources:                make(map[string]interface{})}

	// Open each template and merge things in
	for i := 0; i < len(template_list); i++ {

		file, e := ioutil.ReadFile(template_list[i])
		if e != nil {
			fmt.Printf("File error: %v\n", e)
			os.Exit(1)
		}

		var template CloudFormationTemplate
		json.Unmarshal(file, &template)

		if template.Resources != nil {
			for k, v := range template.Resources {
				outputTemplate.Resources[k] = v
			}
		}

		if template.Parameters != nil {
			for k, v := range template.Parameters {
				outputTemplate.Parameters[k] = v
			}
		}

		if template.Mappings != nil {
			for k, v := range template.Mappings {
				outputTemplate.Mappings[k] = v
			}
		}

		if template.Outputs != nil {
			for k, v := range template.Outputs {
				outputTemplate.Outputs[k] = v
			}
		}
	}

	output, _ := json.MarshalIndent(outputTemplate, "  ", "  ")

	fmt.Printf("%s", output)

}
