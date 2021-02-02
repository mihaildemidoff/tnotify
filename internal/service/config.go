package service

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

const botTokenKey = "bot.token"
const chatIdKey = "chat.id"

func ReadBotToken() string {
	return viper.GetString(botTokenKey)
}

func ReadChatId() int64 {
	return viper.GetInt64(chatIdKey)
}

func WriteBotToken(token string) {
	viper.Set(botTokenKey, token)
	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("Couldn't save bot token %s due to %s\n", token, err)
	}
}

func WriteChatId(chatId string) {
	viper.Set(chatIdKey, chatId)
	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("Couldn't save chat id token %s due to %s\n", chatId, err)
	}
}

func CheckConfigured() bool {
	return viper.Get(chatIdKey) != nil && viper.Get(botTokenKey) != nil
}

func InitConfigFile() {
	// stub init
	homeDir, _ := os.UserHomeDir()
	configHome := homeDir + "/.tnotify"
	configName := "config"
	configType := "yml"
	configPath := filepath.Join(configHome, configName+"."+configType)
	// ----

	viper.AddConfigPath(configHome)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	_, err := os.Stat(configHome)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(configHome, os.ModePerm); err != nil { // perm 0666
			log.Fatalf("Error occurred during config directory creation: %s\n", err)
		}
	}
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		if _, err := os.Create(configPath); err != nil { // perm 0666
			log.Fatalf("Error occurred during config file creation: %s\n", err)
		}
	}
}
