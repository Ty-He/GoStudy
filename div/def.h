#pragma once
#include <stddef.h>

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

// pattern_match.c
enum strategy_of_pattern_match {
  BF,
  KMP
};
slice pattern_match(int s, char const* text, char const* pattern);

// maxsum.c
enum max_sum_strategy {
  MAX_SUB_SUM_DIV,
  MAX_SUB_SUM_DP
};
// flag == 0, use div; else use dp
int max_sum_of_sub_arr(slice s, int flag);

// mode.c
slice get_mode(slice s);

// closest_pair.c
enum {
  CLOSEST_PAIR_REG,
  CLOSEST_PAIR_DIV
};
typedef struct {
  double x, y;
}* node_t;
node_t make_node(double x, double y);
double get_closest_distance(node_t* nodes, size_t n, int flag);

// test.c
void test_kmp_by_file();
void test_kmp_by_chars();
void test_maxsum();
void test_mode();
void test_closest_pair();
void test_closest_pair_with_time();
