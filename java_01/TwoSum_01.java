public class TwoSum_01 {
    public static void main(String[] args) {
        
    }
   

}

//1.双循环 
//时间复杂度：O(n^2)
class Solution {
    public int[] twoSum(int[] nums, int target) {
        for(int i=0;i<nums.length;i++)
        {
            for(int j=i+1;j<nums.length;j++)
            {
                if(nums[i]+nums[j]==target)
                {
                    return new int[]{i,j};
                }
            }
        }
    }
}