# CloudDropletIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disk** | Pointer to **bool** | Disk requests a PERMANENT resize that grows the disk. DO can never resize such a droplet down again, so it defaults false — a CPU/RAM-only change, reversible. | [optional] 
**Id** | Pointer to **string** | ID is the DO droplet id, from the path. Numeric. | [optional] 
**Size** | Pointer to **string** | Size is the target DigitalOcean size slug on resize, e.g. \&quot;s-4vcpu-8gb\&quot;. | [optional] 

## Methods

### NewCloudDropletIn

`func NewCloudDropletIn() *CloudDropletIn`

NewCloudDropletIn instantiates a new CloudDropletIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDropletInWithDefaults

`func NewCloudDropletInWithDefaults() *CloudDropletIn`

NewCloudDropletInWithDefaults instantiates a new CloudDropletIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisk

`func (o *CloudDropletIn) GetDisk() bool`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *CloudDropletIn) GetDiskOk() (*bool, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *CloudDropletIn) SetDisk(v bool)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *CloudDropletIn) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetId

`func (o *CloudDropletIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudDropletIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudDropletIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudDropletIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSize

`func (o *CloudDropletIn) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CloudDropletIn) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CloudDropletIn) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CloudDropletIn) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


