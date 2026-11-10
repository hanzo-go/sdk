# AccountKeyBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | Pointer to **int64** | Month is per calendar month, UTC. | [optional] 
**Total** | Pointer to **int64** | Total is over the key&#39;s whole life. | [optional] 

## Methods

### NewAccountKeyBudget

`func NewAccountKeyBudget() *AccountKeyBudget`

NewAccountKeyBudget instantiates a new AccountKeyBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountKeyBudgetWithDefaults

`func NewAccountKeyBudgetWithDefaults() *AccountKeyBudget`

NewAccountKeyBudgetWithDefaults instantiates a new AccountKeyBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *AccountKeyBudget) GetMonth() int64`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *AccountKeyBudget) GetMonthOk() (*int64, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *AccountKeyBudget) SetMonth(v int64)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *AccountKeyBudget) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetTotal

`func (o *AccountKeyBudget) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccountKeyBudget) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccountKeyBudget) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AccountKeyBudget) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


