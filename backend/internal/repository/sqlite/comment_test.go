package repository_sqlite_test

import (
	"backend/internal/entity"
	repository_sqlite "backend/internal/repository/sqlite"
	"errors"
	"log"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type CommentTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo *repository_sqlite.CommentSQLite
}

func TestCommentTestSuite(t *testing.T) {
	suite.Run(t, new(CommentTestSuite))
}

func (suite *CommentTestSuite) SetupTest() {
	db, err := repository_sqlite.NewSQLiteDB("file::memory:?cache=shared")
	if err != nil {
		suite.Fail("Failed to connect database", err)
	}

	suite.db = db
	suite.repo = repository_sqlite.NewCommentSQLite(db)
}

func (suite *CommentTestSuite) TestCreate() {
	// Create a new comment
	comment := &entity.Comment{CommentBase: entity.CommentBase{
		MediaID:  1,
		SenderID: 1,
		Content:  "Test Comment",
	},
	}

	// Call the Create method
	err := suite.repo.Create(comment)

	// Assert that there are no errors
	suite.NoError(err)

	// Assert that the comment was created successfully
	var createdComment entity.Comment
	result := suite.db.First(&createdComment, comment.ID)
	suite.NoError(result.Error)
	comment.CommentBase.CreatedAt = createdComment.CommentBase.CreatedAt
	comment.CommentBase.UpdatedAt = createdComment.CommentBase.UpdatedAt
	suite.Equal(comment, &createdComment)

	// Attempt to create the same comment again
	err = suite.repo.Create(comment)

	// Assert that the appropriate error is returned
	suite.Equal(entity.ErrCommentAlreadyExists, err)
}

func (suite *CommentTestSuite) TestUpdate() {
	// Create a new comment
	comment := &entity.Comment{CommentBase: entity.CommentBase{
		MediaID:  1,
		SenderID: 1,
		Content:  "Test Comment",
	},
	}

	// Call the Create method to create the comment
	err := suite.repo.Create(comment)
	suite.NoError(err)

	// Update the comment
	comment.Content = "Updated Comment"
	err = suite.repo.Update(comment)
	suite.NoError(err)

	// Retrieve the updated comment from the database
	var updatedComment entity.Comment
	result := suite.db.First(&updatedComment, comment.ID)
	suite.NoError(result.Error)

	// Assert that the comment was updated successfully
	suite.Equal(comment.Content, updatedComment.Content)
}

func (suite *CommentTestSuite) TestUpdate_NotFound() {
	// Create a new comment
	comment := &entity.Comment{CommentBase: entity.CommentBase{
		ID:       100,
		MediaID:  1,
		SenderID: 1,
		Content:  "Test Comment",
	},
	}

	// Attempt to update a non-existent comment
	err := suite.repo.Update(comment)

	// Assert that the appropriate error is returned
	suite.Equal(entity.ErrCommentNotFound, err)
}
func (suite *CommentTestSuite) TestDelete() {
	// Create a new comment
	comment := &entity.Comment{CommentBase: entity.CommentBase{
		MediaID:  1,
		SenderID: 1,
		Content:  "Test Comment",
	},
	}

	// Call the Create method to create the comment
	err := suite.repo.Create(comment)
	suite.NoError(err)

	// Delete the comment
	err = suite.repo.Delete(comment.ID, comment.SenderID)
	suite.NoError(err)

	// Attempt to retrieve the deleted comment from the database
	var deletedComment entity.Comment
	result := suite.db.First(&deletedComment, comment.ID)

	// Assert that the comment is not found in the database
	suite.Error(result.Error)
	suite.True(errors.Is(result.Error, gorm.ErrRecordNotFound))
}

func (suite *CommentTestSuite) TestDelete_NotFound() {
	// Attempt to delete a non-existent comment
	err := suite.repo.Delete(100, 100)

	// Assert that the appropriate error is returned
	suite.Equal(entity.ErrCommentNotFound, err)
}
func (suite *CommentTestSuite) TestLoad() {
	// Create a user for testing
	user := entity.User{ID: 1, UserRegister: entity.UserRegister{UserLogin: entity.UserLogin{Login: "test 1"}, Fullname: "Test 1", Email: "test1"}}
	result := suite.db.Create(&user)
	suite.NoError(result.Error)
	view := entity.UserView{}
	err := copier.Copy(&view, &user)
	suite.NoError(err)

	// Create a new comment
	comment := &entity.Comment{CommentBase: entity.CommentBase{
		MediaID:  1,
		SenderID: 1,
		Content:  "Test Comment",
	},
	}
	err = suite.repo.Create(comment)
	suite.NoError(err)

	// Call the Load method to retrieve the comment
	loadedComment, err := suite.repo.Load(comment.ID, comment.SenderID)
	loadedComment.Sender.CreatedAt = loadedComment.Sender.CreatedAt.Local()
	log.Println(loadedComment)
	suite.NoError(err)

	// Assert that the loaded comment is not nil
	suite.NotNil(loadedComment)

	// Assert that the loaded comment matches the created comment
	suite.Equal(comment.ID, loadedComment.ID)
	suite.Equal(comment.MediaID, loadedComment.MediaID)
	suite.Equal(comment.SenderID, loadedComment.SenderID)
	suite.Equal(comment.Content, loadedComment.Content)
	suite.Equal(view, loadedComment.Sender)
}

func (suite *CommentTestSuite) TestLoad_NotFound() {
	// Attempt to load a non-existent comment
	loadedComment, err := suite.repo.Load(100, 100)

	// Assert that the appropriate error is returned
	suite.Equal(entity.ErrCommentNotFound, err)
	suite.Nil(loadedComment)
}
func (suite *CommentTestSuite) TestLoadAll() {
	// Create a user for testing
	result := suite.db.Create(&entity.User{ID: 2, UserRegister: entity.UserRegister{UserLogin: entity.UserLogin{Login: "test 2"}, Fullname: "Test 2", Email: "test2"}})
	suite.NoError(result.Error)
	result = suite.db.Create(&entity.User{ID: 3, UserRegister: entity.UserRegister{UserLogin: entity.UserLogin{Login: "test 3"}, Fullname: "Test 3", Email: "test3"}})
	suite.NoError(result.Error)

	// Create comments for testing
	comments := []entity.Comment{
		{CommentBase: entity.CommentBase{MediaID: 2, SenderID: 2, Content: "Test Comment 1"}},
		{CommentBase: entity.CommentBase{MediaID: 2, SenderID: 3, Content: "Test Comment 2"}},
		{CommentBase: entity.CommentBase{MediaID: 1, SenderID: 2, Content: "Test Comment 3"}},
	}
	for _, comment := range comments {
		err := suite.repo.Create(&comment)
		suite.NoError(err)
	}

	// Call the LoadAll method
	loadedComments, err := suite.repo.LoadAll(2)
	suite.NoError(err)

	bases := []entity.CommentBase{
		comments[0].CommentBase,
		comments[1].CommentBase,
	}

	// Assert that the loaded comments match the expected comments
	suite.Len(*loadedComments, 2)
	suite.Contains(bases, comments[0].CommentBase)
	suite.Contains(bases, comments[1].CommentBase)
	suite.NotContains(bases, comments[2].CommentBase)
}
