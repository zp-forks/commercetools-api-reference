// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// StoreUpdateAction uses action as discriminator attribute
type StoreUpdateAction interface{}

func mapDiscriminatorStoreUpdateAction(input interface{}) (StoreUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "setName":
		new := StoreSetNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Store is a standalone struct
type Store struct {
	Version        int              `json:"version"`
	Name           *LocalizedString `json:"name,omitempty"`
	LastModifiedAt time.Time        `json:"lastModifiedAt"`
	Key            string           `json:"key"`
	ID             string           `json:"id"`
	CreatedAt      time.Time        `json:"createdAt"`
}

// StoreDraft is a standalone struct
type StoreDraft struct {
	Name *LocalizedString `json:"name"`
	Key  string           `json:"key"`
}

// StoreKeyReference implements the interface KeyReference
type StoreKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreKeyReference) MarshalJSON() ([]byte, error) {
	type Alias StoreKeyReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "store", Alias: (*Alias)(&obj)})
}

// StorePagedQueryResponse is of type PagedQueryResponse
type StorePagedQueryResponse struct {
	Total   int     `json:"total,omitempty"`
	Offset  int     `json:"offset"`
	Count   int     `json:"count"`
	Results []Store `json:"results"`
}

// StoreReference implements the interface Reference
type StoreReference struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
	Obj *Store `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreReference) MarshalJSON() ([]byte, error) {
	type Alias StoreReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "store", Alias: (*Alias)(&obj)})
}

// StoreSetNameAction implements the interface StoreUpdateAction
type StoreSetNameAction struct {
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

// StoreUpdate is of type Update
type StoreUpdate struct {
	Version int                 `json:"version"`
	Actions []StoreUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StoreUpdate) UnmarshalJSON(data []byte) error {
	type Alias StoreUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorStoreUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
