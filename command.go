package main

import "fmt"

type Command interface {
	Execute()
}

type Attack struct {
}

func (a *Attack) Execute() {
	fmt.Println("Attack!")
}

type MoveUnits struct {
}

func (m *MoveUnits) Execute() {
	fmt.Println("Move your units!")
}

type ChangePlan struct {
}

func (c *ChangePlan) Execute() {
	fmt.Println("Change of a plan!")
}

type Army struct {
	number int32
}

func (a *Army) FollowOrder(order Command) {
	order.Execute()
}

type Rapporteur struct {
	commands map[int32][]Command
}

func (r *Rapporteur) IssueOrders(army *Army, c Command) {
	if r.commands == nil {
		r.commands = make(map[int32][]Command)
	}
	r.commands[army.number] = append(r.commands[army.number], c)
}

func (r *Rapporteur) SendIssues(a *Army) {
	fmt.Printf("Sending orders to army #%d \n", a.number)
	orders := r.commands[a.number]
	for _, command := range orders {
		a.FollowOrder(command)
	}
}

func main() {
	army1 := &Army{number: 1}
	army2 := &Army{number: 2}

	rapporteur := &Rapporteur{}
	rapporteur.IssueOrders(army1, &Attack{})
	rapporteur.IssueOrders(army1, &MoveUnits{})
	rapporteur.SendIssues(army1)
	rapporteur.IssueOrders(army2, &ChangePlan{})
	rapporteur.SendIssues(army2)
}
