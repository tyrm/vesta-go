package speech

import "testing"

func TestVoiceHash(t *testing.T) {
	hashes := []struct {
		voice string
		phrase string
		vhash string
	}{
		{"jane", "i love pancakes", "1346893e256e5f112ee98288528ae4e219551837df6a40d6054a803f4d9b3669"},
		{"john", "i love pancakes", "e928303065786d62c24dc18d322f4fe72041864fe46374a90642a8da15a4cddd"},
		{"john", "good morning", "65509af59b701372659866ca08544480041d1f97042cf20ee68eedb97f8e155d"},
	}

	for _, hash := range hashes {
		newVhash := voiceHash(hash.voice, hash.phrase)
		if newVhash != hash.vhash {
			t.Errorf("Hash of (%s+%s) was incorrect, got: %s, want: %s.", hash.voice, hash.phrase, newVhash, hash.vhash)
		}
	}
}