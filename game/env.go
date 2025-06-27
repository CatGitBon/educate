package game

// Игра
// Обьекты: игрок, локации, дейтсвия, предметы
type player struct {
	age            int
	gender         string
	name           string
	condition      string
	backpackExists bool
	backpack       []string
}

type item struct {
	name     string
	cantains []item
}

type location struct {
	name       string
	temprature float64
	items      []item
}

type plays struct {
	location string
	player   string
}

type door struct {
	state bool
}

// Правила
func rulesMoving() map[string][]string {
	var mov = make(map[string][]string)

	mov["room"] = []string{"hallway"}
	mov["kitchen"] = []string{"hallway"}
	mov["hallway"] = []string{"kitchen", "room", "street"}
	mov["street"] = []string{"hallway", "home"}

	return mov
}

func rulesApplyOn() map[string][]string {
	var apl = make(map[string][]string)

	apl["door"] = []string{"key"}

	return apl
}

func functionDependenceOnAction() map[string]func([]string, map[string]location) string {
	var rul = make(map[string]func([]string, map[string]location) string)

	rul["put_on"] = funcPutOn
	rul["take_it"] = funcTakeIt
	rul["go_to"] = funcGoTo
	rul["looking"] = funcLooking
	rul["apply_on"] = funcApplyOn
	rul["breakfast"] = breakfast

	return rul
}
