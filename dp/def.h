#pragma once
#include <stddef.h> // for size_t

// slice.c
// implement dynamic array
typedef int slice_value_type;
typedef struct slice_t {
  slice_value_type* data;
  size_t len;
  size_t cap;
}* slice;

slice make_slice(size_t len);
slice make_slice_with_cap(size_t len, size_t cap);
slice make_view_slice(int* arr, size_t n);
void destroy(slice s);
slice append(slice s, slice_value_type v);
slice append_slice(slice s, slice other);
void destroy_view_slice(slice s);

// solution.c
slice_value_type maxSubArraySum(slice s);
slice_value_type getMaxValue(slice value, slice weight, slice_value_type bag_capacity);
slice_value_type longestCommonSubquence(slice s1, slice s2);
slice_value_type minCostInMatrixMul(slice p);
