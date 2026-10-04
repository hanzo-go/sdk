# DeclareResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**Declaration**](Declaration.md) |  | [optional] 
**Build** | Pointer to [**BuildRef**](BuildRef.md) |  | [optional] 
**Declaration** | Pointer to [**DeclareResult**](DeclareResult.md) |  | [optional] 
**Notice** | Pointer to **string** |  | [optional] 

## Methods

### NewDeclareResp

`func NewDeclareResp() *DeclareResp`

NewDeclareResp instantiates a new DeclareResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclareRespWithDefaults

`func NewDeclareRespWithDefaults() *DeclareResp`

NewDeclareRespWithDefaults instantiates a new DeclareResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *DeclareResp) GetApp() Declaration`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *DeclareResp) GetAppOk() (*Declaration, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *DeclareResp) SetApp(v Declaration)`

SetApp sets App field to given value.

### HasApp

`func (o *DeclareResp) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetBuild

`func (o *DeclareResp) GetBuild() BuildRef`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *DeclareResp) GetBuildOk() (*BuildRef, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *DeclareResp) SetBuild(v BuildRef)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *DeclareResp) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetDeclaration

`func (o *DeclareResp) GetDeclaration() DeclareResult`

GetDeclaration returns the Declaration field if non-nil, zero value otherwise.

### GetDeclarationOk

`func (o *DeclareResp) GetDeclarationOk() (*DeclareResult, bool)`

GetDeclarationOk returns a tuple with the Declaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaration

`func (o *DeclareResp) SetDeclaration(v DeclareResult)`

SetDeclaration sets Declaration field to given value.

### HasDeclaration

`func (o *DeclareResp) HasDeclaration() bool`

HasDeclaration returns a boolean if a field has been set.

### GetNotice

`func (o *DeclareResp) GetNotice() string`

GetNotice returns the Notice field if non-nil, zero value otherwise.

### GetNoticeOk

`func (o *DeclareResp) GetNoticeOk() (*string, bool)`

GetNoticeOk returns a tuple with the Notice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotice

`func (o *DeclareResp) SetNotice(v string)`

SetNotice sets Notice field to given value.

### HasNotice

`func (o *DeclareResp) HasNotice() bool`

HasNotice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


