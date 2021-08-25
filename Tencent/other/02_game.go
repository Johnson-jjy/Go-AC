//package main
//
//import "fmt"
//
///**
// * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
// *
// * 输入参数 Ax 角色A的血量，Ay 角色A 的攻击力，Az A的攻击CD，Aw 角色A的恢复力
//输入参数 Bx 角色B的血量，By 角色B 的攻击力，Bz B的攻击CD，Bw 角色B的恢复力
// * @param Ax int整型 角色A的血量上限
// * @param Ay int整型 角色A的攻击力
// * @param Az int整型 A的攻击CD
// * @param Aw int整型 角色A的每回合恢复血量值
// * @param Bx int整型 角色B的血量上限
// * @param By int整型 角色B的攻击力
// * @param Bz int整型 B的攻击CD
// * @param Bw int整型 角色B的每回合恢复血量值
// * @return int整型
//*/
//func PK( Ax int ,  Ay int ,  Az int ,  Aw int ,  Bx int ,  By int ,  Bz int ,  Bw int ) int {
//	// write code here
//	Aeach := Ay - (Az + 1) * Bw
//	Beach := By - (Bz + 1) * Aw
//	if Aeach <= 0 && Beach <= 0 {
//		return 4
//	}
//	if Aeach <= 0 {
//		return 2
//	}
//	if Beach <= 0 {
//		return 1
//	}
//	if Aeach == Beach && Ax == Bx {
//		return 3
//	}
//
//	curA := Ax
//	curB := Bx
//	tA := 0
//	tB := 0
//	flagA := false
//	flagB := false
//
//	for  {
//		// A攻击
//		curA = min(Ax, curA + Aw)
//		curB = min(Bx, curB + Bw)
//
//		if tA == 0 {
//			curB -= Ay
//			if curB <= 0 {
//				flagB = true
//			}
//		}
//		tA++
//		if tA > Az {
//			tA = 0
//		}
//
//		if tB == 0 {
//			curA -= By
//			if curA <= 0 {
//				flagA = true
//			}
//		}
//
//		tB++
//		if tB > Bz {
//			tB = 0
//		}
//
//		fmt.Println(curA, curB)
//		if flagA || flagB {
//			break
//		}
//	}
//
//	if flagA && flagB {
//		return 3
//	} else if flagA {
//		return 2
//	}
//	return 1
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//
//
//func main() {
//	fmt.Println(PK(25,10,1,3,30,5,0,4))
//}