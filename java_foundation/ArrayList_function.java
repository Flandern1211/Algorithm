//动态数组

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
public class ArrayList_function {
    public static void main(String[] args) {
        ArrayList<Integer> nums= new ArrayList<>();
        ArrayList<Integer> nums2= new ArrayList<>(Arrays.asList(1,3,5,7,9));
        //ArrayList的常用方法

        //初始化ArrayList,指定大小10和初始值全为0
        ArrayList<Integer> nums3= new ArrayList<>(Collections.nCopies(10,0));

        //在数值尾部添加元素
        nums3.add(2);
        System.out.println(nums3);

        //在指定位置添加元素
        nums3.add(1,4);
        System.out.println(nums3);

        //获取数组的长度
        System.out.println(nums3.size());

        //判断数组是否为空
        System.out.println(nums3.isEmpty());
        
        //获取指定位置的元素
        System.out.println(nums3.get(1));

        // 根据元素值查找索引，时间复杂度 O(N)
        int index = nums3.indexOf(666);

        //获取最后一位的元素
        System.out.println(nums3.get(nums3.size()-1));

        //删除最后一位元素
        nums3.remove(nums3.size()-1);
        System.out.println(nums3);

        //修改指定位置的值
         System.out.println("-----------------------");
        nums3.set(2,11);
        System.out.println(nums3.set(2,11));
        System.out.println(nums3);

        //交换指定位置数值
        Collections.swap(nums3, 0, 1);
        System.out.println(nums3);


        // 遍历数组
        for(int num : nums3) {
            System.out.println(num + " ");
        }
        System.out.println();
        System.out.println(nums3);
    }
}