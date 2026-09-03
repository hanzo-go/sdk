# LegalHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; when the subsystem is serving. | [optional] 
**Templates** | Pointer to **int64** | Templates is how many built-in templates the catalog carries. | [optional] 

## Methods

### NewLegalHealth

`func NewLegalHealth() *LegalHealth`

NewLegalHealth instantiates a new LegalHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalHealthWithDefaults

`func NewLegalHealthWithDefaults() *LegalHealth`

NewLegalHealthWithDefaults instantiates a new LegalHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *LegalHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LegalHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LegalHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LegalHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplates

`func (o *LegalHealth) GetTemplates() int64`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *LegalHealth) GetTemplatesOk() (*int64, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *LegalHealth) SetTemplates(v int64)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *LegalHealth) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


