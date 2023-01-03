package spot

//go:generate mockgen -source $GOFILE -package mock -destination mock/$GOFILE

type Repository interface {
	Get(Identifier) (Spot, error)
	BulkGet([]Identifier) (map[Identifier]Spot, error)
	Save(Spot) error
	BulkSave([]Spot) error
	Delete(Identifier) error
	BulkDelete([]Identifier) error
	NextIdentifier() (Identifier, error)
	NextIdentifiers(uint) ([]Identifier, error)
}
