// running test on linux, because use some syscall to file io
#ifdef __linux__
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#endif

#include <stdlib.h> // for free
#include <stdio.h>
#include <string.h>
#include <time.h>
#include "def.h"

void test_kmp_by_chars()
{
  slice s = pattern_match(KMP, "aabaabaaf", "aabaaf");
  // slice s = pattern_match(KMP, "aabababc", "ab");
  // for (int i = 0; i < s->len; ++i)
  //   printf("%d ", s->data[i]);
  // printf("\n");
  destroy(s);
}

void test_kmp_by_file()
{
#ifdef __linux__
  printf("test_kmp_by_file running:\n");
  int fd = open("./def.h", O_RDONLY);
  if (fd == -1) {
    perror("open");
    return;
  }

  char buf[1024];
  memset(buf, 0, sizeof(buf));

  int n = read(fd, buf, sizeof(buf));
  if (n < 0) {
    perror("read");
    return;
  }

  char target[] = "slice";
  slice s1 = pattern_match(BF, buf, target);
  slice s2 = pattern_match(KMP, buf, target);

  if (s1->len != s2->len) {
    printf("pattern_match failed\n");
    return;
  }

  int m = strlen(target);
  for (int i = 0; i < s1->len; ++i) {

    if (s1->data[i] != s2->data[i]) {
      printf("pattern_match failed\n");
      return;
    }

    printf("%d -> ", s1->data[i]);
    for (int j = 0; j < m; ++j) {
      printf("%c", *(buf + j + s1->data[i]));
    }
    printf("\n");

    for (int j = 0; j < m; j++) {
      if (*(buf + s1->data[i] + j) != target[j]) {
        printf("not match charnot match char:%c\n", *(buf + s1->data[i] + j));
        printf("index = %d\n", s1->data[i]);
        printf("pattern_match failed\n");
        return;
      }
    }
  }

  printf("pattern_match pass\n");
#endif
}

void test_maxsum() {
  int a[] = {-2, 11, -4, 13, -5, -2};
  int b[] = {-6, 2, 4, -7, 5, 3, 2, -1, 6, -9, 10, -2};
  slice s = make_view_slice(a, sizeof(a) / sizeof(int));

  int ans = max_sum_of_sub_arr(s, MAX_SUB_SUM_DIV);
  if (ans != 20) {
    printf("test_maxsum failed, ans = %d\n", ans);
    destroy_view_slice(s);
    return;
  }
  destroy_view_slice(s);
  printf("test_maxsum pass 1\n");

  s = make_view_slice(b, sizeof(b) / sizeof(int));
  if ((ans = max_sum_of_sub_arr(s, MAX_SUB_SUM_DP)) != 16) {
    printf("test_maxsum failed, ans = %d\n", ans);
    destroy_view_slice(s);
    return;
  }
  destroy_view_slice(s);
  printf("test_maxsum pass 2\n");
}

void test_mode() {
  int a[] = {2, 2, 9, 1, 5, 6, 4, 1, 3, 8, 2};
  slice s = make_view_slice(a, sizeof(a) / sizeof(int));
  slice r = get_mode(s);
  printf("get_mode: %d -> %d\n", r->data[0], r->data[1]);

  if (r->data[0] != 2 && r->data[1] != 3) {
    printf("test_mode failed\n");
  } else {
    printf("test_mode pass\n");
  }

  destroy(r);
  destroy_view_slice(s);
}

void test_closest_pair()
{
#ifdef __linux__
  int fd = open("./test_closest_pair", O_RDONLY);
  if (fd == -1) {
    perror("open");
    return;
  }
  node_t nodes[11];
  char buf[4];
  for (int i = 0; i < 11; i++) {
    if (read(fd, buf, 4) != 4) {
      printf("test_closest_pair: data err\n");
      for (int j = 0; j < i; j++) free(nodes[j]);
      close(fd);
      return;
    }
    nodes[i] = make_node(buf[0] - '0', buf[2] - '0');
    printf("node %d: x = %lf, y = %lf\n", i, nodes[i]->x, nodes[i]->y);
  }

  double dis1 = get_closest_distance(nodes, 11, CLOSEST_PAIR_REG);
  double dis2 = get_closest_distance(nodes, 11, CLOSEST_PAIR_DIV);

  printf("dis1(reg) = %lf, dis2(div) = %lf\n", dis1, dis2);

  if (dis1 != dis2 || dis1 != 1.) {
    printf("test_closest_pair failed\n");
  } else {
    printf("test_closest_pair pass\n");
  }

  close(fd);
  for (int i = 0; i < 11; i++) free(nodes[i]);
#endif
}

void test_closest_pair_with_time()
{
  srand((unsigned int)time(NULL));
  const int n = 100000;
  node_t nodes[n];
  for (int i = 0; i < n; i++)
    nodes[i] = make_node(rand(), rand());
  clock_t begin_time = clock();
  double dis = get_closest_distance(nodes, n, CLOSEST_PAIR_DIV);
  clock_t end_time = clock();

  double time_taken = (double)(end_time - begin_time) / CLOCKS_PER_SEC;

  printf("pair size = %d, time_taken = %lf second\n", n, time_taken);
}
