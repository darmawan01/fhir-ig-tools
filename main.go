package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"ig_tools/models"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var (
	templatePath string
	resourcePath string
	resourceOut  string
)

func init() {
	flag.StringVar(&templatePath, "template", "", "Specify the Implementation Guide")
	flag.StringVar(&resourcePath, "resource", "", "Specify the Resource destination")
	flag.StringVar(&resourceOut, "out", "", "Specify the Out Resource file name")
	flag.Parse()

	if templatePath == "" || resourcePath == "" {
		panic("Please pass template and resource file path.")
	}
}

func main() {
	template, err := excelize.OpenFile(templatePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	resource, err := ioutil.ReadFile(resourcePath)
	if err != nil {
		fmt.Printf("Error when opening file: %s ", err.Error())
	}

	var resourceData models.BaseModel
	err = json.Unmarshal(resource, &resourceData)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	rows := template.GetRows("Sheet1")
	for _, row := range rows {
		if len(row) >= 2 {
			for i, item := range resourceData.Differential.Element {
				if item.Path == row[0] {
					resourceData.Differential.Element[i].Definition = row[1]

					if len(row[2]) > 0 {
						val, err := strconv.Atoi(row[2])
						if err != nil {
							log.Println("Invalid Min value")
						}
						resourceData.Differential.Element[i].Min = val
					}

					if len(row[3]) > 0 {
						resourceData.Differential.Element[i].Max = row[3]
					}

					if len(row[4]) > 0 {
						resourceData.Differential.Element[i].Short = row[4]
					}

					if len(row[5]) > 0 {
						resourceData.Differential.Element[i].Comment = row[5]
					}
				}
			}
		}
	}

	outFilePath := "out.json"
	if resourceOut != "" {
		outFilePath = resourceOut
	}

	file, _ := json.MarshalIndent(resourceData, "", " ")
	if err = ioutil.WriteFile(outFilePath, file, 0644); err != nil {
		panic(err)
	}

}
