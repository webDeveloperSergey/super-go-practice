package output

import "github.com/fatih/color"

func PrintError(msg any) {
	switch t := msg.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Неизвестная ошибка")
	}
}