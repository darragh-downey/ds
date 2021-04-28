"""
See https://www.geeksforgeeks.org/binary-heap/


"""

import math
import sys
from heapq import heappush, heappop, heapify

class MinHeap:

    def __init__(self) -> None:
        self.heap = []

    
    def __str__(self) -> str:
        return str(self.heap)
    

    def parent(self, idx) -> int:
        return math.floor((idx-1)/2)


    def insert_key(self, key) -> None:
        heappush(self.heap, key)

    
    def decrease_key(self, idx: int, new_value: int) -> None:
        self.heap[idx] = new_value
        while idx != 0 and self.heap[self.parent(idx)] > self.heap[idx]:
            self.heap[idx], self.heap[self.parent(idx)] = (self.heap[self.parent(idx)], self.heap[idx])


    def extract_min(self) -> int:
        return heappop(self.heap)

    
    def delete_key(self, idx) -> None:
        self.decrease_key(idx, float("-inf"))
        self.extract_min()

    
    def get_min(self) -> int:
        return self.heap[0]


class MaxHeap:
    """
    A Max Heap is a complete binary tree. A Max Heap is typically represented as an array.
    The root element will be at Arr[0]. Below shows indexes of other nodes for the i-th node,
    i.e., Arr[i]:
    Arr[(i-1)/2] Returns the parent node
    Arr[(2*i)+1] Returns the left child node
    Arr[(2*i)+2] Returns the right child node
    """
    
    def __init__(self, max_size) -> None:
        self.max_size = max_size
        self.size = 0
        self.heap = [0] * (self.max_size + 1)
        self.heap[0] = sys.maxsize
        self.FRONT = 1

    
    def __str__(self) -> str:
        return str(self.heap)
    

    def parent(self, idx) -> int:
        return math.floor((idx)//2)

    
    def left_child(self, idx):
        return 2 * idx


    def right_child(self, idx):
        return (2 * idx) + 1
        

    def is_leaf(self, idx):
        if idx >= (self.size // 2) and idx <= self.size:
            return True
        return False


    def swap(self, fpos, spos):
        self.heap[fpos], self.heap[spos] = (self.heap[spos], self.heap[fpos])


    def max_heapify(self, idx):
        if not self.is_leaf(idx):
            if (self.heap[idx] < self.heap[self.left_child(idx)] or 
            self.heap[idx] < self.heap[self.right_child(idx)]):

                if (self.heap[self.left_child(idx)] > self.heap[self.right_child(idx)]):
                    self.swap(idx, self.left_child(idx))
                    self.max_heapify(self.left_child(idx))
                else:
                    self.swap(idx, self.right_child(idx))
                    self.max_heapify(self.right_child(idx))


    def insert_key(self, key) -> None:
        if self.size >= self.max_size:
            return
        
        self.size += 1
        self.heap[self.size] = key

        current = self.size

        while self.heap[current] > self.heap[self.parent(current)]:
            self.swap(current, self.parent(current))
            current = self.parent(current)

    
    def decrease_key(self, idx: int, new_value: int) -> None:
        self.heap[idx] = new_value
        while idx != 0 and self.heap[self.parent(idx)] < self.heap[idx]:
            self.heap[idx], self.heap[self.parent(idx)] = (self.heap[self.parent(idx)], self.heap[idx])


    def extract_max(self) -> int:
        popped = self.heap[self.FRONT]
        self.heap[self.FRONT] = self.heap[self.size]
        self.size -= 1
        self.max_heapify(self.FRONT)

        return popped

    
    def delete_key(self, idx) -> None:
        self.decrease_key(idx, float("-inf"))
        self.extract_min()

    
    def get_max(self) -> int:
        return self.heap[0]


def main():
    min_heap = MinHeap()
    min_heap.insert_key(3)
    min_heap.insert_key(2)
    min_heap.insert_key(15)
    min_heap.insert_key(5)
    min_heap.insert_key(4)
    min_heap.insert_key(45)
    
    print(min_heap.extract_min())
    print(min_heap.get_min())
    min_heap.decrease_key(2, 1)
    print(min_heap.get_min())

    print(min_heap)


if __name__ == '__main__':
    main()