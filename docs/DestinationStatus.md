# DestinationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]string** |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Fields** | Pointer to [**[]DestinationField**](DestinationField.md) |  | [optional] 
**Live** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pixel** | Pointer to **bool** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Secrets** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDestinationStatus

`func NewDestinationStatus() *DestinationStatus`

NewDestinationStatus instantiates a new DestinationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationStatusWithDefaults

`func NewDestinationStatusWithDefaults() *DestinationStatus`

NewDestinationStatusWithDefaults instantiates a new DestinationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *DestinationStatus) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *DestinationStatus) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *DestinationStatus) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *DestinationStatus) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCategory

`func (o *DestinationStatus) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DestinationStatus) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DestinationStatus) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *DestinationStatus) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfig

`func (o *DestinationStatus) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DestinationStatus) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DestinationStatus) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DestinationStatus) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConnected

`func (o *DestinationStatus) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *DestinationStatus) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *DestinationStatus) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *DestinationStatus) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetEnabled

`func (o *DestinationStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DestinationStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DestinationStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DestinationStatus) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFields

`func (o *DestinationStatus) GetFields() []DestinationField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DestinationStatus) GetFieldsOk() (*[]DestinationField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DestinationStatus) SetFields(v []DestinationField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DestinationStatus) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetLive

`func (o *DestinationStatus) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *DestinationStatus) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *DestinationStatus) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *DestinationStatus) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetName

`func (o *DestinationStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DestinationStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPixel

`func (o *DestinationStatus) GetPixel() bool`

GetPixel returns the Pixel field if non-nil, zero value otherwise.

### GetPixelOk

`func (o *DestinationStatus) GetPixelOk() (*bool, bool)`

GetPixelOk returns a tuple with the Pixel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixel

`func (o *DestinationStatus) SetPixel(v bool)`

SetPixel sets Pixel field to given value.

### HasPixel

`func (o *DestinationStatus) HasPixel() bool`

HasPixel returns a boolean if a field has been set.

### GetPlatform

`func (o *DestinationStatus) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DestinationStatus) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DestinationStatus) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DestinationStatus) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSecrets

`func (o *DestinationStatus) GetSecrets() []string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *DestinationStatus) GetSecretsOk() (*[]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *DestinationStatus) SetSecrets(v []string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *DestinationStatus) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


