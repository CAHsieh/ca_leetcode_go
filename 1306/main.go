package main

import "fmt"

func main() {
	// fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
	// fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0))
	// fmt.Println(canReach([]int{3, 0, 2, 1, 2}, 2))
	// fmt.Println(canReach([]int{0, 1}, 1))
	// fmt.Println(canReach([]int{4, 4, 1, 3, 0, 3}, 2))
	fmt.Println(canReach([]int{719, 622, 532, 746, 476, 295, 285, 472, 712, 283, 808, 140, 730, 334, 215, 509, 573, 121, 54, 430, 791, 41, 351, 548, 38, 108, 415, 490, 393, 183, 798, 423, 112, 172, 791, 195, 68, 489, 803, 703, 248, 705, 213, 757, 473, 382, 334, 693, 6, 414, 223, 352, 718, 727, 403, 702, 13, 171, 256, 71, 645, 94, 159, 83, 513, 119, 10, 33, 64, 179, 635, 492, 87, 133, 767, 781, 182, 289, 636, 155, 729, 216, 64, 728, 649, 802, 149, 321, 179, 662, 195, 143, 299, 7, 630, 33, 527, 706, 726, 752, 755, 101, 732, 663, 794, 24, 799, 780, 438, 707, 482, 237, 252, 107, 659, 527, 652, 636, 48, 388, 405, 42, 247, 191, 654, 324, 6, 314, 649, 249, 289, 382, 137, 808, 455, 774, 571, 789, 176, 634, 762, 266, 799, 54, 126, 41, 681, 802, 157, 148, 745, 265, 777, 436, 233, 455, 337, 606, 239, 338, 508, 322, 210, 482, 534, 245, 618, 589, 567, 639, 355, 736, 534, 113, 588, 240, 795, 367, 245, 249, 641, 783, 701, 469, 521, 518, 59, 528, 250, 634, 135, 13, 645, 739, 531, 102, 36, 291, 22, 541, 482, 153, 533, 664, 559, 784, 707, 28, 297, 630, 3, 606, 237, 216, 39, 793, 517, 194, 92, 506, 63, 526, 55, 504, 295, 185, 110, 35, 9, 527, 8, 54, 259, 498, 229, 684, 579, 619, 409, 330, 187, 60, 112, 180, 477, 24, 313, 190, 180, 807, 115, 788, 238, 599, 464, 160, 464, 662, 809, 300, 788, 658, 137, 630, 363, 321, 706, 434, 358, 534, 257, 195, 226, 473, 191, 223, 282, 518, 378, 339, 34, 644, 231, 523, 547, 544, 491, 263, 683, 528, 797, 587, 753, 445, 450, 783, 537, 249, 374, 546, 662, 149, 394, 202, 571, 562, 524, 587, 606, 645, 100, 193, 37, 329, 650, 92, 462, 131, 623, 510, 257, 118, 434, 493, 721, 748, 280, 248, 515, 232, 41, 166, 644, 112, 455, 760, 633, 50, 488, 589, 611, 117, 649, 791, 385, 67, 200, 305, 285, 733, 471, 468, 755, 582, 359, 543, 366, 206, 74, 24, 20, 395, 192, 199, 801, 33, 658, 474, 341, 500, 781, 367, 714, 576, 669, 327, 508, 325, 796, 608, 38, 656, 710, 219, 59, 481, 735, 475, 425, 400, 690, 541, 806, 488, 246, 735, 310, 762, 454, 15, 550, 74, 289, 351, 84, 486, 81, 161, 341, 500, 629, 224, 364, 182, 309, 530, 539, 713, 116, 511, 428, 392, 524, 681, 599, 120, 658, 266, 592, 184, 76, 160, 284, 490, 771, 74, 398, 336, 155, 502, 356, 268, 427, 98, 12, 232, 383, 381, 563, 10, 634, 669, 650, 79, 298, 734, 730, 803, 201, 142, 35, 704, 405, 788, 678, 171, 407, 314, 770, 697, 741, 649, 227, 346, 80, 790, 620, 504, 306, 601, 764, 490, 115, 266, 657, 463, 475, 116, 390, 396, 538, 178, 130, 602, 496, 196, 56, 382, 252, 663, 696, 343, 775, 341, 427, 69, 242, 354, 658, 568, 281, 21, 625, 3, 499, 551, 569, 744, 0, 398, 586, 645, 32, 600, 537, 477, 679, 335, 779, 405, 540, 563, 443, 629, 477, 164, 591, 21, 719, 255, 241, 649, 602, 247, 713, 218, 349, 599, 53, 55, 773, 187, 529, 460, 621, 558, 56, 699, 335, 666, 177, 354, 607, 145, 580, 529, 367, 678, 712, 756, 405, 52, 169, 144, 275, 95, 496, 45, 705, 378, 385, 6, 795, 45, 463, 63, 511, 222, 81, 683, 807, 468, 142, 125, 697, 238, 358, 765, 195, 747, 636, 504, 451, 121, 544, 692, 5, 774, 89, 357, 240, 48, 673, 443, 539, 632, 111, 224, 575, 22, 108, 277, 85, 0, 456, 783, 410, 519, 27, 500, 570, 35, 576, 231, 293, 463, 307, 229, 341, 36, 274, 262, 170, 709, 232, 149, 156, 223, 797, 408, 562, 796, 394, 320, 324, 710, 520, 654, 12, 674, 617, 432, 365, 379, 250, 217, 699, 267, 197, 354, 423, 365, 312, 253, 535, 174, 800, 430, 320, 217, 652, 129, 650, 448, 387, 399, 390, 185, 709, 539, 241, 474, 70, 756, 45, 616, 397, 317, 252, 372, 48, 306, 21, 554, 725, 186, 422, 717, 392, 683, 810, 752, 369, 421, 537, 13, 736, 144, 79, 536, 177, 648, 231, 788, 677, 428, 446, 571, 162, 562, 147, 100, 416, 723, 49, 731, 727, 625, 644, 43, 688, 0, 399, 8, 174, 281, 137, 751, 116, 256, 452, 427, 129, 602, 277, 169, 213, 108, 595, 803, 412, 727, 81, 648, 334, 433, 659, 519, 273, 103, 184, 758, 652, 776, 605, 385, 603, 318, 577, 185, 584, 491, 684, 251, 604, 560, 567, 515, 112, 256, 446, 116, 741, 801, 338, 581, 571, 465, 78, 403, 10, 566, 264, 247, 43, 808, 123, 267, 271, 180, 666}, 734))
}

// func canReach(arr []int, start int) bool {

// 	queue := []int{start}
// 	seen := make(map[int]bool)
// 	seen[start] = true

// 	for len(queue) > 0 {
// 		current := queue[0]

// 		if arr[current] == 0 {
// 			return true
// 		}

// 		queue = queue[1:]

// 		left := current - arr[current]
// 		right := current + arr[current]

// 		_, seenLeft := seen[left]
// 		_, seenRight := seen[right]

// 		if left >= 0 && !seenLeft {
// 			queue = append(queue, left)
// 			seen[left] = true
// 		}

// 		if right < len(arr) && !seenRight {
// 			queue = append(queue, right)
// 			seen[right] = true
// 		}

// 	}

// 	return false
// }

func canReach(arr []int, start int) bool {
	return dfs(arr, []int{}, start)
}
func dfs(arr, traveled []int, current int) bool {

	if arr[current] == 0 {
		return true
	}

	left := current - arr[current]
	right := current + arr[current]

	leftResult := false
	if left >= 0 {

		isTraveled := false

		for _, t := range traveled {
			if left == t {
				isTraveled = true
				break
			}
		}

		if !isTraveled {
			nTraveled := make([]int, len(traveled)+1)
			copy(nTraveled, append(traveled, left))
			leftResult = dfs(arr, nTraveled, left)
		}
	}

	rightResult := false
	if right < len(arr) {

		isTraveled := false

		for _, t := range traveled {
			if right == t {
				isTraveled = true
				break
			}
		}
		if !isTraveled {
			nTraveled := make([]int, len(traveled)+1)
			copy(nTraveled, append(traveled, right))
			rightResult = dfs(arr, nTraveled, right)
		}
	}

	return leftResult || rightResult
}
