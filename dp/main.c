#include <stdio.h>
#include <string.h>
#include "def.h"

void test_1();
void test_2();
void test_3();
void test_4();

int
main(int argc, char* argv[])
{
  if (argc != 2) {
    printf("useage dp_exe [option]\n");
    return 0;
  }

  if (!strcmp(argv[1], "help")) {
    printf("1. maxSubArraySum\n");
    printf("2. 0-1bag\n");
    printf("3. longestCommonSubquence\n");
    printf("4. maxtrixMul\n");
    return 0;
  }
  int op = argv[1][0] - '0';
  switch (op) {
    case 1:
      test_1();
      break;
    case 2:
      test_2();
      break;
    case 3:
      test_3();
      break;
    case 4:
      test_4();
      break;
    default:
      printf("Unkonw test\n");
  }
  return 0;
}

void println(const char* str, slice s) {
  printf("%s ", str);
  for (int i = 0; i < s->len; i++) {
    printf("%d ", s->data[i]);
  }
  printf("\n");
}

void test_1()
{
  int a[] = {-2, 11, -4, 13, -5, -2};
  int b[] = {-6, 2, 4, -7, 5, 3, 2, -1, 6, -9, 10, -2};
  slice s = make_view_slice(a, sizeof(a) / sizeof(a[0]));
  println("test case1:", s);
  if (maxSubArraySum(s) != 20) {
    printf("fail!\n");
  } else {
    printf("pass: ans = 20\n");
  }

  destroy_view_slice(s);

  s = make_view_slice(b, sizeof(b) / sizeof(b[0]));
  println("test case2:", s);
  if (maxSubArraySum(s) != 16) {
    printf("fail!\n");
  } else {
    printf("pass: ans = 16\n");
  }

  destroy_view_slice(s);
}

void test_2()
{
  int v[] = {4, 5, 2, 3, 1};
  int w[] = {3, 5, 1, 2, 2};
  slice val = make_view_slice(v, sizeof(v) / sizeof(v[0]));
  slice wei = make_view_slice(w, sizeof(w) / sizeof(w[0]));

  println("value array:", val);
  println("weight array:", wei);
  int ans;
  if ((ans = getMaxValue(val, wei, 8)) != 10) {
    printf("fail: ans = %d\n", ans);
  } else {
    printf("pass: ans = %d\n", ans);
  }

  destroy_view_slice(val);
  destroy_view_slice(wei);
}

void test_3()
{
  int a[] = {3, 2, 1, 4, 5};
  int b[] = {1, 2, 3, 4, 5};
  slice s1 = make_view_slice(a, sizeof(a) / sizeof(a[0]));
  slice s2 = make_view_slice(b, sizeof(b) / sizeof(b[0]));
  println("s1 array:", s1);
  println("s2 array:", s2);
  int ans;
  if ((ans = longestCommonSubquence(s1, s2)) != 3) {
    printf("fail: ans = %d\n", ans);
  } else {
    printf("pass: ans = %d\n", ans);
  }

  destroy_view_slice(s1);
  destroy_view_slice(s2);
}

void test_4() 
{
  int a[] = {30, 35, 15, 5, 10, 20, 25};
  slice s = make_view_slice(a, sizeof(a) / sizeof(a[0]));
  println("matrix chain:", s);
  int ans;
  if ((ans = minCostInMatrixMul(s)) != 15125) {
    printf("fail: ans = %d\n", ans);
  } else {
    printf("pass: ans = %d\n", ans);
  } 
  destroy_view_slice(s);
}

