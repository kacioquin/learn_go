package controllers

import (
	"bytes"
	"fmt"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/kacioquin/learn_go/api_gin/api/entities"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (t *tweetController) FindAll(ctx *gin.Context) {
	logger.Print("Hello, log file!")
	fmt.Print(&buf)
	ctx.JSON(http.StatusOK, t.tweets)
}

func (t *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		return
	}

	t.tweets = append(t.tweets, *tweet)

	ctx.JSON(http.StatusOK, tweet)
}

func (t *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for idx, tweet := range t.tweets {
		if tweet.ID == id {
			t.tweets = append(t.tweets[0:idx], t.tweets[idx+1:]...)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Tweet not found",
	})

}
