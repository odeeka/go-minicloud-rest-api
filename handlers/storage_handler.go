// Manage the incoming HTTP requests
package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/odeeka/go-minicloud-rest-api/models"
	"github.com/odeeka/go-minicloud-rest-api/services"
)

func ListStorages(context *gin.Context) {
	storages, err := models.GetAllStorages()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retriev all Storages"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Fetched all Storages from database", "Storages": storages})
}

func CreateStorage(context *gin.Context) {
	var storage models.Storage

	err := context.ShouldBindJSON(&storage)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err})
		return
	}

	// Generate storage name if it's not defined
	if storage.Name == "" {

		// with timestamp
		//timestamp := time.Now().Unix()
		//storage.Name = fmt.Sprintf("minicloud-storage-%d", timestamp)
		// with custom UUID
		storage.Name = fmt.Sprintf("minicloud-storage-%s", uuid.New().String()[:8])
	}

	err = services.StartStorageVolume(&storage)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to start/create the storage volume", "error": err.Error()})
		return
	}

	// If the containers runs the metadata will be inserted
	err = storage.InsertStorage()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create storage metadata into database.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Storage created and stored in database", "Storage": storage})
}

func GetStorage(context *gin.Context) {
	storageId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse Storage id.", "error": err.Error()})
		return
	}

	storage, err := models.GetStorageByID(storageId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch target Storage.", "error": err.Error()})
		return
	}

	if storage == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Storage not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Storage": storage})
}

func DeleteStorage(context *gin.Context) {

	storageId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse storage id.", "error": err.Error()})
		return
	}

	storage, err := models.GetStorageByID(storageId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the storage.", "error": err.Error()})
		return
	}

	// Check the attachment of storage to container
	// Don't delete if it's attached and the container is running
	// Stop and remove Docker container if exists
	// if vm.ContainerID != "" {
	// 	err = services.StopAndRemoveContainer(vm.ContainerID)
	// 	if err != nil {
	// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not remove VM (simulated container).", "error": err})
	// 		return
	// 	}
	// }

	err = storage.DeleteStorageByID()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the storage.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Storage deleted successfully with ID: " + strconv.FormatInt(storage.ID, 10)})
}

func UpdateStorageSize(context *gin.Context) {
	storageId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse storage id.", "error": err.Error()})
		return
	}

	storage, err := models.GetStorageByID(storageId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the storage.", "error": err.Error()})
		return
	}

	if storage == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Storage not found with ID: " + strconv.FormatInt(storage.ID, 10)})
		return
	}

	var updatedStorage models.Storage
	err = context.ShouldBindJSON(&updatedStorage)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	updatedStorage.ID = storageId
	err = updatedStorage.UpdateStorageSizeByID()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update storage size.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Storage size updated successfully!", "Storage": updatedStorage})
}

func AttachStorageToVM(context *gin.Context) {
	storageId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse storage id.", "error": err.Error()})
		return
	}

	vmId, err := strconv.ParseInt(context.Param("vmid"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse VM id.", "error": err.Error()})
		return
	}

	storage, err := models.GetStorageByID(storageId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the storage.", "error": err.Error()})
		return
	}

	if storage == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Storage not found with ID: " + strconv.FormatInt(storage.ID, 10)})
		return
	}

	vm, err := models.GetVMByID(vmId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the VM.", "error": err.Error()})
		return
	}

	if vm == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "VM not found with ID: " + strconv.FormatInt(vm.ID, 10)})
		return
	}

	storage.VMID = &vmId
	err = storage.AttachStorageByID()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not attach storage to VM.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Storage attached successfully!", "Storage": storage})
}

func DetachStorageFromVM(context *gin.Context) {
	storageId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse storage id.", "error": err.Error()})
		return
	}

	vmId, err := strconv.ParseInt(context.Param("vmid"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse VM id.", "error": err.Error()})
		return
	}

	storage, err := models.GetStorageByID(storageId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the storage.", "error": err.Error()})
		return
	}

	if storage == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Storage not found with ID: " + strconv.FormatInt(storage.ID, 10)})
		return
	}

	vm, err := models.GetVMByID(vmId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the VM.", "error": err.Error()})
		return
	}

	if vm == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "VM not found with ID: " + strconv.FormatInt(vm.ID, 10)})
		return
	}

	storage.ID = storageId
	err = storage.DetachStorageByID()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not detach storage to VM.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Storage detached successfully!", "Storage": storage})
}
