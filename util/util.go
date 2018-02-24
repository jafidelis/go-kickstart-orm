package util

import (
	"crypto/md5"
	"encoding/json"
	"log"
	"os"
	"strings"
	"unicode"

	"io"

	"encoding/hex"

	"fmt"

	"github.com/go-kickstart-orm/model/entity"
)

const keyword128 = "38512b786c584f2e6d487a767d"

//GetConfiguration - return the application configuration
func GetConfiguration() (entity.Configuration, error) {
	config := entity.Configuration{}
	file, err := os.Open("./config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}

//CheckErr - default treatment
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//ConvertCamelCaseToUnderscore - example camelCase to camel_case
func ConvertCamelCaseToUnderscore(str string) string {
	var newStr string
	for _, char := range str {
		if unicode.IsUpper(char) && char != ' ' {
			newStr += "_" + strings.ToLower(string(char))
			continue
		}
		newStr += strings.ToLower(string(char))
	}
	return newStr
}

//EncriptyMd5 generate MD5 with keyword
func EncriptyMd5(s string) (string, error) {
	hasher := md5.New()
	_, err := io.WriteString(hasher, s)
	if err != nil {
		return "", err
	}
	fmt.Println("crypto", hasher)

	_, err = io.WriteString(hasher, keyword128)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
