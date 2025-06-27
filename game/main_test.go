package game

import (
	"fmt"
	"strings"
	"testing"
)

func TestGame0(t *testing.T) {

	fn := functionDependenceOnAction()

	for caseNum, commands := range game0cases {

		mp := InitGame()

		if caseNum == 0 {
			Pl.location = "kitchen"
		} else if caseNum == 1 {
			Pl.location = "kitchen"
		}

		for _, item := range commands {

			answer := ""
			parseString := strings.Split(item.command, " ")
			var params []string

			if len(parseString) > 0 {
				// Берем команду
				commandRu := parseString[0]
				commandEn := EngineParseRuToEn(commandRu)

				// Собираем параметры
				for i := 1; i < len(parseString); i++ {
					params = append(params, parseString[i])
				}

				// Выполняем команду
				if handler, ok := fn[commandEn]; ok {
					answer = handler(params, mp)

					fmt.Println(answer)
				}

			}

			if answer != item.answer {
				t.Error("case:", caseNum, item.step,
					"\n\tcmd:", item.command,
					"\n\tresult:  ", answer,
					"\n\texpected:", item.answer)
			}
		}
	}

}
