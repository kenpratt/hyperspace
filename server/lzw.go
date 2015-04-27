package main

import "log"

// Compress a string to a list of output symbols.
func LzwCompress(uncompressed []byte) string {
	// Build the dictionary.
	dictSize := 256
	dictionary := make(map[string]rune)
	for i := 0; i < 256; i++ {
		dictionary[string(i)] = RuneFor(i)
	}

	w := ""
	result := make([]rune, 0)
	for _, c := range uncompressed {
		wc := w + string(c)
		if _, ok := dictionary[wc]; ok {
			w = wc
		} else {
			result = append(result, dictionary[w])
			// Add wc to the dictionary.
			dictionary[wc] = RuneFor(dictSize)
			dictSize++
			w = string(c)
		}
	}

	// Output the code for w.
	if w != "" {
		result = append(result, dictionary[w])
	}
	return string(result)
}

// Decompress a list of output ks to a string.
func LzwDecompress(raw string) []byte {
	compressed := []rune(raw)

	// Build the dictionary.
	dictSize := 256
	dictionary := make(map[rune]string)
	for i := 0; i < 256; i++ {
		dictionary[RuneFor(i)] = string(i)
	}

	w := string(compressed[0])
	result := w
	for _, k := range compressed[1:] {
		var entry string
		if x, ok := dictionary[k]; ok {
			entry = x
		} else if k == RuneFor(dictSize) {
			entry = w + w[:1]
		} else {
			log.Println("Error compressing to LZW", k)
			return nil
		}

		result += entry

		// Add w+entry[0] to the dictionary.
		dictionary[RuneFor(dictSize)] = w + entry[:1]
		dictSize++

		w = entry
	}
	return []byte(result)
}

// skip unicode surrogates (http://en.wikipedia.org/wiki/Mapping_of_Unicode_characters#Surrogates)
func RuneFor(i int) rune {
	if i >= 0xD800 {
		i += (0xDFFF - 0xD800 + 1)
	}
	if i >= 0xFFFE {
		i++
	}
	if i >= 0xFFFF {
		i++
	}
	return rune(i)
}
