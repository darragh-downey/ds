# See https://www.geeksforgeeks.org/binary-search-tree-set-1-search-and-insertion/

class Node:

    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None


def insert(root, key):
    if root is None:
        return Node(key)
    
    if key < root.value:
        root.left = insert(root.left, key)
    else:
        root.right = insert(root.right, key)
            
    return root


def minValueNode(node):
    current = node

    while current.left is not None:
        current = current.left
    
    return current


def del_node(root, key):
    """
    Given a binary search tree and a key, this function deletes the 
    key and returns the new root
    """
    # Base case
    if root is None:
        return root
    
    # Recursive calls for ancestors of node to be deleted
    if key < root.value:
        root.left = del_node(root.left, key)
        return root
    elif key > root.value:
        root.right = del_node(root.right, key)
        return root

    # We reach this point when the root is the node to be deleted

    # If root node is a leaf node
    if root.left is None and root.right is None:
        return None
    
    # If a single child is empty
    if root.left is None:
        temp = root.right
        root = None
        return temp
    elif root.right is None:
        temp = root.left
        root = None
        return temp
    
    # Node with two children
    # both children exist
    successor_Parent = root

    # find successor
    successor = root.right

    while successor.left != None:
        successor_Parent = successor
        successor = successor.left
    
    # delete successor. Since successor is always the left child
    # of its parent we can safely make successor's right right child 
    # as left of its parent. If there is no succ, then assign succ->right
    # to succParent->right
    if successor_Parent != root:
        successor_Parent.left = successor.right
    else:
        successor_Parent.right = successor.right

    # copy successor data to root
    root.value = successor.value
    
    return root


def inorder(root):
    if root is not None:
        inorder(root.left)
        print(root.value)
        inorder(root.right)


def search(root, key):
    if root is None or root.value == key:
        return root
    
    if root.value < key:
        return search(root.right, key)
    
    return search(root.left, key)




def main():
    r = None
    r = insert(r, 50)
    r = insert(r, 30)
    r = insert(r, 20)
    r = insert(r, 40)
    r = insert(r, 70)
    r = insert(r, 60)
    r = insert(r, 80)
 
    # Print inoder traversal of the BST
    inorder(r)

    print("Delete 20")
    r = del_node(r, 20)
    inorder(r)

    print("Delete 30")
    r = del_node(r, 30)
    inorder(r)

    print("Delete 50")
    r = del_node(r, 50)
    inorder(r)


if __name__ == '__main__':
    main()