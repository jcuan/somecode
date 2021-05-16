class Solution:
    def fib0(self, n: int) -> int:
        data = [0, 1]
        if n < 2:
            return data[n]
        for index in range(2, n+1):
            data[index%2] = (data[0] + data[1])
            index += 1
        if n % 2 == 0:
            return data[0] % 1000000007
        else:
            return data[1] % 1000000007

    def fib(self, n: int) -> int:
        a = 0
        b = 1
        for _ in range(n):
            a, b = b, a+b
        return a % 1000000007

    def numWays(self, n: int) -> int:
        a=1
        b=1
        for _ in range(n):
            a,b = b,a+b
        return a % 1000000007
        
if __name__ == '__main__':
    datas = [0,1,2,3]
    for data in datas:
        print(Solution().numWays(data))