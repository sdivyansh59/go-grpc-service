//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/sdivyansh59/go-grpc-service/internal/rocket Store
package rocket

import "context"

// Rocket - should contain the definition of our
// rocket
type Rocket struct {
	ID string 
	Name string
	Type string
	Flights int
}

// Store - defines the interface we expect 
// out database implementation to follow
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - out rocket service, responmsible for
// updating the rocket inventory
type Service struct {
	Store Store
}

// New - returns a new instance of our rocket Service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketBYID- retieves a rocket baed on the ID from the store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{},err
	}
	return rkt, nil
}

// InsertRocket - insert a new rocket into the store.
func (s Service) InsertRocket (rkt Rocket) (Rocket, error) {
	rkt, err :=  s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket - delete a rocket from our inventory
func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}

