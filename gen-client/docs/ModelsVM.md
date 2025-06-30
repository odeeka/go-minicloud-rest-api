# ModelsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | Pointer to **string** | ContainerID is the Docker container ID simulating the VM. required: false | [optional] 
**Cpu** | Pointer to **float32** | CPU represents the number of virtual CPUs allocated to the VM. example: 1.5 | [optional] 
**Env** | Pointer to **map[string]string** | Env contains environment variables passed to the container. example: {\&quot;ENV\&quot;: \&quot;dev\&quot;} | [optional] 
**Id** | Pointer to **int32** | ID is the unique identifier of the virtual machine. required: true | [optional] 
**Image** | Pointer to **string** | Image used to start the virtual machine (e.g., \&quot;nginx\&quot;, \&quot;ubuntu\&quot;). required: true | [optional] 
**Memory** | Pointer to **int32** | Memory in megabytes allocated to the VM. example: 1024 | [optional] 
**Name** | Pointer to **string** | Name of the virtual machine. required: true | [optional] 
**Ports** | Pointer to **[]int32** | Ports to expose on the container (e.g., [80, 443]). example: [80, 443] | [optional] 

## Methods

### NewModelsVM

`func NewModelsVM() *ModelsVM`

NewModelsVM instantiates a new ModelsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsVMWithDefaults

`func NewModelsVMWithDefaults() *ModelsVM`

NewModelsVMWithDefaults instantiates a new ModelsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *ModelsVM) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ModelsVM) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ModelsVM) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ModelsVM) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetCpu

`func (o *ModelsVM) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ModelsVM) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ModelsVM) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ModelsVM) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetEnv

`func (o *ModelsVM) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ModelsVM) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ModelsVM) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ModelsVM) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetId

`func (o *ModelsVM) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsVM) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsVM) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsVM) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *ModelsVM) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ModelsVM) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ModelsVM) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ModelsVM) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMemory

`func (o *ModelsVM) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ModelsVM) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ModelsVM) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ModelsVM) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetName

`func (o *ModelsVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsVM) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *ModelsVM) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ModelsVM) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ModelsVM) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ModelsVM) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


