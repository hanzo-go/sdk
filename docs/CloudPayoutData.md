# CloudPayoutData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**CloudAdminAuthorView**](CloudAdminAuthorView.md) | Author is the author record after the payout, with the balances updated. | [optional] 
**Payout** | Pointer to [**CloudPayoutView**](CloudPayoutView.md) | Payout is the recorded payout, including where it settled. | [optional] 

## Methods

### NewCloudPayoutData

`func NewCloudPayoutData() *CloudPayoutData`

NewCloudPayoutData instantiates a new CloudPayoutData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPayoutDataWithDefaults

`func NewCloudPayoutDataWithDefaults() *CloudPayoutData`

NewCloudPayoutDataWithDefaults instantiates a new CloudPayoutData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *CloudPayoutData) GetAuthor() CloudAdminAuthorView`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CloudPayoutData) GetAuthorOk() (*CloudAdminAuthorView, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CloudPayoutData) SetAuthor(v CloudAdminAuthorView)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *CloudPayoutData) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetPayout

`func (o *CloudPayoutData) GetPayout() CloudPayoutView`

GetPayout returns the Payout field if non-nil, zero value otherwise.

### GetPayoutOk

`func (o *CloudPayoutData) GetPayoutOk() (*CloudPayoutView, bool)`

GetPayoutOk returns a tuple with the Payout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayout

`func (o *CloudPayoutData) SetPayout(v CloudPayoutView)`

SetPayout sets Payout field to given value.

### HasPayout

`func (o *CloudPayoutData) HasPayout() bool`

HasPayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


