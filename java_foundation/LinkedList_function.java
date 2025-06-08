import java.util.Arrays;
import java.util.LinkedList;
public class LinkedList_function {
    public static void main(String[] args) {
        // 初始化LinkedList
        LinkedList<Integer> lst=new LinkedList<>(Arrays.asList(1,3,5,7,9));

         // LinkedList的常用方法
         //增
         lst.add(2); // 在数值尾部添加元素
         System.out.println(lst);
         lst.addFirst(4); // 在首部添加元素
         System.out.println(lst);
         lst.addLast(6); // 在尾部添加元素
         System.out.println(lst);
         lst.add(2, 8); // 在指定位置添加元素
         System.out.println(lst);

         //删
         lst.remove(2); // 删除指定位置的元素
         System.out.println(lst);
         lst.removeFirst(); // 删除首部元素
         System.out.println(lst);
         lst.removeLast(); // 删除尾部元素
         System.out.println(lst);

         //改
        lst.set(1, 10); // 修改指定位置的值
        System.out.println(lst);

        //查
        System.out.println(lst.isEmpty()); // 判断LinkedList是否为空
        System.out.println(lst.size()); // 获取LinkedList的长度
        System.out.println(lst.get(1)); // 获取指定位置的元素
        System.out.println(lst.getFirst()); // 获取首部元素
        System.out.println(lst.getLast()); // 获取尾部元素

        //遍历
        for(int val : lst) {
            System.out.print(val + " ");
        }
        System.out.println();
    }
}
    
