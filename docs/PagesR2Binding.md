# PagesR2Binding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the R2 bucket this binding points at, by bucket name — R2 addresses buckets by name where KV and D1 use ids. The binding name the Worker code reads it as is the map key. | [optional] 

## Methods

### NewPagesR2Binding

`func NewPagesR2Binding() *PagesR2Binding`

NewPagesR2Binding instantiates a new PagesR2Binding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagesR2BindingWithDefaults

`func NewPagesR2BindingWithDefaults() *PagesR2Binding`

NewPagesR2BindingWithDefaults instantiates a new PagesR2Binding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PagesR2Binding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PagesR2Binding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PagesR2Binding) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PagesR2Binding) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


