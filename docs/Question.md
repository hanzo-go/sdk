# Question

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** | formatted figure ($…) | [optional] 
**Id** | Pointer to **string** | the source transaction id it concerns | [optional] 
**Kind** | Pointer to **string** | outlier|reversal|roundoff|uncosted|overdrawn | [optional] 
**PostedAt** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** | the specific question to ask the founder | [optional] 

## Methods

### NewQuestion

`func NewQuestion() *Question`

NewQuestion instantiates a new Question object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionWithDefaults

`func NewQuestionWithDefaults() *Question`

NewQuestionWithDefaults instantiates a new Question object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Question) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Question) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Question) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Question) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAmount

`func (o *Question) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Question) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Question) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Question) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetId

`func (o *Question) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Question) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Question) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Question) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Question) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Question) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Question) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Question) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPostedAt

`func (o *Question) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *Question) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *Question) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *Question) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### GetText

`func (o *Question) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Question) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Question) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Question) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


