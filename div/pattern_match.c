#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "def.h"


// BF Algorithm
slice
bf(char const* text, char const* pattern)
{
  slice s = NULL; 
  unsigned long text_len = strlen(text), pattern_len = strlen(pattern);
  int j = 0; // index of pattern
  for (int i = 0; i < text_len; ++i) {
    if (j == pattern_len) {
      s = append(s, i - pattern_len);
      --i;
      j = 0;
    } else if (text[i] != pattern[j]) {
      i -= j;
      j = 0;
    } else {
      ++j;
    }
  }
  return s;
}

// KMP Algorithm
slice
kmp(char const* text, char const* pattern)
{
  unsigned long text_len = strlen(text), pattern_len = strlen(pattern);
  slice prefix_m = make_slice(pattern_len);
  slice_value_type* prefix = prefix_m->data;
  prefix[0] = 0;
  // i is suffer index, j is prefix index
  for (int j = 0, i = 1; i < pattern_len; ++i) {
    while (j > 0 && pattern[i] != pattern[j])
      j = prefix[i];
    if (pattern[i] == pattern[j]) ++j;

    prefix[i] = j;
  }

  for (int i = 0; i < prefix_m->len; ++i)
    printf("%d -> %d\n", i, prefix[i]);

  slice s = NULL;
  for (int i = 0, j = 0; i < text_len;) {
    if (text[i] != pattern[j]) {
      // i -= j;
      if (j == 0) {
        ++i;
      } else {
        j = prefix[j - 1];
      }
    } else {
      ++i;
      ++j;
    }

    if (j == pattern_len) {
      s = append(s, i - pattern_len);
      // printf("ans = %lu\n", i - pattern_len);
      if (i == text_len - 1) break;
      j = 0;
    }
  }

  destroy(prefix_m);
  return s;
}

slice
pattern_match(int s, char const* text, char const* pattern)
{
  switch (s) {
case BF:
  return bf(text, pattern);
case KMP:
  return kmp(text, pattern);
default:
  return NULL;
  }
}
