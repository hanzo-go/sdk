# BooksQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the chart number the questioned entry posted to, where one applies. | [optional] 
**Amount** | Pointer to **string** | Amount is the figure that makes the question concrete, already FORMATTED for display with its currency symbol — a string, not cents, and not for arithmetic. | [optional] 
**Id** | Pointer to **string** | ID is the source transaction the question is about, so answering it leads straight back to the entry that raised it. | [optional] 
**Kind** | Pointer to **string** | Kind is what looked wrong: outlier (a charge far above the usual), reversal (a posting undone), roundoff (a balancing plug big enough to be worth explaining), uncosted (revenue booked with no cost matched to it), or overdrawn (a wallet spent past its balance). | [optional] 
**PostedAt** | Pointer to **string** | PostedAt anchors the question in time — when the entry it concerns posted. | [optional] 
**Text** | Pointer to **string** | Text is the question itself, written for a founder to answer directly. | [optional] 

## Methods

### NewBooksQuestion

`func NewBooksQuestion() *BooksQuestion`

NewBooksQuestion instantiates a new BooksQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksQuestionWithDefaults

`func NewBooksQuestionWithDefaults() *BooksQuestion`

NewBooksQuestionWithDefaults instantiates a new BooksQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *BooksQuestion) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BooksQuestion) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BooksQuestion) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *BooksQuestion) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAmount

`func (o *BooksQuestion) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BooksQuestion) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BooksQuestion) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BooksQuestion) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetId

`func (o *BooksQuestion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BooksQuestion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BooksQuestion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BooksQuestion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *BooksQuestion) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BooksQuestion) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BooksQuestion) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *BooksQuestion) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPostedAt

`func (o *BooksQuestion) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *BooksQuestion) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *BooksQuestion) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *BooksQuestion) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### GetText

`func (o *BooksQuestion) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *BooksQuestion) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *BooksQuestion) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *BooksQuestion) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


