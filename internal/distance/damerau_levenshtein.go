package distance

func DistanceDamerauLevenshtein(s1, s2 string) int {
	dp := make([][]int, len(s1)+1)

	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := 0; i < len(s1)+1; i++ {
		dp[i][0] = i
	}

	for i := 0; i < len(s2)+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			cost := 1
			if s1[i-1] == s2[j-1] {
				cost = 0
			}
			dp[i][j] = min(
				dp[i-1][j]+1,      // deletion
				dp[i][j-1]+1,      // insertion
				dp[i-1][j-1]+cost, // substitution
			)
			if i > 1 && j > 1 && s1[i-1] == s2[j-2] && s1[i-2] == s2[j-1] {
				dp[i][j] = min(dp[i][j], dp[i-2][j-2]+1) // transposition
			}
			// fmt.Println(dp)
		}
	}
	return dp[len(s1)][len(s2)]
}

func min(a int, others ...int) int {
	m := a
	for _, i := range others {
		if i < m {
			m = i
		}
	}
	return m
}
