// Manage the incoming HTTP requests
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/odeeka/go-minicloud-rest-api/models"
	"github.com/odeeka/go-minicloud-rest-api/services"
)

func ListVMs(context *gin.Context) {
	vms, err := models.GetAllVms()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retriev all VMs"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Fetched all VMS from database", "VMS": vms})
}

func CreateVM(context *gin.Context) {
	var vm models.VM

	err := context.ShouldBindJSON(&vm)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err})
		return
	}

	// Here create and start the Docker container to simulate the VM
	err = services.StartContainer(&vm)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to start the VM", "error": err.Error()})
		return
	}

	// If the containers runs the metadata will be inserted
	err = vm.InsertVM()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create VM metadata into database.", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "VM created and stored in database", "VM": vm})
}

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

func UpdateVM(context *gin.Context) {
	vmId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse VM id.", "error": err})
		return
	}

	//vm, err := models.GetVMByID(vmId)

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

	updatedVM.ID = vmId
	err = updatedVM.UpdateVMByID()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update VM.", "error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "VM updated successfully!", "VM": updatedVM})
}
