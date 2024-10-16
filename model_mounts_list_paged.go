/*
dofusdude

# A project for you - the developer. The all-in-one toolbelt for your next Ankama related project.  ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [PHP](https://github.com/dofusdude/dofusdude-php) - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering game assets to high-res images with up to 800x800 px.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 0.9.3
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the MountsListPaged type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MountsListPaged{}

// MountsListPaged struct for MountsListPaged
type MountsListPaged struct {
	Links *LinksPaged `json:"_links,omitempty"`
	Items []MountListEntry `json:"items,omitempty"`
}

// NewMountsListPaged instantiates a new MountsListPaged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMountsListPaged() *MountsListPaged {
	this := MountsListPaged{}
	return &this
}

// NewMountsListPagedWithDefaults instantiates a new MountsListPaged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMountsListPagedWithDefaults() *MountsListPaged {
	this := MountsListPaged{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MountsListPaged) GetLinks() LinksPaged {
	if o == nil || IsNil(o.Links) {
		var ret LinksPaged
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MountsListPaged) GetLinksOk() (*LinksPaged, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MountsListPaged) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksPaged and assigns it to the Links field.
func (o *MountsListPaged) SetLinks(v LinksPaged) {
	o.Links = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *MountsListPaged) GetItems() []MountListEntry {
	if o == nil || IsNil(o.Items) {
		var ret []MountListEntry
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MountsListPaged) GetItemsOk() ([]MountListEntry, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MountsListPaged) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MountListEntry and assigns it to the Items field.
func (o *MountsListPaged) SetItems(v []MountListEntry) {
	o.Items = v
}

func (o MountsListPaged) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MountsListPaged) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableMountsListPaged struct {
	value *MountsListPaged
	isSet bool
}

func (v NullableMountsListPaged) Get() *MountsListPaged {
	return v.value
}

func (v *NullableMountsListPaged) Set(val *MountsListPaged) {
	v.value = val
	v.isSet = true
}

func (v NullableMountsListPaged) IsSet() bool {
	return v.isSet
}

func (v *NullableMountsListPaged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMountsListPaged(val *MountsListPaged) *NullableMountsListPaged {
	return &NullableMountsListPaged{value: val, isSet: true}
}

func (v NullableMountsListPaged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMountsListPaged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


