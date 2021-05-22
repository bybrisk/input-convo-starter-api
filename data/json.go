package data

import (
	"encoding/json"
	"io"
)	

func (d *InitialConversationResponse) InitialConversationResponseToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *InitialConversationRequest) FromJSONToInitialConversationRequest (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func (d *ActionHandlerResponse) ActionHandlerResponseToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *ActionHandlerRequest) FromJSONToActionHandlerRequest (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}