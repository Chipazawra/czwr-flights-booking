package session

import (
	"testing"
)

func Test_session_Auth(t *testing.T) {
	type fields struct {
		url string
		appid string
		secret string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "positive test",
			fields: fields{
				url: "https://test.api.amadeus.com/v1/security/oauth2/token",
				appid: "68ipF36GBniX4THOJBYovQ564gG86cjj",
				secret: "JBVIKTlKjcWCBAPy",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := session{
				url:   tt.fields.url,
				appid:  tt.fields.appid,
				secret:  tt.fields.secret,
			}
			if err := s.Init(); (err != nil) != tt.wantErr {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
