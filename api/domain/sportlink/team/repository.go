package team

type Repository interface {
	Save(entity Entity) error
}
