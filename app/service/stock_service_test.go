package service

import (
	"fmt"
	"quick-go/app/entity"
	"quick-go/app/models"
	"quick-go/test/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetSpuStock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	stockList := []models.Stock{models.Stock{AppID: "app1g5f4c7r8050"}}

	mockRepo := mocks.NewMockIStockRepo(mockCtrl)
	mockRepo.EXPECT().GetStockDetail(gomock.Any(), "app1g5f4c7r8050", "bclass_5fc9fba1b61f8_zqoTfU").
		Return(stockList, nil)

	service := NewStockService(mockRepo)

	assert := assert.New(t)

	var tests = []struct {
		name     string
		input    *entity.GetSpuStockReq
		expected interface{}
	}{
		{
			name: "mock test",
			input: &entity.GetSpuStockReq{
				AppID: "app1g5f4c7r8050",
				SpuID: "bclass_5fc9fba1b61f8_zqoTfU",
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := service.GetSpuStock(tt.input)
			assert.Nil(err)
			fmt.Print(res)
		})

	}

}
