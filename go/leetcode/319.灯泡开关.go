/*
 * @lc app=leetcode.cn id=319 lang=golang
 *
 * [319] 灯泡开关
 *
 * https://leetcode-cn.com/problems/bulb-switcher/description/
 *
 * algorithms
 * Medium (43.63%)
 * Likes:    69
 * Dislikes: 0
 * Total Accepted:    6K
 * Total Submissions: 13.7K
 * Testcase Example:  '3'
 *
 * 初始时有 n 个灯泡关闭。 第 1 轮，你打开所有的灯泡。 第 2 轮，每两个灯泡你关闭一次。 第 3
 * 轮，每三个灯泡切换一次开关（如果关闭则开启，如果开启则关闭）。第 i 轮，每 i 个灯泡切换一次开关。 对于第 n 轮，你只切换最后一个灯泡的开关。
 * 找出 n 轮后有多少个亮着的灯泡。
 * 
 * 示例:
 * 
 * 输入: 3
 * 输出: 1 
 * 解释: 
 * 初始时, 灯泡状态 [关闭, 关闭, 关闭].
 * 第一轮后, 灯泡状态 [开启, 开启, 开启].
 * 第二轮后, 灯泡状态 [开启, 关闭, 开启].
 * 第三轮后, 灯泡状态 [开启, 关闭, 关闭]. 
 * 
 * 你应该返回 1，因为只有一个灯泡还亮着。
 * 
 * 
 */

/* 
 * 注意审题，是n个灯泡经过n轮，先取一个n看看，比如n=16：
 * 第1个灯泡，第1轮改变状态，一直亮着
 * 第2个灯泡，第1、2轮改变状态，灭掉
 * 第3个灯泡，第1、3轮改变状态，灭掉
 * 第4个灯泡，第1、2、4轮改变状态，亮着
 * ...
 * 第13个灯泡，第1、13轮会切换状态，最后灭了
 * 第14个灯泡，第1、2、7、14轮会切换状态，最后灭了
 * 第15个灯泡，第1、3、5、15轮会切换状态，最后灭了
 * 第16个灯泡，第1、2、4、8、16轮会切换状态，也就是找出所有因子，最后是亮着到的
 * 对于第k个灯泡来说，也就是找出1~k的因子，数量为奇数则亮着
 * 再看下规律，只有1、4、9、16亮着，也就是完全平方数才亮着，就改为求n以内的完全平方数数量
 */


// @lc code=start
func bulbSwitch(n int) int {
	res := 1
	for res * res <= n {
		res++
	}
	return res - 1
}
// @lc code=end

