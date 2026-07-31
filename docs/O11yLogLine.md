# O11yLogLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | RFC3339 UTC timestamp. | [optional] 
**TsNano** | Pointer to **int64** | Nanosecond epoch cursor. Pass the response nextCursor back as sinceNs to tail. | [optional] 
**Severity** | Pointer to **string** | INFO, WARN, ERROR, etc. | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** | infra is the raw stdout stream (SuperAdmin only); request is the per-org request log. | [optional] 

## Methods

### NewO11yLogLine

`func NewO11yLogLine() *O11yLogLine`

NewO11yLogLine instantiates a new O11yLogLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yLogLineWithDefaults

`func NewO11yLogLineWithDefaults() *O11yLogLine`

NewO11yLogLineWithDefaults instantiates a new O11yLogLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *O11yLogLine) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *O11yLogLine) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *O11yLogLine) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *O11yLogLine) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetTsNano

`func (o *O11yLogLine) GetTsNano() int64`

GetTsNano returns the TsNano field if non-nil, zero value otherwise.

### GetTsNanoOk

`func (o *O11yLogLine) GetTsNanoOk() (*int64, bool)`

GetTsNanoOk returns a tuple with the TsNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsNano

`func (o *O11yLogLine) SetTsNano(v int64)`

SetTsNano sets TsNano field to given value.

### HasTsNano

`func (o *O11yLogLine) HasTsNano() bool`

HasTsNano returns a boolean if a field has been set.

### GetSeverity

`func (o *O11yLogLine) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *O11yLogLine) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *O11yLogLine) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *O11yLogLine) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetBody

`func (o *O11yLogLine) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *O11yLogLine) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *O11yLogLine) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *O11yLogLine) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetSource

`func (o *O11yLogLine) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *O11yLogLine) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *O11yLogLine) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *O11yLogLine) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


