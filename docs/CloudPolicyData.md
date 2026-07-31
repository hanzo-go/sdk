# CloudPolicyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**CloudSharePolicy**](CloudSharePolicy.md) | Policy is the revenue-share configuration as stored. | [optional] 

## Methods

### NewCloudPolicyData

`func NewCloudPolicyData() *CloudPolicyData`

NewCloudPolicyData instantiates a new CloudPolicyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPolicyDataWithDefaults

`func NewCloudPolicyDataWithDefaults() *CloudPolicyData`

NewCloudPolicyDataWithDefaults instantiates a new CloudPolicyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *CloudPolicyData) GetPolicy() CloudSharePolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CloudPolicyData) GetPolicyOk() (*CloudSharePolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CloudPolicyData) SetPolicy(v CloudSharePolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CloudPolicyData) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


