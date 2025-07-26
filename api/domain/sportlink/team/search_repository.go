package team

type SearchRepository interface {
	FindByID(id ID) (Entity, error)
}
