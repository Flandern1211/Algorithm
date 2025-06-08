import java.util.NoSuchElementException;

public class CycleArray01<T>{
    private T[] arr;
    private int size;
    private int start;
    private int end;
    private int count;

    public CycleArray01(){
        this(1);
    }

    public CycleArray01(int size){
        this.size=0;
        this.arr=(T[]) new Object[size];
        this.start=0;
        this.end=0;
        this.count=0;
    }

    //******************************************************* */
    //自定义扩缩容机制
    public void resize(int newSize){
        //创建一个新的自定义容量的泛型数组
        T[] newArr = (T[]) new Object[newSize];
        for(int i=0;i<count;i++){
            //start%size=start;加一个i使其变换成=start+0，start+1 start+2等等
            //使用%size的作用是使其成为环形，当start+i=size时自动回到0索引处并且会继续往下加1索引
            //索引范围始终为[0-size-1],正好可以使用%size实现
            //******************************** */
            newArr[i]=arr[(start+i)%size];
            //******************************** */
        }
        arr=newArr;

        start=0;
        //在新的数组中头索引数据已经在0处在循环中
        end=count;
        //为索引因为newSize肯定比size大，直接根据长度即可得出
        size=newSize;
    //******************************************************* */

    }
    //判断是否数组满了没，用在add中
    public boolean isFull(){
        return count==size;
    }

    public void isEmpty(){
        if(size==0)
        {
            throw new  NoSuchElementException();
        }
    }

    public void addFirst(T e){
        if(isFull()) resize(2*size); 
        //先往左移一位，start-1，然后因为是环形数组要在start=0时往左移就为添加到最后一位为，发现0-1=-1加size使其成为size-1正好为最后一位的索引，，并且这个逻辑始终使得新的start在旧的前一位
        start=(start-1+size)%size;
        arr[start]=e;      
        count++;   
    }

    public void addLast(T e){
        if(isFull()) resize(2*size); 
        //先往左移一位，start-1，然后因为是环形数组要在start=0时往左移就为添加到最后一位为，发现0-1=-1加size使其成为size-1正好为最后一位的索引，，并且这个逻辑始终使得新的start在旧的前一位
        arr[end]=e; 
        end=(end+1)%size;     
        count++;

      
    }

    public T removeFirst(){
          isEmpty();

        T val=arr[start];
        arr[start]=null;
        start=(start+1)%size;
        count--;
        if(count>0&&count/size<=4)
            resize(size/2);
        return val;
    
    } 

    public T removeLast(){
          isEmpty();

       
        end=(start-1+size)%size;//end是最后元素的后一位，先将其左移，再将其标志位null便可以删除最后一位
        T val=arr[end];
        arr[end]=null;
        count--;
        if(count>0&&count/size<=4)
            resize(size/2);
        return val;
    
    } 


    //查
     // 获取数组头部元素，时间复杂度 O(1)
    public T getFirst() {
        isEmpty();
           
        
        return arr[start];
    }

    // 获取数组尾部元素，时间复杂度 O(1)
    public T getLast() {
        isEmpty();
        // end 是开区间，指向的是下一个元素的位置，所以要减 1
        return arr[(end - 1 + size) % size];
    }
}