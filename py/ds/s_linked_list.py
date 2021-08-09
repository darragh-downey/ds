
class Node:
    def __init__(self, data) -> None:
        self.next = None
        self.data = data


class SLinkedList:
    def __init__(self) -> None:
        self.head = None


    def append(self, data) -> None:
        if self.head is None:
            # empty list make new node self.head
            self.head = Node(data)
            return
        
        last = self.head
        while (last.next):
            # iterate to the end
            last = last.next
        # append new node to the end of the linked list
        last.next = Node(data)
            
    
    def insert_after(self, prev, data) -> None:
        if prev is None:
            # fail early
            print("Node must not be None")
            return

        new_node = Node(data)
        # point new_node.next to self.head reference 
        new_node.next = prev.next
        # point prev.next to new_node reference
        prev.next = new_node        


    def prepend(self, data) -> None:
        new_node = Node(data)
        # point new_node.next to self.head reference
        new_node.next = self.head
        # set self.head to new_node reference
        self.head = new_node


    def delete(self, key) -> None:
        # store the head node
        temp = self.head
        # if head node itself holds the key to be deleted
        if temp is not None:
            if temp.data == key:
                self.head = temp.next
                temp = None
                return
        
        # search for the key to be deleted, keep track of the 
        # previous node as we need to change 'prev.next'
        while temp is not None:
            if temp.data == key:
                break
            prev = temp
            temp = temp.next

        # if data was not present in linked list
        if temp == None:
            return
        
        # unlink the node from the linked list
        prev.next = temp.next
        # to be cleared up by the garbage collector
        temp = None