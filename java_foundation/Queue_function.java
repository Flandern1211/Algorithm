import java.util.LinkedList;
import java.util.Queue;
public class Queue_function {
    public static void main(String[] args) {
        //初始化队列
        Queue<Integer> queue= new LinkedList<>();
        //队列的常用方法
        //增
        queue.offer(10); // 在队列尾部添加元素
        queue.offer(20); 
        queue.offer(30); 
        System.out.println(queue); // 输出队列内容

        //删
        queue.poll(); // 删除队头元素
        System.out.println(queue); // 输出队列内容

        //查
        System.out.println(queue.isEmpty()); // 判断队列是否为空
        System.out.println(queue.size()); // 获取队列的长度
        System.out.println(queue.peek()); // 获取队头元素但不删除
        
        


    }
        
        
       
}