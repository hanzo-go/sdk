# CloudKeyTypeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type is the key class to act on: \&quot;secret\&quot; (sk-, session-equivalent, belongs on a server) or \&quot;publishable\&quot; (pk-, org-identifying, safe in a browser bundle). Omitted means secret, which is what every existing caller means. | [optional] 

## Methods

### NewCloudKeyTypeIn

`func NewCloudKeyTypeIn() *CloudKeyTypeIn`

NewCloudKeyTypeIn instantiates a new CloudKeyTypeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudKeyTypeInWithDefaults

`func NewCloudKeyTypeInWithDefaults() *CloudKeyTypeIn`

NewCloudKeyTypeInWithDefaults instantiates a new CloudKeyTypeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CloudKeyTypeIn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudKeyTypeIn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudKeyTypeIn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudKeyTypeIn) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


