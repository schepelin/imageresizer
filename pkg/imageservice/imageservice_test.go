package imageservice

import (
	"github.com/golang/mock/gomock"
	"github.com/schepelin/imageresizer/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"context"
	"time"
	"image"
	"image/color"
)

func initMockersForImageService(t *testing.T) (*mocks.MockClocker, *mocks.MockStorage, *mocks.MockHasher, *mocks.MockConverter) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCloker := mocks.NewMockClocker(mockCtrl)
	mockStorage := mocks.NewMockStorage(mockCtrl)
	mockHasher := mocks.NewMockHasher(mockCtrl)
	mockConverter := mocks.NewMockConverter(mockCtrl)

	return mockCloker, mockStorage, mockHasher, mockConverter
}

func createSampleImage() image.Image {
	sampleImg := image.NewRGBA(image.Rect(0, 0, 10, 10))
	sampleImg.Set(1, 1, color.RGBA{255, 0, 0, 255})
	return sampleImg
}

func TestNewImageService(t *testing.T) {
	mockCloker, mockStorage, mockHasher, mockConverter := initMockersForImageService(t)

	is := NewImageService(mockStorage, mockCloker, mockHasher, mockConverter)

	assert.Equal(t, is.Converter, mockConverter)
	assert.Equal(t, is.Clock, mockCloker)
	assert.Equal(t, is.Storage, mockStorage)
	assert.Equal(t, is.Hash, mockHasher)
}

func TestImageService_Create(t *testing.T) {
	mockCloker, mockStorage, mockHasher, mockConverter := initMockersForImageService(t)
	ctx := context.TODO()
	is := NewImageService(mockStorage, mockCloker, mockHasher, mockConverter)
	rawByte := []byte{42, 10, 15}
	expectedHash := "42"
	expectedCteatedAt := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	expectedImage := createSampleImage()
	mockHasher.EXPECT().Gen(&rawByte).Return(expectedHash)
	mockCloker.EXPECT().Now().Return(expectedCteatedAt)
	mockConverter.EXPECT().Transform(&rawByte).Return(expectedImage, nil)
	mockStorage.EXPECT().Create(ctx, gomock.Any()).Return(nil)

	imgObj, _ := is.Create(ctx, &rawByte)

	assert.Equal(t, imgObj.Id, expectedHash)
	assert.Equal(t, imgObj.Image, expectedImage)
	assert.Equal(t, imgObj.CreatedAt, expectedCteatedAt)

}
