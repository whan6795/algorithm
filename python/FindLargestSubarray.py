from typing import Tuple


class Solution:
    def __init__(self, input_list: list[int]):
        self.data = input_list

    def find_max_crossing_subarray(self, p: int, q: int, r: int) -> Tuple[int, int, int]:
        # p为左索引，q为中心值，r为右索引,寻找跨过节点q的最大子数组
        left_sum, right_sum = float('-INF'), float('-INF')
        left_index, right_index = -1, -1
        sum_num = 0
        for i in range(q - 1, p - 1, -1):
            sum_num += self.data[i]
            if sum_num > left_sum:
                left_sum = sum_num
                left_index = i
        sum_num = 0
        for j in range(q, r):
            sum_num += self.data[j]
            if sum_num > right_sum:
                right_sum = sum_num
                right_index = j
        return left_index, right_index, left_sum + right_sum

    def find_max_subarray(self, p: int, q: int) -> Tuple[int, int, int]:
        if p == q - 1:
            return p, q, self.data[p]
        mid = int((p + q) / 2)
        left_p, left_q, left_sum = self.find_max_subarray(p, mid)
        right_p, right_q, right_sum = self.find_max_subarray(mid, q)
        mid_p, mid_q, mid_sum = self.find_max_crossing_subarray(p, mid, q)
        if left_sum >= right_sum and left_sum >= mid_sum:
            return left_p, left_q, left_sum
        elif right_sum >= mid_sum and right_sum >= left_sum:
            return right_p, right_q, right_sum
        return mid_p, mid_q, mid_sum


a = [13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7]
solution = Solution(a)
p, q, sum_num = solution.find_max_subarray(0, len(a))
print(p, q, sum_num)
