# LegalDocumentPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]LegalDocumentSummary**](LegalDocumentSummary.md) | Data are the documents, WITHOUT their rendered content — fetch one to read it. | [optional] 
**Disclaimer** | Pointer to **string** | Disclaimer is the boundary made visible on the wire. | [optional] 

## Methods

### NewLegalDocumentPage

`func NewLegalDocumentPage() *LegalDocumentPage`

NewLegalDocumentPage instantiates a new LegalDocumentPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalDocumentPageWithDefaults

`func NewLegalDocumentPageWithDefaults() *LegalDocumentPage`

NewLegalDocumentPageWithDefaults instantiates a new LegalDocumentPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LegalDocumentPage) GetData() []LegalDocumentSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LegalDocumentPage) GetDataOk() (*[]LegalDocumentSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LegalDocumentPage) SetData(v []LegalDocumentSummary)`

SetData sets Data field to given value.

### HasData

`func (o *LegalDocumentPage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisclaimer

`func (o *LegalDocumentPage) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *LegalDocumentPage) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *LegalDocumentPage) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *LegalDocumentPage) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


