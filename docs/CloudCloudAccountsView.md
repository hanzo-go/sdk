# CloudCloudAccountsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]CloudCloudAccountView**](CloudCloudAccountView.md) | Accounts is every account this org has linked, across all providers. Empty when it has linked none. | [optional] 

## Methods

### NewCloudCloudAccountsView

`func NewCloudCloudAccountsView() *CloudCloudAccountsView`

NewCloudCloudAccountsView instantiates a new CloudCloudAccountsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCloudAccountsViewWithDefaults

`func NewCloudCloudAccountsViewWithDefaults() *CloudCloudAccountsView`

NewCloudCloudAccountsViewWithDefaults instantiates a new CloudCloudAccountsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *CloudCloudAccountsView) GetAccounts() []CloudCloudAccountView`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *CloudCloudAccountsView) GetAccountsOk() (*[]CloudCloudAccountView, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *CloudCloudAccountsView) SetAccounts(v []CloudCloudAccountView)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *CloudCloudAccountsView) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


