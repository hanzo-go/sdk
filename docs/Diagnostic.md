# Diagnostic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **map[string]interface{}** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Range** | Pointer to [**Range**](Range.md) |  | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewDiagnostic

`func NewDiagnostic() *Diagnostic`

NewDiagnostic instantiates a new Diagnostic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticWithDefaults

`func NewDiagnosticWithDefaults() *Diagnostic`

NewDiagnosticWithDefaults instantiates a new Diagnostic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Diagnostic) GetCode() map[string]interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Diagnostic) GetCodeOk() (*map[string]interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Diagnostic) SetCode(v map[string]interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *Diagnostic) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *Diagnostic) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Diagnostic) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Diagnostic) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Diagnostic) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRange

`func (o *Diagnostic) GetRange() Range`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Diagnostic) GetRangeOk() (*Range, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Diagnostic) SetRange(v Range)`

SetRange sets Range field to given value.

### HasRange

`func (o *Diagnostic) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSeverity

`func (o *Diagnostic) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Diagnostic) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Diagnostic) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Diagnostic) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *Diagnostic) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Diagnostic) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Diagnostic) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Diagnostic) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


