#include <limits.h>
#include <stdlib.h>
#include <string.h>
// #include <stdio.h>
#include "def.h"

#define max(x, y) (x > y ? x : y)

// +++++++++++++++ 1.max sub array +++++++++++++++++
// dp[i] include s[i], could get max sum
// dp[i] = max(dp[i - 1], 0) + s[i] 
// ans = max{dp[i]...}
slice_value_type
maxSubArraySum(slice s)
{
  slice dp = make_slice(s->len);
  dp->data[0] = s->data[0];
  slice_value_type ans = dp->data[0];
  for (int i = 1; i < s->len; i++) {
    dp->data[i] = max(dp->data[i - 1], 0) + s->data[i];
    ans = max(ans, dp->data[i]);
  }
  destroy(dp);
  return ans;
}

// +++++++++++++++ 2.0-1bag +++++++++++++++++
// dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - weight[i]] + value[i])
slice_value_type
getMaxValue(slice value, slice weight, slice_value_type bag_capacity)
{
  if (value->len <= 0) return 0;
  slice dp[2];
  dp[0] = make_slice(bag_capacity + 1);
  dp[1] = make_slice(bag_capacity + 1);

  slice_value_type ans = 0;
  // dp[0]->data[0] = 0;
  if (weight->data[0] <= bag_capacity) {
    dp[0]->data[weight->data[0]] = value->data[0];
    ans = value->data[0];
  }

  for (int i = 1; i < value->len; i++) {
    // printf("v[i] = %d w[i] = %d\n", value->data[i], weight->data[i]);
    for (int j = weight->data[i]; j <= bag_capacity; j++) {
      dp[i & 1]->data[j] = max(dp[i - 1 & 1]->data[j], dp[i - 1 & 1]->data[j - weight->data[i]] + value->data[i]);
      // printf("i = %d, j = %d: max(%d, %d)\n", i, j, dp[i - 1 & 1]->data[j], dp[i - 1 & 1]->data[j - weight->data[i]] + value->data[i]);
      ans = max(ans, dp[i & 1]->data[j]);
    }
  }

  destroy(dp[0]);
  destroy(dp[1]);
  return ans;
}


// +++++++++++++++ 3.max-sub-length +++++++++++++++++
slice_value_type
longestCommonSubquence(slice s1, slice s2)
{
  size_t m = s1->len, n = s2->len;
  if (m <= 0 || n <= 0) return 0;
  
  // dp[i][0] = dp[0][j] = 0
  int** dp = malloc(sizeof(slice_value_type*) * (m + 1));
  for (int i = 0; i <= m; i++) {
    dp[i] = malloc(sizeof(slice_value_type) * (n + 1));
    dp[i][0] = 0;
  }
  memset(dp[0], 0, sizeof(slice_value_type) * (n + 1));

  for (int i = 1; i <= m; i++) {
    for (int j = 1; j <= n; j++) {
      dp[i][j] = s1->data[i - 1] == s2->data[j - 1] ? dp[i - 1][j - 1] + 1 :
        max(dp[i - 1][j], dp[i][j - 1]);
    }
  }
  
  slice_value_type ans = dp[m][n];
  for (int i = 0; i < m; i++) free(dp[i]);
  free(dp);

  return ans;
}


// +++++++++++++++ 4.maxtrix-mul +++++++++++++++++
#define min(x, y) (x < y ? x : y)

slice_value_type
minCostInMatrixMul(slice p)
{
  slice_value_type **dp = malloc(sizeof(void*) * p->len);
  for (int i = 0; i < p->len; i++) {
    dp[i] = malloc(sizeof(slice_value_type) * p->len);
    bzero(dp[i], sizeof(slice_value_type) * p->len);
  }
  size_t n = p->len - 1;

  for (int len = 2; len <= n; len++) {
    for (int i = 1; i + len - 1 <= n; i++) {
      int j = i + len - 1;
      dp[i][j] = INT_MAX;
      for (int k = i; k < j; k++) {
        dp[i][j] = min(dp[i][j], dp[i][k] + dp[k + 1][j] + p->data[i - 1] * p->data[k] * p->data[j]);
      }
    }
  }

  slice_value_type ans = dp[1][n];
  for (int i = 0; i < p->len; i++) free(dp[i]);
  free(dp);
  return ans;
}
