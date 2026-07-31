# CloudLogLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** | INFO | WARN | ERROR | ... | [optional] 
**Source** | Pointer to **string** | \&quot;infra\&quot; (stdout) | \&quot;request\&quot; (org request log) | [optional] 
**Ts** | Pointer to **string** | RFC3339 (UTC) | [optional] 
**TsNano** | Pointer to **int32** | nanosecond cursor | [optional] 

## Methods

### NewCloudLogLine

`func NewCloudLogLine() *CloudLogLine`

NewCloudLogLine instantiates a new CloudLogLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLogLineWithDefaults

`func NewCloudLogLineWithDefaults() *CloudLogLine`

NewCloudLogLineWithDefaults instantiates a new CloudLogLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *CloudLogLine) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CloudLogLine) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CloudLogLine) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CloudLogLine) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetSeverity

`func (o *CloudLogLine) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CloudLogLine) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CloudLogLine) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CloudLogLine) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *CloudLogLine) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudLogLine) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudLogLine) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudLogLine) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTs

`func (o *CloudLogLine) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CloudLogLine) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CloudLogLine) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CloudLogLine) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetTsNano

`func (o *CloudLogLine) GetTsNano() int32`

GetTsNano returns the TsNano field if non-nil, zero value otherwise.

### GetTsNanoOk

`func (o *CloudLogLine) GetTsNanoOk() (*int32, bool)`

GetTsNanoOk returns a tuple with the TsNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsNano

`func (o *CloudLogLine) SetTsNano(v int32)`

SetTsNano sets TsNano field to given value.

### HasTsNano

`func (o *CloudLogLine) HasTsNano() bool`

HasTsNano returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


