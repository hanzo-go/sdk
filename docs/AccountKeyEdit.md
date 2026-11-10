# AccountKeyEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to [**AccountKeyBudget**](AccountKeyBudget.md) | Budget replaces the key&#39;s spend caps, in cents; a zero drops that cap. | [optional] 
**Expires** | Pointer to **string** | Expires replaces when the key stops working, RFC 3339. An empty string means never. | [optional] 
**Id** | Pointer to **string** | ID is the key to edit, from the path. | [optional] 
**Limit** | Pointer to **[]string** | Limit replaces what the key may reach. An empty list makes it unrestricted. | [optional] 
**Name** | Pointer to **string** | Name relabels the key. | [optional] 
**Rate** | Pointer to **int64** | Rate replaces the key&#39;s requests a minute; zero drops the limit. | [optional] 

## Methods

### NewAccountKeyEdit

`func NewAccountKeyEdit() *AccountKeyEdit`

NewAccountKeyEdit instantiates a new AccountKeyEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountKeyEditWithDefaults

`func NewAccountKeyEditWithDefaults() *AccountKeyEdit`

NewAccountKeyEditWithDefaults instantiates a new AccountKeyEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *AccountKeyEdit) GetBudget() AccountKeyBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *AccountKeyEdit) GetBudgetOk() (*AccountKeyBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *AccountKeyEdit) SetBudget(v AccountKeyBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *AccountKeyEdit) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetExpires

`func (o *AccountKeyEdit) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AccountKeyEdit) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AccountKeyEdit) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *AccountKeyEdit) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetId

`func (o *AccountKeyEdit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountKeyEdit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountKeyEdit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountKeyEdit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLimit

`func (o *AccountKeyEdit) GetLimit() []string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AccountKeyEdit) GetLimitOk() (*[]string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AccountKeyEdit) SetLimit(v []string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AccountKeyEdit) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *AccountKeyEdit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountKeyEdit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountKeyEdit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountKeyEdit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRate

`func (o *AccountKeyEdit) GetRate() int64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *AccountKeyEdit) GetRateOk() (*int64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *AccountKeyEdit) SetRate(v int64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *AccountKeyEdit) HasRate() bool`

HasRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


