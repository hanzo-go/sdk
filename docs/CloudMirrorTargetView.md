# CloudMirrorTargetView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is RFC 3339 UTC. | [optional] 
**Host** | Pointer to **string** | Host is the target&#39;s lowercased hostname, taken from URL and never the body. | [optional] 
**Id** | Pointer to **string** | ID is the target&#39;s identifier (\&quot;mir_…\&quot;), the handle to remove it by. | [optional] 
**Repo** | Pointer to **string** | Repo is the repo whose advanced refs are pushed downstream. | [optional] 
**Url** | Pointer to **string** | URL is the canonical https remote, with any embedded credentials stripped. | [optional] 

## Methods

### NewCloudMirrorTargetView

`func NewCloudMirrorTargetView() *CloudMirrorTargetView`

NewCloudMirrorTargetView instantiates a new CloudMirrorTargetView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMirrorTargetViewWithDefaults

`func NewCloudMirrorTargetViewWithDefaults() *CloudMirrorTargetView`

NewCloudMirrorTargetViewWithDefaults instantiates a new CloudMirrorTargetView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudMirrorTargetView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudMirrorTargetView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudMirrorTargetView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudMirrorTargetView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHost

`func (o *CloudMirrorTargetView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudMirrorTargetView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudMirrorTargetView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudMirrorTargetView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *CloudMirrorTargetView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudMirrorTargetView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudMirrorTargetView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudMirrorTargetView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRepo

`func (o *CloudMirrorTargetView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudMirrorTargetView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudMirrorTargetView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudMirrorTargetView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetUrl

`func (o *CloudMirrorTargetView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudMirrorTargetView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudMirrorTargetView) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudMirrorTargetView) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


