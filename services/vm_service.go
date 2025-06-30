package services

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/odeeka/go-minicloud-rest-api/models"
)

// StartContainer simulates the VM by running a Docker container for the VM configuration
func StartContainer(vm *models.VM) error {
	// Base command to run
	args := []string{"run", "-d"}

	// Port redirections: only valid between 1–65535
	for _, port := range vm.Ports {
		if port > 0 && port <= 65535 {
			// Map host port to container's port 80 (default for nginx)
			//portMapping := fmt.Sprintf("%d:%d", port, port)
			portMapping := fmt.Sprintf("%d:%d", port, 80)
			args = append(args, "-p", portMapping)
		}
	}

	// Set the container name (not used random name)
	args = append(args, "--name", fmt.Sprintf("%s", vm.Name))

	// Memory limit (converted to string with 'm' suffix)
	args = append(args, "--memory", fmt.Sprintf("%dm", vm.Memory))

	// CPU limit (in decimal format, e.g., 0.5 means 50% of one CPU core)
	args = append(args, "--cpus", fmt.Sprintf("%.2f", vm.CPU))

	// Environment variables
	for key, value := range vm.Env {
		envMapping := fmt.Sprintf("%s=%s", key, value)
		args = append(args, "-e", envMapping)
	}

	// Base image
	args = append(args, vm.Image)

	// (Debug) Output the full command with arguments to terminal
	fmt.Println("Running Docker with args:", args)

	// Run docker command
	cmd := exec.Command("docker", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker run failed: %w, output: %s", err, strings.TrimSpace(string(output)))
	}

	// Save container ID into VM
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

// UpdateContainer simulates the VM by running a Docker container for the VM configuration
func UpdateContainer(vm *models.VM) error {

	// Base command for update
	args := []string{"update"}

	// Memory limit (converted to string with 'm' suffix)
	args = append(args, "--memory", fmt.Sprintf("%dm", vm.Memory))
	args = append(args, "--memory-swap", fmt.Sprintf("%dm", vm.Memory))

	// CPU limit (in decimal format, e.g., 0.5 means 50% of one CPU core)
	args = append(args, "--cpus", fmt.Sprintf("%.2f", vm.CPU))
	args = append(args, vm.Name)

	// (Debug) Show the command in terminal
	fmt.Println("Updating Docker with args:", args)

	// Run the command
	cmd := exec.Command("docker", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker update failed: %w, output: %s", err, strings.TrimSpace(string(output)))
	}

	return nil
}
