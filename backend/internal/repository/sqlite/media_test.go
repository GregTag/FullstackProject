package repository_sqlite_test

import (
	"backend/internal/entity"
	repository_sqlite "backend/internal/repository/sqlite"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type MediaTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo *repository_sqlite.MediaSQLite
}

func TestMediaTestSuite(t *testing.T) {
	suite.Run(t, new(MediaTestSuite))
}

func (suite *MediaTestSuite) SetupTest() {
	db, err := repository_sqlite.NewSQLiteDB("file::memory:?cache=shared")
	if err != nil {
		suite.Fail("Failed to connect database", err)
	}

	suite.db = db
	suite.repo = repository_sqlite.NewMediaSQLite(db)
}

func (suite *MediaTestSuite) TestCreate() {
	// Prepare test data
	media := &entity.Media{
		Title:       "Test Media 0",
		Description: "Test Description 0",
		Category:    "Serials",
		ReleaseYear: 2020,
	}

	// Test Create method
	err := suite.repo.Create(media)
	suite.NoError(err)

	// Verify the media is created
	createdMedia, err := suite.repo.Get(media.ID)
	suite.NoError(err)
	suite.Equal(media, createdMedia)

	// Test creating media with the same ID
	err = suite.repo.Create(media)
	suite.Equal(entity.ErrMediaAlreadyExists, err)

	// Test creating media with different ID
	media2 := &entity.Media{
		Title:       "Test Media 1",
		Description: "Test Description 1",
		Category:    "Movies",
		ReleaseYear: 2021,
	}
	err = suite.repo.Create(media2)
	suite.NoError(err)

	// Verify the second media is created
	createdMedia2, err := suite.repo.Get(media2.ID)
	suite.NoError(err)
	suite.Equal(media2, createdMedia2)
}
func (suite *MediaTestSuite) TestUpdate() {
	// Prepare test data
	media := &entity.Media{
		Title:       "Test Media",
		Description: "Test Description",
		Category:    "Movies",
		ReleaseYear: 2022,
	}
	err := suite.repo.Create(media)
	suite.NoError(err)

	// Update the media
	media.Title = "Updated Media"
	media.Description = "Updated Description"
	media.Category = "Serials"
	media.ReleaseYear = 2023
	err = suite.repo.Update(media)
	suite.NoError(err)

	// Verify the media is updated
	updatedMedia, err := suite.repo.Get(media.ID)
	suite.NoError(err)
	suite.Equal(media, updatedMedia)

	// Update non-existing media
	nonExistingMedia := &entity.Media{
		ID:          100,
		Title:       "Non-existing Media",
		Description: "Non-existing Description",
		Category:    "Non-existing Category",
		ReleaseYear: 2024,
	}
	err = suite.repo.Update(nonExistingMedia)
	suite.Equal(entity.ErrMediaNotFound, err)
}
func (suite *MediaTestSuite) TestDelete() {
	// Prepare test data
	media := &entity.Media{
		Title:       "Test Media",
		Description: "Test Description",
		Category:    "Movies",
		ReleaseYear: 2022,
	}
	err := suite.repo.Create(media)
	suite.NoError(err)

	// Test Delete method
	err = suite.repo.Delete(media.ID)
	suite.NoError(err)

	// Verify the media is deleted
	_, err = suite.repo.Get(media.ID)
	suite.Equal(entity.ErrMediaNotFound, err)

	// Test deleting non-existing media
	err = suite.repo.Delete(100)
	suite.Equal(entity.ErrMediaNotFound, err)
}
func (suite *MediaTestSuite) TestGet() {
	// Prepare test data
	media := &entity.Media{
		Title:       "Test Media",
		Description: "Test Description",
		Category:    "Movies",
		ReleaseYear: 2022,
	}
	err := suite.repo.Create(media)
	suite.NoError(err)

	// Test Get method with existing media
	retrievedMedia, err := suite.repo.Get(media.ID)
	suite.NoError(err)
	suite.Equal(media, retrievedMedia)

	// Test Get method with non-existing media
	nonExistingMediaID := uint(100)
	nonExistingMedia, err := suite.repo.Get(nonExistingMediaID)
	suite.Equal(entity.ErrMediaNotFound, err)
	suite.Nil(nonExistingMedia)
}
func (suite *MediaTestSuite) TestSearch() {
	// Prepare test data
	media1 := &entity.Media{
		Title:           "Untest Media 1",
		Description:     "Test Description 1",
		Category:        "Movies",
		ReleaseYear:     2020,
		NumberOfRatings: 1,
	}
	err := suite.repo.Create(media1)
	suite.NoError(err)

	media2 := &entity.Media{
		Title:           "Untest Media 2",
		Description:     "Test Description 2",
		Category:        "Serials",
		ReleaseYear:     2021,
		NumberOfRatings: 1,
	}
	err = suite.repo.Create(media2)
	suite.NoError(err)

	media3 := &entity.Media{
		Title:           "Untest Media 3",
		Description:     "Test Description 3",
		Category:        "Movies",
		ReleaseYear:     2022,
		NumberOfRatings: 1,
	}
	err = suite.repo.Create(media3)
	suite.NoError(err)

	// Test Search method with filter
	filter := &entity.Filter{
		Title:    "Untest Media",
		Category: "Movies",
		YearFrom: 2020,
		YearTo:   2022,
	}
	results, err := suite.repo.Search(filter)
	suite.NoError(err)
	suite.Len(*results, 2)
	suite.Contains(*results, entity.MediaResult{
		ID:     media1.ID,
		Title:  media1.Title,
		Image:  media1.Image,
		Rating: float32(media1.CumulativeRating) / float32(media1.NumberOfRatings),
	})
	suite.Contains(*results, entity.MediaResult{
		ID:     media3.ID,
		Title:  media3.Title,
		Image:  media3.Image,
		Rating: float32(media3.CumulativeRating) / float32(media3.NumberOfRatings),
	})

	// Test Search method with empty filter
	emptyFilter := &entity.Filter{
		Title: "Untest Media",
	}
	results, err = suite.repo.Search(emptyFilter)
	suite.NoError(err)
	suite.Len(*results, 3)
	suite.Contains(*results, entity.MediaResult{
		ID:     media1.ID,
		Title:  media1.Title,
		Image:  media1.Image,
		Rating: float32(media1.CumulativeRating) / float32(media1.NumberOfRatings),
	})
	suite.Contains(*results, entity.MediaResult{
		ID:     media2.ID,
		Title:  media2.Title,
		Image:  media2.Image,
		Rating: float32(media2.CumulativeRating) / float32(media2.NumberOfRatings),
	})
	suite.Contains(*results, entity.MediaResult{
		ID:     media3.ID,
		Title:  media3.Title,
		Image:  media3.Image,
		Rating: float32(media3.CumulativeRating) / float32(media3.NumberOfRatings),
	})
}
