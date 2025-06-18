package services

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/odeeka/go-minicloud-rest-api/models"
)

// StartContainer simulates the VM by running a Docker container for the VM configuration
func StartContainer(vm *models.VM) error {
	// Alap parancs és image
	args := []string{"run", "-d"}

	// Portok átirányítása: csak érvényes 1–65535 portokat engedünk
	for _, port := range vm.Ports {
		if port > 0 && port <= 65535 {
			portMapping := fmt.Sprintf("%d:%d", port, port)
			args = append(args, "-p", portMapping)
		}
	}

	// Environment változók
	for key, value := range vm.Env {
		envMapping := fmt.Sprintf("%s=%s", key, value)
		args = append(args, "-e", envMapping)
	}

	// Image
	args = append(args, vm.Image)

	// (Debug) Kiírjuk a parancsot, amit futtatunk
	fmt.Println("Running Docker with args:", args)

	// Parancs futtatása
	cmd := exec.Command("docker", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker run failed: %w, output: %s", err, strings.TrimSpace(string(output)))
	}

	// Container ID mentése a VM-be
	vm.ContainerID = strings.TrimSpace(string(output))

	return nil
}

// StopAndRemoveContainer stops and removes a Docker container by ID
func StopAndRemoveContainer(containerID string) error {
	// Stop
	stopCmd := exec.Command("docker", "stop", containerID)
	stopOut, stopErr := stopCmd.CombinedOutput()
	if stopErr != nil {
		return fmt.Errorf("failed to stop container: %w, output: %s", stopErr, strings.TrimSpace(string(stopOut)))
	}

	// Remove
	rmCmd := exec.Command("docker", "rm", containerID)
	rmOut, rmErr := rmCmd.CombinedOutput()
	if rmErr != nil {
		return fmt.Errorf("failed to remove container: %w, output: %s", rmErr, strings.TrimSpace(string(rmOut)))
	}

	return nil
}
