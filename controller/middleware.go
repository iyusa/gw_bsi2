package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ussidev/permata_trx/common"
	"github.com/ussidev/permata_trx/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("Middleware > start %s\n", c.FullPath())
		var signature = c.Request.Header.Get("signature")
		var instCode = c.Request.Header.Get("x-institution-code")

		apiKey, ok := model.ApiKeys[instCode]
		if !ok {
			fmt.Printf("Middleware > invalid institution code %s", instCode)
			handleError(c, fmt.Errorf("invalid institution code %s", instCode))
			return
		}
		organizationName := model.OrgNames[instCode] //model.GetOrgName(instCode)
		fmt.Printf("Middleware > InstCode: %s, apiKey: %s, orgName %s\n", instCode, apiKey, organizationName)

		var body, err = c.GetRawData()
		if err != nil {
			fmt.Println("Middleware > Invalid body")
			handleError(c, fmt.Errorf("invalid body"))
			return
		}

		decrypted, err := decryptBody(body, apiKey)
		if err != nil {
			fmt.Printf("Middleware > error decrypt %v\n", err)
			handleError(c, err)
			return
		}

		err = validateSignature(c, signature, apiKey, string(decrypted))
		if err != nil {
			fmt.Printf("Middleware > error signature %v\n", err)
			handleError(c, err)
			return
		}

		// balikin body ke request
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(decrypted))

		// set api key & orgname sebagai global
		c.Set("institution_code", instCode)
		c.Set("api_key", apiKey)
		c.Set("organization_name", organizationName)
		c.Next()
	}
}

func decryptBody(body []byte, apiKey string) (decryptedBody []byte, err error) {
	apiKeyByte := []byte(apiKey)

	decryptedBody, err = common.AESDecrypt(body, apiKeyByte, apiKeyByte)
	// decText, err := aesLib.Decrypt(apiKeyByte, apiKeyByte, string(body))
	// decryptedBody = []byte(decText)
	return
}

func validateSignature(c *gin.Context, signature, apiKey, decryptedBody string) (err error) {
	path := c.Request.URL.Path //c.FullPath()

	fmt.Printf("Will calculate signature, path: %s, apiKey: %s, \nbody: %s\n", path, apiKey, decryptedBody)
	tts := fmt.Sprintf("%s:%s:%s", apiKey, path, strings.Trim(decryptedBody, " "))
	tts = strings.Trim(tts, " ")
	calculated := common.CreateSignature(tts)
	if calculated != signature {
		fmt.Printf("Calculated signature: %s\n", calculated)
		fmt.Printf("Expected   signature: %s\n", signature)
		err = fmt.Errorf("invalid signature")
	}
	return
}
