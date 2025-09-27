使字符串平衡的最小交换次数
1
2
3
1963. 使字符串平衡的最小交换次数 - 力扣（LeetCode）
      https://leetcode.cn/problems/minimum-number-of-swaps-to-make-the-string-balanced/
1963. 使字符串平衡的最小交换次数-CSDN博客
      https://blog.csdn.net/Lucy_wzw/article/details/149149543
1963. 使字符串平衡的最小交换次数 - 归鸿唱晚 - 博客园
      https://www.cnblogs.com/Frank-Hong/p/15318427.html
      要将一个由等量的 '[' 和 ']' 组成的字符串调整为平衡字符串，需要通过最少的交换次数修复括号顺序。以下是解决此问题的高效方法。

贪心算法：统计最大不平衡度

该方法通过一次遍历计算字符串的不平衡程度，并基于此确定最少交换次数。

步骤

初始化两个变量： balance：表示当前的括号平衡度。 max_unbalance：记录遍历过程中最大的负平衡值。

遍历字符串： 遇到 '[' 时，balance += 1。 遇到 ']' 时，balance -= 1。 如果 balance < 0，更新 max_unbalance = max(max_unbalance, -balance)。

计算最少交换次数： 最少交换次数公式为： (max_unbalance + 1) // 2

代码实现

class Solution:
    def minSwaps(self, s: str) -> int:
        balance = 0
        max_unbalance = 0

        for ch in s:
            if ch == '[':
                balance += 1
            else: # ch == ']'
                balance -= 1
            if balance < 0:
                max_unbalance = max(max_unbalance, -balance)
        
        return (max_unbalance + 1) // 2
复制
示例

输入：

s = "]]][[["
复制
输出：

2
复制
解释：

初始不平衡度为 3（即需要修复的右括号数量）。

每次交换修复两个括号，因此需要 (3 + 1) // 2 = 2 次交换。

时间与空间复杂度

时间复杂度：O(n)，仅需一次遍历。

空间复杂度：O(1)，无需额外数据结构。

此方法高效且适用于大规模数据，是解决此类问题的最佳选择。


```go
func minSwaps(s string) int {
    cnt, mincnt := 0, 0
    for _, ch := range s {
        if ch == '[' {
            cnt++
        } else {
            cnt--
            mincnt = min(mincnt, cnt)
        }
    }
    return (-mincnt + 1) / 2
}
```