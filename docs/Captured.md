# Captured

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answered** | Pointer to **string** | Answered is the label the caller was answered: the org&#39;s capability&#39;s when one answered, else Kai&#39;s own. | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCaptured

`func NewCaptured() *Captured`

NewCaptured instantiates a new Captured object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapturedWithDefaults

`func NewCapturedWithDefaults() *Captured`

NewCapturedWithDefaults instantiates a new Captured object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswered

`func (o *Captured) GetAnswered() string`

GetAnswered returns the Answered field if non-nil, zero value otherwise.

### GetAnsweredOk

`func (o *Captured) GetAnsweredOk() (*string, bool)`

GetAnsweredOk returns a tuple with the Answered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswered

`func (o *Captured) SetAnswered(v string)`

SetAnswered sets Answered field to given value.

### HasAnswered

`func (o *Captured) HasAnswered() bool`

HasAnswered returns a boolean if a field has been set.

### GetKey

`func (o *Captured) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Captured) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Captured) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Captured) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabels

`func (o *Captured) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Captured) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Captured) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Captured) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *Captured) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Captured) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Captured) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Captured) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


