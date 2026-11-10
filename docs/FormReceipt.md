# FormReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **bool** | Completed is whether the response is now submitted whole. | [optional] 
**Id** | Pointer to **string** | ID is the response&#39;s id. | [optional] 
**SubmittedAt** | Pointer to **string** | SubmittedAt is when it was completed, RFC 3339. | [optional] 
**Thanks** | Pointer to [**FormScreen**](FormScreen.md) | Thanks is the form&#39;s thank-you screen, for a completed response. | [optional] 

## Methods

### NewFormReceipt

`func NewFormReceipt() *FormReceipt`

NewFormReceipt instantiates a new FormReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormReceiptWithDefaults

`func NewFormReceiptWithDefaults() *FormReceipt`

NewFormReceiptWithDefaults instantiates a new FormReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *FormReceipt) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *FormReceipt) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *FormReceipt) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *FormReceipt) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetId

`func (o *FormReceipt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormReceipt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormReceipt) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormReceipt) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *FormReceipt) GetSubmittedAt() string`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *FormReceipt) GetSubmittedAtOk() (*string, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *FormReceipt) SetSubmittedAt(v string)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *FormReceipt) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetThanks

`func (o *FormReceipt) GetThanks() FormScreen`

GetThanks returns the Thanks field if non-nil, zero value otherwise.

### GetThanksOk

`func (o *FormReceipt) GetThanksOk() (*FormScreen, bool)`

GetThanksOk returns a tuple with the Thanks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThanks

`func (o *FormReceipt) SetThanks(v FormScreen)`

SetThanks sets Thanks field to given value.

### HasThanks

`func (o *FormReceipt) HasThanks() bool`

HasThanks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


