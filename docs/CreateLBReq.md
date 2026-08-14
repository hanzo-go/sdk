# CreateLBReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwardingRules** | Pointer to [**[]FwdRule**](FwdRule.md) | ForwardingRules are the listen→backend port mappings. Empty defaults to plain HTTP 80→80, the same default DigitalOcean&#39;s own console applies. | [optional] 
**Name** | Pointer to **string** | Name is the FRIENDLY name, a DNS-safe slug of at most 40 characters. The physical DigitalOcean name is derived from it and the caller&#39;s org. | [optional] 
**Region** | Pointer to **string** | Region is the DigitalOcean region slug (nyc3, sfo3, …). Required. | [optional] 
**Size** | Pointer to **string** | Size is the DigitalOcean size slug. Empty takes DO&#39;s default. | [optional] 
**Type** | Pointer to **string** | Type is the DigitalOcean load-balancer type. Empty takes DO&#39;s default (REGIONAL). | [optional] 

## Methods

### NewCreateLBReq

`func NewCreateLBReq() *CreateLBReq`

NewCreateLBReq instantiates a new CreateLBReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLBReqWithDefaults

`func NewCreateLBReqWithDefaults() *CreateLBReq`

NewCreateLBReqWithDefaults instantiates a new CreateLBReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardingRules

`func (o *CreateLBReq) GetForwardingRules() []FwdRule`

GetForwardingRules returns the ForwardingRules field if non-nil, zero value otherwise.

### GetForwardingRulesOk

`func (o *CreateLBReq) GetForwardingRulesOk() (*[]FwdRule, bool)`

GetForwardingRulesOk returns a tuple with the ForwardingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingRules

`func (o *CreateLBReq) SetForwardingRules(v []FwdRule)`

SetForwardingRules sets ForwardingRules field to given value.

### HasForwardingRules

`func (o *CreateLBReq) HasForwardingRules() bool`

HasForwardingRules returns a boolean if a field has been set.

### GetName

`func (o *CreateLBReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLBReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLBReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLBReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *CreateLBReq) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateLBReq) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateLBReq) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateLBReq) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSize

`func (o *CreateLBReq) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateLBReq) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateLBReq) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateLBReq) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *CreateLBReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateLBReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateLBReq) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateLBReq) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


