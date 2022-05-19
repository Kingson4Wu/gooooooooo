package sort

//快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。

//选择一个元素作为“基准”(pivot)

func QuickSortTopk_1(nums []int, k int) []int {
	quickSortTopk(nums, 0, len(nums)-1, k)
	return nums[0:k]
}

func quickSortTopk(nums []int, start, end int, k int) {
	if start >= end {
		return
	}

	mid := partition2(nums, start, end)

	if (mid + 1) > k {
		quickSortTopk(nums, start, mid, k)

	} else if (mid + 1) == k {
		return

	} else {
		quickSortTopk(nums, mid+1, end, k)
	}

}

func partition2(nums []int, start, end int) int {
	temp := nums[start]
	low := start
	high := end

	for low < high {
		for low < high && temp <= nums[high] {
			high--
		}
		if low < high {
			nums[low] = nums[high]
		}

		for low < high && temp > nums[low] {
			low++
		}

		if low < high {
			nums[high] = nums[low]
		}
	}

	nums[low] = temp

	return low
}

/**
运行时间：4ms
超过26.90% 用Go提交的代码
占用内存：1260KB
超过9.65%用Go提交的代码
*/

//原文链接：https://blog.csdn.net/guidao13/article/details/86430483

/**

class Solution {
public:
    int partition(vector<int> &input, int l, int r) {
        int pivot = input[r-1];
        int i = l;
        for (int j=l; j<r-1; ++j) {
            if (input[j] < pivot) {
                swap(input[i++], input[j]);
            }
        }
        swap(input[i], input[r-1]);
        return i;

    }
    vector<int> GetLeastNumbers_Solution(vector<int> input, int k) {
        vector<int> ret;
        if (k==0 || k > input.size()) return ret;
         int l = 0, r = input.size();
        while (l < r) {
            int p = partition(input, l, r);
            if (p+1 == k) {
                return vector<int>({input.begin(), input.begin()+k});
            }
            if (p+1 < k) {
                l = p + 1;
            }
            else {
                r = p;
            }

        }
        return ret;
    }
};
*/
