package types_test

import (
	"testing"

	"github.com/onomyprotocol/reserve/testutil/sample"
	"github.com/onomyprotocol/reserve/x/reserve/types"
)

func TestCreateDenomProposal_ValidateBasic(t *testing.T) { //nolint:dupl // test template
	const denom1 = "denom1"

	type fields struct {
		Sender      string
		Title       string
		Description string
		Denom       string
		Rate        []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "positive",
			fields: fields{
				Sender:      sample.AccAddress(),
				Title:       "title",
				Description: "desc",
				Denom:       "denom1",
				Rate:        []string{"1", "1"},
			},
		},
		{
			name: "negative_invalid_sender",
			fields: fields{
				Sender:      "invalid-sender",
				Title:       "title",
				Description: "desc",
				Denom:       "denom",
				Rate:        []string{"1", "1"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &types.CreateDenomProposal{
				Sender:      tt.fields.Sender,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Denom:       tt.fields.Denom,
				Rate:        tt.fields.Rate,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
