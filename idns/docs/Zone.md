# Zone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Hosted zone id | [optional] 
**Name** | Pointer to **string** | Hosted zone name | [optional] 
**Domain** | Pointer to **string** | Hosted zone domain | [optional] 
**IsActive** | Pointer to **bool** | If hosted zone is active | [optional] 
**Retry** | Pointer to **NullableInt32** |  | [optional] 
**NxTtl** | Pointer to **NullableInt32** |  | [optional] 
**SoaTtl** | Pointer to **NullableInt32** |  | [optional] 
**Refresh** | Pointer to **NullableInt32** |  | [optional] 
**Expiry** | Pointer to **NullableInt32** |  | [optional] 
**Nameservers** | Pointer to **[]string** | List of nameservers | [optional] 

## Methods

### NewZone

`func NewZone() *Zone`

NewZone instantiates a new Zone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneWithDefaults

`func NewZoneWithDefaults() *Zone`

NewZoneWithDefaults instantiates a new Zone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Zone) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Zone) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Zone) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Zone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Zone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Zone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Zone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Zone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomain

`func (o *Zone) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Zone) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Zone) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Zone) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetIsActive

`func (o *Zone) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Zone) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Zone) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Zone) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetRetry

`func (o *Zone) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *Zone) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *Zone) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *Zone) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### SetRetryNil

`func (o *Zone) SetRetryNil(b bool)`

 SetRetryNil sets the value for Retry to be an explicit nil

### UnsetRetry
`func (o *Zone) UnsetRetry()`

UnsetRetry ensures that no value is present for Retry, not even an explicit nil
### GetNxTtl

`func (o *Zone) GetNxTtl() int32`

GetNxTtl returns the NxTtl field if non-nil, zero value otherwise.

### GetNxTtlOk

`func (o *Zone) GetNxTtlOk() (*int32, bool)`

GetNxTtlOk returns a tuple with the NxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxTtl

`func (o *Zone) SetNxTtl(v int32)`

SetNxTtl sets NxTtl field to given value.

### HasNxTtl

`func (o *Zone) HasNxTtl() bool`

HasNxTtl returns a boolean if a field has been set.

### SetNxTtlNil

`func (o *Zone) SetNxTtlNil(b bool)`

 SetNxTtlNil sets the value for NxTtl to be an explicit nil

### UnsetNxTtl
`func (o *Zone) UnsetNxTtl()`

UnsetNxTtl ensures that no value is present for NxTtl, not even an explicit nil
### GetSoaTtl

`func (o *Zone) GetSoaTtl() int32`

GetSoaTtl returns the SoaTtl field if non-nil, zero value otherwise.

### GetSoaTtlOk

`func (o *Zone) GetSoaTtlOk() (*int32, bool)`

GetSoaTtlOk returns a tuple with the SoaTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaTtl

`func (o *Zone) SetSoaTtl(v int32)`

SetSoaTtl sets SoaTtl field to given value.

### HasSoaTtl

`func (o *Zone) HasSoaTtl() bool`

HasSoaTtl returns a boolean if a field has been set.

### SetSoaTtlNil

`func (o *Zone) SetSoaTtlNil(b bool)`

 SetSoaTtlNil sets the value for SoaTtl to be an explicit nil

### UnsetSoaTtl
`func (o *Zone) UnsetSoaTtl()`

UnsetSoaTtl ensures that no value is present for SoaTtl, not even an explicit nil
### GetRefresh

`func (o *Zone) GetRefresh() int32`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *Zone) GetRefreshOk() (*int32, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *Zone) SetRefresh(v int32)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *Zone) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### SetRefreshNil

`func (o *Zone) SetRefreshNil(b bool)`

 SetRefreshNil sets the value for Refresh to be an explicit nil

### UnsetRefresh
`func (o *Zone) UnsetRefresh()`

UnsetRefresh ensures that no value is present for Refresh, not even an explicit nil
### GetExpiry

`func (o *Zone) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *Zone) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *Zone) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *Zone) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### SetExpiryNil

`func (o *Zone) SetExpiryNil(b bool)`

 SetExpiryNil sets the value for Expiry to be an explicit nil

### UnsetExpiry
`func (o *Zone) UnsetExpiry()`

UnsetExpiry ensures that no value is present for Expiry, not even an explicit nil
### GetNameservers

`func (o *Zone) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *Zone) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *Zone) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *Zone) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### SetNameserversNil

`func (o *Zone) SetNameserversNil(b bool)`

 SetNameserversNil sets the value for Nameservers to be an explicit nil

### UnsetNameservers
`func (o *Zone) UnsetNameservers()`

UnsetNameservers ensures that no value is present for Nameservers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


