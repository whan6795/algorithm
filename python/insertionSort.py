class Solution:
    def insertSort(self, input_list: list[int]) -> list[int]:
        # 从列表第二个元素开始，每个元素与前一个做对比，若小于前一个元素则交换位置，并继续向前比较，直到元素大于前一个元素或者到达列表开始位置
        for i in range(1, len(input_list)):
            keyword = input_list[i]
            j = i - 1
            while j >= 0 and input_list[j] > keyword:
                input_list[j + 1] = input_list[j]
                j -= 1
            input_list[j + 1] = keyword
        return input_list


# input_list = [1, 3, 5, 76, 3, 33]
# input_list = []
# input_list = [1]
input_list = [33, 2, 4, 3]
result = Solution().insertSort(input_list)
print(result)
