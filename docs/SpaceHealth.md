# SpaceHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error is why the probe is degraded, in plain words. Absent when it is not. | [optional] 
**Presign** | Pointer to **bool** | Presign is whether upload and download URLs can be minted, which needs a PUBLIC endpoint on top of the credentials. False does not make the surface degraded — listing spaces, drives and folders still works, only the bytes cannot be reached. | [optional] 
**Ready** | Pointer to **bool** | Ready is whether this deployment can serve drive and file operations at all: true only when object-store credentials are configured. | [optional] 
**Service** | Pointer to **string** | Service names the subsystem this probe is for. Always \&quot;space\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; when the store is reachable in principle, \&quot;degraded\&quot; when it is not. It is the field to read; the HTTP status carries the same fact for a caller that only looks at the code. | [optional] 

## Methods

### NewSpaceHealth

`func NewSpaceHealth() *SpaceHealth`

NewSpaceHealth instantiates a new SpaceHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceHealthWithDefaults

`func NewSpaceHealthWithDefaults() *SpaceHealth`

NewSpaceHealthWithDefaults instantiates a new SpaceHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SpaceHealth) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SpaceHealth) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SpaceHealth) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SpaceHealth) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPresign

`func (o *SpaceHealth) GetPresign() bool`

GetPresign returns the Presign field if non-nil, zero value otherwise.

### GetPresignOk

`func (o *SpaceHealth) GetPresignOk() (*bool, bool)`

GetPresignOk returns a tuple with the Presign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresign

`func (o *SpaceHealth) SetPresign(v bool)`

SetPresign sets Presign field to given value.

### HasPresign

`func (o *SpaceHealth) HasPresign() bool`

HasPresign returns a boolean if a field has been set.

### GetReady

`func (o *SpaceHealth) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *SpaceHealth) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *SpaceHealth) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *SpaceHealth) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetService

`func (o *SpaceHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SpaceHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SpaceHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SpaceHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *SpaceHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpaceHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpaceHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpaceHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


