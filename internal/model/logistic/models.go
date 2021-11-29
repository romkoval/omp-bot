package logistic

import "fmt"

type Group struct {
	Id   uint64
	Name string
}

func (g *Group) String() string {
	return fmt.Sprintf("Logistic.group id:%d, name:%s", g.Id, g.Name)
}
