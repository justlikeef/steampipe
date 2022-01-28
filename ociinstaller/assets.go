package ociinstaller

import (
	"context"

	"github.com/turbot/steampipe/constants"
)

// InstallAssets Installs the Steampipe Postgres foreign data wrapper from an OCI image
func InstallAssets(ctx context.Context, dest string) error {
	imageRef := constants.AssetsImageRef
	tempDir := NewTempDir(imageRef)
	defer tempDir.Delete()

	// download the blobs.
	imageDownloader := NewOciDownloader()
	image, err := imageDownloader.Download(ctx, imageRef, "assets", tempDir.Path)
	if err != nil {
		return err
	}

	// install the files
	if err = installAssetsFiles(image, tempDir.Path, dest); err != nil {
		return err
	}

	return nil
}

func installAssetsFiles(image *SteampipeImage, tempdir string, dest string) error {

	return nil
}
