package cmd

import "testing"

func Test_infoFull(t *testing.T) {
	type args struct {
		cveID string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Get full CVE data",
			args{cveID: "CVE-2019-11510"},
			"a JSON value",
		},
	}
	for _, tt := range tests {

		//TODO: Convert the slice of bytes to an array
		t.Run(tt.name, func(t *testing.T) {
			// if got := infoFull(tt.args.cveID); got == "" {
			// 	t.Errorf("infoFull() = %v, want %v", got, tt.want)
			// }
		})
	}
}
