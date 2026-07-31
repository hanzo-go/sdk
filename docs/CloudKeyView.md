# CloudKeyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is RFC 3339 UTC. | [optional] 
**Fingerprint** | Pointer to **string** | Fingerprint is the key&#39;s SHA256 fingerprint (\&quot;SHA256:…\&quot;), globally unique and the handle SSH auth resolves a presented key by. | [optional] 
**Id** | Pointer to **string** | ID is the key&#39;s identifier (\&quot;gitkey_…\&quot;), the handle to delete it by. | [optional] 
**PublicKey** | Pointer to **string** | PublicKey is the canonical OpenSSH authorized-key line as stored. | [optional] 
**Title** | Pointer to **string** | Title is the key&#39;s label — the caller&#39;s, or the comment on the key line. | [optional] 

## Methods

### NewCloudKeyView

`func NewCloudKeyView() *CloudKeyView`

NewCloudKeyView instantiates a new CloudKeyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudKeyViewWithDefaults

`func NewCloudKeyViewWithDefaults() *CloudKeyView`

NewCloudKeyViewWithDefaults instantiates a new CloudKeyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudKeyView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudKeyView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudKeyView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudKeyView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFingerprint

`func (o *CloudKeyView) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CloudKeyView) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CloudKeyView) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *CloudKeyView) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetId

`func (o *CloudKeyView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudKeyView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudKeyView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudKeyView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPublicKey

`func (o *CloudKeyView) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CloudKeyView) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CloudKeyView) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *CloudKeyView) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetTitle

`func (o *CloudKeyView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudKeyView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudKeyView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudKeyView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


