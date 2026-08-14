# CreateLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is an optional vanity code; it must be free across the whole directory, and omitting it mints a random one. Body-only. | [optional] 
**Label** | Pointer to **string** | Label is cosmetic — trimmed, stripped of control characters, capped — and never part of a code. Body-only: the URL cannot supply it. | [optional] 

## Methods

### NewCreateLinkRequest

`func NewCreateLinkRequest() *CreateLinkRequest`

NewCreateLinkRequest instantiates a new CreateLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLinkRequestWithDefaults

`func NewCreateLinkRequestWithDefaults() *CreateLinkRequest`

NewCreateLinkRequestWithDefaults instantiates a new CreateLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateLinkRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateLinkRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateLinkRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateLinkRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *CreateLinkRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateLinkRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateLinkRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateLinkRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


