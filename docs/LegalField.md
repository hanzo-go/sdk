# LegalField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the identifier the body substitutes ({{.key}}) and the key a generation&#39;s data map must carry. snake_case by convention across the built-ins — effective_date, company_name, governing_law. An override whose body references a key no Field declares is refused on save. | [optional] 
**Label** | Pointer to **string** | Label is the human prompt for whoever fills the value in — \&quot;Governing law (state)\&quot;. It never reaches the rendered document; only Key does. | [optional] 

## Methods

### NewLegalField

`func NewLegalField() *LegalField`

NewLegalField instantiates a new LegalField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalFieldWithDefaults

`func NewLegalFieldWithDefaults() *LegalField`

NewLegalFieldWithDefaults instantiates a new LegalField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *LegalField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LegalField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LegalField) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *LegalField) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *LegalField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LegalField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LegalField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LegalField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


