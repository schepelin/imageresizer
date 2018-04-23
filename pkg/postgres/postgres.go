package postgres
//
//import (
//	"bytes"
//	"crypto/md5"
//	"database/sql"
//	"encoding/hex"
//	"image/png"
//	"log"
//
//	_ "github.com/lib/pq"
//	"github.com/schepelin/imageresizer/pkg/resizer"
//)
//
//type ImageService struct {
//	DB     *sql.DB
//	Logger *log.Logger
//}
//
//func (is *ImageService) Create(b []byte) (*resizer.Image, error) {
//	buf := bytes.NewBuffer(b)
//	img, err := png.Decode(buf)
//	if err != nil {
//		is.Logger.Panic("Could not decode png ", err)
//		return nil, err
//	}
//	hash := md5.New()
//	if _, err := hash.Write(b); err != nil {
//		return nil, err
//	}
//	imgObj := resizer.Image{
//		Id:  hex.EncodeToString(hash.Sum(nil)),
//		Image: img,
//	}
//	is.DB.Query(`INSERT INTO images(hash, data) VALUES($1, $2)`, hash, b)
//
//	return &imgObj, nil
//}
//
//
//func (is *ImageService) Get(imgId string) (*resizer.Image, error) {
//	var rawData []byte
//
//
//	err := is.DB.QueryRow("SELECT data FROM images WHERE hash=?", imgId).Scan(&rawData)
//	if err != nil {
//		return nil, err
//	}
//	img, err := png.Decode(bytes.NewReader(rawData))
//	if err != nil {
//		return nil, err
//	}
//
//	return &resizer.Image{
//		Id: imgId,
//		Image: img,
//	}, nil
//}
//
//
//func New(dbConnect string, logger *log.Logger) *ImageService {
//	db, err := sql.Open("postgres", dbConnect)
//	if err != nil {
//		logger.Fatal("could not connect to the database ", err)
//	}
//
//	return &ImageService{
//		DB:     db,
//		Logger: logger,
//	}
//}
