# EmailMXCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Reason describes the failure or skip reason when Status is not pass. | [optional] 
**Records** | Pointer to **[]string** | Records lists the discovered mail exchanger hostnames. | [optional] 
**Status** | Pointer to **string** | Status reports the check outcome: pass, fail, unknown, or skipped. | [optional] 

## Methods

### NewEmailMXCheck

`func NewEmailMXCheck() *EmailMXCheck`

NewEmailMXCheck instantiates a new EmailMXCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMXCheckWithDefaults

`func NewEmailMXCheckWithDefaults() *EmailMXCheck`

NewEmailMXCheckWithDefaults instantiates a new EmailMXCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *EmailMXCheck) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EmailMXCheck) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EmailMXCheck) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EmailMXCheck) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRecords

`func (o *EmailMXCheck) GetRecords() []string`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *EmailMXCheck) GetRecordsOk() (*[]string, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *EmailMXCheck) SetRecords(v []string)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *EmailMXCheck) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetStatus

`func (o *EmailMXCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailMXCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailMXCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailMXCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


