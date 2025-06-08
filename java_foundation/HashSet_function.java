import java.util.Arrays;
import java.util.HashSet;

public class HashSet_function {
    public static void main(String[] args) {
        // 初始化哈希集合
        HashSet<Integer> hashset = new HashSet<>(Arrays.asList(1, 2, 3, 4));

        // 检查哈希集合是否为空，输出：false
        System.out.println(hashset.isEmpty());

        // 获取哈希集合的大小，输出：4
        System.out.println(hashset.size());

        // 查找指定元素是否存在
        // 输出：Element 3 found.
        if(hashset.contains(3)) {
            System.out.println("Element 3 found.");
        } else {
            System.out.println("Element 3 not found.");
        }

        // 插入一个新的元素
        hashset.add(5);

        // 删除一个元素
        hashset.remove(2);
        // 输出：Element 2 not found.
        if(hashset.contains(2)) {
            System.out.println("Element 2 found.");
        } else {
            System.out.println("Element 2 not found.");
        }

        // 遍历哈希集合
        // 输出（顺序可能不同）：
        // 1
        // 3
        // 4
        // 5
        for(int element : hashset) {
            System.out.println(element);
        }
    }
}