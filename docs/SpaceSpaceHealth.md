# SpaceSpaceHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error is why the probe is degraded, in plain words. Absent when it is not. | [optional] 
**Presign** | Pointer to **bool** | Presign is whether upload and download URLs can be minted, which needs a PUBLIC endpoint on top of the credentials. False does not make the surface degraded — listing spaces, drives and folders still works, only the bytes cannot be reached. | [optional] 
**Ready** | Pointer to **bool** | Ready is whether this deployment can serve drive and file operations at all: true only when object-store credentials are configured. | [optional] 
**Service** | Pointer to **string** | Service names the subsystem this probe is for. Always \&quot;space\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; when the store is reachable in principle, \&quot;degraded\&quot; when it is not. It is the field to read; the HTTP status carries the same fact for a caller that only looks at the code. | [optional] 

## Methods

### NewSpaceSpaceHealth

`func NewSpaceSpaceHealth() *SpaceSpaceHealth`

NewSpaceSpaceHealth instantiates a new SpaceSpaceHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceSpaceHealthWithDefaults

`func NewSpaceSpaceHealthWithDefaults() *SpaceSpaceHealth`

NewSpaceSpaceHealthWithDefaults instantiates a new SpaceSpaceHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SpaceSpaceHealth) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SpaceSpaceHealth) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SpaceSpaceHealth) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SpaceSpaceHealth) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPresign

`func (o *SpaceSpaceHealth) GetPresign() bool`

GetPresign returns the Presign field if non-nil, zero value otherwise.

### GetPresignOk

`func (o *SpaceSpaceHealth) GetPresignOk() (*bool, bool)`

GetPresignOk returns a tuple with the Presign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresign

`func (o *SpaceSpaceHealth) SetPresign(v bool)`

SetPresign sets Presign field to given value.

### HasPresign

`func (o *SpaceSpaceHealth) HasPresign() bool`

HasPresign returns a boolean if a field has been set.

### GetReady

`func (o *SpaceSpaceHealth) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *SpaceSpaceHealth) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *SpaceSpaceHealth) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *SpaceSpaceHealth) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetService

`func (o *SpaceSpaceHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SpaceSpaceHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SpaceSpaceHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SpaceSpaceHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *SpaceSpaceHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpaceSpaceHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpaceSpaceHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpaceSpaceHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


