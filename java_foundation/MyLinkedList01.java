//双向链表
import java.util.NoSuchElementException;



public class MyLinkedList01<E>{
    //先定义虚拟头尾节点，final保证引用不会变
     final private Node<E> head, tail;
     //元素个数
     int size;

    //内部静态类 1.why?(语雀)
    private static class Node<E>{
        E val;
        Node<E> next;
        Node<E> prev;

        Node(E val){// 构造方法，初始化节点值
            this.val=val;
        }
    }
    
    public MyLinkedList(){// 构造方法，初始化节点值
        this.head=new Node<>(null);//创建头哨兵节点（不存在）
        this.tail=new Node<>(null);//创建尾哨兵节点（不存在）
        head.next=tail;//头的下一个是尾
        tail.prev=head;//尾的上一个是头
        this.size=0;
    }

    //*******增******* */

    //头前增
    public void addFirst(E e)
    {
        Node<E> newNode=new Node<>(e);
        Node<E> cur=head;
        newNode.next=cur.next;
        newNode.prev=cur;
        cur.next=newNode;
        newNode.next.prev=newNode;
        size++;
    }

    //尾增
    public void addLast(E e){
        Node<E> newNode=new Node<E>(e);
        Node<E> end=tail.prev;
        end.next=newNode;
        newNode.next=tail;
        tail.prev=newNode;
        newNode.prev=end;
        size++;
    }

    //中间增
    public void add(int index,E e){
        //太过繁琐，分成两种情况，一种在index前，一种在index后，其实可以都在前，减少代码的数量
        // checkPositionIndex(index);
        
        // Node<E> newNode=new Node<>(e);
        // if(distance(index))
        // {   Node<E> cur=head.next;
        //    for(int i=0;i<index-1;i++)
        //    {
        //         cur=cur.next;
        //    }
        //    newNode.prev=cur;
        //    newNode.next=cur.next;
        //    cur.next.prev=newNode;
        //    cur.next=newNode;     
        // }
        // else{
        //     Node<E> cur1=tail.prev;
        //    for(int i=size-1;i>index;i--)
        //    {    
        //         cur1=cur1.prev;
        //    }
        //       newNode.next=cur1;
        //       cur1.prev.next=newNode;
        //       newNode.prev=cur1.prev;
        //       cur1.prev=newNode;           
        // }
        // size++;
        Node<E> newNode=new Node<>(e);
        Node<E> cur;
        if(distance(index)){
            cur=head;
            for(int i=0;i<index;i++)
            {
                cur=cur.next;
            }          
        }
        else{
            cur=tail;
            for(int i=size;i>index;i--)//从尾遍历到index节点
            {
                cur=cur.prev;
            }
            //获取index前一个节点
            cur=cur.prev;
        }
        //插入节点
        newNode.next=cur.next;
        cur.next.prev=newNode;
        cur.next=newNode;
        newNode.prev=cur;
        size++;
    }



    //删
    //头删
    public E removeFirst(){
        if(size<1)
            throw new NoSuchElementException();
        Node<E> cur=head.next;
        head.next=cur.next;
        cur.next.prev=head;
        E val=cur.val;
        cur.next=null;
        cur.prev=null;
        size--;
        return val;
    }

    //尾删
    public E removeLast(){
        if(size<1)
            throw new NoSuchElementException();
        Node<E> cur=tail.prev;
        Node<E> temp=cur.prev;
        temp.next=tail;
        tail.prev=temp;
        E val=cur.val;
        cur.next=null;
        cur.prev=null;
        size--;
        return val;
    }
    //中间删
    public E remove(int index){
         checkElementIndex(index);
        Node<E> cur;
        if(distance(index)){
            cur=head;
            for(int i=0;i<index;i++)
            {
                cur=cur.next;
            }          
        }
        else{
            cur=tail;
            for(int i=size;i>index;i--)//从尾遍历到index节点
            {
                cur=cur.prev;
            }
            //获取index前一个节点
            cur=cur.prev;
        }
        //删除节点
        Node<E> temp=cur.next;
        cur.next=temp.next;
        temp.prev=cur;
        E val=temp.val;
        temp.next=null;
        temp.prev=null;
        size--;  
        return val;
    }

    //查

    public E get(int index){
        checkElementIndex(index);
        Node<E> temp =getNode(index);
        return temp.val;
    }


    //改
     public E set(int index,E e){
        checkElementIndex(index);
        Node<E> temp =getNode(index);
        E oldval=temp.val;
        temp.val=e;
        return oldval;
    }


    //方法
    //1.size
    public int size(){
        return size;
    }
    //2.isEmprty
    public boolean isEmpty(){
        return size==0;
    }

    //3.检查索引的有效性
    public void checkElementIndex(int index){
        if(index<0||index>=size)//有效下标的范围是[0,size-1]
        throw new IndexOutOfBoundsException("Index"+index);
    }

    public void checkPositionIndex(int index){
        if(index<0||index>size)//插入的范围是[0,size]
        throw new IndexOutOfBoundsException("Index"+index);
    }

    //4.判断index距头尾节点的距离,true从头
    public boolean distance(int index){
        //return index<=size/2;
        if(index<=size/2) return true;
        return false;
        
    }

    //5.getNode,取得节点
    public Node<E> getNode(int index){
        checkElementIndex(index);
        Node<E> cur;
        if(distance(index)){
             cur =head;
            for (int i = 0; i <= index; i++){
                cur=cur.next;
            }
        }
        else{
             cur=tail;
            for (int i = size; i> index; i++){
                cur=cur.prev;
            }
      
        }
         return cur;  }
    
    }



