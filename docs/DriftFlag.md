# DriftFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kind is which finding this is, one of stale, un-rolled, floating-declared, floating-running, no-release or zero-assets. It is what code matches on, and the kinds are independent — one row can carry several at once. | [optional] 
**Message** | Pointer to **string** | Message is the finding in words, naming the tags that produced it (\&quot;running v1.2.3 has not rolled to declared v1.2.4\&quot;). For display: match on Kind. | [optional] 
**Severity** | Pointer to **string** | Severity is this ONE finding&#39;s weight — yellow for stale and un-rolled, red for the other four. It is a constant of the kind (severityOf), never a judgement about the row, so the same kind always weighs the same. | [optional] 

## Methods

### NewDriftFlag

`func NewDriftFlag() *DriftFlag`

NewDriftFlag instantiates a new DriftFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriftFlagWithDefaults

`func NewDriftFlagWithDefaults() *DriftFlag`

NewDriftFlagWithDefaults instantiates a new DriftFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *DriftFlag) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DriftFlag) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DriftFlag) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DriftFlag) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMessage

`func (o *DriftFlag) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DriftFlag) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DriftFlag) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DriftFlag) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSeverity

`func (o *DriftFlag) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *DriftFlag) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *DriftFlag) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *DriftFlag) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


