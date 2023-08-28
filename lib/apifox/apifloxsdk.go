package apifox

import (
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/spf13/viper"
	"log"
	"strings"
)

const BASE_URL = "https://api.apifox.cn/api/v1"

type ImportOpenApiResponse struct {
	Success      bool   `json:"success"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type ImportOpenApiRequest struct {
	ImportFormat        string `json:"importFormat"`
	Data                string `json:"data"`
	SchemaOverwriteMode string `json:"schemaOverwriteMode"`
	ApiOverwriteMode    string `json:"apiOverwriteMode"`
	SyncApiFolder       bool   `json:"syncApiFolder"`
}

func getAccessToken() string {

	viper.SetConfigName("config")           // name of config file (without extension)
	viper.SetConfigType("yaml")             // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/apifoxcli/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.apifoxcli") // call multiple times to add many search paths
	viper.AddConfigPath(".")                // optionally look for config in the working directory
	err := viper.ReadInConfig()             // Find and read the config file
	if err != nil {                         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	accessToken := viper.GetString("accessToken")

	if accessToken == "" {
		panic("config accessToken required.")

	}

	return accessToken

}

func ImportOpenApi(projectId string, content string) {

	accessToken := getAccessToken()

	url := BASE_URL + "/projects/" + projectId + "/import-data"

	content = strings.Replace(content, "\n", "", -1)
	content = strings.Replace(content, " ", "", -1)

	r := ImportOpenApiRequest{
		ImportFormat:        "openapi",
		SchemaOverwriteMode: "name",
		ApiOverwriteMode:    "methodAndPath",
		SyncApiFolder:       true,
		Data:                content,
	}

	var response ImportOpenApiResponse

	gout.POST(url).Debug(false).SetJSON(&r).
		SetHeader(gout.H{
			"X-Apifox-Version": "2022-11-16",
			"Authorization":    "Bearer " + accessToken,
		}).
		BindJSON(&response).Do()

	if !response.Success {

		log.Println(response.ErrorCode, response.ErrorMessage)
	} else {
		log.Println("import success")

	}

}
