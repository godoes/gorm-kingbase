package kingbase

import "testing"

func TestBuildUrl(t *testing.T) {
	type args struct {
		user     string
		password string
		host     string
		port     int
		dbname   string
		options  []map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "options",
			args: args{
				user:     "kes",
				password: "abc!@#$%^",
				host:     "127.0.0.1",
				port:     54321,
				dbname:   "kingbase",
				options:  []map[string]string{{"search_path": "public"}},
			},
			want: "kingbase://kes:abc%21%40%23$%25%5E@127.0.0.1:54321/kingbase?search_path=public",
		},
		{
			name: "multiple-host",
			args: args{
				user:     "kes",
				password: "abc!@#$%^",
				host:     "127.0.0.1,192.168.1.2,192.168.1.3",
				port:     54321,
				dbname:   "gorm",
			},
			want: "kingbase://kes:abc%21%40%23$%25%5E@127.0.0.1,192.168.1.2,192.168.1.3:54321/gorm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildUrl(tt.args.user, tt.args.password, tt.args.host, tt.args.port, tt.args.dbname, tt.args.options...); got != tt.want {
				t.Errorf("BuildUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
