# O11yO11yRoleCreateOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yO11yCreated**](O11yO11yCreated.md) | Data carries the new role&#39;s id. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yRoleCreateOut

`func NewO11yO11yRoleCreateOut() *O11yO11yRoleCreateOut`

NewO11yO11yRoleCreateOut instantiates a new O11yO11yRoleCreateOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yRoleCreateOutWithDefaults

`func NewO11yO11yRoleCreateOutWithDefaults() *O11yO11yRoleCreateOut`

NewO11yO11yRoleCreateOutWithDefaults instantiates a new O11yO11yRoleCreateOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yRoleCreateOut) GetData() O11yO11yCreated`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yRoleCreateOut) GetDataOk() (*O11yO11yCreated, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yRoleCreateOut) SetData(v O11yO11yCreated)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yRoleCreateOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yRoleCreateOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yRoleCreateOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yRoleCreateOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yRoleCreateOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


