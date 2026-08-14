# PolicyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**SharePolicy**](SharePolicy.md) | Policy is the revenue-share configuration as stored. | [optional] 

## Methods

### NewPolicyData

`func NewPolicyData() *PolicyData`

NewPolicyData instantiates a new PolicyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDataWithDefaults

`func NewPolicyDataWithDefaults() *PolicyData`

NewPolicyDataWithDefaults instantiates a new PolicyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *PolicyData) GetPolicy() SharePolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PolicyData) GetPolicyOk() (*SharePolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PolicyData) SetPolicy(v SharePolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PolicyData) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


