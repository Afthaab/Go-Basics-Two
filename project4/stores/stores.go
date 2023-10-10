package stores

type Storer interface {
	Create(usr User) error
	Update(usr User) error
	Delete(usr User) error
}

var StorerInterface Storer

type Service struct {
	Storer
}

func NewService(S Storer) Service {
	return Service{
		Storer: S,
	}
}
