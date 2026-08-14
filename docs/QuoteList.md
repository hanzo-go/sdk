# QuoteList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Offer**](Offer.md) | Results is one quote per name, priced RETAIL — this deployment&#39;s markup is already applied and the wholesale cost is never on the wire. | [optional] 

## Methods

### NewQuoteList

`func NewQuoteList() *QuoteList`

NewQuoteList instantiates a new QuoteList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteListWithDefaults

`func NewQuoteListWithDefaults() *QuoteList`

NewQuoteListWithDefaults instantiates a new QuoteList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *QuoteList) GetResults() []Offer`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *QuoteList) GetResultsOk() (*[]Offer, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *QuoteList) SetResults(v []Offer)`

SetResults sets Results field to given value.

### HasResults

`func (o *QuoteList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


