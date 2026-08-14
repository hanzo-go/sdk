# PayoutData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**AdminAuthorView**](AdminAuthorView.md) | Author is the author record after the payout, with the balances updated. | [optional] 
**Payout** | Pointer to [**PayoutView**](PayoutView.md) | Payout is the recorded payout, including where it settled. | [optional] 

## Methods

### NewPayoutData

`func NewPayoutData() *PayoutData`

NewPayoutData instantiates a new PayoutData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutDataWithDefaults

`func NewPayoutDataWithDefaults() *PayoutData`

NewPayoutDataWithDefaults instantiates a new PayoutData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *PayoutData) GetAuthor() AdminAuthorView`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PayoutData) GetAuthorOk() (*AdminAuthorView, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PayoutData) SetAuthor(v AdminAuthorView)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *PayoutData) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetPayout

`func (o *PayoutData) GetPayout() PayoutView`

GetPayout returns the Payout field if non-nil, zero value otherwise.

### GetPayoutOk

`func (o *PayoutData) GetPayoutOk() (*PayoutView, bool)`

GetPayoutOk returns a tuple with the Payout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayout

`func (o *PayoutData) SetPayout(v PayoutView)`

SetPayout sets Payout field to given value.

### HasPayout

`func (o *PayoutData) HasPayout() bool`

HasPayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


