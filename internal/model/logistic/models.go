package logistic

import "fmt"

type Group struct {
	Id    uint64
	Title string
}

func (g *Group) String() string {
	return fmt.Sprintf("Logistic.group id:%d, title:%s", g.Id, g.Title)
}
