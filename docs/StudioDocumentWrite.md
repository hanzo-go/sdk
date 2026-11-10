# StudioDocumentWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** |  | [optional] 
**Ir** | Pointer to [**StudioIR**](StudioIR.md) |  | [optional] 
**Kind** | Pointer to **string** | Kind is doc or deck. | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewStudioDocumentWrite

`func NewStudioDocumentWrite() *StudioDocumentWrite`

NewStudioDocumentWrite instantiates a new StudioDocumentWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioDocumentWriteWithDefaults

`func NewStudioDocumentWriteWithDefaults() *StudioDocumentWrite`

NewStudioDocumentWriteWithDefaults instantiates a new StudioDocumentWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *StudioDocumentWrite) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *StudioDocumentWrite) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *StudioDocumentWrite) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *StudioDocumentWrite) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetIr

`func (o *StudioDocumentWrite) GetIr() StudioIR`

GetIr returns the Ir field if non-nil, zero value otherwise.

### GetIrOk

`func (o *StudioDocumentWrite) GetIrOk() (*StudioIR, bool)`

GetIrOk returns a tuple with the Ir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIr

`func (o *StudioDocumentWrite) SetIr(v StudioIR)`

SetIr sets Ir field to given value.

### HasIr

`func (o *StudioDocumentWrite) HasIr() bool`

HasIr returns a boolean if a field has been set.

### GetKind

`func (o *StudioDocumentWrite) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StudioDocumentWrite) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StudioDocumentWrite) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StudioDocumentWrite) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetShared

`func (o *StudioDocumentWrite) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *StudioDocumentWrite) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *StudioDocumentWrite) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *StudioDocumentWrite) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetTitle

`func (o *StudioDocumentWrite) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StudioDocumentWrite) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StudioDocumentWrite) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StudioDocumentWrite) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


