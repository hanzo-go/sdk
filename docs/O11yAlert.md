# O11yAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**EndsAt** | Pointer to **time.Time** |  | [optional] 
**GeneratorURL** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**StartsAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewO11yAlert

`func NewO11yAlert() *O11yAlert`

NewO11yAlert instantiates a new O11yAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAlertWithDefaults

`func NewO11yAlertWithDefaults() *O11yAlert`

NewO11yAlertWithDefaults instantiates a new O11yAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *O11yAlert) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *O11yAlert) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *O11yAlert) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *O11yAlert) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetEndsAt

`func (o *O11yAlert) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *O11yAlert) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *O11yAlert) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *O11yAlert) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetGeneratorURL

`func (o *O11yAlert) GetGeneratorURL() string`

GetGeneratorURL returns the GeneratorURL field if non-nil, zero value otherwise.

### GetGeneratorURLOk

`func (o *O11yAlert) GetGeneratorURLOk() (*string, bool)`

GetGeneratorURLOk returns a tuple with the GeneratorURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorURL

`func (o *O11yAlert) SetGeneratorURL(v string)`

SetGeneratorURL sets GeneratorURL field to given value.

### HasGeneratorURL

`func (o *O11yAlert) HasGeneratorURL() bool`

HasGeneratorURL returns a boolean if a field has been set.

### GetLabels

`func (o *O11yAlert) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *O11yAlert) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *O11yAlert) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *O11yAlert) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetStartsAt

`func (o *O11yAlert) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *O11yAlert) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *O11yAlert) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *O11yAlert) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


