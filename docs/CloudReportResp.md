# CloudReportResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | Pointer to **int32** | Accepted is how many samples passed validation. Every one of them was accepted, or the whole report was refused — there is no partial success. | [optional] 
**Stored** | Pointer to **bool** | Stored is whether the warehouse actually persisted them. False means the datastore was unavailable and the poll of history was lost; the request still succeeded, so a device retries without being blocked. | [optional] 

## Methods

### NewCloudReportResp

`func NewCloudReportResp() *CloudReportResp`

NewCloudReportResp instantiates a new CloudReportResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudReportRespWithDefaults

`func NewCloudReportRespWithDefaults() *CloudReportResp`

NewCloudReportRespWithDefaults instantiates a new CloudReportResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *CloudReportResp) GetAccepted() int32`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *CloudReportResp) GetAcceptedOk() (*int32, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *CloudReportResp) SetAccepted(v int32)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *CloudReportResp) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### GetStored

`func (o *CloudReportResp) GetStored() bool`

GetStored returns the Stored field if non-nil, zero value otherwise.

### GetStoredOk

`func (o *CloudReportResp) GetStoredOk() (*bool, bool)`

GetStoredOk returns a tuple with the Stored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStored

`func (o *CloudReportResp) SetStored(v bool)`

SetStored sets Stored field to given value.

### HasStored

`func (o *CloudReportResp) HasStored() bool`

HasStored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


