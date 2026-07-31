# CloudRegisterKeyReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | Pointer to **string** | PublicKey is one OpenSSH authorized-key line (\&quot;ssh-ed25519 AAAA… you@host\&quot;). Required; a line that does not parse is refused and never stored. | [optional] 
**Title** | Pointer to **string** | Title labels the key in the console. Max 256 chars; when omitted the comment on the key line is used. | [optional] 

## Methods

### NewCloudRegisterKeyReq

`func NewCloudRegisterKeyReq() *CloudRegisterKeyReq`

NewCloudRegisterKeyReq instantiates a new CloudRegisterKeyReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRegisterKeyReqWithDefaults

`func NewCloudRegisterKeyReqWithDefaults() *CloudRegisterKeyReq`

NewCloudRegisterKeyReqWithDefaults instantiates a new CloudRegisterKeyReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *CloudRegisterKeyReq) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CloudRegisterKeyReq) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CloudRegisterKeyReq) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *CloudRegisterKeyReq) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetTitle

`func (o *CloudRegisterKeyReq) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudRegisterKeyReq) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudRegisterKeyReq) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudRegisterKeyReq) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


