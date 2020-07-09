#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

long long solve(vector<long long> a) {
  sort(a.begin(), a.end());
  const long long max = 1e18;
  long long ans = 1;
  for (int i = 0; i < a.size(); i++) {
    if (a[i] == 0) {
      return 0;
    } else if (a[i] > max / ans) {
      return -1;
    }
    ans *= a[i];
  }
  return ans;
}

int main() {
  int n;
  cin >> n;
  vector<long long> a(n);
  for (int i = 0; i < n; i++) {
    cin >> a[i];
  }
  cout << solve(a) << endl;
  return 0;
}
