# BooksAskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | From is the RFC3339 start of the metric window. Empty means all time, treated as a single reporting period (see monthsBetween). | [optional] 
**Question** | Pointer to **string** | Question is the plain-language question about the org&#39;s books, e.g. \&quot;what is my MRR?\&quot;. Longer than 2000 characters is truncated, never refused. | [optional] 
**To** | Pointer to **string** | To is the RFC3339 end of the metric window. Empty means up to now. | [optional] 

## Methods

### NewBooksAskRequest

`func NewBooksAskRequest() *BooksAskRequest`

NewBooksAskRequest instantiates a new BooksAskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksAskRequestWithDefaults

`func NewBooksAskRequestWithDefaults() *BooksAskRequest`

NewBooksAskRequestWithDefaults instantiates a new BooksAskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *BooksAskRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BooksAskRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BooksAskRequest) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BooksAskRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetQuestion

`func (o *BooksAskRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *BooksAskRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *BooksAskRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *BooksAskRequest) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetTo

`func (o *BooksAskRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BooksAskRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BooksAskRequest) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *BooksAskRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


