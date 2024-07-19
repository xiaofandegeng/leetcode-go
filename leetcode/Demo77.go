package main

var (
	path []int
	res  [][]int
)

func combine(n int, k int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	dfs(n, k, 1)
	return res
}

func dfs(n int, k int, start int) {
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i <= n; i++ {
		if n-i+1 < k-len(path) {
			break
		}
		path = append(path, i)

	}
}
