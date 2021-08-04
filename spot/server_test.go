package spot

import (
	"go-binance"
	"context"
	"reflect"
	"testing"
)

func TestServerTimeService_Do(t *testing.T) {
	type fields struct {
		Client    *binance.Client
		ApiKey    string
		ApiSecret string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		{"TestServerTimeService_Do", fields{
			ApiKey:    "",
			ApiSecret: "",
		}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := binance.NewClient(tt.fields.ApiKey, tt.fields.ApiSecret)
			got, err := NewServerTimeService(c).Do(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServerTime error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetServerTime got = %v, want %v", got, tt.want)
			}
		})
	}
}
