# CloudActivationReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activate** | Pointer to **[]string** | Activate switches these tool names on for the caller&#39;s org and project. | [optional] 
**Deactivate** | Pointer to **[]string** | Deactivate switches these tool names off. | [optional] 

## Methods

### NewCloudActivationReq

`func NewCloudActivationReq() *CloudActivationReq`

NewCloudActivationReq instantiates a new CloudActivationReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudActivationReqWithDefaults

`func NewCloudActivationReqWithDefaults() *CloudActivationReq`

NewCloudActivationReqWithDefaults instantiates a new CloudActivationReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivate

`func (o *CloudActivationReq) GetActivate() []string`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *CloudActivationReq) GetActivateOk() (*[]string, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *CloudActivationReq) SetActivate(v []string)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *CloudActivationReq) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### GetDeactivate

`func (o *CloudActivationReq) GetDeactivate() []string`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *CloudActivationReq) GetDeactivateOk() (*[]string, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *CloudActivationReq) SetDeactivate(v []string)`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *CloudActivationReq) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


