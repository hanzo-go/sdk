# Verdict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | Pointer to [**[]DriftFlag**](DriftFlag.md) | Flags are the findings behind the severity, in detection order: floating-declared, floating-running, stale, un-rolled, then the release-artifact ones. Always present — &#x60;[]&#x60; for a row that runs what it declares, never null. | [optional] 
**Severity** | Pointer to **string** | Severity is the roll-up over Flags — red if any flag is red, else yellow if any is yellow, else ok. It is the column a board sorts and filters on, and \&quot;ok\&quot; is exactly what no flags means. | [optional] 

## Methods

### NewVerdict

`func NewVerdict() *Verdict`

NewVerdict instantiates a new Verdict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerdictWithDefaults

`func NewVerdictWithDefaults() *Verdict`

NewVerdictWithDefaults instantiates a new Verdict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *Verdict) GetFlags() []DriftFlag`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Verdict) GetFlagsOk() (*[]DriftFlag, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Verdict) SetFlags(v []DriftFlag)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Verdict) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetSeverity

`func (o *Verdict) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Verdict) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Verdict) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Verdict) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


