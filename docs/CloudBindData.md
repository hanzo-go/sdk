# CloudBindData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundAnchorSigner** | Pointer to **string** | BoundAnchorSigner is the EVM address now signing anchors. Fund it for gas. | [optional] 
**ChainId** | Pointer to **int32** | ChainID is the EVM chain the signer is bound for. | [optional] 
**Org** | Pointer to **string** | Org is the org whose treasury wallet was resolved. | [optional] 

## Methods

### NewCloudBindData

`func NewCloudBindData() *CloudBindData`

NewCloudBindData instantiates a new CloudBindData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBindDataWithDefaults

`func NewCloudBindDataWithDefaults() *CloudBindData`

NewCloudBindDataWithDefaults instantiates a new CloudBindData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundAnchorSigner

`func (o *CloudBindData) GetBoundAnchorSigner() string`

GetBoundAnchorSigner returns the BoundAnchorSigner field if non-nil, zero value otherwise.

### GetBoundAnchorSignerOk

`func (o *CloudBindData) GetBoundAnchorSignerOk() (*string, bool)`

GetBoundAnchorSignerOk returns a tuple with the BoundAnchorSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundAnchorSigner

`func (o *CloudBindData) SetBoundAnchorSigner(v string)`

SetBoundAnchorSigner sets BoundAnchorSigner field to given value.

### HasBoundAnchorSigner

`func (o *CloudBindData) HasBoundAnchorSigner() bool`

HasBoundAnchorSigner returns a boolean if a field has been set.

### GetChainId

`func (o *CloudBindData) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CloudBindData) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CloudBindData) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *CloudBindData) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetOrg

`func (o *CloudBindData) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudBindData) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudBindData) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudBindData) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


