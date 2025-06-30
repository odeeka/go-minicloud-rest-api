# ModelsAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier | [optional] 
**Password** | Pointer to **string** | Password (stored hashed, omit in response) | [optional] 
**Username** | Pointer to **string** | Username for login | [optional] 

## Methods

### NewModelsAccount

`func NewModelsAccount() *ModelsAccount`

NewModelsAccount instantiates a new ModelsAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsAccountWithDefaults

`func NewModelsAccountWithDefaults() *ModelsAccount`

NewModelsAccountWithDefaults instantiates a new ModelsAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsAccount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassword

`func (o *ModelsAccount) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsAccount) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsAccount) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelsAccount) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *ModelsAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsAccount) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelsAccount) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


