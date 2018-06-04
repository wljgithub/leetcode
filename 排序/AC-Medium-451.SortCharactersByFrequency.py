'''
Given a string, sort it in decreasing order based on the frequency of characters.

Example 1:

Input:
"tree"

Output:
"eert"

Explanation:
'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
Example 2:

Input:
"cccaaa"

Output:
"cccaaa"

Explanation:
Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
Note that "cacaca" is incorrect, as the same characters must be together.
Example 3:

Input:
"Aabb"

Output:
"bbAa"

Explanation:
"bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.
'''
#字母区分大小写，出现次数相同次序无所谓

class Solution(object):
	def frequencySort(self, s):
		"""
		:type s: str
		:rtype: str
		"""
		#用桶排序，先将数据装进桶里，根据桶里的数据进行排序，然后再组合桶里的数据

		#把数据装进桶里
		d = {}
		for i in s:
			if i in d:
				d[i] += 1
			else:
				d[i] = 1

		#根据桶里的数据个数进行排序
		new_d=sorted(zip(d.keys(),d.values()),key=lambda x :x[1],reverse=True)

	    #将桶里的数据按个数的大小进行组合
		output=''
		for i in new_d:
			output+=i[0]*i[1]
		return output


o=Solution()
print(o.frequencySort('tree'))


