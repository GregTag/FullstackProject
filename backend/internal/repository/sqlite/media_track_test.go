package repository_sqlite_test

import (
	"backend/internal/entity"
	repository_sqlite "backend/internal/repository/sqlite"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type MediaTrackTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo *repository_sqlite.MediaTrackSQLite
}

func TestMediaTrackTestSuite(t *testing.T) {
	suite.Run(t, new(MediaTrackTestSuite))
}

func (suite *MediaTrackTestSuite) SetupTest() {
	db, err := repository_sqlite.NewSQLiteDB("file::memory:?cache=shared")
	if err != nil {
		suite.Fail("Failed to connect database", err)
	}

	suite.db = db
	suite.repo = repository_sqlite.NewMediaTrackSQLite(db)
}
func (suite *MediaTrackTestSuite) TestCreate() {
	// Create a new media track
	track := &entity.MediaTrack{
		Media: entity.Media{
			ID:               1,
			Title:            "Test Media",
			CumulativeRating: 100,
			NumberOfRatings:  10,
			NumberOfTracks:   5,
		},
		MediaTrackBase: entity.MediaTrackBase{
			Rating: 4,
		},
	}

	// Call the Create method
	err := suite.repo.Create(track)

	// Assert that there are no errors
	suite.NoError(err)

	// Assert that the media track was created successfully

	createdTrack, err := suite.repo.Get(track.MediaTrackBase.UserID, track.MediaTrackBase.MediaID)
	suite.NoError(err)
	suite.Equal(track, createdTrack)

	// Attempt to create the same media track again
	err = suite.repo.Create(track)

	// Assert that the appropriate error is returned
	suite.Equal(entity.ErrMediaTrackAlreadyExists, err)

	// Assert that the media track count is updated correctly
	var updatedMedia entity.Media
	result := suite.db.First(&updatedMedia, track.Media.ID)
	suite.NoError(result.Error)
	suite.Equal(track.Media.NumberOfTracks, updatedMedia.NumberOfTracks)
	suite.Equal(track.Media.CumulativeRating, updatedMedia.CumulativeRating)
	suite.Equal(track.Media.NumberOfRatings, updatedMedia.NumberOfRatings)
}
func (suite *MediaTrackTestSuite) TestUpdate() {
	// Create a new media track
	track := &entity.MediaTrack{
		Media: entity.Media{
			ID:               1,
			Title:            "Test Media",
			CumulativeRating: 100,
			NumberOfRatings:  10,
			NumberOfTracks:   5,
		},
		MediaTrackBase: entity.MediaTrackBase{
			UserID:  1,
			MediaID: 1,
			Rating:  4,
		},
	}

	// Call the Create method to create the media track
	err := suite.repo.Create(track)
	suite.NoError(err)

	// Update the media track
	track.Rating = 5
	err = suite.repo.Update(track)
	suite.NoError(err)

	// Retrieve the updated media track from the database
	updatedTrack, err := suite.repo.Get(track.UserID, track.MediaID)
	suite.NoError(err)

	// Assert that the media track was updated successfully
	suite.Equal(track.MediaTrackBase.Rating, updatedTrack.MediaTrackBase.Rating)
	suite.Equal(track.Media.CumulativeRating, updatedTrack.Media.CumulativeRating)
	suite.Equal(track.Media.NumberOfRatings, updatedTrack.Media.NumberOfRatings)
	suite.Equal(track.Media.NumberOfTracks, updatedTrack.Media.NumberOfTracks)
}

func (suite *MediaTrackTestSuite) TestUpdate_NotFound() {
	// Create a new media track
	track := &entity.MediaTrack{
		Media: entity.Media{
			ID:               100,
			Title:            "Test Media",
			CumulativeRating: 100,
			NumberOfRatings:  10,
			NumberOfTracks:   5,
		},
		MediaTrackBase: entity.MediaTrackBase{
			UserID:  1,
			MediaID: 100,
			Rating:  4,
		},
	}

	// Attempt to update a non-existent media track
	err := suite.repo.Update(track)

	// Assert that the appropriate error is returned
	suite.Equal(entity.ErrMediaTrackNotFound, err)
}
func (suite *MediaTrackTestSuite) TestDelete() {
	// Create a new media track
	track := &entity.MediaTrack{
		Media: entity.Media{
			ID:               1,
			Title:            "Test Media",
			CumulativeRating: 100,
			NumberOfRatings:  10,
			NumberOfTracks:   5,
		},
		MediaTrackBase: entity.MediaTrackBase{
			UserID:  1,
			MediaID: 1,
			Rating:  4,
		},
	}

	// Call the Create method to create the media track
	err := suite.repo.Create(track)
	suite.NoError(err)

	// Call the Delete method
	err = suite.repo.Delete(track.UserID, track.MediaID)
	suite.NoError(err)

	// Attempt to retrieve the deleted media track from the database
	_, err = suite.repo.Get(track.MediaTrackBase.UserID, track.MediaTrackBase.MediaID)

	// Assert that the media track is not found in the database
	suite.Error(err)
	suite.True(errors.Is(err, entity.ErrMediaTrackNotFound))

	// Assert that the media track count is updated correctly
	var updatedMedia entity.Media
	result := suite.db.First(&updatedMedia, track.Media.ID)
	suite.NoError(result.Error)
	suite.Equal(track.Media.NumberOfTracks-1, updatedMedia.NumberOfTracks)
	suite.Equal(track.Media.CumulativeRating-uint64(track.Rating), updatedMedia.CumulativeRating)
	suite.Equal(track.Media.NumberOfRatings-1, updatedMedia.NumberOfRatings)
}

func (suite *MediaTrackTestSuite) TestDelete_NotFound() {
	// Attempt to delete a non-existent media track
	err := suite.repo.Delete(100, 100)

	// Assert that the appropriate error is returned
	suite.Equal(entity.ErrMediaTrackNotFound, err)
}
