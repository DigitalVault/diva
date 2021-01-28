package db

type DbIf interface {
	init() (err error)
}
