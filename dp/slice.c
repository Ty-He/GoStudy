#include <stdlib.h>
#include <string.h>
#include "def.h"

slice make_slice(size_t len)
{
  slice s = malloc(sizeof(struct slice_t));
  if (s == NULL) {
    return NULL;
  }
  s->data = malloc(len * sizeof(slice_value_type));
  bzero(s->data, len * sizeof(slice_value_type));
  s->len = len;
  s->cap = len;
  return s;
}

slice make_slice_with_cap(size_t len, size_t cap)
{
  if (len > cap) {
    return NULL;
  }
  slice s = make_slice(cap);
  s->len = len;
  return s;
}

slice make_view_slice(int* arr, size_t n)
{
  slice s = malloc(sizeof(struct slice_t));
  if (s == NULL) {
    return NULL;
  }
  s->data = arr;
  s->len = s->cap = n;
  return s;
}

void destroy(slice s)
{
  if (s) {
    free(s->data);
    free(s);
  }
}

void destroy_view_slice(slice s)
{
  if (s) free(s);
}

slice append(slice s, slice_value_type v)
{
  if (s == NULL) {
    s = make_slice(1);
    s->data[0] = v;
  }
  else if (s->len < s->cap) {
    s->data[s->len++] = v;
  }
  else {
    size_t newcap = s->cap + s->cap / 2 + 1;
    slice ns = make_slice_with_cap(s->len + 1, newcap);
    memcpy(ns->data, s->data, s->len * sizeof(slice_value_type));
    ns->data[s->len] = v;
    destroy(s);
    s = ns;
  }
  return s;
}

slice append_slice(slice s, slice other) {
  if (other == NULL) {
    return s;
  }
  for (int i = 0; i < other->len; ++i)
    s = append(s, other->data[i]);
  return s;
}
