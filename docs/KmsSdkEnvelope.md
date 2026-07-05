# KmsSdkEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**V** | **int32** | Envelope version. Only 1 is accepted. | 
**Id** | [**KmsSdkEnvelopeIdentity**](KmsSdkEnvelopeIdentity.md) |  | 
**Ts** | **int64** | Unix seconds at signing. Rejected outside ±5m of server time. | 
**Nonce** | **string** | Caller-fresh anti-replay nonce (typically 16 random bytes, base64). Reuse within the window is rejected. | 
**Op** | **int32** | Signed operation. 0x0040 get, 0x0041 put (also rotate), 0x0042 list, 0x0043 delete, 0x0050 sign, 0x0051 verify.  | 
**Req** | **map[string]interface{}** | Inner request JSON. Shape by op — secrets:{path,name,env[,value]}; sign:{validator_id,key_type,message}; verify adds {signature}.  | 
**Sig** | **string** | ML-DSA-65 signature (base64) over the canonical digest. | 

## Methods

### NewKmsSdkEnvelope

`func NewKmsSdkEnvelope(v int32, id KmsSdkEnvelopeIdentity, ts int64, nonce string, op int32, req map[string]interface{}, sig string, ) *KmsSdkEnvelope`

NewKmsSdkEnvelope instantiates a new KmsSdkEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsSdkEnvelopeWithDefaults

`func NewKmsSdkEnvelopeWithDefaults() *KmsSdkEnvelope`

NewKmsSdkEnvelopeWithDefaults instantiates a new KmsSdkEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetV

`func (o *KmsSdkEnvelope) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *KmsSdkEnvelope) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *KmsSdkEnvelope) SetV(v int32)`

SetV sets V field to given value.


### GetId

`func (o *KmsSdkEnvelope) GetId() KmsSdkEnvelopeIdentity`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsSdkEnvelope) GetIdOk() (*KmsSdkEnvelopeIdentity, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsSdkEnvelope) SetId(v KmsSdkEnvelopeIdentity)`

SetId sets Id field to given value.


### GetTs

`func (o *KmsSdkEnvelope) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *KmsSdkEnvelope) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *KmsSdkEnvelope) SetTs(v int64)`

SetTs sets Ts field to given value.


### GetNonce

`func (o *KmsSdkEnvelope) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *KmsSdkEnvelope) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *KmsSdkEnvelope) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetOp

`func (o *KmsSdkEnvelope) GetOp() int32`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *KmsSdkEnvelope) GetOpOk() (*int32, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *KmsSdkEnvelope) SetOp(v int32)`

SetOp sets Op field to given value.


### GetReq

`func (o *KmsSdkEnvelope) GetReq() map[string]interface{}`

GetReq returns the Req field if non-nil, zero value otherwise.

### GetReqOk

`func (o *KmsSdkEnvelope) GetReqOk() (*map[string]interface{}, bool)`

GetReqOk returns a tuple with the Req field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReq

`func (o *KmsSdkEnvelope) SetReq(v map[string]interface{})`

SetReq sets Req field to given value.


### GetSig

`func (o *KmsSdkEnvelope) GetSig() string`

GetSig returns the Sig field if non-nil, zero value otherwise.

### GetSigOk

`func (o *KmsSdkEnvelope) GetSigOk() (*string, bool)`

GetSigOk returns a tuple with the Sig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSig

`func (o *KmsSdkEnvelope) SetSig(v string)`

SetSig sets Sig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


