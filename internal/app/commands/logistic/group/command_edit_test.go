package group

import "testing"

func Test_parseEditGroupArgs(t *testing.T) {

	tests := []struct {
		name        string
		args        string
		wantGroupId uint64
		wantTitle   string
		wantErr     bool
	}{
		{"correct one", "1 Some Title", 1, "Some Title", false},
		{"incorrect id", "-1 Some Title", 0, "", true},
		{"incorrect arg", "TitleNoId", 0, "", true},
		{"empty arg", "", 0, "", true},
		{"Non int id", "Hello Title", 0, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGroupId, gotTitle, err := parseEditGroupArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseEditGroupArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotGroupId != tt.wantGroupId {
				t.Errorf("parseEditGroupArgs() gotGroupId = %v, want %v", gotGroupId, tt.wantGroupId)
			}
			if gotTitle != tt.wantTitle {
				t.Errorf("parseEditGroupArgs() gotTitle = %v, want %v", gotTitle, tt.wantTitle)
			}
		})
	}
}
