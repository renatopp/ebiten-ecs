package debug

import (
	"encoding/json"
	"fmt"
	"strings"

	sk "github.com/renatopp/skald"
)

func DescribeEntityInstance(e *sk.EntityInstance, pad ...int) string {
	space := 0
	if len(pad) > 0 {
		space = pad[0]
	}

	padding := strings.Repeat(" ", space)

	def := e.Definition()

	result := ""
	result += fmt.Sprintf("%s(%d) Entity of type (%d)\n", padding, e.Id(), def.Id())
	for _, c := range def.Components() {
		result += DescribeComponentInstance(c, e.GetComponent(c.Id()), space+2)
	}
	return result
}

func DescribeComponentInstance(def sk.IComponent, inst any, pad ...int) string {
	space := 0
	if len(pad) > 0 {
		space = pad[0]
	}

	padding := strings.Repeat(" ", space)
	b, err := json.Marshal(inst)
	r := string(b)
	if err != nil {
		r = "[could not be serialized]"
	}
	return fmt.Sprintf("%s(%d) %s: %s\n", padding, def.Id(), def.Name(), r)
}
