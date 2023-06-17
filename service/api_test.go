package service

import (
	"context"
	"github.com/autobaza/auto_catalog/repository/mocks"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewService(t *testing.T) {
	var logger log.Logger
	var repo = mocks.NewRepository(t)
	srv := NewService(logger, repo)
	assert.NotNil(t, srv)
}

func TestServiceMethods(t *testing.T) {
	var logger log.Logger
	var repo = mocks.NewRepository(t)
	repo.On("GetCarTypes").Return(nil).Once()
	repo.On("GetCarMarks", "1").Return(nil).Once()
	repo.On("GetCarModels", "1").Return(nil).Once()
	srv := NewService(logger, repo)
	assert.Nil(t, srv.ListCarTypes(context.Background()))
	assert.Nil(t, srv.ListCarMarks(context.Background(), "1"))
	assert.Nil(t, srv.ListCarModels(context.Background(), "1"))
}
