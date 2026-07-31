# CloudConnectionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account names the connected external account, when the provider reports one. | [optional] 
**Provider** | Pointer to **string** | Provider is the connector this answer is about. | [optional] 
**Status** | Pointer to **string** | Status is the connection state: connected or disconnected. | [optional] 

## Methods

### NewCloudConnectionOut

`func NewCloudConnectionOut() *CloudConnectionOut`

NewCloudConnectionOut instantiates a new CloudConnectionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudConnectionOutWithDefaults

`func NewCloudConnectionOutWithDefaults() *CloudConnectionOut`

NewCloudConnectionOutWithDefaults instantiates a new CloudConnectionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudConnectionOut) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudConnectionOut) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudConnectionOut) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudConnectionOut) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetProvider

`func (o *CloudConnectionOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudConnectionOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudConnectionOut) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudConnectionOut) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *CloudConnectionOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudConnectionOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudConnectionOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudConnectionOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


