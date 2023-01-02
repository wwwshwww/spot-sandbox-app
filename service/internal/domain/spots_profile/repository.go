package spots_profile

//go:generate mockgen -source $GOFILE -package mock -destination mock/$GOFILE

type Repository interface {
	Get(Identifier) (SpotsProfile, error)
	Save(SpotsProfile) error
	Delete(Identifier) error
	NextIdentifier() (Identifier, error)
}
