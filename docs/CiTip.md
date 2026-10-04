# CiTip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **time.Time** |  | [optional] 
**Build** | Pointer to [**CiCheck**](CiCheck.md) |  | [optional] 
**Sha** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewCiTip

`func NewCiTip() *CiTip`

NewCiTip instantiates a new CiTip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiTipWithDefaults

`func NewCiTipWithDefaults() *CiTip`

NewCiTipWithDefaults instantiates a new CiTip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *CiTip) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *CiTip) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *CiTip) SetAt(v time.Time)`

SetAt sets At field to given value.

### HasAt

`func (o *CiTip) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBuild

`func (o *CiTip) GetBuild() CiCheck`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *CiTip) GetBuildOk() (*CiCheck, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *CiTip) SetBuild(v CiCheck)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *CiTip) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetSha

`func (o *CiTip) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *CiTip) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *CiTip) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *CiTip) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetTitle

`func (o *CiTip) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CiTip) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CiTip) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CiTip) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


