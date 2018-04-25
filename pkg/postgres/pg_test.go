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
	"v/gopkg.in/DATA-DOG/go-sqlmock.v1@v1.3.0-gopkgin-v1.3.0"
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
	_, err := NewPostgresStorage("postgres://test:test@localhost/test?sslmode=disable")
	assert.NoError(t, err)
	// TODO: How to check ps
}

func TestPostgresStorage_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error("error creating mock database")

	}
	defer db.Close()
	ps := PostgresStorage{db}
	ctx := context.TODO()
	imgModel := storage.ImageModel{
		Id:        "100500",
		Raw:       getRawImageSample(),
		CreatedAt: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	mock.ExpectExec("INSERT INTO images").WithArgs(imgModel.Id, string(imgModel.Raw), imgModel.CreatedAt)
	ps.Create(ctx, &imgModel)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresStorage_Get(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	ps := PostgresStorage{db}
	ctx := context.TODO()
	imgId := "42"
	expectedRaw := []byte{10, 42, 15}
	expectedCreatedAt := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	rows := sqlmock.NewRows([]string{"raw", "created_at"}).AddRow(
		string(expectedRaw),
		expectedCreatedAt,
	)
	mock.ExpectQuery(
		"SELECT raw, created_at FROM images WHERE id=?",
	).WithArgs(imgId).WillReturnRows(rows)

	imgModel, err := ps.Get(ctx, imgId)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, imgModel.Id, imgId)
	assert.Equal(t, imgModel.Raw, expectedRaw)
	assert.Equal(t, imgModel.CreatedAt, expectedCreatedAt)
}

func TestPostgresStorage_Delete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	ps := PostgresStorage{db}
	ctx := context.TODO()
	imgId := "42"

	mock.ExpectExec(
		"DELETE FROM images WHERE id=?",
	).WithArgs(imgId).WillReturnResult(sqlmock.NewResult(100500, 1))

	ps.Delete(ctx, imgId)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresStorage_CreateWithDbHelper(t *testing.T) {
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
	assert.Equal(t, []byte(actualRaw), imgModel.Raw)

	assert.Equal(t, actualCreatedAt, imgModel.CreatedAt)
}

func TestPostgresStorage_CreateWithRealDB(t *testing.T) {
	const dbConnect string = "postgres://localhost/imageres_dev?sslmode=disable"
	migrationsHistory := &migrate.AssetMigrationSource{
		Asset:    migrations.Asset,
		AssetDir: migrations.AssetDir,
		Dir:      "migrations",
	}

	testDbName := "imgres_dev_test_" + fmt.Sprintf("%v", time.Now().UnixNano())

	ps, err := NewPostgresStorage(dbConnect)
	if err != nil {
		t.Error("Could not connect to the databse")
	}
	originalDb := ps.DB
	defer originalDb.Close()
	originalDb.Exec("CREATE DATABASE " + testDbName + " ENCODING 'UTF8'")
	db, err := sql.Open("postgres", "postgres://localhost/"+testDbName+"?sslmode=disable")

	if err != nil {
		t.Error(err)
	}

	_, err = migrate.Exec(db, "postgres", migrationsHistory, migrate.Up)
	if err != nil {
		t.Error(err)
	}

	ps.DB = db
	ctx := context.TODO()
	imgModel := storage.ImageModel{
		Id:        "100500",
		Raw:       getRawImageSample(),
		CreatedAt: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	ps.Create(ctx, &imgModel)
	db.Close()
	_, err = originalDb.Exec("DROP DATABASE " + testDbName)
	if err != nil {
		t.Error(err)
	}

}
