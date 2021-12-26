import random


class Solution:
    def __init__(self, input_list: list[float]):
        self.data = input_list

    def merge(self, p: int, q: int, r: int):
        # p为左序列最左侧索引，q为做序列最右侧索引，r为右序列最右侧序列，merge函数目的为合并已排序好的序列data[p, q]和data[q+1, r],使得data[p, r]也符合排序要求
        l_list = self.data[p: q]
        r_list = self.data[q: r]
        l_list.append(float("INF"))
        r_list.append(float("INF"))
        l_index, r_index = 0, 0
        for i in range(p, r):
            if l_list[l_index] > r_list[r_index]:
                self.data[i] = r_list[r_index]
                r_index += 1
            else:
                self.data[i] = l_list[l_index]
                l_index += 1

    def mergeSort(self, p, r):
        if p < r - 1:
            q = int((p + r) / 2)
            self.mergeSort(p, q)
            self.mergeSort(q, r)
            self.merge(p, q, r)


a = [random.randint(0, 99) for i in range(16)]
print("input data is: ", a)
solution = Solution(a)
solution.mergeSort(0, 16)
print("sorted data is: ", solution.data)
