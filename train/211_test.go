package train

import "testing"

func TestWordDictionary_Search(t *testing.T) {
	trie := Constructor208()
	trie.Insert("at")
	trie.Insert("bat")
	type fields struct {
		trie Trie
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test1",
			fields: fields{
				trie: trie,
			},
			args: args{
				word: ".at",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &WordDictionary{
				trie: tt.fields.trie,
			}
			if got := this.Search(tt.args.word); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
