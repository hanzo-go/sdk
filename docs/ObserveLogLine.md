# ObserveLogLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **string** | RFC3339 timestamp (UTC). | [optional] 
**TsNano** | Pointer to **int64** | Nanosecond cursor. | [optional] 
**Severity** | Pointer to **string** | INFO | WARN | ERROR | ... | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** | \&quot;infra\&quot; (stdout) or \&quot;request\&quot; (org request log). | [optional] 

## Methods

### NewObserveLogLine

`func NewObserveLogLine() *ObserveLogLine`

NewObserveLogLine instantiates a new ObserveLogLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObserveLogLineWithDefaults

`func NewObserveLogLineWithDefaults() *ObserveLogLine`

NewObserveLogLineWithDefaults instantiates a new ObserveLogLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *ObserveLogLine) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ObserveLogLine) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ObserveLogLine) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ObserveLogLine) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetTsNano

`func (o *ObserveLogLine) GetTsNano() int64`

GetTsNano returns the TsNano field if non-nil, zero value otherwise.

### GetTsNanoOk

`func (o *ObserveLogLine) GetTsNanoOk() (*int64, bool)`

GetTsNanoOk returns a tuple with the TsNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsNano

`func (o *ObserveLogLine) SetTsNano(v int64)`

SetTsNano sets TsNano field to given value.

### HasTsNano

`func (o *ObserveLogLine) HasTsNano() bool`

HasTsNano returns a boolean if a field has been set.

### GetSeverity

`func (o *ObserveLogLine) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ObserveLogLine) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ObserveLogLine) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ObserveLogLine) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetBody

`func (o *ObserveLogLine) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ObserveLogLine) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ObserveLogLine) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ObserveLogLine) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetSource

`func (o *ObserveLogLine) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ObserveLogLine) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ObserveLogLine) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ObserveLogLine) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


