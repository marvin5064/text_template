package main

import (
	"fmt"
	"os"
	"text/template"
)

type Data struct {
	SpecialId    string
	ContestantId string
	LineId       int
}

func main() {
	fmt.Println("Begin:")
	templateSample := `
	current_total_array_number: {{ len (.) }}
	{{ $total_number := len (.)}}
	{{ range $index, $element := .}}
	{{ if not $index }}{
	"bets":	[
	{{end}}		{
				"accpetBetterLine": true,
				"oddsFormat": "DECIMAL",
				"ContestantId": "{{ .ContestantId }}",
				"uniqueRequestId": "request_for_money",
				"SpecialId": "{{ .SpecialId }}",
				"winRiskStake": "RISK",
				"LineId": {{ .LineId }}
			Extra info:
				current_index: {{ $index }}
				current_total: {{ $total_number }}
			}{{if lt $index (minus1 $total_number) }},{{ else }}
		]
	}{{end}}{{end}}
	`
	fmt.Println("The template for testing:", templateSample)
	t, err := template.New("pinnacle_bet").Funcs(template.FuncMap{
		"minus1": func(x int) int {
			return x - 1
		},
	}).Parse(templateSample)
	if err != nil {
		panic(err)
	}

	data := defineDataForTemplate()
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func defineDataForTemplate() []Data {
	return []Data{
		Data{
			SpecialId:    "SpecialId 1",
			ContestantId: "ContestantId 1",
			LineId:       1,
		},
		Data{
			SpecialId:    "SpecialId 2",
			ContestantId: "ContestantId 2",
			LineId:       2,
		},
		Data{
			SpecialId:    "SpecialId 3",
			ContestantId: "ContestantId 3",
			LineId:       3,
		},
	}
}
