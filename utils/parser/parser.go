package parser

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/takez0o/honestwork-api/utils/loggersentry"
)

func Parse(content string) string {
	fmt.Println("Parsing:", content)
	m := map[string]interface{}{}
	err := json.Unmarshal([]byte(content), &m)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error())
		log.Fatal(err)
	}
	return parseMapping(m, "")
}

func parseMapping(m map[string]interface{}, carry string) string {
	for _, value := range m {
		switch v := value.(type) {
		case []interface{}:
			for _, sub_v := range v {
				if sub_v.(map[string]interface{})["content"] != nil {
					for key, value := range sub_v.(map[string]interface{}) {
						if key == "type" && (value == "paragraph" || value == "heading") {
							for _, j := range sub_v.(map[string]interface{})["content"].([]interface{}) {
								for z, u := range j.(map[string]interface{}) {
									if z == "text" {
										carry += u.(string)
									}
								}
							}
						} else if key == "type" && value == "bulletList" {
							for _, i := range sub_v.(map[string]interface{})["content"].([]interface{}) {
								for _, p := range i.(map[string]interface{})["content"].([]interface{}) {
									for _, t := range p.(map[string]interface{})["content"].([]interface{}) {
										for x, y := range t.(map[string]interface{}) {
											if x == "text" {
												carry += y.(string)
											}
										}
									}
								}
							}

						}
					}
				}
			}
		}
	}
	return carry
}
