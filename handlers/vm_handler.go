// Manage the incoming HTTP requests
package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/odeeka/go-minicloud-rest-api/models"
	"github.com/odeeka/go-minicloud-rest-api/services"
)

// ListVMs godoc
// @Summary      List all VMs
// @Description  Retrieves a list of all virtual machines
// @Tags         vms
// @Security BearerAuth
// @Accept json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]string
// @Router       /vms [get]
func ListVMs(context *gin.Context) {
	vms, err := models.GetAllVms()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retriev all VMs"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Fetched all VMS from database", "VMS": vms})
}

// CreateVM handles the POST request to create a new virtual machine simulation
// CreateVM godoc
// @Summary      Create a new VM
// @Description  Creates and stores metadata for a new virtual machine simulation
// @Tags         vms
// @Accept       json
// @Produce      json
// @Param        vm  body      models.VM  true  "VM to create"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /vms [post]
func CreateVM(context *gin.Context) {
	var vm models.VM

	// Parse the incoming JSON payload into the VM model
	err := context.ShouldBindJSON(&vm)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err})
		return
	}

	// Start a background container to simulate the virtual machine
	// This is the place to plug in different VM simulation technologies like Docker, VirtualBox, or others
	// If you only want to simulate the VM at the database level and do not need to start an actual service,
	// you can comment out the following line
	err = services.StartContainer(&vm)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to start the VM", "error": err.Error()})
		return
	}

	// Store the VM metadata in the database after the container has started
	err = vm.InsertVM()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create VM metadata into database.", "error": err})
		return
	}

	// Return success response with the VM details
	context.JSON(http.StatusCreated, gin.H{"message": "VM created and stored in database", "VM": vm})
}

// GetVM godoc
// @Summary      Get VM by ID
// @Description  Retrieves a single virtual machine by its ID
// @Tags         vms
// @Produce      json
// @Param        id   path      int  true  "VM ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /vms/{id} [get]
func GetVM(context *gin.Context) {
	vmId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse VM id.", "error": err.Error()})
		return
	}

	vm, err := models.GetVMByID(vmId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch target VM.", "error": err.Error()})
		return
	}
	if vm == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "VM not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"vm": vm})
}

// DeleteVM godoc
// @Summary      Delete VM
// @Description  Deletes a VM and removes its associated container
// @Tags         vms
// @Produce      json
// @Param        id   path      int  true  "VM ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /vms/{id} [delete]
func DeleteVM(context *gin.Context) {

	vmId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse VM id.", "error": err})
		return
	}

	vm, err := models.GetVMByID(vmId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the VM.", "error": err})
		return
	}

	// Stop and remove Docker container if exists
	if vm.ContainerID != "" {
		err = services.StopAndRemoveContainer(vm.ContainerID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not remove VM (simulated container).", "error": err})
			return
		}
	}

	err = vm.DeleteVMByID()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the VM.", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "VM deleted successfully with ID: " + strconv.FormatInt(vm.ID, 10)})
}

// UpdateVM godoc
// @Summary      Update VM
// @Description  Updates VM metadata or restarts container if necessary
// @Tags         vms
// @Accept       json
// @Produce      json
// @Param        id   path      int       true  "VM ID"
// @Param        vm   body      models.VM true  "Updated VM"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /vms/{id} [put]
func UpdateVM(context *gin.Context) {
	vmId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse VM id.", "error": err})
		return
	}

	// Get the current VM metadata
	vm, err := models.GetVMByID(vmId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the VM.", "error": err})
		return
	}

	var updatedVM models.VM
	err = context.ShouldBindJSON(&updatedVM)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err})
		return
	}

	needsRecreate := false

	fmt.Println("Needs recreate: ", needsRecreate)
	fmt.Println("VM image: ", vm.Image)
	fmt.Println("VM ports: ", vm.Ports)
	fmt.Println("VM env: ", vm.Env)

	if vm.Image != updatedVM.Image {
		needsRecreate = true
	}

	// Normalize the default values (client based data vs Terraform input)
	if updatedVM.Ports != nil && len(updatedVM.Ports) == 0 {
		updatedVM.Ports = nil
	}

	if updatedVM.Env != nil && len(updatedVM.Env) == 0 {
		updatedVM.Env = nil
	}

	if !reflect.DeepEqual(vm.Ports, updatedVM.Ports) {
		needsRecreate = true
	}

	if !reflect.DeepEqual(vm.Env, updatedVM.Env) {
		needsRecreate = true
	}

	fmt.Println("Needs recreate: ", needsRecreate)
	// Perform a live update of the container when only CPU or memory values change.
	// If image, ports, or environment variables are modified, the container should be recreated.
	// This function simulates VM updates using container technology (e.g., Docker, VirtualBox, etc.).
	if needsRecreate {

		err = services.StopAndRemoveContainer(vm.ContainerID)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to stop & remove the VM", "error": err.Error()})
		}

		err = services.StartContainer(&updatedVM)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to start the VM", "error": err.Error()})
			return
		}

	} else {
		updatedVM.ContainerID = vm.ContainerID // Keep the container ID
		err = services.UpdateContainer(&updatedVM)
	}
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update the VM", "error": err.Error()})
		return
	}

	// Update the database with new data
	updatedVM.ID = vmId
	err = updatedVM.UpdateVMByID()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the metadata of VM in database", "error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "VM updated successfully!", "VM": updatedVM})
}
