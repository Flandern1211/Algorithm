import java.util.HashMap;
import java.util.HashSet;
public class ContainsDuplicate_02{

}

//1 双循环
//时间复杂度：O(n^2)
//测试案例太多就不通过
class Solution1 {
    public boolean containsDuplicate(int[] nums) {
        for(int  i=0;i<nums.length;i++)
        {
            for(int j=i+1;j<nums.length;j++)
            {
                if(nums[i]==nums[j])
                {
                    return true;
                }
            }
        }
        return false;
    }
}

//2 哈希集合
//时间复杂度：O(n)
class Solution2 {
    public boolean containsDuplicate(int[] nums) {
        HashSet<Integer> set=new HashSet<>();
        for(int num:nums)
        {
            if(set.contains(num)){
                return true;
            }
            set.add(num);
        }
        return false;
    }
}


//哈希表
//时间复杂度：O(n)
class Solution3 {
    public boolean containsDuplicate(int[] nums) {
        HashMap<Integer,Integer> count= new HashMap<>();
        for(int num:nums)
        {
            count.put(num,count.getOrDefault(num, 0)+1);//getOrDefault方法：查找num，找不到返回0，找到返回num
            //这里就是先遍历使所有数字次数皆记为0+1，然后再次出现的会变成多个
        }
        for(int num:nums)
        {
            if(count.get(num)>1)//重复可能有很多，没有确定数
            {
                return true;
            }
        }
        return false;
    }
}