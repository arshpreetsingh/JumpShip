# This problem is related to sliding window
nums = [1, 3, -1, -3, 5, 3, 6, 7]
k = 3
def maxSumSubarray(nums, k):
    max_sum = 0
    current_sum = 0
    start = 0
    for end in range(len(nums)):
        current_sum+=nums[end]
        print("current-sum",current_sum)
        if (end-start+1)==k:
            print("now entering into values")
            max_sum = max(current_sum,max_sum)
            current_sum-=nums[start]
            start+=1
    return max_sum

result = maxSumSubarray(nums, k)
print(result)  # Output: 16