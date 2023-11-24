package repository_sqlite_test

import (
	"backend/internal/entity"
	repository_sqlite "backend/internal/repository/sqlite"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo *repository_sqlite.UserSQLite
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (suite *UserTestSuite) SetupTest() {
	db, err := repository_sqlite.NewSQLiteDB("file::memory:?cache=shared")
	if err != nil {
		suite.Fail("Failed to connect database", err)
	}

	suite.db = db
	suite.repo = repository_sqlite.NewUserSQLite(db)
}

func (suite *UserTestSuite) TestCreate() {
	// Create a new user
	user := &entity.User{
		UserRegister: entity.UserRegister{
			UserLogin: entity.UserLogin{
				Login:    "testuser0",
				Password: "testpassword",
			},
			Fullname: "Test User",
			Email:    "test0@example.com",
		},
	}

	// Call the Create method
	err := suite.repo.Create(user)

	// Assert that there are no errors
	suite.NoError(err)

	// Assert that the user was created successfully
	var createdUser entity.User
	result := suite.db.First(&createdUser, user.ID)
	suite.NoError(result.Error)
	suite.Equal(user, &createdUser)

	// Attempt to create the same user again
	err = suite.repo.Create(user)

	// Assert that the appropriate error is returned
	suite.Equal(entity.ErrUserAlreadyExists, err)
}
func (suite *UserTestSuite) TestUpdate() {
	// Create a new user
	user := &entity.User{
		UserRegister: entity.UserRegister{
			UserLogin: entity.UserLogin{
				Login:    "testuser1",
				Password: "testpassword",
			},
			Fullname: "Test User",
			Email:    "test1@example.com",
		},
	}

	// Call the Create method to insert the user into the database
	err := suite.repo.Create(user)
	suite.NoError(err)

	// Update the user's email
	user.Email = "updated@example.com"

	// Call the Update method
	err = suite.repo.Update(user)

	// Assert that there are no errors
	suite.NoError(err)

	// Retrieve the updated user from the database
	var updatedUser entity.User
	result := suite.db.First(&updatedUser, user.ID)
	suite.NoError(result.Error)

	// Assert that the user's email has been updated
	suite.Equal(user.Email, updatedUser.Email)

	// Attempt to update a non-existing user
	nonExistingUser := &entity.User{
		ID: 9999,
	}
	err = suite.repo.Update(nonExistingUser)

	// Assert that the appropriate error is returned
	suite.Equal(entity.ErrUserNotFound, err)
}

func (suite *UserTestSuite) TestDelete() {
	// Create a new user
	user := &entity.User{
		UserRegister: entity.UserRegister{
			UserLogin: entity.UserLogin{
				Login:    "testuser2",
				Password: "testpassword",
			},
			Fullname: "Test User",
			Email:    "test2@example.com",
		},
	}

	// Add the user to the database
	err := suite.repo.Create(user)
	suite.NoError(err)

	// Delete the user
	err = suite.repo.Delete(user.ID)

	// Assert that there are no errors
	suite.NoError(err)

	// Attempt to delete the same user again
	err = suite.repo.Delete(user.ID)

	// Assert that the appropriate error is returned
	suite.Equal(entity.ErrUserNotFound, err)
}
func (suite *UserTestSuite) TestGet() {
	// Create a new user
	user := &entity.User{
		UserRegister: entity.UserRegister{
			UserLogin: entity.UserLogin{
				Login:    "testuser3",
				Password: "testpassword",
			},
			Fullname: "Test User",
			Email:    "test3@example.com",
		},
	}

	// Add the user to the database
	err := suite.repo.Create(user)
	suite.NoError(err)

	// Call the Get method to retrieve the user
	retrievedUser, err := suite.repo.Get(user.ID)

	// Assert that there are no errors
	suite.NoError(err)

	// Assert that the retrieved user matches the original user
	suite.Equal(user, retrievedUser)

	// Attempt to retrieve a non-existing user
	nonExistingUser, err := suite.repo.Get(9999)

	// Assert that the appropriate error is returned
	suite.Equal(entity.ErrUserNotFound, err)
	suite.Nil(nonExistingUser)
}
func (suite *UserTestSuite) TestGetByLogin() {
	// Create a new user
	user := &entity.User{
		UserRegister: entity.UserRegister{
			UserLogin: entity.UserLogin{
				Login:    "testuser4",
				Password: "testpassword",
			},
			Fullname: "Test User",
			Email:    "test4@example.com",
		},
	}

	// Add the user to the database
	err := suite.repo.Create(user)
	suite.NoError(err)

	// Call the GetByLogin method to retrieve the user
	retrievedUser, err := suite.repo.GetByLogin(user.Login)

	// Assert that there are no errors
	suite.NoError(err)

	// Assert that the retrieved user matches the original user
	suite.Equal(user, retrievedUser)

	// Attempt to retrieve a non-existing user
	nonExistingUser, err := suite.repo.GetByLogin("nonexistinguser")

	// Assert that the appropriate error is returned
	suite.Equal(entity.ErrUserNotFound, err)
	suite.Nil(nonExistingUser)
}
func (suite *UserTestSuite) TestGetViewByLogin() {
	// Create a new user
	user := &entity.User{
		UserRegister: entity.UserRegister{
			UserLogin: entity.UserLogin{
				Login:    "testuser5",
				Password: "testpassword",
			},
			Fullname: "Test User",
			Email:    "test5@example.com",
		},
	}

	// Add the user to the database
	err := suite.repo.Create(user)
	suite.NoError(err)

	// Call the GetViewByLogin method to retrieve the user view
	retrievedUserView, err := suite.repo.GetViewByLogin(user.Login)

	// Assert that there are no errors
	suite.NoError(err)

	// Assert that the retrieved user view matches the original user
	expectedUserView := &entity.UserView{
		ID:       user.ID,
		Login:    user.Login,
		Fullname: user.Fullname,
	}
	suite.Equal(expectedUserView, retrievedUserView)

	// Attempt to retrieve a non-existing user view
	nonExistingUserView, err := suite.repo.GetViewByLogin("nonexistinguser")

	// Assert that the appropriate error is returned
	suite.Equal(entity.ErrUserNotFound, err)
	suite.Nil(nonExistingUserView)
}
func (suite *UserTestSuite) TestFillTracks() {
	// Create a new user
	user := &entity.User{
		UserRegister: entity.UserRegister{
			UserLogin: entity.UserLogin{
				Login:    "testuser6",
				Password: "testpassword",
			},
			Fullname: "Test User",
			Email:    "test6@example.com",
		},
	}

	// Add the user to the database
	err := suite.repo.Create(user)
	suite.NoError(err)

	// Create a new user view
	userView := &entity.UserView{
		ID:       user.ID,
		Login:    user.Login,
		Fullname: user.Fullname,
	}

	// Call the FillTracks method to fill the tracks for the user view
	userInfo, err := suite.repo.FillTracks(userView)
	suite.NoError(err)

	// Assert that there are no errors
	suite.NoError(err)

	// Assert that the user info contains the correct user view
	suite.Equal(userView, &userInfo.User)

	// Assert that the user info contains the correct tracks
	expectedTracks, err := repository_sqlite.NewMediaTrackSQLite(suite.db).LoadAll(userView.ID)
	suite.NoError(err)
	suite.Equal(expectedTracks, &userInfo.Tracks)
}
