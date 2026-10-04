# LegalLegalHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; when the subsystem is serving. | [optional] 
**Templates** | Pointer to **int64** | Templates is how many built-in templates the catalog carries. | [optional] 

## Methods

### NewLegalLegalHealth

`func NewLegalLegalHealth() *LegalLegalHealth`

NewLegalLegalHealth instantiates a new LegalLegalHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalLegalHealthWithDefaults

`func NewLegalLegalHealthWithDefaults() *LegalLegalHealth`

NewLegalLegalHealthWithDefaults instantiates a new LegalLegalHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *LegalLegalHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LegalLegalHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LegalLegalHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LegalLegalHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplates

`func (o *LegalLegalHealth) GetTemplates() int64`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *LegalLegalHealth) GetTemplatesOk() (*int64, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *LegalLegalHealth) SetTemplates(v int64)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *LegalLegalHealth) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


