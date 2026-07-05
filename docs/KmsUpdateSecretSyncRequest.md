# KmsUpdateSecretSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**IsAutoSyncEnabled** | Pointer to **bool** |  | [optional] 
**DestinationConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewKmsUpdateSecretSyncRequest

`func NewKmsUpdateSecretSyncRequest() *KmsUpdateSecretSyncRequest`

NewKmsUpdateSecretSyncRequest instantiates a new KmsUpdateSecretSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsUpdateSecretSyncRequestWithDefaults

`func NewKmsUpdateSecretSyncRequestWithDefaults() *KmsUpdateSecretSyncRequest`

NewKmsUpdateSecretSyncRequestWithDefaults instantiates a new KmsUpdateSecretSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KmsUpdateSecretSyncRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsUpdateSecretSyncRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsUpdateSecretSyncRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmsUpdateSecretSyncRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsAutoSyncEnabled

`func (o *KmsUpdateSecretSyncRequest) GetIsAutoSyncEnabled() bool`

GetIsAutoSyncEnabled returns the IsAutoSyncEnabled field if non-nil, zero value otherwise.

### GetIsAutoSyncEnabledOk

`func (o *KmsUpdateSecretSyncRequest) GetIsAutoSyncEnabledOk() (*bool, bool)`

GetIsAutoSyncEnabledOk returns a tuple with the IsAutoSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoSyncEnabled

`func (o *KmsUpdateSecretSyncRequest) SetIsAutoSyncEnabled(v bool)`

SetIsAutoSyncEnabled sets IsAutoSyncEnabled field to given value.

### HasIsAutoSyncEnabled

`func (o *KmsUpdateSecretSyncRequest) HasIsAutoSyncEnabled() bool`

HasIsAutoSyncEnabled returns a boolean if a field has been set.

### GetDestinationConfig

`func (o *KmsUpdateSecretSyncRequest) GetDestinationConfig() map[string]interface{}`

GetDestinationConfig returns the DestinationConfig field if non-nil, zero value otherwise.

### GetDestinationConfigOk

`func (o *KmsUpdateSecretSyncRequest) GetDestinationConfigOk() (*map[string]interface{}, bool)`

GetDestinationConfigOk returns a tuple with the DestinationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationConfig

`func (o *KmsUpdateSecretSyncRequest) SetDestinationConfig(v map[string]interface{})`

SetDestinationConfig sets DestinationConfig field to given value.

### HasDestinationConfig

`func (o *KmsUpdateSecretSyncRequest) HasDestinationConfig() bool`

HasDestinationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


