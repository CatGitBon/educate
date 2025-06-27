package game

import (
	"fmt"
	"strings"
)

// Осмотреться
func funcLooking(arg []string, initWorld map[string]location) string {

	// Нужно тут взять глобальную переменную где сейчас игрок
	location := Pl.location
	// Нужно определять локацию куда можно пойти
	mov := rulesMoving()

	// Если комната пустая, то

	text := ""
	switch location {
	case "room":

		if len(initWorld[location].items) == 0 {
			return fmt.Sprintf(
				"пустая комната. можно пройти - %s",
				strings.Join(EngineParseEnToRuMap(mov[initWorld[location].name]), ", "),
			)
		}

		text = fmt.Sprintf(
			"%s. можно пройти - %s",
			getTranslateAllItem(initWorld[location].items),
			strings.Join(EngineParseEnToRuMap(mov[initWorld[location].name]), ", "),
		)
	case "kitchen":
		text = fmt.Sprintf(
			"ты находишься на кухне, %s, надо собрать рюкзак и идти в универ. можно пройти - %s",
			getTranslateAllItem(initWorld[location].items),
			strings.Join(EngineParseEnToRuMap(mov[initWorld[location].name]), ", "),
		)
	case "hallway":

		text = fmt.Sprintf(
			"ничего интересного. можно пройти - %s",
			strings.Join(EngineParseEnToRuMap(mov[initWorld[location].name]), ", "),
		)

	case "street":

		// Костыль на тему - коридор это дом
		where := strings.Join(EngineParseEnToRuMap(mov[initWorld[location].name]), ", ")

		if where == "коридор" {
			where = "домой"
		}

		text = fmt.Sprintf(
			"на улице весна. можно пройти - %s",
			where,
		)

	default:
		text = "Где ты в этом мире, путник"
	}

	return text
}

// Идти куда-то
func funcGoTo(arg []string, initWorld map[string]location) string {

	// Нужно тут взять глобальную переменную где сейчас игрок
	locationFrom := Pl.location
	locationTo := EngineParseRuToEn(arg[0])
	// Нужно определять локацию куда можно пойти
	mov := rulesMoving()

	// Проверка сможем мы идти туда или нет
	to := mov[locationFrom]
	if !in_array(locationTo, to) {
		return fmt.Sprintf("нет пути в %s", EngineParseEnToRu(locationTo))
	}

	// На вывод текста
	text := ""
	// Переходим в локацию
	switch locationTo {
	case "room":
		text = fmt.Sprintf(
			"ты в своей комнате. можно пройти - %s",
			strings.Join(EngineParseEnToRuMap(mov[initWorld[locationTo].name]), ", "),
		)
		Pl.location = locationTo
	case "kitchen":
		text = fmt.Sprintf(
			"кухня, ничего интересного. можно пройти - %s",
			strings.Join(EngineParseEnToRuMap(mov[initWorld[locationTo].name]), ", "),
		)
		Pl.location = locationTo
	case "hallway":
		text = fmt.Sprintf(
			"ничего интересного. можно пройти - %s",
			strings.Join(EngineParseEnToRuMap(mov[initWorld[locationTo].name]), ", "),
		)
		Pl.location = locationTo

		if locationFrom == "street" {
			Door.state = false
		}

	case "street":
		if !Door.state {
			text = "дверь закрыта"
		} else {

			// Костыль на тему - коридор это дом
			where := strings.Join(EngineParseEnToRuMap(mov[initWorld[locationTo].name]), ", ")

			if where == "коридор" {
				where = "домой"
			}

			text = fmt.Sprintf(
				"на улице весна. можно пройти - %s",
				where,
			)

			Pl.location = locationTo
		}
	}

	return text
}

// Взять что-то
func funcTakeIt(arg []string, initWorld map[string]location) string {

	// Здесь мы берем инвентарь если у нас есть рюкзак
	if !Player.backpackExists {
		return "некуда класть"
	}

	item := EngineParseRuToEn(arg[0])

	loc := initWorld[Pl.location]

	items := getAllItems(initWorld[loc.name].items)

	if !in_array(item, items) {
		return "нет такого"
	}

	// Нужно взять что то -> изменить состояние юзера
	Player.backpack = append(Player.backpack, item)

	// Удаляем из комнаты предмет
	deleteItemFromLocation(item, initWorld)

	return fmt.Sprintf("предмет добавлен в инвентарь: %s", EngineParseEnToRu(item))
}

// Применить что-то
func funcApplyOn(arg []string, initWorld map[string]location) string {

	// Тут мы делаем взаимодействие предметами из инвентаря
	itemWho := EngineParseRuToEn(arg[0])
	itemWhom := EngineParseRuToEn(arg[1])

	if Player.backpackExists && len(Player.backpack) == 0 {
		return fmt.Sprintf("нет предмета в инвентаре - %s", EngineParseEnToRu(itemWho))
	}

	// Далее если все таки есть какой то предмет, то ищем его в рюкзаке и проверяем подойдет ли он
	if !in_array(itemWho, Player.backpack) {
		return fmt.Sprintf("нет предмета в инвентаре - %s", EngineParseEnToRu(itemWho))
	}

	// Правила применения предметов в мире
	apl := rulesApplyOn()

	// Если все-таки предмет есть, то применяем его
	// Ищем то что можем применить для предмета по правилам
	whoItems := apl[itemWhom]

	// Если нельзя применить по правилам мира один предмет к другому
	if !in_array(itemWho, whoItems) {
		return "не к чему применить"
	}

	// Если все-таки все успешно и можно применить, то
	Door.state = true

	// Изменить состояние рюкзака - оказывается ключи остаются в замке
	// deleteItemFromBackpack(itemWho, &Player.backpack)

	return fmt.Sprintf("%s открыта", EngineParseEnToRu(itemWhom))
}

// Надеть что-то
func funcPutOn(arg []string, initWorld map[string]location) string {

	item := EngineParseRuToEn(arg[0])

	// Нужно надеть что то -> изменить состояние юзера
	Player.backpackExists = true

	// fmt.Println(item)

	// Изменить состояние мира - рюкзака больше нет в комнате
	deleteItemFromLocation(item, initWorld)

	return fmt.Sprintf("вы надели: %s", EngineParseEnToRu(item))
}

func breakfast(arg []string, initWorld map[string]location) string {
	return "неизвестная команда"
}
