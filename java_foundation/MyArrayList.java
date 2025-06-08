
import java.util.NoSuchElementException;

public class MyArrayList<E>{
    //真正存储数据的底层数组
    private E[] data;
    //记录元素的个数
    private int size;
    //默认初始容量
    private static final int INIT_CAP=1;

    //初始化构造器
    //无参
    public MyArrayList()
    {
        this(INIT_CAP);
    }
    //带参
    public MyArrayList(int initCapacity)
    {
        data=(E[])new Object[initCapacity];//创建一个指定容量的Object数组，并强制类型转换为E[]，用于存储数据。
        size=0; 
    }

//增

//在尾部插入元素
    public void addLast(E e){
        int cap=data.length;
        if(size==cap)
        {
            resize(2*cap);
        }
        data[size]=e;
        size++;
    }

//在任意位置插入
    public void add(int index,E e){
        //先检查index是否超出范围index<=size,相当于在元素间隙中插入值
        checkPositionIndex(index); // 直接检查，非法则抛异常
    
        // 先扩容检查
        if (size == data.length) {
            resize(2 * data.length);
        }
        
        // 统一移动元素：从后向前移位
        for (int i = size; i > index; i--) {
            data[i] = data[i - 1];
        }
        
        data[index] = e;
        size++;

    }
//在头部插入
    public void addFirst(E e){
        add(0,e);
    }
    

//删
    //删除指定索引元素
    public E remove(int index){
       // 检查索引越界
        checkElementIndex(index);

        int cap = data.length;
        // 可以缩容，节约空间
      

        E deletedVal = data[index];

        // 搬移数据 data[index+1..] -> data[index..]
        for (int i = index + 1; i < size; i++) {
            data[i - 1] = data[i];
        }

        data[size - 1] = null;
        size--;
        if (size <= cap / 4) {
            resize(cap / 2);
        }


        return deletedVal;
      
    }
    //删除最后的元素
    public E remove(){
        if(size==0){
            throw new NoSuchElementException();}
        int cap=data.length;
      
       E dely=data[size-1];
       data[size-1]=null;
       size--;
       if(size<=(cap/4))
        {
            resize(cap/2);
        }
       return dely;
    }

    //删除第一个
    public E removeFirst(){
         return remove(0);
    }

//查
    //查找指定索引
    public E get(int index){
        checkElementIndex(index);
            return data[index];
        
    }

//改: 要返回老数据
    public E set(int index,E e){
        checkElementIndex(index);

        E olddata=data[index];
        data[index]=e;
        return olddata;
        
    }


//工具+方法
    //size
 public int size() {
        return size;
    }
    //isEmprty
public boolean isEmpty() {
    return size == 0;
    }


    //扩容+缩容
    private void resize(int newCap)
    {
        //先开辟一片指定大小的内存，类为Object转化为E
        E[] temp=(E[])new Object[newCap];
        for(int j=0;j<size;j++)//范围要看好0-size 不是到newcap
        {
            temp[j]=data[j];
        }
        data=temp;
    }
    //检查元素间隙的数量，是否可以插入元素，增使用(0<=index<=size)
    private void checkPositionIndex(int index){
        if(index<0||index>size)
            throw new IndexOutOfBoundsException("Index " + index);
       
    }
    //检查索引处是否存在元素，只看索引，删，查，改使用(0<=index<size)
     private void checkElementIndex(int index){
         if(index<0||index>=size)
            throw new IndexOutOfBoundsException("Index " + index);
    }

}