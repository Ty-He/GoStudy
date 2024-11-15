#include <math.h>
#include <limits.h>
#define __USE_GNU
#include <stdlib.h>
// #include <stdio.h>
#include "def.h"

static double
calc_distance(node_t nx, node_t ny)
{
  return sqrt(pow(nx->x - ny->x, 2) + pow(nx->y - ny->y, 2));
}

static double
closest_dis_reg(node_t* nodes, size_t n)
{
  double distance = INT_MAX;
  for (int i = 0; i < n; i++) {
    for (int j = i + 1; j < n; j++) {
      double cur = calc_distance(nodes[i], nodes[j]);
      if (cur < distance) distance = cur;
    }
  }

  return distance;
}

static int
cmp_node_by_y(const void* p1, const void* p2, void* context)
{
  slice_value_type
  i = *(slice_value_type*)p1, j = *(slice_value_type*)p2;
  node_t* nodeset = *(node_t**)context;
  return nodeset[i]->y - nodeset[j]->y;
}

// this func compile only by gcc for qsort_r()
// if try to running on other, return -1
static double
closest_dis_div(node_t* nodes, int l, int r)
{
  if (l == r) return INT_MAX;
  else if (l + 1 == r) return calc_distance(nodes[l], nodes[r]);

  int mid = (l + r) >> 1;
  double d1 = closest_dis_div(nodes, l, mid);
  double d2 = closest_dis_div(nodes, mid + 1, r);

  double min_dis = d1 < d2 ? d1 : d2;
  
  slice s = NULL;
  for (int i = l; i <= r; i++) {
    if (fabs(nodes[i]->x - nodes[mid]->x) < min_dis)
      s = append(s, i);
  }

#ifdef __GNUC__
  // sort by y greater
  // user qsort_r() is better, __USE_GNU
  if (s->data) {
    // printf("before qsort: s.len = %lu\n", s->len);
    qsort_r(s->data, s->len, sizeof(slice_value_type), cmp_node_by_y, &nodes);
    // printf("after qsort: s.len = %lu\n", s->len);
  }
#else 
  destroy(s);
  return -1;
#endif

  for (int i = 0; i < s->len; i++) {
    for (int j = i + 1; j < s->len && nodes[s->data[j]]->y - nodes[s->data[i]]->y < min_dis; j++) {
      double cur = calc_distance(nodes[s->data[i]], nodes[s->data[j]]);
      if (cur < min_dis) min_dis = cur;
    }
  }

  destroy(s);
  return min_dis;
}

static int
compare_node(const void* arg1, const void* arg2)
{
  node_t nx = *(node_t*)arg1, ny = *(node_t*)arg2;
  if (nx->x != ny->x) return nx->y - ny->y;
  return nx->x - ny->x;
}

double
get_closest_distance(node_t* nodes, size_t n, int flag)
{
  switch (flag) {
    case CLOSEST_PAIR_REG:
      return closest_dis_reg(nodes, n);
    case CLOSEST_PAIR_DIV:
      qsort(nodes, n, sizeof(node_t), compare_node);
      return closest_dis_div(nodes, 0, n - 1);
    default:
      return 0.;
  }
}

node_t make_node(double x, double y) {
  node_t p;
  p = malloc(sizeof(*p));
  if (p == NULL) return NULL;
  p->x = x;
  p->y = y;
  return p;
}
