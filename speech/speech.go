package speech

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"sync"

	"github.com/go-redis/redis"
)

type Phrase struct {
	RedisAddr      string
	RedisPassword  string
}

var redisClient  *redis.Client
var speakingLock *sync.Mutex
var voiceID      *string

func InitSpeech(voice string, address string, password string) {
	// Connect to Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       0,       // use default DB
	})

	voiceID = &voice
}

func sayWithID(phrase string, voiceID string) {
	//voiceHash := voiceHash(voiceID, phrase)
}

func voiceHash(voiceID string, phrase string) string {
	// Concat into std m
	msg := fmt.Sprintf("polly|%s|%s", strings.ToLower(voiceID), strings.ToLower(phrase))

	h := sha256.New()
	h.Write([]byte(msg))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func CloseSpeech() {

}