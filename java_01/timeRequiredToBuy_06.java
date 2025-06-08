//2073

import java.util.LinkedList;
import java.util.Queue;

public class timeRequiredToBuy_06{

}



//想的有问题，感觉这样很难解出来，想直接把数组全部放到队列里实现，变化太多了
class Solution1 {
    public int timeRequiredToBuy(int[] tickets, int k) {
          Queue<Integer> q = new LinkedList<>();
          int times=0;
          //while(tickets[k]){
            if((q.peek()-1)!=0&&!q.isEmpty())
            {
                q.offer(q.peek()-1);
                q.poll();
                times++;
                k--;
            }
            else{
                q.poll();
            }

          }
        
    }


//队列实现，但队列中是索引
//时间复杂度：O(n)
class Solution2 {
    public int timeRequiredToBuy(int[] tickets, int k) {
       Queue<Integer> index = new LinkedList<>();
        for(int i=0;i<tickets.length;i++)//将所有索引依次填入index队列
        {
            index.offer(i);
        }
        int times=0;
        while(tickets[k]!=0)//循环条件，k索引的票数不为零
        {
            if(tickets[index.peek()]>1&&!index.isEmpty())
            //大于1才行，大于0的条件判断会让票数已经为0的索引再到队尾一次
            {
               
                (tickets[index.peek()])--;
                index.offer(index.peek());
                index.poll();
                times++;
            }
            else{
                (tickets[index.peek()])--;
                index.poll();
                times++;
            }
        }
        return times;
    }
    
}

//
class Solution3 {
    public int timeRequiredToBuy(int[] tickets, int k) {
      Queue<Integer> index = new LinkedList<>();
       for(int i=0;i<tickets.length;i++)//将所有索引依次填入index队列
        {
            index.offer(i);
        }
        int times=0;
        while(!index.isEmpty())
        {
            int font=index.poll();
            times++;
            tickets[font]--;
            if(font==k&&tickets[font]==0)
            {
               return times;
            }
            if(tickets[font]==0){
                continue;
            }
            index.offer(font);

        }
        return times; // 语法占位符（实际在循环内部返回）
    }
}