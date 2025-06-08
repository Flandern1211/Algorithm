
import java.util.HashMap;

public class SingleNumber_03{

}

//1.双重循环
//时间复杂度：O(n^2)
class Solution1 {
    public int singleNumber(int[] nums) {
        for (int i = 0; i < nums.length; i++) {
            boolean flag=false;
            for(int j=0;j<nums.length;j++)//不能从i+1开始，因为如果前面有重复的检查不出来
            {
                if(nums[i]==nums[j]&&i!=j)//将i!=j放在if判断里而不是直接在for循环里，因为直接放在for里面第一个数据j==i永远成立，for直接结束，无法往下进行判断
                {
                    flag=true;//flag标识该数据已有重复
                    break;
                }
                
            }
            if(!flag)
                {
                    return nums[i];
                }
        }
       return -1; 
    }
}

//哈希表实现
//时间复杂度：O(n)
class Solution2 {
    public int singleNumber(int[] nums) {
        HashMap<Integer,Integer> count= new HashMap<>();
        for(int num:nums)
        {
            count.put(num,count.getOrDefault(num, 0)+1);//getOrDefault方法：查找num，找不到返回0，找到返回num
            //这里就是先遍历使所有数字次数皆记为0+1，然后再次出现的会变成1+1=2
        }
        for(int num:nums)
        {
            if(count.get(num)==1)
            {
                return num;
            }
        }
        return -1;
    }
}