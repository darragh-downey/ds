
class DNode:
    def __init__(self, data) -> None:
        self.prev = None
        self.next = None
        self.data = data


class DLinkedList:
    def __init__(self) -> None:
        self.head = None


    def prepend(self, data) -> None:
        if self.head is None:
            self.head = DNode(data)
            return
        
        new_node = DNode(data)
        self.head.prev = new_node
        new_node.next = self.head
        self.head = self.head.prev

    def insertAfter(self, prev, data) -> None:
        if prev is None:
            

        
