package main

import (
	"fmt"
	"os"
	"regexp"
)

type Config struct {
	VoiceID        string

	RedisAddr      string
	RedisPassword  string
}

func CollectConfig() (config Config) {
	var missingEnv []string

	// REDIS_ADDR
	var envRedisAddress string = os.Getenv("REDIS_ADDR")

	if envRedisAddress == "" {
		config.RedisAddr = "localhost"
	} else {
		config.RedisAddr = envRedisAddress
	}

	// REDIS_PASSWORD
	config.RedisPassword = os.Getenv("REDIS_PASSWORD")

	// REDIS_ADDR
	var envVoiceID string = os.Getenv("REDIS_ADDR")

	if envVoiceID == "" {
		config.VoiceID = "Raveena"
	} else {
		config.VoiceID = envVoiceID
	}

	// Validation
	if len(missingEnv) > 0 {
		var msg string = fmt.Sprintf("Environment variables missing: %v", missingEnv)
		logger.Criticalf(msg)
		panic(fmt.Sprint(msg))
	}

	return
}

func DecodeEngine(engine string) (dialect string, args string) {
	pgRe, err := regexp.Compile(`postgresql://([\w]*):([\w\-.~:/?#\[\]!$&'()*+,;=]*)@([\w.]*)/([\w]*)`)
	if err != nil {
		logger.Criticalf("Regex compile error: %s", err)
		panic("PANIC!")
	}

	if pgRe.MatchString(engine) {
		dialect = "postgres"
		match := pgRe.FindStringSubmatch(engine)
		args = fmt.Sprintf("host=%s user=%s dbname=%s password=%s", match[3], match[1], match[4], match[2])
	} else {
		panic(fmt.Sprint("Could not parse DB_ENGINE"))
	}

	return
}