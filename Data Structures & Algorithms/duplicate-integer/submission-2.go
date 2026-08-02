func hasDuplicate(nums []int) bool {
    hashset := make(map[int]struct{}) // map is like dictionary, but since we only use the key to check duplicate, we use key=int and value=struct{}, because struct{} in Go basically claim 0 space (Okb) so it's efficient
    for i:=0; i<len(nums); i++ {
        if _, exists := hashset[nums[i]]; exists {
            return true
        }
        hashset[nums[i]] = struct{}{}
    }
    return false
}
