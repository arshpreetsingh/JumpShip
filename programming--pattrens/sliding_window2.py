#Smallest Subarray with a Given Sum: Given an array of positive integers and a target sum,
# find the length of the smallest contiguous subarray whose sum is greater than or equal to the target sum.

def smallestSubarrayWithSum(nums, targetSum):
    start = 0
    currentSum = 0
    minLength = float('inf')  # Initialize minLength to positive infinity

    for end in range(len(nums)):
        currentSum += nums[end]

        while currentSum >= targetSum:
            minLength = min(minLength, end - start + 1)
            currentSum -= nums[start]
            start += 1

    return minLength if minLength != float('inf') else 0


# Test the function
nums = [2, 3, 1, 2, 4, 3]
targetSum = 7
result = smallestSubarrayWithSum(nums, targetSum)
print(result)  # Output: 2

