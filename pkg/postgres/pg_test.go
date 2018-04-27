package postgres

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/schepelin/imageresizer/pkg/migrations"
	"github.com/schepelin/imageresizer/pkg/storage"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"image"
	"image/color"
	"image/png"
	"testing"
	"time"
	"v/github.com/rubenv/sql-migrate@v0.0.0-20180217203553-081fe17d19ff"
)

func getRawImageSample() []byte {
	sampleImg := image.NewRGBA(image.Rect(0, 0, 10, 10))
	sampleImg.Set(1, 1, color.RGBA{255, 0, 0, 255})

	buf := new(bytes.Buffer)
	err := png.Encode(buf, sampleImg)
	if err != nil {
		fmt.Println("failed to create buffer", err)
	}

	return buf.Bytes()
}

type dependencies struct {
	db *sql.DB
}

func getImageRowDataDummy() (string, []byte, time.Time) {
	imgId := "42010"
	raw := []byte{42, 10, 15}
	createdAt := time.Date(1970, time.January, 1, 0, 0, 1, 0, time.FixedZone("", 0))

	return imgId, raw, createdAt
}

func preparator(t *testing.T, deps *dependencies) func() {
	var dbName string
	var err error
	dbName, deps.db, err = prepareTestDB()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	return func() {
		if err := deps.db.Close(); err != nil {
			t.Fatal(err)
		}
		if err := dropTestDB(dbName); err != nil {
			t.Fatal(err)
		}
	}
}

func prepareTestDB() (string, *sql.DB, error) {
	dbName, err := createTempDB()
	if err != nil {
		return dbName, nil, err
	}

	db, err := sql.Open("postgres", "postgres://localhost/"+dbName+"?sslmode=disable")
	if err != nil {
		return dbName, nil, err
	}
	migrate.SetTable("migrations")
	migrationsHistory := &migrate.AssetMigrationSource{
		Asset:    migrations.Asset,
		AssetDir: migrations.AssetDir,
		Dir:      "migrations",
	}
	_, err = migrate.Exec(db, "postgres", migrationsHistory, migrate.Up)
	if err != nil {
		return dbName, nil, err
	}

	return dbName, db, nil
}

func createTempDB() (string, error) {
	dbName := "imageres_test_" + fmt.Sprintf("%v", time.Now().UnixNano())

	db, err := sql.Open("postgres", "postgres://localhost/imageres_dev?sslmode=disable")
	if err != nil {
		return dbName, err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE DATABASE "` + dbName + `"`)
	if err != nil {
		return dbName, err
	}

	return dbName, nil

}

func dropTestDB(dbName string) error {

	db, err := sql.Open("postgres", "postgres://localhost/imageres_dev?sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`DROP DATABASE "` + dbName + `"`)
	return err
}

func TestNewPostgresStorage(t *testing.T) {
	var deps dependencies
	defer preparator(t, &deps)()
	ps := NewPostgresStorage(deps.db)
	assert.NoError(t, ps.DB.Ping())
}

func TestPostgresStorage_Get(t *testing.T) {
	var deps dependencies
	var err error
	defer preparator(t, &deps)()
	ps := PostgresStorage{deps.db}

	ctx := context.TODO()
	imgId, expectedRaw, expectedCreatedAt := getImageRowDataDummy()
	ps.DB.Exec("INSERT INTO images(id, raw, created_at) VALUES ($1, $2, $3)", imgId, string(expectedRaw), expectedCreatedAt)

	imgModel, err := ps.Get(ctx, imgId)

	assert.NoError(t, err)
	assert.Equal(t, imgId, imgModel.Id)
	assert.Equal(t, expectedRaw, imgModel.Raw)
	assert.Equal(t, expectedCreatedAt, imgModel.CreatedAt)

}

func TestPostgresStorage_Delete(t *testing.T) {
	var deps dependencies
	var err error
	defer preparator(t, &deps)()
	ps := PostgresStorage{deps.db}
	ctx := context.TODO()

	imgId, raw, createdAt := getImageRowDataDummy()
	ps.DB.Exec("INSERT INTO images(id, raw, created_at) VALUES ($1, $2, $3)", imgId, raw, createdAt)

	err = ps.Delete(ctx, imgId)

	var selectedId string
	err = ps.DB.QueryRow("SELECT id from images where id=$1", imgId).Scan(&selectedId)
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestPostgresStorage_Create(t *testing.T) {
	var deps dependencies
	defer preparator(t, &deps)()
	ps := PostgresStorage{deps.db}
	ctx := context.TODO()
	imgModel := storage.ImageModel{
		Id:        "100500",
		Raw:       getRawImageSample(),
		CreatedAt: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
	}

	ps.Create(ctx, &imgModel)

	var actualRaw string
	var actualCreatedAt time.Time

	err := ps.DB.QueryRow("SELECT raw, created_at FROM images where id=$1", imgModel.Id).Scan(
		&actualRaw,
		&actualCreatedAt,
	)
	assert.NoError(t, err)
	assert.Equal(t, imgModel.Raw, []byte(actualRaw))

	assert.Equal(t, imgModel.CreatedAt, actualCreatedAt)
}


func TestPostgresStorage_CreateResizeJob(t *testing.T) {
	var deps dependencies
	defer preparator(t, &deps)()
	ps := PostgresStorage{deps.db}
	ctx := context.TODO()
	imgId, raw, imgCreatedAt := getImageRowDataDummy()
	ps.DB.Exec("INSERT INTO images(id, raw, created_at) VALUES ($1, $2, $3)", imgId, raw, imgCreatedAt)

	req := storage.ResizeJobRequest{
		ImgId: imgId,
		Width: 10,
		Height: 20,
	}
	resp, err := ps.CreateResizeJob(ctx, &req)
	assert.NoError(t, err)
	assert.Equal(t, storage.StatusCreated, resp.Status)

	var jobId uint64
	var createdAt time.Time
	err = ps.DB.QueryRow(
		"SELECT id, created_at from resize_jobs where image_id=$1 AND width=$2 AND height=$3",
		imgId, req.Width, req.Height,
	).Scan(&jobId, &createdAt)

	assert.NoError(t, err)
	assert.Equal(t, jobId, resp.Id)
	assert.Equal(t, createdAt, resp.CreatedAt)

}

func TestPostgresStorage_CreateResizeJob_ThereIsNoImage(t *testing.T) {
	var deps dependencies
	defer preparator(t, &deps)()
	ps := PostgresStorage{deps.db}
	ctx := context.TODO()
	req := storage.ResizeJobRequest{
		ImgId: "there_is_no_such_image",
		Width: 10,
		Height: 20,
	}
	_, err := ps.CreateResizeJob(ctx, &req)
	assert.Equal(t, storage.ErrNoImageFound, err)
}
