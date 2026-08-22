# Completion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | Detail is the server&#39;s short elaboration, typically the type or signature. | [optional] 
**Kind** | Pointer to **int32** | Kind is the LSP CompletionItemKind number (2 method, 3 function, 5 field, 6 variable, …), passed through as the protocol spells it. | [optional] 
**Label** | Pointer to **string** | Label is the text a client would insert, and what an editor lists. | [optional] 

## Methods

### NewCompletion

`func NewCompletion() *Completion`

NewCompletion instantiates a new Completion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionWithDefaults

`func NewCompletionWithDefaults() *Completion`

NewCompletionWithDefaults instantiates a new Completion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *Completion) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Completion) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Completion) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Completion) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetKind

`func (o *Completion) GetKind() int32`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Completion) GetKindOk() (*int32, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Completion) SetKind(v int32)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Completion) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *Completion) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Completion) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Completion) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Completion) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


