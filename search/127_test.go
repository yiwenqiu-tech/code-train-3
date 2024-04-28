package search

import "testing"

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "cog"},
			},
			want: 5,
		},
		{
			name: "test2",
			args: args{
				beginWord: "hot",
				endWord:   "dog",
				wordList:  []string{"hot", "dog", "dot"},
			},
			want: 3,
		},
		{
			name: "test3",
			args: args{
				beginWord: "hbo",
				endWord:   "qbx",
				wordList:  []string{"abo", "hco", "hbw", "ado", "abq", "hcd", "hcj", "hww", "qbq", "qby", "qbz", "qbx", "qbw"},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength2(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
