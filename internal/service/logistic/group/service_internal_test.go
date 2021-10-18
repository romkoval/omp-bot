package group

import (
	"reflect"
	"testing"

	"github.com/ozonmp/omp-bot/internal/model/logistic"
)

func TestDummyGroupService_Remove(t *testing.T) {
	type args struct {
		id uint64
	}
	s := NewDummyGroupService()
	length := s.Size()
	tests := []struct {
		name    string
		args    args
		wantErr bool
		length  int
	}{
		{"del out of seq", args{10}, true, length},
		{"del in seq", args{1}, false, length - 1},
		{"del already deleted", args{1}, true, length - 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := s.Remove(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DummyGroupService.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.length != s.Size() {
				t.Errorf("DummyGroupService.Remove() length %d expected: %d", s.Size(), tt.length)
			}
		})
	}
}

func TestDummyGroupService_Update(t *testing.T) {
	type args struct {
		grp logistic.Group
		Id  uint64
	}
	s := NewDummyGroupService()
	length := s.Size()
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		length   int
		expected *logistic.Group
	}{
		{
			"Update out of seq",
			args{logistic.Group{Title: "Updated"}, uint64(10)},
			true,
			length,
			nil,
		},
		{
			"Update in seq",
			args{logistic.Group{Title: "Updated"}, uint64(5)},
			false,
			length,
			&logistic.Group{Id: 5, Title: "Updated"},
		},
		{
			"Update first",
			args{logistic.Group{Title: "Updated"}, uint64(1)},
			false,
			length,
			&logistic.Group{Id: 1, Title: "Updated"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Update(tt.args.Id, tt.args.grp); (err != nil) != tt.wantErr {
				t.Errorf("DummyGroupService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.length != s.Size() {
				t.Errorf("DummyGroupService.Update() length %d expected: %d", s.Size(), tt.length)
			}
			if tt.wantErr {
				return
			}
			got, err := s.storage.SelectOne(tt.args.Id)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(*tt.expected, *got) {
				t.Errorf("DummyGroupService.Update() = %v, want %v", *got, *tt.expected)
			}
		})
	}
}

func TestDummyGroupService_Create(t *testing.T) {
	s := NewDummyGroupService()
	allgrp, err := s.List(0, 100)
	if err != nil {
		t.Error(err)
	}
	last_id := allgrp[s.Size()-1].Id
	t.Run("add one", func(t *testing.T) {
		expected := append(allgrp, logistic.Group{Id: last_id + 1, Title: "New item"})
		if id, err := s.Create(logistic.Group{Title: "New item"}); err != nil {
			t.Error(err)
		} else {
			if id != last_id+1 {
				t.Errorf("invalid next id:%d", id)
			}
			got, err := s.storage.SelectAll(uint64(0), uint64(100))
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("DummyGroupService.Create() = %v, want %v", got, expected)
			}
		}
	})
}
