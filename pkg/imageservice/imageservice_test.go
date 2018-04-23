package imageservice

import (
	"bytes"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/schepelin/imageresizer/pkg/mocks"
	"github.com/schepelin/imageresizer/pkg/storage"
	"github.com/stretchr/testify/assert"
	"image"
	"image/color"
	"image/png"
	"testing"
	"time"
)

func initMockersForImageService(t *testing.T) (*mocks.MockClocker, *mocks.MockStorage, *mocks.MockHasher, *mocks.MockConverter) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClocker := mocks.NewMockClocker(mockCtrl)
	mockStorage := mocks.NewMockStorage(mockCtrl)
	mockHasher := mocks.NewMockHasher(mockCtrl)
	mockConverter := mocks.NewMockConverter(mockCtrl)

	return mockClocker, mockStorage, mockHasher, mockConverter
}

func createSampleImage() image.Image {
	sampleImg := image.NewRGBA(image.Rect(0, 0, 10, 10))
	sampleImg.Set(1, 1, color.RGBA{255, 0, 0, 255})
	return sampleImg
}

func imageToByte(img image.Image) []byte {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		fmt.Println("failed to create buffer", err)
	}
	return buf.Bytes()
}

func TestNewImageService(t *testing.T) {
	mockClocker, mockStorage, mockHasher, mockConverter := initMockersForImageService(t)

	is := NewImageService(mockStorage, mockClocker, mockHasher, mockConverter)

	assert.Equal(t, is.Converter, mockConverter)
	assert.Equal(t, is.Clock, mockClocker)
	assert.Equal(t, is.Storage, mockStorage)
	assert.Equal(t, is.Hash, mockHasher)
}

func TestImageService_Create(t *testing.T) {
	mockClocker, mockStorage, mockHasher, mockConverter := initMockersForImageService(t)

	ctx := context.TODO()
	is := NewImageService(mockStorage, mockClocker, mockHasher, mockConverter)
	rawByte := []byte{42, 10, 15}
	expectedId := "42"
	expectedCteatedAt := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	expectedImage := createSampleImage()

	mockHasher.EXPECT().Gen(&rawByte).Return(expectedId)
	mockClocker.EXPECT().Now().Return(expectedCteatedAt)
	mockConverter.EXPECT().Transform(&rawByte).Return(expectedImage, nil)
	mockStorage.EXPECT().Create(ctx, gomock.Any()).Return(nil)

	imgObj, _ := is.Create(ctx, &rawByte)

	assert.Equal(t, imgObj.Id, expectedId)
	assert.Equal(t, imgObj.Image, expectedImage)
	assert.Equal(t, imgObj.CreatedAt, expectedCteatedAt)

}

func TestImageService_Read(t *testing.T) {
	mockClocker, mockStorage, mockHasher, mockConverter := initMockersForImageService(t)
	expectedImage := createSampleImage()
	expectedImageRaw := imageToByte(expectedImage)
	imageModelId := "42"
	storageReturnValue := storage.ImageModel{
		Id:        imageModelId,
		Raw:       expectedImageRaw,
		CreatedAt: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	ctx := context.TODO()
	is := NewImageService(mockStorage, mockClocker, mockHasher, mockConverter)

	mockConverter.EXPECT().Transform(gomock.Any()).Return(expectedImage, nil)
	mockStorage.EXPECT().Get(ctx, imageModelId).Return(&storageReturnValue, nil)

	imgObj, _ := is.Read(ctx, imageModelId)
	assert.Equal(t, imgObj.Id, imageModelId)
	assert.Equal(t, imgObj.Image, expectedImage)
	assert.Equal(t, imgObj.CreatedAt, storageReturnValue.CreatedAt)
}

func TestImageService_Delete(t *testing.T) {
	mockClocker, mockStorage, mockHasher, mockConverter := initMockersForImageService(t)
	ctx := context.TODO()
	is := NewImageService(mockStorage, mockClocker, mockHasher, mockConverter)
	imgId := "42"

	mockStorage.EXPECT().Delete(ctx, "42").Return(nil).Times(1)
	err := is.Delete(ctx, imgId)
	assert.NoError(t, err)

}