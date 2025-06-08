import java.util.Stack;

public class Stack_function {
    public static void main(String[] args) {
        // 初始化一个空的整型栈 s
        Stack<Integer> s = new Stack<>();

        // 向栈顶添加元素
        s.push(10);
        s.push(20);
        s.push(30);

        // 检查栈是否为空，输出：false
        System.out.println(s.isEmpty());

        // 获取栈的大小，输出：3
        System.out.println(s.size());

        // 获取栈顶元素，输出：30
        System.out.println(s.peek());

        // 删除栈顶元素
        s.pop();

        // 输出新的栈顶元素：20
        System.out.println(s.peek());
    }
}