# CloudAuthoredPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is when the plugin was last built, Unix seconds. | [optional] 
**Id** | Pointer to **string** | ID is the plugin&#39;s id within the org, and the id a delete addresses. | [optional] 
**Name** | Pointer to **string** | Name is the plugin&#39;s name: one lowercase path segment, the id it runs by. | [optional] 
**Org** | Pointer to **string** | Org is the org that built the plugin — the validated caller&#39;s. | [optional] 
**Provider** | Pointer to **string** | Provider is the connectors provider whose credential this plugin uses at run time. Absent for a plugin that needs none. The credential itself is never here — it stays under KMS custody in the connectors plane. | [optional] 
**Source** | Pointer to **string** | Source is the TypeScript as authored (or as generated from a spec). | [optional] 

## Methods

### NewCloudAuthoredPlugin

`func NewCloudAuthoredPlugin() *CloudAuthoredPlugin`

NewCloudAuthoredPlugin instantiates a new CloudAuthoredPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAuthoredPluginWithDefaults

`func NewCloudAuthoredPluginWithDefaults() *CloudAuthoredPlugin`

NewCloudAuthoredPluginWithDefaults instantiates a new CloudAuthoredPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudAuthoredPlugin) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAuthoredPlugin) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAuthoredPlugin) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAuthoredPlugin) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudAuthoredPlugin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAuthoredPlugin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAuthoredPlugin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAuthoredPlugin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudAuthoredPlugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAuthoredPlugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAuthoredPlugin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAuthoredPlugin) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *CloudAuthoredPlugin) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudAuthoredPlugin) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudAuthoredPlugin) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudAuthoredPlugin) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProvider

`func (o *CloudAuthoredPlugin) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudAuthoredPlugin) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudAuthoredPlugin) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudAuthoredPlugin) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSource

`func (o *CloudAuthoredPlugin) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudAuthoredPlugin) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudAuthoredPlugin) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudAuthoredPlugin) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


