# LegalTemplateCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]LegalTemplateView**](LegalTemplateView.md) | Data is the catalog, metadata and merge fields only — never the template bodies, which are fetched one at a time. | [optional] 
**Disclaimer** | Pointer to **string** | Disclaimer is the boundary made visible on the wire: Hanzo Legal is document tooling, not legal advice. | [optional] 

## Methods

### NewLegalTemplateCatalog

`func NewLegalTemplateCatalog() *LegalTemplateCatalog`

NewLegalTemplateCatalog instantiates a new LegalTemplateCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalTemplateCatalogWithDefaults

`func NewLegalTemplateCatalogWithDefaults() *LegalTemplateCatalog`

NewLegalTemplateCatalogWithDefaults instantiates a new LegalTemplateCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LegalTemplateCatalog) GetData() []LegalTemplateView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LegalTemplateCatalog) GetDataOk() (*[]LegalTemplateView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LegalTemplateCatalog) SetData(v []LegalTemplateView)`

SetData sets Data field to given value.

### HasData

`func (o *LegalTemplateCatalog) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisclaimer

`func (o *LegalTemplateCatalog) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *LegalTemplateCatalog) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *LegalTemplateCatalog) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *LegalTemplateCatalog) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


