package game

type ParserReverce interface {
	parseReverce(locate string) string
}

type PRlocation struct{}
type PRaction struct{}
type PRitem struct{}

func ParseReverce(p ParserReverce, val string) string {
	return p.parseReverce(val)
}

func (pl PRlocation) parseReverce(locate string) string {
	switch locate {
	case "room":
		return "комната"
	case "kitchen":
		return "кухня"
	case "hallway":
		return "коридор"
	case "street":
		return "улица"
	default:
		return "unknown"
	}
}

func (pl PRaction) parseReverce(action string) string {
	switch action {
	case "go_to":
		return "идти"
	case "apply_on":
		return "применить"
	case "looking":
		return "осмотреться"
	case "take_it":
		return "взять"
	case "put_on":
		return "надеть"
	case "breakfast":
		return "завтракать"
	default:
		return "unknown"
	}
}

func (pl PRitem) parseReverce(item string) string {
	switch item {
	case "key":
		return "ключи"
	case "backpack":
		return "рюкзак"
	case "notes":
		return "конспекты"
	case "door":
		return "дверь"
	case "table":
		return "на столе"
	case "chair":
		return "на стуле"
	case "tea":
		return "чай"
	case "phone":
		return "телефон"
	default:
		return "unknown"
	}
}
