package services

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/odeeka/go-minicloud-rest-api/models"
)

func StartStorageVolume(storage *models.Storage) error {

	volumeName := storage.Name

	cmd := exec.Command("docker", "volume", "create", volumeName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker volume failed: %w, output: %s", err, strings.TrimSpace(string(output)))
	}

	return nil
}
