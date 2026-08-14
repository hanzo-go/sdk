# KeyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is RFC 3339 UTC. | [optional] 
**Fingerprint** | Pointer to **string** | Fingerprint is the key&#39;s SHA256 fingerprint (\&quot;SHA256:…\&quot;), globally unique and the handle SSH auth resolves a presented key by. | [optional] 
**Id** | Pointer to **string** | ID is the key&#39;s identifier (\&quot;gitkey_…\&quot;), the handle to delete it by. | [optional] 
**PublicKey** | Pointer to **string** | PublicKey is the canonical OpenSSH authorized-key line as stored. | [optional] 
**Title** | Pointer to **string** | Title is the key&#39;s label — the caller&#39;s, or the comment on the key line. | [optional] 

## Methods

### NewKeyView

`func NewKeyView() *KeyView`

NewKeyView instantiates a new KeyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyViewWithDefaults

`func NewKeyViewWithDefaults() *KeyView`

NewKeyViewWithDefaults instantiates a new KeyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KeyView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KeyView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KeyView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KeyView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFingerprint

`func (o *KeyView) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *KeyView) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *KeyView) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *KeyView) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetId

`func (o *KeyView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPublicKey

`func (o *KeyView) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *KeyView) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *KeyView) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *KeyView) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetTitle

`func (o *KeyView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *KeyView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *KeyView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *KeyView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


