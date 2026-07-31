# CloudServiceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Hosts** | Pointer to **[]string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**WaitlistMode** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudServiceInput

`func NewCloudServiceInput() *CloudServiceInput`

NewCloudServiceInput instantiates a new CloudServiceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudServiceInputWithDefaults

`func NewCloudServiceInputWithDefaults() *CloudServiceInput`

NewCloudServiceInputWithDefaults instantiates a new CloudServiceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CloudServiceInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudServiceInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudServiceInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudServiceInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudServiceInput) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudServiceInput) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudServiceInput) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudServiceInput) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHosts

`func (o *CloudServiceInput) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *CloudServiceInput) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *CloudServiceInput) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *CloudServiceInput) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetService

`func (o *CloudServiceInput) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudServiceInput) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudServiceInput) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CloudServiceInput) HasService() bool`

HasService returns a boolean if a field has been set.

### GetWaitlistMode

`func (o *CloudServiceInput) GetWaitlistMode() bool`

GetWaitlistMode returns the WaitlistMode field if non-nil, zero value otherwise.

### GetWaitlistModeOk

`func (o *CloudServiceInput) GetWaitlistModeOk() (*bool, bool)`

GetWaitlistModeOk returns a tuple with the WaitlistMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistMode

`func (o *CloudServiceInput) SetWaitlistMode(v bool)`

SetWaitlistMode sets WaitlistMode field to given value.

### HasWaitlistMode

`func (o *CloudServiceInput) HasWaitlistMode() bool`

HasWaitlistMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


