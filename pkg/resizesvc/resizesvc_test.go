package resizesvc

import (
	"github.com/golang/mock/gomock"
	"github.com/schepelin/imageresizer/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"image"
	"image/color"
	"testing"
)

func createSampleImage() image.Image {
	sampleImg := image.NewRGBA(image.Rect(0, 0, 10, 10))
	sampleImg.Set(1, 1, color.RGBA{255, 0, 0, 255})
	return sampleImg
}

// TODO: apply defer hack to eliminate boilerpalte code
func TestNewResizeService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockResizeStorage := mocks.NewMockResizeStorage(mockCtrl)
	mockConverter := mocks.NewMockConverter(mockCtrl)
	mockPubSub := mocks.NewMockPublisherConsumer(mockCtrl)

	rs := NewResizeService(mockResizeStorage, mockConverter, mockPubSub)
	assert.Equal(t, mockConverter, rs.Converter)
	assert.Equal(t, mockResizeStorage, rs.Storage)
}

//func TestResizeService_ResizeAsync(t *testing.T) {
//	mockCtrl := gomock.NewController(t)
//	defer mockCtrl.Finish()
//
//	mockStorage := mocks.NewMockResizeStorage(mockCtrl)
//	mockConverter := mocks.NewMockConverter(mockCtrl)
//	mockPubSub := mocks.NewMockPublisherConsumer(mockCtrl)
//	ctx := context.TODO()
//
//	rs := NewResizeService(mockStorage, mockConverter, mockPubSub)
//
//	req := resizer.ResizeServiceRequest{
//		JobId:  100500,
//		RawImg: []byte{1, 2, 3},
//		Width:  20,
//		Height: 10,
//	}
//	//imgSample := createSampleImage()
//	//resizeRaw := []byte{4, 5, 6}
//	//storageReq := storage.ResizeResultRequest{req.JobId, resizeRaw}
//
//	mockPubSub.EXPECT().PublishResizeJob(ctx, req.JobId).Return(nil)
//	//gomock.InOrder(
//	//	mockConverter.EXPECT().Transform(&req.RawImg).Return(imgSample, nil),
//	//	mockConverter.EXPECT().Resize(&imgSample, req.Width, req.Height).Return(resizeRaw, nil),
//	//	mockStorage.EXPECT().WriteResizeJobResult(ctx, &storageReq).Return(nil),
//	//)
//
//	err := rs.RunResizeWorker(ctx, &req)
//	assert.NoError(t, err)
//}
