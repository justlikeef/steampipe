package reportassets

import (
	"context"
	_ "embed"
	"encoding/json"
	"log"
	"os"

	"github.com/turbot/steampipe/constants"
	"github.com/turbot/steampipe/ociinstaller"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe/filepaths"
)

func Ensure(ctx context.Context) error {
	// load report assets versions.json
	versionFile, err := LoadReportAssetVersionFile()
	if err != nil {
		return err
	}
	if versionFile.Version == constants.AssetsVersion {
		return nil
	}

	reportAssetsPath := filepaths.ReportAssetsPath()
	return ociinstaller.InstallAssets(ctx, reportAssetsPath)

	return nil
}

type ReportAssetsVersionFile struct {
	Version string `json:"version"`
}

func LoadReportAssetVersionFile() (*ReportAssetsVersionFile, error) {
	versionFilePath := filepaths.ReportAssetsVersionFilePath()
	if !helpers.FileExists(versionFilePath) {

		return &ReportAssetsVersionFile{}, nil
	}

	file, _ := os.ReadFile(versionFilePath)
	var versionFile ReportAssetsVersionFile
	if err := json.Unmarshal(file, &versionFile); err != nil {
		log.Println("[ERROR]", "Error while reading report assets version file", err)
		return nil, err
	}

	return &versionFile, nil

}
