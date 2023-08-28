package apifox

import (
	"os"
	"testing"
)

func TestImportOpenApi(t *testing.T) {

	projectId := "3214855"
	docsPath := "./swagger.json"

	content, _ := os.ReadFile(docsPath)

	ImportOpenApi(projectId, string(content))

}
