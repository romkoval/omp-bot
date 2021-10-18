package logistic

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/ozonmp/omp-bot/internal/model/logistic"
)

type GroupStorage struct {
	groupIdSeq    uint64
	groupEntities []logistic.Group
}

func NewStorage() *GroupStorage {
	return &GroupStorage{
		groupIdSeq: 5,
		groupEntities: []logistic.Group{
			{Id: 1, Title: "one"},
			{Id: 2, Title: "two"},
			{Id: 3, Title: "three"},
			{Id: 4, Title: "four"},
			{Id: 5, Title: "five"},
		},
	}
}

func (s *GroupStorage) SelectAll(cursor uint64, limit uint64) ([]logistic.Group, error) {
	if cursor >= uint64(len(s.groupEntities)) {
		return []logistic.Group{},
			fmt.Errorf("cursor out of bounds, max cursor value is %d",
				len(s.groupEntities)-1)
	}
	maxsz := math.Min(float64(len(s.groupEntities)), float64(cursor+limit))
	return s.groupEntities[cursor:int(maxsz)], nil
}

func (s *GroupStorage) Create(grp *logistic.Group) uint64 {
	s.groupIdSeq += 1
	grp.Id = s.groupIdSeq
	s.groupEntities = append(s.groupEntities, *grp)
	return s.groupIdSeq
}

func (s *GroupStorage) Remove(groupID uint64) error {
	size := len(s.groupEntities)
	ind, err := s.getIndex(groupID)
	if err != nil {
		return err
	}
	copy(s.groupEntities[ind:], s.groupEntities[ind+1:])
	s.groupEntities = s.groupEntities[0 : size-1]
	log.Printf("deleted item at %d, %#v", ind, s.groupEntities)
	return nil
}

func (s *GroupStorage) Update(groupID uint64, grp *logistic.Group) error {
	if ind, err := s.getIndex(groupID); err != nil {
		return err
	} else {
		grp.Id = s.groupEntities[ind].Id
		s.groupEntities[ind] = *grp
	}
	return nil
}

func (s *GroupStorage) SelectOne(groupID uint64) (*logistic.Group, error) {
	if ind, err := s.getIndex(groupID); err != nil {
		return nil, err
	} else {
		return &s.groupEntities[ind], nil
	}
}

func (s *GroupStorage) Size() int {
	return len(s.groupEntities)
}

func (s *GroupStorage) getIndex(groupID uint64) (int, error) {
	size := len(s.groupEntities)
	ind := sort.Search(size, func(ind int) bool {
		return s.groupEntities[ind].Id >= groupID
	})
	log.Printf("found index %d by id:%d, len:%d\n", ind, groupID, size)
	if ind == size || groupID != s.groupEntities[ind].Id {
		return 0, fmt.Errorf("group not found by id: %d", groupID)
	}
	return ind, nil
}
