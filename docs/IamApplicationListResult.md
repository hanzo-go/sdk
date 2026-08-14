# IamApplicationListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]IamApplication**](IamApplication.md) |  | [optional] 

## Methods

### NewIamApplicationListResult

`func NewIamApplicationListResult() *IamApplicationListResult`

NewIamApplicationListResult instantiates a new IamApplicationListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamApplicationListResultWithDefaults

`func NewIamApplicationListResultWithDefaults() *IamApplicationListResult`

NewIamApplicationListResultWithDefaults instantiates a new IamApplicationListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *IamApplicationListResult) GetApplications() []IamApplication`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *IamApplicationListResult) GetApplicationsOk() (*[]IamApplication, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *IamApplicationListResult) SetApplications(v []IamApplication)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *IamApplicationListResult) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


