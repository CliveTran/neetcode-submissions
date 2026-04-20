func getConcatenation(nums []int) []int {
    return append(nums, nums...)
}

// we can create a n*2 array and assign 2 values at the same time 
// int[i]=value and int[i+arrLenght]=value
// so the Time=O(N/2) and Space=O(N)