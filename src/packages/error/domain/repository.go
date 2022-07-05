package errlog

type Repository interface {
	Save(data Entity) error
	FindPaginated(page uint32) ([]Entity, error)
}
