#include "limits.h"
#include "def.h"

static int
div(slice s, int l, int r)
{
  if (l == r) {
    return s->data[l];
  }

  int mid = (l + r) >> 1;
  int left_val = div(s, l, mid);
  int right_val = div(s, mid + 1, r);
  int max_val = left_val > right_val ? left_val : right_val;

  // cross
  int left_cross_sum = INT_MIN;
  for (int i = mid, sum = 0; i >= l; i--) {
    sum += s->data[i];  
    if (sum > left_cross_sum) left_cross_sum = sum;
  }
  int right_cross_sum = INT_MIN;
  for (int i = mid, sum = 0; i <= r; i++) {
    sum += s->data[i];  
    if (sum > right_cross_sum) right_cross_sum = sum;
  }

  int cross = left_cross_sum + right_cross_sum - s->data[mid];
  
  return max_val > cross ? max_val : cross;
}

static int
dp(slice s) {
  int ans = INT_MIN, sum = 0;
  for (int i = 0; i < s->len; i++) {
    sum += s->data[i];
    if (sum > ans) ans = sum;
    if (sum < 0) sum = 0;
  }
  return ans;
}

int max_sum_of_sub_arr(slice s, int flag)
{
  if (flag == MAX_SUB_SUM_DP) {
    return dp(s);
  }
  if (s->len == 0) return INT_MIN;
  return div(s, 0, s->len - 1);
}
