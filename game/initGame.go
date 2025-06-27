package game

import (
	"fmt"
	"strings"
)

var Pl = plays{
	location: "kitchen",
	player:   "Tom",
}
var Door = door{
	state: false,
}
var Player = player{
	age:            21,
	gender:         "male",
	name:           "Tom",
	condition:      "good",
	backpack:       []string{},
	backpackExists: false,
}

func InitGame() map[string]location {

	// // Инициализация мира
	var room location = location{
		name:       "room",
		temprature: 34.6,
		items: []item{
			{
				name: "table",
				cantains: []item{
					{
						name:     "key",
						cantains: []item{},
					},
					{
						name:     "notes",
						cantains: []item{},
					},
				},
			},
			{
				name: "chair",
				cantains: []item{
					{
						name:     "backpack",
						cantains: []item{},
					},
				},
			},
		},
	}
	var kitchen location = location{
		name:       "kitchen",
		temprature: 35,
		items: []item{
			{
				name: "table",
				cantains: []item{
					{
						name:     "tea",
						cantains: []item{},
					},
				},
			},
		},
	}
	var hallway location = location{
		name:       "hallway",
		temprature: 30,
		items:      []item{},
	}
	var street location = location{
		name:       "street",
		temprature: 25,
		items:      []item{},
	}
	var mp = make(map[string]location)

	mp["room"] = room
	mp["kitchen"] = kitchen
	mp["hallway"] = hallway
	mp["street"] = street

	return mp
}

func Play(test []gameCase, initWorld map[string]location) {

	fn := functionDependenceOnAction()

	for _, gc := range test {

		parseString := strings.Split(gc.command, " ")
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
				answer := handler(params, initWorld)

				fmt.Println(answer)
			}

		}

	}
}
