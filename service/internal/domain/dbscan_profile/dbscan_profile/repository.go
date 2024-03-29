package dbscan_profile

//go:generate mockgen -source $GOFILE -package mock -destination mock/$GOFILE

type Repository interface {
	Get(Identifier) (DbscanProfile, error)
	BulkGet([]Identifier) (map[Identifier]DbscanProfile, error)
	Save(DbscanProfile) error
	Delete(Identifier) error
	NextIdentifier() (Identifier, error)
}
