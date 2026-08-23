# CdResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]CDApp**](CDApp.md) |  | [optional] 

## Methods

### NewCdResp

`func NewCdResp() *CdResp`

NewCdResp instantiates a new CdResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdRespWithDefaults

`func NewCdRespWithDefaults() *CdResp`

NewCdRespWithDefaults instantiates a new CdResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *CdResp) GetApplications() []CDApp`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *CdResp) GetApplicationsOk() (*[]CDApp, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *CdResp) SetApplications(v []CDApp)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *CdResp) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


