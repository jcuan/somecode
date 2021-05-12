## 思路

字符串可能的种类256种
```
1. 分配一个长度为256长度的字符串locations，用来记录某个字符在字符串中出现的位置， 输入为字符串为s
2. subLen初始化为0，用来记录子串长度，resLen用来记录当前查到的最长无重复字符子串的长度。
    循环遍历字符串中的每个字符s[0],..,s[i],...,s[strLen]，记录查找开始的位置subStart，如果某个字符s[i]在location中有记录, 得到该记录n=locations[s[i]]（代表在字符串中的位置）
    如果n<subStart，说明字符出现的位置在本次子串查找的前边，更新s[i]=n，subLen=subLen+1，继续比较下一个字符s[i+1]
    如果n>=subStart， 说明字符出现在本次子串的查找中，也就是子串s[subStart],...,s[i]中包含了重复的字符
        如果subLen>resLen，则resLen=subLen。
        subStart=i， 
```