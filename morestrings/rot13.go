package morestrings

// ReverseRunes returns its argument string reversed rune-wise left to right.
func Rot13(r rune) rune {
  switch {
  case r >= 'A' && r <= 'Z':
    return 'A' + (r-'A'+13)%26
  case r >= 'a' && r <= 'z':
    return 'a' + (r-'a'+13)%26
  }
  return r
}
