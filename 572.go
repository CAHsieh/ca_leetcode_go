package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if nil == s && nil == t {
		return true
	}
	if nil == s || nil == t {
		return false
	}

	same := isSame(s, t)
	if !same {
		same = isSubtree(s.Left, t)
		if !same {
			same = isSubtree(s.Right, t)
		}
	}

	return same
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if nil == s && nil == t {
		return true
	} else if nil == s || nil == t {
		return false
	}
	if s.Val != t.Val {
		return false
	} else {
		return (isSame(s.Left, t.Left) && isSame(s.Right, t.Right))
	}
}
