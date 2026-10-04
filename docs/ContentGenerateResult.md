# ContentGenerateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctype** | Pointer to **string** | the marketing type the draft was filed as | [optional] 
**Name** | Pointer to **string** | the new document&#39;s name — its address for every later call | [optional] 
**Status** | Pointer to **string** | always \&quot;draft\&quot;; the lifecycle owns the initial state | [optional] 

## Methods

### NewContentGenerateResult

`func NewContentGenerateResult() *ContentGenerateResult`

NewContentGenerateResult instantiates a new ContentGenerateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentGenerateResultWithDefaults

`func NewContentGenerateResultWithDefaults() *ContentGenerateResult`

NewContentGenerateResultWithDefaults instantiates a new ContentGenerateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctype

`func (o *ContentGenerateResult) GetDoctype() string`

GetDoctype returns the Doctype field if non-nil, zero value otherwise.

### GetDoctypeOk

`func (o *ContentGenerateResult) GetDoctypeOk() (*string, bool)`

GetDoctypeOk returns a tuple with the Doctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctype

`func (o *ContentGenerateResult) SetDoctype(v string)`

SetDoctype sets Doctype field to given value.

### HasDoctype

`func (o *ContentGenerateResult) HasDoctype() bool`

HasDoctype returns a boolean if a field has been set.

### GetName

`func (o *ContentGenerateResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentGenerateResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentGenerateResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentGenerateResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ContentGenerateResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContentGenerateResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContentGenerateResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContentGenerateResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


