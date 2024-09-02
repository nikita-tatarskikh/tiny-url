package url_hash

import "testing"

func TestURLHashImpl_HashURL(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "success",
			input: "Hello world",
			want:  "333",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := URLHashImpl{}
			if got := h.HashURL(tt.input); got != tt.want {
				t.Errorf("HashURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
