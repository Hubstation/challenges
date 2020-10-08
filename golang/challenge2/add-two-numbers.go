package challenge2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val   int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	v , n  : =  0 , 0

	for {
		// Add operation on the current bit
		v, n = add(l1, l2, n)
		temp.Val = v

		// go to the next
		l1 = next(l1)
		l2  =  next ( l2 )
		// If the next digit of the two numbers is nil, then the bitwise addition operation ends
		if l1 == nil && l2 == nil {
			break
		}

		// prepare the node for the next operation
		temp.Next = &ListNode{}
		temp = temp.Next
	}

	// n == 1 means that the last addition operation is carried, and another node needs to be added.
	if n == 1 {
		temp.Next = &ListNode{Val: n}
	}

	return result
}

// next goes to the next digit of l.
func next(l *ListNode) *ListNode {
	if  l  ! =  nil {
		return l.Next
	}
	return nil
}

func  add ( n1 , n2  * ListNode , I  int ) ( v , n  int ) {
	if  n1  ! =  nil {
		v  + =  n1 . Val
	}

	if n2 != nil {
		v  + =  n2 . Val
	}

	v  + =  i

	if v > 9 {
		v  - =  10
		n = 1
	}

	return
}
