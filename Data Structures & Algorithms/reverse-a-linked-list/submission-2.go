

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    current := head

    for current != nil {
        next := current.Next   // lưu node tiếp theo
        current.Next = prev    // đảo chiều
        prev = current         // dời prev
        current = next         // dời current
    }

    return prev
}
