package warcalc

import (
	"freecocoa/src/models"
	"math"
)

func noNeg(x int) int {
	if x >= 0 {
		return x
	}

	return 0
}

func createRange(n1, n2 int) []int {
	elems := n2 - n1

	if elems <= 0 {
		return []int{}
	}

	arr := make([]int, elems)

	for i := 0; i < elems; i++ {
		arr[i] = n1 + i
	}

	return arr
}

func newt(n, k int) int {
	l1 := 1

	for _, i := range createRange(k+1, n+1) {
		l1 *= i
	}

	for _, i := range createRange(1, n-k+1) {
		l1 /= i
	}

	return l1
}

func remhp(atthp, defhp, attfp, deffp int, attpo, defpo float64) float64 {
	p := attpo / (attpo + defpo)
	k := int(float64(defhp / attfp))
	l := int(float64(atthp / deffp))

	x := 0.0

	for hp := 0; hp < k; hp++ {
		y := float64(newt(hp+l, hp)) * math.Pow(p, float64(hp)) * math.Pow(1-p, float64(l))
		x += y
	}

	y := 0.0

	for al := 0; al < l; al++ {
		y += float64(newt(k-1+al, al)) * math.Pow(p, float64(k-1)) * math.Pow(1-p, float64(al)) * p
	}

	x += y
	_ = x

	return y
}

func winprob(p float64, hp1, hp2, fp1, fp2 int) float64 {
	dp := make([][]float64, hp1+1)
	for i := range dp {
		dp[i] = make([]float64, hp2+1)
	}

	for i := 0; i < hp1; i++ {
		for j := 0; j < hp2; j++ {
			if i == 0 {
				dp[i][j] = 0
				continue
			}

			if j == 0 {
				dp[i][j] = 1
				continue
			}

			dp[i][j] = p*dp[noNeg(i-fp2)][j] + (1-p)*dp[i][noNeg(j-fp1)]
		}
	}

	return 100 * dp[hp1][hp2]
}

func exphp(lp float64, hp1, hp2, fp1, fp2 int) float64 {
	dh := make([][]float64, hp1+1)
	for i := range dh {
		dh[i] = make([]float64, hp2+2)
	}

	for i := 0; i < hp1; i++ {
		for j := 0; j < hp2; j++ {
			if i == 0 {
				dh[i][j] = 0
				continue
			}

			if j == 0 {
				dh[i][j] = float64(i)
				continue
			}

			dh[i][j] = lp*dh[noNeg(i-fp2)][j] + (1-lp)*dh[i][noNeg(j-fp1)]
		}
	}

	return dh[hp1][hp2]
}

func Warcalc(avd models.FinalStats) models.CombatResults {
	astr := avd.Attacker.AP
	ahp := avd.Attacker.HP
	afp := avd.Attacker.FP
	dstr := avd.Defender.DP
	dhp := avd.Defender.HP
	dfp := avd.Defender.FP
	p := astr / (astr + dstr)
	tab := make([]models.CombatResult, 0, 20)
	ndhp := (dhp)

	for i := 0; i < 20; i++ {
		defprob := winprob(p, ndhp, ahp, dfp, afp)
		attprob := 100 - defprob
		dexphp := exphp(p, ndhp, ahp, dfp, afp)
		aexphp := exphp(1.0-p, ahp, ndhp, afp, dfp)
		delta := float64(ndhp) - dexphp
		results := models.CombatResult{
			AttProb:  attprob,
			AttHP:    aexphp,
			DefProb:  defprob,
			DefHP:    dexphp,
			DefDelta: delta,
		}

		tab = append(tab, results)

		ndhp := int(dexphp)
		if ndhp == 0 {
			break
		}
	}

	prob := remhp(ahp, dhp, afp, dfp, astr, dstr)

	return models.CombatResults{
		Prob:    prob,
		Results: tab,
	}
}
