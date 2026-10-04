# DomainQuoteList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]DomainOffer**](DomainOffer.md) | Results is one quote per name, priced RETAIL — this deployment&#39;s markup is already applied and the wholesale cost is never on the wire. | [optional] 

## Methods

### NewDomainQuoteList

`func NewDomainQuoteList() *DomainQuoteList`

NewDomainQuoteList instantiates a new DomainQuoteList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainQuoteListWithDefaults

`func NewDomainQuoteListWithDefaults() *DomainQuoteList`

NewDomainQuoteListWithDefaults instantiates a new DomainQuoteList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *DomainQuoteList) GetResults() []DomainOffer`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DomainQuoteList) GetResultsOk() (*[]DomainOffer, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DomainQuoteList) SetResults(v []DomainOffer)`

SetResults sets Results field to given value.

### HasResults

`func (o *DomainQuoteList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


