package subcmd

import (
	"github.com/gitdog7/gitdog/src/config"
	"github.com/gitdog7/gitdog/src/repo"
	"github.com/spf13/cobra"
	"os"
	"strings"

	log "github.com/google/logger"
)

// loadRepository
func loadRepository(cmd *cobra.Command) *repo.GitHubRepo {
	// parse flags
	workdir, _ := cmd.Flags().GetString("workdir")

	// check working directory exists
	if _, err := os.Stat(workdir); os.IsNotExist(err) {
		// path exists
		log.Fatalf("working directory: %v doesn't exists.", workdir)
	}

	// load config file
	configFileName := workdir + "/" + config.ConfigFileName
	config := config.Load(configFileName)
	if config == nil {
		log.Fatalf("failed to load config file in %v", configFileName)
	}

	// load data
	repository := repo.Load(workdir + "/" + repo.RepoDataFileName)
	return repository
}

func getGraphSaveName(title string) string {
	result := strings.Replace(title, " ", "_", -1)
	result = strings.Replace(result, "/", "|", -1)
	result = strings.Replace(result, "(", "_", -1)
	result = strings.Replace(result, ")", "_", -1)
	return result
}
