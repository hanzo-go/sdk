# CloudMutateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Add** | Pointer to **[]string** | Add is the product ids to turn ON. Each must already be an ACTIVE entitlement of the org&#39;s plan, unless the caller is a platform super admin. | [optional] 
**Remove** | Pointer to **[]string** | Remove is the product ids to turn OFF. Disabling is never gated. | [optional] 

## Methods

### NewCloudMutateReq

`func NewCloudMutateReq() *CloudMutateReq`

NewCloudMutateReq instantiates a new CloudMutateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMutateReqWithDefaults

`func NewCloudMutateReqWithDefaults() *CloudMutateReq`

NewCloudMutateReqWithDefaults instantiates a new CloudMutateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdd

`func (o *CloudMutateReq) GetAdd() []string`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *CloudMutateReq) GetAddOk() (*[]string, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *CloudMutateReq) SetAdd(v []string)`

SetAdd sets Add field to given value.

### HasAdd

`func (o *CloudMutateReq) HasAdd() bool`

HasAdd returns a boolean if a field has been set.

### GetRemove

`func (o *CloudMutateReq) GetRemove() []string`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *CloudMutateReq) GetRemoveOk() (*[]string, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *CloudMutateReq) SetRemove(v []string)`

SetRemove sets Remove field to given value.

### HasRemove

`func (o *CloudMutateReq) HasRemove() bool`

HasRemove returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


