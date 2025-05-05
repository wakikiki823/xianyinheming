package response

import (
	"net/http"
	"time"
	"wenxinhexuan/models"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type TokenResponse struct {
	User      string    `json:"user"`
	Token     string    `json:"token"`
	ExpiresAT time.Time `json:"expiresAT"`
}

type VerificationResponse struct {
	Username string `json:"username"`
}

type SongResponse struct {
	Song_id string `json:"song_id"`
}

type SongsResponse struct {
	Song_ids []models.Song `json:"song_ids"`
}

type UploadToken struct {
	UploadToken string `json:"upload_token"`
}

type BroadcastBehave struct {
	Code   int    `json:"code"`
	Event  string `json:"event"`
	UserID int    `json:"userid"`
	RoomID string `json:"roomid"`
	Score  int    `json:"score"`
}

type ScoreReponse struct {
	UserID   int       `json:"user_id"`
	Username string    `json:"username"`
	Score    int       `json:"score"`
	Title    string    `json:"song_title"`
	Time     time.Time `json:"time"`
}

const (
	ERROR            = 0
	SUCCESS          = 1
	NEW_PLAYER_ENTER = 30001
	PLAYER_EXIT      = 30002
	ROOM_CREATED     = 30003
	UPDATE_SCORE     = 20001
)

func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

// 待定
// func WebSocketResult(Event string, UserID uint, RoomID string)

func OKWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, message, map[string]interface{}{}, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, message, map[string]interface{}{}, c)
}

func OKWithDetailed(message string, data interface{}, c *gin.Context) {
	Result(SUCCESS, message, data, c)
}
