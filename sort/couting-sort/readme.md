# 计数排序

计数排序的核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。

## 算法步骤

1. 找出原数组中元素值最大的，记为max。

2. 创建一个新数组count，其长度是max加1，其元素默认值都为0。

3. 遍历原数组中的元素，以原数组中的元素作为count数组的索引，以原数组中的元素出现次数作为count数组的元素值。

4. 创建结果数组result，起始索引index。

5. 遍历count数组，找出其中元素值大于0的元素，将其对应的索引作为元素值填充到result数组中去，每处理一次，count中的该元素值减1，直到该元素值不大于0，依次处理count中剩下的元素。

6. 返回结果数组result。
