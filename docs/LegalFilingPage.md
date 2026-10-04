# LegalFilingPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]LegalLegalFiling**](LegalLegalFiling.md) | Data are the filing records. | [optional] 
**Disclaimer** | Pointer to **string** | Disclaimer is the boundary made visible on the wire. | [optional] 

## Methods

### NewLegalFilingPage

`func NewLegalFilingPage() *LegalFilingPage`

NewLegalFilingPage instantiates a new LegalFilingPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalFilingPageWithDefaults

`func NewLegalFilingPageWithDefaults() *LegalFilingPage`

NewLegalFilingPageWithDefaults instantiates a new LegalFilingPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LegalFilingPage) GetData() []LegalLegalFiling`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LegalFilingPage) GetDataOk() (*[]LegalLegalFiling, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LegalFilingPage) SetData(v []LegalLegalFiling)`

SetData sets Data field to given value.

### HasData

`func (o *LegalFilingPage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisclaimer

`func (o *LegalFilingPage) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *LegalFilingPage) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *LegalFilingPage) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *LegalFilingPage) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


