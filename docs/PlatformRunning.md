# PlatformRunning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** | Digest is the digest the kubelet resolved that image to, \&quot;\&quot; when no pod has reported one yet. | [optional] 
**Image** | Pointer to **string** | Image is the image the app&#39;s container runs, as the pod spec names it. When pods disagree — a rollout in flight — it is the one most of them run. | [optional] 
**Pods** | Pointer to **int64** | Pods is how many pods carry the app&#39;s selector. | [optional] 
**Ready** | Pointer to **int64** | Ready is how many of them report every container ready. | [optional] 

## Methods

### NewPlatformRunning

`func NewPlatformRunning() *PlatformRunning`

NewPlatformRunning instantiates a new PlatformRunning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformRunningWithDefaults

`func NewPlatformRunningWithDefaults() *PlatformRunning`

NewPlatformRunningWithDefaults instantiates a new PlatformRunning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *PlatformRunning) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *PlatformRunning) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *PlatformRunning) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *PlatformRunning) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetImage

`func (o *PlatformRunning) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PlatformRunning) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PlatformRunning) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PlatformRunning) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetPods

`func (o *PlatformRunning) GetPods() int64`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *PlatformRunning) GetPodsOk() (*int64, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *PlatformRunning) SetPods(v int64)`

SetPods sets Pods field to given value.

### HasPods

`func (o *PlatformRunning) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetReady

`func (o *PlatformRunning) GetReady() int64`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *PlatformRunning) GetReadyOk() (*int64, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *PlatformRunning) SetReady(v int64)`

SetReady sets Ready field to given value.

### HasReady

`func (o *PlatformRunning) HasReady() bool`

HasReady returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


