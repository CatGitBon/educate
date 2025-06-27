package game

type Parser interface {
	parse(locate string) string
}

type Plocation struct{}
type Paction struct{}
type Pitem struct{}

func Parse(p Parser, val string) string {
	return p.parse(val)
}

func (pl Plocation) parse(locate string) string {
	switch locate {
	case "комната":
		return "room"
	case "кухня":
		return "kitchen"
	case "коридор":
		return "hallway"
	case "улица":
		return "street"
	default:
		return "unknown"
	}
}

func (pl Paction) parse(action string) string {
	switch action {
	case "идти":
		return "go_to"
	case "применить":
		return "apply_on"
	case "осмотреться":
		return "looking"
	case "взять":
		return "take_it"
	case "надеть":
		return "put_on"
	case "завтракать":
		return "breakfast"
	default:
		return "unknown"
	}
}

func (pl Pitem) parse(item string) string {
	switch item {
	case "ключи":
		return "key"
	case "рюкзак":
		return "backpack"
	case "конспекты":
		return "notes"
	case "дверь":
		return "door"
	case "на столе":
		return "table"
	case "на стуле":
		return "chair"
	case "чай":
		return "tea"
	case "телефон":
		return "phone"
	default:
		return "unknown"
	}
}
