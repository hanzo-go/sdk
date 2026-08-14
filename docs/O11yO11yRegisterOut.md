# O11yO11yRegisterOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yO11yUser**](O11yO11yUser.md) | Data is the user. The runtime answers register with the same user shape the identity face reads, so it is the ONE O11yUser (identity.go) — a created user is a user, and the document names it once. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yRegisterOut

`func NewO11yO11yRegisterOut() *O11yO11yRegisterOut`

NewO11yO11yRegisterOut instantiates a new O11yO11yRegisterOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yRegisterOutWithDefaults

`func NewO11yO11yRegisterOutWithDefaults() *O11yO11yRegisterOut`

NewO11yO11yRegisterOutWithDefaults instantiates a new O11yO11yRegisterOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yRegisterOut) GetData() O11yO11yUser`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yRegisterOut) GetDataOk() (*O11yO11yUser, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yRegisterOut) SetData(v O11yO11yUser)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yRegisterOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yRegisterOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yRegisterOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yRegisterOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yRegisterOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


