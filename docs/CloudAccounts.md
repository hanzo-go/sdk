# CloudAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**CloudSourceState**](CloudSourceState.md) | Account is the state of the caller&#39;s own linked-account side. | [optional] 
**Hanzo** | Pointer to [**CloudSourceState**](CloudSourceState.md) | Hanzo is the state of the org&#39;s Hanzo-routed side. | [optional] 
**Rows** | Pointer to [**[]CloudTotalView**](CloudTotalView.md) | Rows is the two row sets CONCATENATED, never summed — each row says which side it came from. A percent is not money and a provider&#39;s own spend is not a Hanzo charge, so adding them would produce a number that means nothing. | [optional] 

## Methods

### NewCloudAccounts

`func NewCloudAccounts() *CloudAccounts`

NewCloudAccounts instantiates a new CloudAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccountsWithDefaults

`func NewCloudAccountsWithDefaults() *CloudAccounts`

NewCloudAccountsWithDefaults instantiates a new CloudAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudAccounts) GetAccount() CloudSourceState`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudAccounts) GetAccountOk() (*CloudSourceState, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudAccounts) SetAccount(v CloudSourceState)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudAccounts) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetHanzo

`func (o *CloudAccounts) GetHanzo() CloudSourceState`

GetHanzo returns the Hanzo field if non-nil, zero value otherwise.

### GetHanzoOk

`func (o *CloudAccounts) GetHanzoOk() (*CloudSourceState, bool)`

GetHanzoOk returns a tuple with the Hanzo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanzo

`func (o *CloudAccounts) SetHanzo(v CloudSourceState)`

SetHanzo sets Hanzo field to given value.

### HasHanzo

`func (o *CloudAccounts) HasHanzo() bool`

HasHanzo returns a boolean if a field has been set.

### GetRows

`func (o *CloudAccounts) GetRows() []CloudTotalView`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *CloudAccounts) GetRowsOk() (*[]CloudTotalView, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *CloudAccounts) SetRows(v []CloudTotalView)`

SetRows sets Rows field to given value.

### HasRows

`func (o *CloudAccounts) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


