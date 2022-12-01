package entiity

import "time"

type Student struct {
	Id, Name, Email, Address, Gender string
	Birth_date                       time.Time
	TakenCredit                      int
}
