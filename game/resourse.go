package game

import (
	"fmt"
	"strings"
)

// Получить переведенные на русский язык предметы и связанные с ними предметы
func getTranslateAllItem(items []item) string {

	mp := make(map[string][]string)

	var ar []string
	for _, item := range items {

		for _, it := range item.cantains {

			ar = append(ar, it.name)
		}
		// Перевод в Русский
		ar = EngineParseEnToRuMap(ar)
		name := EngineParseEnToRu(item.name)

		mp[name] = ar

	}

	// Здесь мы собираем Русский текст в одно целое для вывода
	var parts []string
	for name, contains := range mp {

		if len(contains) > 0 {
			parts = append(parts, fmt.Sprintf("%s: %s", name, strings.Join(contains, ", ")))
		} else {
			parts = append(parts, name)
		}
	}
	s := strings.Join(parts, ", ")

	return s
}

// Получить все конечные элементы, такие как ключи и тд
func getAllItems(items []item) []string {
	mp := make(map[string][]string)

	var ar []string
	for _, item := range items {

		for _, it := range item.cantains {

			ar = append(ar, it.name)
		}
		// Перевод в Русский
		// ar = EngineParseEnToRuMap(ar)
		// name := EngineParseEnToRu(item.name)

		mp[item.name] = ar

	}

	// Здесь мы собираем Русский текст в одно целое для вывода
	var parts []string
	for _, contains := range mp {
		parts = append(parts, contains...)
	}

	return parts
}

func in_array[T comparable](val T, arr []T) bool {

	for i := 0; i < len(arr); i++ {
		if val == arr[i] {
			return true
		}
	}

	return false
}

// Движок, транслейт ru to en
func EngineParseRuToEn(val string) string {

	var mov = Plocation{}
	var takeIt = Paction{}
	var applyOn = Pitem{}

	var arr []Parser = []Parser{mov, takeIt, applyOn}

	a := ""
	for _, i := range arr {
		a = Parse(i, val)

		if a != "unknown" {
			break
		}
	}

	return a
}

// Движок, транслейт en to ru
func EngineParseEnToRu(val string) string {

	var mov = PRlocation{}
	var takeIt = PRaction{}
	var applyOn = PRitem{}

	var arr []ParserReverce = []ParserReverce{mov, takeIt, applyOn}

	a := ""
	for _, i := range arr {
		a = ParseReverce(i, val)

		if a != "unknown" {
			break
		}
	}

	return a
}

// Движок, транслейт en to ru
func EngineParseEnToRuMap(vals []string) []string {

	var mov = PRlocation{}
	var takeIt = PRaction{}
	var applyOn = PRitem{}

	var arr []ParserReverce = []ParserReverce{mov, takeIt, applyOn}

	var ar []string

	for _, i := range arr {
		for _, val := range vals {

			v := ParseReverce(i, val)
			if v != "unknown" {
				// fmt.Println(val)
				ar = append(ar, v)
			}

		}

	}

	return ar
}

// Когда я что то забираю из комнаты, я с ней взаимодействую поэтому нужно взаимодейтсвовать
func deleteItemFromLocation(itemVal string, initWorld map[string]location) {

	// Создаем копию
	var it []item
	// var itemCopy item

	// Если предмет удален и больше парент предмет не содержит предметов то нужно удалить и его
	for _, val := range initWorld[Pl.location].items {
		var filteredCantains []item

		for _, in := range val.cantains {
			if itemVal != in.name {
				filteredCantains = append(filteredCantains, in)
			}
		}

		// Если после фильтрации есть дочерние элементы, добавляем
		if len(filteredCantains) > 0 {
			itemCopy := item{
				name:     val.name,
				cantains: filteredCantains,
			}
			it = append(it, itemCopy)
		}
		// Если filteredCantains пуст, то пропускаем — элемент полностью удалён
	}

	loc := initWorld[Pl.location]
	loc.items = it
	initWorld[Pl.location] = loc

	// fmt.Println(initWorld[Pl.location])
}

func deleteItemFromBackpack(itemVal string, backpack *[]string) {

	var a []string
	for _, i := range *backpack {
		if i != itemVal {
			a = append(a, i)
		}
	}

	*backpack = a
}
