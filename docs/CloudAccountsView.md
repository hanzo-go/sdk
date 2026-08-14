# CloudAccountsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]CloudAccountView**](CloudAccountView.md) | Accounts is every account this org has linked, across all providers. Empty when it has linked none. | [optional] 

## Methods

### NewCloudAccountsView

`func NewCloudAccountsView() *CloudAccountsView`

NewCloudAccountsView instantiates a new CloudAccountsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccountsViewWithDefaults

`func NewCloudAccountsViewWithDefaults() *CloudAccountsView`

NewCloudAccountsViewWithDefaults instantiates a new CloudAccountsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *CloudAccountsView) GetAccounts() []CloudAccountView`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *CloudAccountsView) GetAccountsOk() (*[]CloudAccountView, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *CloudAccountsView) SetAccounts(v []CloudAccountView)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *CloudAccountsView) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


