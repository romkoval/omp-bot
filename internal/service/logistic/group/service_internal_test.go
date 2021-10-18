package group

import (
	"reflect"
	"testing"

	"github.com/ozonmp/omp-bot/internal/model/logistic"
)

func TestService_Del(t *testing.T) {
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
				t.Errorf("DummyGroupService.Del() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.length != s.Size() {
				t.Errorf("DummyGroupService.Del() length %d expected: %d", s.Size(), tt.length)
			}
		})
	}
}

func TestService_Add(t *testing.T) {
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
				t.Errorf("DummyGroupService.Add() = %v, want %v", got, expected)
			}
		}
	})
}
