//单链表

import java.util.NoSuchElementException;

public class MyLinkedList02<E>{
    int size;
    private Node<E> head,tail;

    private static class Node<E>{
        E val;
        Node<E> next;

        Node(E val){
        this.val=val;
        this.next=null;
    }
    }

    public MyLinkedList02() {
        this.head=new Node<E>(null);//     虚拟节点的命名方式
        tail=head;
        this.size=0;

    }

    //增
    public void addFirst(E e){
        Node<E> newNode=new Node<>(e);
        if(size==0) 
        {
            head.next=newNode;
            tail=newNode;
        }
        else{
            newNode.next=head.next;
            head.next=newNode;
           
        }
        size++;
    }
    public void addLast(E e){
        Node<E> newNode=new Node<>(e);
        if(size==0)
        {
            head.next=newNode;
            tail=newNode;
        }
        else{
            tail.next=newNode;
            tail=newNode;
        }
        size++;
    }
    
    public void add(int index,E e){
        //***************************** *
        //使用getpreNode会有风险，index==0时引用不到
        if(index==0)  addFirst(e);
        //****************************
        Node<E> prev=getPreNode(index);
        Node<E> newNode= new Node<>(e);
        newNode.next=prev.next;
        prev.next=newNode;
        size++;
    }

    public E removeFirst(){
        if(size==0) throw new NoSuchElementException();
        E val=head.next.val;
        head.next=head.next.next;
        size--;
        return val;
    }

    public E removeLast(){
        if(size==0) throw new NoSuchElementException();
        Node<E> prev=getPreNode(size-1);
        
        E val=tail.val;
        prev.next=null;
        tail=prev;
        size--;
        return val;
    }
    
    public E remove(int index){
        if(size==0) throw new NoSuchElementException();
        Node<E> prev=getPreNode(index);
        E val= prev.next.val;
        Node<E> temp=prev.next;
        prev.next=prev.next.next;
        temp.next=null;
        size--;
        return val;
    }

    public E get(int index){
        checkElementIndex(index);
            
        Node<E> prev=getPreNode(index);
        return prev.next.val;
    }

    public void set(int index,E e){
         Node<E> prev=getPreNode(index);
         prev.next.val=e;
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

  

    //5.getNode,取得上一个节点
    public Node<E> getPreNode(int index){
        checkElementIndex(index);
        Node<E> cur = head;
        for (int i = 0; i < index; i++){
            cur = cur.next;
        }
        return cur;
    }
    
}



