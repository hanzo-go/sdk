# LinkMint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to [**CodeView**](CodeView.md) | Link is the link just minted, with its full shareable URL. Its funnel counters all start at zero — nothing has clicked or signed up through it yet. | [optional] 

## Methods

### NewLinkMint

`func NewLinkMint() *LinkMint`

NewLinkMint instantiates a new LinkMint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkMintWithDefaults

`func NewLinkMintWithDefaults() *LinkMint`

NewLinkMintWithDefaults instantiates a new LinkMint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *LinkMint) GetLink() CodeView`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *LinkMint) GetLinkOk() (*CodeView, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *LinkMint) SetLink(v CodeView)`

SetLink sets Link field to given value.

### HasLink

`func (o *LinkMint) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


