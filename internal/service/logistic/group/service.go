package group

import (
	"github.com/ozonmp/omp-bot/internal/model/logistic"
	storage "github.com/ozonmp/omp-bot/internal/storage/logistic"
)

type Service interface {
	Describe(groupID uint64) (*logistic.Group, error)
	List(cursor uint64, limit uint64) ([]logistic.Group, error)
	Create(logistic.Group) (uint64, error)
	Update(groupID uint64, group logistic.Group) error
	Remove(groupID uint64) (bool, error)
}

type DummyGroupService struct {
	storage *storage.GroupStorage
}

func NewDummyGroupService() *DummyGroupService {
	return &DummyGroupService{
		storage: storage.NewStorage(),
	}
}

func (s *DummyGroupService) List(cursor uint64, limit uint64) ([]logistic.Group, error) {
	return s.storage.SelectAll(cursor, limit)
}

func (s *DummyGroupService) Describe(groupID uint64) (*logistic.Group, error) {
	if grp, err := s.storage.SelectOne(groupID); err != nil {
		return nil, err
	} else {
		return grp, nil
	}
}

func (s *DummyGroupService) Create(grp logistic.Group) (uint64, error) {
	return s.storage.Create(&grp), nil
}

func (s *DummyGroupService) Remove(groupID uint64) (bool, error) {
	if err := s.storage.Remove(groupID); err != nil {
		return false, err
	}
	return true, nil
}

func (s *DummyGroupService) Update(groupID uint64, grp logistic.Group) error {
	if err := s.storage.Update(groupID, &grp); err != nil {
		return err
	}
	return nil
}

func (s *DummyGroupService) Size() int {
	return s.storage.Size()
}
