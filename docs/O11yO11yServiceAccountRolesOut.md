# O11yO11yServiceAccountRolesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]O11yO11yRole**](O11yO11yRole.md) | Data holds the roles. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yServiceAccountRolesOut

`func NewO11yO11yServiceAccountRolesOut() *O11yO11yServiceAccountRolesOut`

NewO11yO11yServiceAccountRolesOut instantiates a new O11yO11yServiceAccountRolesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yServiceAccountRolesOutWithDefaults

`func NewO11yO11yServiceAccountRolesOutWithDefaults() *O11yO11yServiceAccountRolesOut`

NewO11yO11yServiceAccountRolesOutWithDefaults instantiates a new O11yO11yServiceAccountRolesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yServiceAccountRolesOut) GetData() []O11yO11yRole`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yServiceAccountRolesOut) GetDataOk() (*[]O11yO11yRole, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yServiceAccountRolesOut) SetData(v []O11yO11yRole)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yServiceAccountRolesOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yServiceAccountRolesOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yServiceAccountRolesOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yServiceAccountRolesOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yServiceAccountRolesOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


