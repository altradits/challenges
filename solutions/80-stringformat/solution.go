package piscine

import (
	"fmt"
	"strings"
)

// Format substitutes %s, %d and %f verbs with args in order. Unknown verbs
// and verbs without a matching arg are kept literally; extra args are ignored.
func Format(template string, args ...interface{}) string {
	var b strings.Builder
	arg := 0
	for i := 0; i < len(template); i++ {
		if template[i] != '%' || i+1 >= len(template) {
			b.WriteByte(template[i])
			continue
		}
		verb := template[i+1]
		if verb != 's' && verb != 'd' && verb != 'f' {
			b.WriteByte(template[i])
			continue
		}
		if arg >= len(args) {
			b.WriteByte(template[i])
			continue
		}
		b.WriteString(fmt.Sprintf("%"+string(verb), args[arg]))
		arg++
		i++
	}
	return b.String()
}
