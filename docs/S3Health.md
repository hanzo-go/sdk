# S3Health

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error is why the probe is degraded, in plain words. Absent when it is not. | [optional] 
**Presign** | Pointer to **bool** | Presign is whether presigned upload and download URLs can be minted, which needs a PUBLIC endpoint on top of the credentials. False does not make the surface degraded — listing and creating still work.  NOT omitempty. On the ready path the untyped probe wrote this key unconditionally, and omitting a false one would turn what a healthy-but-presignless deployment reports from \&quot;presign: false\&quot; into silence — two different facts. The degraded body now carries it too, which is the one delta of typing this: the probe answers ONE shape under both of its statuses, which is what makes a single declared Out honest. | [optional] 
**Ready** | Pointer to **bool** | Ready is whether this deployment can serve object operations at all: true only when admin credentials are configured. | [optional] 
**Service** | Pointer to **string** | Service names the subsystem this probe is for. Always \&quot;s3\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; when the store is reachable in principle, \&quot;degraded\&quot; when it is not. It is the field to read; the HTTP status carries the same fact for a caller that only looks at the code. | [optional] 

## Methods

### NewS3Health

`func NewS3Health() *S3Health`

NewS3Health instantiates a new S3Health object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3HealthWithDefaults

`func NewS3HealthWithDefaults() *S3Health`

NewS3HealthWithDefaults instantiates a new S3Health object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *S3Health) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *S3Health) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *S3Health) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *S3Health) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPresign

`func (o *S3Health) GetPresign() bool`

GetPresign returns the Presign field if non-nil, zero value otherwise.

### GetPresignOk

`func (o *S3Health) GetPresignOk() (*bool, bool)`

GetPresignOk returns a tuple with the Presign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresign

`func (o *S3Health) SetPresign(v bool)`

SetPresign sets Presign field to given value.

### HasPresign

`func (o *S3Health) HasPresign() bool`

HasPresign returns a boolean if a field has been set.

### GetReady

`func (o *S3Health) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *S3Health) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *S3Health) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *S3Health) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetService

`func (o *S3Health) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *S3Health) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *S3Health) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *S3Health) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *S3Health) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *S3Health) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *S3Health) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *S3Health) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


