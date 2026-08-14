# O11yO11yAccountsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yGettableAccounts**](O11yGettableAccounts.md) | Data holds the connected accounts. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yAccountsOut

`func NewO11yO11yAccountsOut() *O11yO11yAccountsOut`

NewO11yO11yAccountsOut instantiates a new O11yO11yAccountsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yAccountsOutWithDefaults

`func NewO11yO11yAccountsOutWithDefaults() *O11yO11yAccountsOut`

NewO11yO11yAccountsOutWithDefaults instantiates a new O11yO11yAccountsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yAccountsOut) GetData() O11yGettableAccounts`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yAccountsOut) GetDataOk() (*O11yGettableAccounts, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yAccountsOut) SetData(v O11yGettableAccounts)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yAccountsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yAccountsOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yAccountsOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yAccountsOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yAccountsOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


