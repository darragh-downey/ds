import ctypes

class Node:
    def __init__(self, data) -> None:
        self.data = data
        # XOR of next and previous nodes
        self.npx = 0


class XorLinkedList:
    def __init__(self) -> None:
        self.head = None
        self.__nodes = []

    # returns the XORed value of the node addresses
    def XOR(self, a, b):
        return a ^ b

    def insert(self, data):
        node = Node(data)

        node.npx = id(self.head)

        if self.head is not None:
            # head.npx
            self.head.npx = self.XOR(id(node), self.head.npx)

        self.__nodes.append(node)

        # change head
        self.head = node