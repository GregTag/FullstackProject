package main

import (
	"backend/internal/entity"
	"backend/internal/repository"
	repository_sqlite "backend/internal/repository/sqlite"
	"fmt"
)

func checkUser(repository entity.UserRepository) {
	fmt.Println("###### Checking User Repository")

	// Create a new user
	user := &entity.User{
		UserRegister: entity.UserRegister{
			UserLogin: entity.UserLogin{
				Login:    "testUser",
				Password: "testpassword",
			},
			Fullname: "Test User",
			Email:    "test@test.com",
		},
		Avatar: "testAvatar",
	}
	err := repository.Create(user)
	if err != nil {
		fmt.Println("Failed to create user:", err)
	} else {
		fmt.Println("User created successfully:", *user)
	}

	// Get the user by login
	retrievedUser, err := repository.GetByLogin("testUser")
	if err != nil {
		fmt.Println("Failed to get user by login:", err)
	} else {
		fmt.Println("User retrieved successfully:", *retrievedUser)
	}

	// Update the user
	user.Avatar = "updatedAvatar"
	err = repository.Update(user)
	if err != nil {
		fmt.Println("Failed to create user:", err)
	} else {
		fmt.Println("User updated successfully:", *user)
	}

	retrievedUser, err = repository.GetByLogin("testUser")
	if err != nil {
		fmt.Println("Failed to get user by login:", err)
	} else {
		fmt.Println("User retrieved successfully:", *retrievedUser)
	}

	// Delete the user
	err = repository.Delete(retrievedUser.ID)
	if err != nil {
		fmt.Println("Failed to delete user:", err)
	} else {
		fmt.Println("User deleted successfully")
	}

	fmt.Println("###### Ended Checking User Repository")
}

func checkMedia(repository entity.MediaRepository) {
	fmt.Println("###### Checking Media Repository")

	// Create a new media
	media := &entity.Media{
		Title:            "Test Media",
		Description:      "This is a test media",
		Image:            "testImage",
		Category:         "Test Category",
		CumulativeRating: 0,
		NumberOfRatings:  0,
		NumberOfTracks:   0,
		ReleaseYear:      2023,
		Duration:         0,
		Genres:           []entity.Genre{},
	}
	err := repository.Create(media)
	if err != nil {
		fmt.Println("Failed to create media:", err)
	} else {
		fmt.Println("Media created successfully:", *media)
	}

	// Load the media by id
	loadedMedia, err := repository.Load(media.ID)
	if err != nil {
		fmt.Println("Failed to load media by id:", err)
	} else {
		fmt.Println("Media loaded successfully:", *loadedMedia)
	}

	// Search the media
	filter := &entity.Filter{
		Title:        "Test Media",
		Category:     "Test Category",
		Genres:       []string{},
		YearFrom:     2023,
		YearTo:       2023,
		DurationFrom: 0,
		DurationTo:   0,
		RatingFrom:   0,
		RatingTo:     0,
	}
	searchResults, err := repository.Search(filter)
	if err != nil {
		fmt.Println("Failed to search media:", err)
	} else {
		fmt.Println("Media search results:", *searchResults)
	}

	// Update the media
	media.Title = "Updated Test Media"
	err = repository.Update(media)
	if err != nil {
		fmt.Println("Failed to update media:", err)
	} else {
		fmt.Println("Media updated successfully:", *media)
	}

	retrievedMedia, err := repository.Get(media.ID)
	if err != nil {
		fmt.Println("Failed to get media by id:", err)
	} else {
		fmt.Println("Media retrieved successfully:", *retrievedMedia)
	}

	// Delete the media
	err = repository.Delete(retrievedMedia.ID)
	if err != nil {
		fmt.Println("Failed to delete media:", err)
	} else {
		fmt.Println("Media deleted successfully")
	}

	fmt.Println("###### Ended Checking Media Repository")
}

func checkComment(repository entity.CommentRepository) {
	fmt.Println("###### Checking Comment Repository")

	// Create a new comment
	comment := &entity.Comment{
		MediaID:  1,
		SenderID: 1,
		Content:  "Test Comment",
	}
	err := repository.Create(comment)
	if err != nil {
		fmt.Println("Failed to create comment:", err)
	} else {
		fmt.Println("Comment created successfully:", *comment)
	}

	// Get the comment by id
	retrievedComment, err := repository.Get(comment.ID)
	if err != nil {
		fmt.Println("Failed to get comment by id:", err)
	} else {
		fmt.Println("Comment retrieved successfully:", *retrievedComment)
	}

	// Load all comments for a media
	comments, err := repository.LoadAll(comment.MediaID)
	if err != nil {
		fmt.Println("Failed to load comments:", err)
	} else {
		fmt.Println("Comments loaded successfully:", *comments)
	}

	// Update the comment
	comment.Content = "Updated Test Comment"
	err = repository.Update(comment)
	if err != nil {
		fmt.Println("Failed to update comment:", err)
	} else {
		fmt.Println("Comment updated successfully:", *comment)
	}

	retrievedComment, err = repository.Get(comment.ID)
	if err != nil {
		fmt.Println("Failed to get comment by id:", err)
	} else {
		fmt.Println("Comment retrieved successfully:", *retrievedComment)
	}

	// Delete the comment
	err = repository.Delete(retrievedComment.ID)
	if err != nil {
		fmt.Println("Failed to delete comment:", err)
	} else {
		fmt.Println("Comment deleted successfully")
	}

	fmt.Println("###### Ended Checking Comment Repository")
}

func checkMediaTrack(repository entity.MediaTrackRepository) {
	fmt.Println("###### Checking MediaTrack Repository")

	// Create a new media track
	mediaTrack := &entity.MediaTrack{
		UserID:      1,
		MediaID:     1,
		Rating:      5,
		TrackStatus: "Test Status",
	}
	err := repository.Create(mediaTrack)
	if err != nil {
		fmt.Println("Failed to create media track:", err)
	} else {
		fmt.Println("MediaTrack created successfully:", *mediaTrack)
	}

	// Get the media track by user_id and media_id
	retrievedMediaTrack, err := repository.Get(mediaTrack.UserID, mediaTrack.MediaID)
	if err != nil {
		fmt.Println("Failed to get media track by user_id and media_id:", err)
	} else {
		fmt.Println("MediaTrack retrieved successfully:", *retrievedMediaTrack)
	}

	// Load all media tracks for a user
	mediaTracks, err := repository.LoadAll(mediaTrack.UserID)
	if err != nil {
		fmt.Println("Failed to load media tracks:", err)
	} else {
		fmt.Println("MediaTracks loaded successfully:", *mediaTracks)
	}

	// Update the media track
	mediaTrack.Rating = 10
	err = repository.Update(mediaTrack)
	if err != nil {
		fmt.Println("Failed to update media track:", err)
	} else {
		fmt.Println("MediaTrack updated successfully:", *mediaTrack)
	}

	retrievedMediaTrack, err = repository.Get(mediaTrack.UserID, mediaTrack.MediaID)
	if err != nil {
		fmt.Println("Failed to get media track by user_id and media_id:", err)
	} else {
		fmt.Println("MediaTrack retrieved successfully:", *retrievedMediaTrack)
	}

	// Delete the media track
	err = repository.Delete(retrievedMediaTrack.UserID, retrievedMediaTrack.MediaID)
	if err != nil {
		fmt.Println("Failed to delete media track:", err)
	} else {
		fmt.Println("MediaTrack deleted successfully")
	}

	fmt.Println("###### Ended Checking MediaTrack Repository")
}

func main() {
	db, err := repository_sqlite.NewSQLiteDB("test.db")
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepository(db)

	checkUser(repo.User)
	checkMedia(repo.Media)
	checkComment(repo.Comment)
	checkMediaTrack(repo.MediaTrack)
}
