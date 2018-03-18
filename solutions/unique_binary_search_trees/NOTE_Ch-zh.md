# 解法要点
1. 假设n=4：dp[0]=1，dp[1]=1，dp[2]=2，dp[3]=5，dp[4]=14，如何得到dp[5]是核心问题
1. 构建一个tree，选择其中一个节点，得出在该节点下可以保留的左left和右left，并将它们相乘
1. 例如选择3，dp[3]=dp[0]*dp[1]+dp[0]*dp[2]+dp[1]*dp[2]=1+2+2=5