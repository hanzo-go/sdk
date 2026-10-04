# LspDiagnostic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **interface{}** |  | [optional] 
**Message** | Pointer to **string** | Message is the problem in the server&#39;s own words, meant to be shown. | [optional] 
**Range** | Pointer to [**LspRange**](LspRange.md) | Range is the span the problem is about. | [optional] 
**Severity** | Pointer to **int64** | Severity is the LSP&#39;s: 1 error, 2 warning, 3 information, 4 hint. A file with only 3s and 4s still compiles. | [optional] 
**Source** | Pointer to **string** | Source is which checker reported it (\&quot;compiler\&quot;, \&quot;go vet\&quot;, a linter&#39;s name), which is what separates a build error from a style opinion. | [optional] 

## Methods

### NewLspDiagnostic

`func NewLspDiagnostic() *LspDiagnostic`

NewLspDiagnostic instantiates a new LspDiagnostic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLspDiagnosticWithDefaults

`func NewLspDiagnosticWithDefaults() *LspDiagnostic`

NewLspDiagnosticWithDefaults instantiates a new LspDiagnostic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *LspDiagnostic) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LspDiagnostic) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LspDiagnostic) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *LspDiagnostic) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *LspDiagnostic) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *LspDiagnostic) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetMessage

`func (o *LspDiagnostic) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LspDiagnostic) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LspDiagnostic) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LspDiagnostic) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRange

`func (o *LspDiagnostic) GetRange() LspRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *LspDiagnostic) GetRangeOk() (*LspRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *LspDiagnostic) SetRange(v LspRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *LspDiagnostic) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSeverity

`func (o *LspDiagnostic) GetSeverity() int64`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *LspDiagnostic) GetSeverityOk() (*int64, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *LspDiagnostic) SetSeverity(v int64)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *LspDiagnostic) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *LspDiagnostic) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LspDiagnostic) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LspDiagnostic) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LspDiagnostic) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


