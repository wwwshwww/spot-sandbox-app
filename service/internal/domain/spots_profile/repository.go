package spots_profile

//go:generate mockgen -source $GOFILE -package mock -destination mock/$GOFILE

type Repository interface {
	Get(Identifier) (SpotProfile, error)
	Save(SpotProfile) error
	Delete(Identifier) error
}
