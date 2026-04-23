func getConcatenation(nums []int) []int {
    ans := make(int[], len(nums) * 2)

    copy(ans, num)

    for (i:= len(nums); i ++; i < len(nums)*2 - 1) {
        ans[i] = ans[i  - len[nums]]
    }
}
