package main

func GetNumOfExpress(express string, desired bool) int64 {
	// write code here
	if !checkIfValid(express) {
		return 0
	}

	n := len(express)
	//store := []byte(express)
	arr1 := make([][]int64, n)
	arr2 := make([][]int64, n)

	for i := 0; i < n; i++ {
		arr1[i] = make([]int64, n)
		arr2[i] = make([]int64, n)
	}

	if express[0] == '1' {
		arr1[0][0] = 1
		arr2[0][0] = 0
	} else {
		arr1[0][0] = 0
		arr2[0][0] = 1
	}

	for j := 2; j < n; j += 2 {
		if express[j] == '1' {
			arr1[j][j] = 1
			arr2[j][j] = 0
		} else {
			arr1[j][j] = 0
			arr2[j][j] = 1
		}

		for i := j - 2; i >= 0; i -= 2 {
			for k := j - 1; k > i; k -= 2 {
				switch express[k] {
				case '&':
					arr1[i][j] += arr1[i][k - 1] * arr1[k + 1][j]
					arr2[i][j] += (arr2[i][k - 1] * arr2[k + 1][j]) + (arr1[i][k - 1] * arr2[k + 1][j]) + (arr2[i][k - 1] * arr1[k + 1][j])
				case '|':
					arr1[i][j] += arr1[i][k - 1] * arr2[k + 1][j] + arr2[i][k - 1] * arr1[k + 1][j] + arr1[i][k - 1]*arr1[k + 1][j]
					arr2[i][j] += arr2[i][k - 1] * arr2[k + 1][j]
				case '^':
					arr1[i][j] += arr1[i][k - 1] * arr2[k + 1][j] + arr2[i][k - 1] * arr1[k + 1][j]
					arr2[i][j] += arr1[i][k - 1] * arr1[k + 1][j] + arr2[i][k - 1] * arr2[k + 1][j]
				}

			}
		}
	}

	if desired {
		return arr1[0][n - 1]
	} else {
		return arr2[0][n - 1]
	}
}

func checkIfValid(express string) bool {
	if len(express) == 0 {
		return false
	}
	if len(express) % 2 == 0 {
		return false
	}

	for i := 0; i < len(express); i += 2 {
		if express[i] < '0' || express[i] > '1' {
			return false
		}
		if i + 1 < len(express) && express[i + 1] != '&' && express[i + 1] != '|' && express[i + 1] != '^' {
			return false
		}
	}

	return true
}

//func getNum(express string, desired bool) int {
//
//}


//arr1[i][j] %= base
//arr2[i][j] %= base