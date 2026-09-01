# Tip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **time.Time** |  | [optional] 
**Build** | Pointer to [**Check**](Check.md) |  | [optional] 
**Sha** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewTip

`func NewTip() *Tip`

NewTip instantiates a new Tip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTipWithDefaults

`func NewTipWithDefaults() *Tip`

NewTipWithDefaults instantiates a new Tip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *Tip) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *Tip) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *Tip) SetAt(v time.Time)`

SetAt sets At field to given value.

### HasAt

`func (o *Tip) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBuild

`func (o *Tip) GetBuild() Check`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *Tip) GetBuildOk() (*Check, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *Tip) SetBuild(v Check)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *Tip) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetSha

`func (o *Tip) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *Tip) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *Tip) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *Tip) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetTitle

`func (o *Tip) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Tip) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Tip) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Tip) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


