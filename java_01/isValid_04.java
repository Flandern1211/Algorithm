//20 

import java.util.HashMap;
import java.util.Stack;

public class isValid_04{

}



//大错特错

//无法处理顺序问题：只检查数量相等无法确保括号顺序正确（例如"([)]"会误判为有效）

//错误的字符串比较：使用==比较字符串（应使用.equals()），但更根本的是应该使用字符(char)处理

//逻辑错误：使用if-else if结构导致只要任意一种括号匹配就返回true

//类型错误：字符串不能直接遍历为String（应使用char）
class Solution2 {
    public boolean isValid(String s) {
        HashMap<Character,Integer> count=new HashMap<>();
        char a1='(';
        char b1='[';
        char c1='{';
        char a2=')';
        char b2=']';
        char c2='}';
        
        for (char s1 : s.toCharArray())
        {
            if(s1==a1)
            {
                count.put(a1,count.getOrDefault(a1,0)+1 );
            }
             if(s1==b1)
            {
                count.put(b1,count.getOrDefault(b1,0)+1 );
            }
             if(s1==c1)
            {
                count.put(c1,count.getOrDefault(c1,0)+1 );
            }
            if(s1==a2)
            {
                count.put(a2,count.getOrDefault(a2,0)+1 );
            }
             if(s1==b2)
            {
                count.put(b2,count.getOrDefault(b2,0)+1 );
            }
             if(s1==c2)
            {
                count.put(c2,count.getOrDefault(c2,0)+1 );
            }
        }
        if(count.get(a1)==count.get(a2)&&count.get(b1)==count.get(b2)&&count.get(c1)==count.get(c2))
        {
            return true;
        }
       
        return false;
    }
}


//2.栈
//时间复杂度：O（n）
//先顾后再看前
class Solution1 {
    public boolean isValid(String s) {
        Stack<Character> stk=new Stack<>();
        for(char c:s.toCharArray())//遍历字符串的正确方法，没有字符遍历就不会进行，没必要考虑这种情况
        {
            if(c=='('||c=='['||c=='{')
            {
                stk.push(c);
            }
            else{
                if(!stk.empty()&&stkof(c)==stk.peek())//
                stk.pop();
                else{
                    return false;
                }
            }
        }
        return stk.isEmpty() ; //栈中还存在元素，说明没有与之对应的括号
    }
    

    char stkof(char c)//秒呀。根据函数的返回特性来进行  对括号对应的解
    {
    if(c==')') return '(';
    if(c==']') return '[';
    return '}';
    }
}


