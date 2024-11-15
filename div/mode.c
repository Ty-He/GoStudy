#include "def.h"

static slice
div(slice s, int l, int r)
{
  if (l == r) {
    slice r = make_slice(2);
    r->data[0] = s->data[l];
    r->data[1] = 1;
    return r;
  }

  int mid = (l + r) >> 1;
  slice leftmode = div(s, l, mid);
  slice rightmode = div(s, mid + 1, r);

  // same num
  if (leftmode->data[0] == rightmode->data[0]) {
    leftmode->data[1] += rightmode->data[1];
    destroy(rightmode);
    return leftmode;
  }
  
  for (int i = mid + 1; i <= r; i++) {
    if (s->data[i] == leftmode->data[0]) ++leftmode->data[1];
  }
  for (int i = mid; i >= l; i--) {
    if (s->data[i] == rightmode->data[0]) ++rightmode->data[1];
  }

  if (leftmode->data[1] > rightmode->data[1]) {
    destroy(rightmode);
    return leftmode;
  } else {
    destroy(leftmode);
    return rightmode;
  }
}

// mode.c
slice get_mode(slice s)
{
  return div(s, 0, s->len - 1);
}
