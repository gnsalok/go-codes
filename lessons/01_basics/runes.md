# Bytes, Runes, and UTF-8

A Go string is an immutable sequence of bytes. Source text is commonly UTF-8, but
the string type does not require valid UTF-8.

```go
text := "Go語"

fmt.Println(len(text))         // bytes: 5
fmt.Println(utf8.RuneCountInString(text)) // runes: 3

for byteOffset, r := range text {
	fmt.Printf("%d: %c\n", byteOffset, r)
}
```

`text[i]` returns a byte. A `range` loop decodes UTF-8 and returns the byte offset
and rune for each decoded code point. Converting to `[]rune` makes rune indexing
possible but allocates and still does not model user-perceived characters: a
grapheme cluster can contain multiple code points.

Use bytes for protocols and storage formats defined in bytes. Use runes when the
operation is defined in Unicode code points. Use a Unicode segmentation library
when the requirement is phrased in user-perceived characters.

## Interview prompts

1. Why can `len(s)` differ from the number of printed characters?
2. What happens when `range` encounters invalid UTF-8?
3. When is converting a large string to `[]rune` an unacceptable tradeoff?
