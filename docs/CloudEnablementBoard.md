# CloudEnablementBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Betas** | Pointer to [**[]CloudUserEnablementItem**](CloudUserEnablementItem.md) | Betas are the subset of Items the caller&#39;s org may still opt into. | [optional] 
**Items** | Pointer to [**[]CloudUserEnablementItem**](CloudUserEnablementItem.md) | Items is every managed item, each resolved for the caller&#39;s org. | [optional] 
**Org** | Pointer to **string** | Org is the org this view was resolved for; empty for a caller with no validated principal, who sees only the generally-available items. | [optional] 

## Methods

### NewCloudEnablementBoard

`func NewCloudEnablementBoard() *CloudEnablementBoard`

NewCloudEnablementBoard instantiates a new CloudEnablementBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEnablementBoardWithDefaults

`func NewCloudEnablementBoardWithDefaults() *CloudEnablementBoard`

NewCloudEnablementBoardWithDefaults instantiates a new CloudEnablementBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBetas

`func (o *CloudEnablementBoard) GetBetas() []CloudUserEnablementItem`

GetBetas returns the Betas field if non-nil, zero value otherwise.

### GetBetasOk

`func (o *CloudEnablementBoard) GetBetasOk() (*[]CloudUserEnablementItem, bool)`

GetBetasOk returns a tuple with the Betas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetas

`func (o *CloudEnablementBoard) SetBetas(v []CloudUserEnablementItem)`

SetBetas sets Betas field to given value.

### HasBetas

`func (o *CloudEnablementBoard) HasBetas() bool`

HasBetas returns a boolean if a field has been set.

### GetItems

`func (o *CloudEnablementBoard) GetItems() []CloudUserEnablementItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CloudEnablementBoard) GetItemsOk() (*[]CloudUserEnablementItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CloudEnablementBoard) SetItems(v []CloudUserEnablementItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *CloudEnablementBoard) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOrg

`func (o *CloudEnablementBoard) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudEnablementBoard) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudEnablementBoard) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudEnablementBoard) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


