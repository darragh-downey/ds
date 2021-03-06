
// Derived from original BST implementation:
// https://github.com/dogHere/rust-bst/blob/master/src/tree.rs

mod ds {
    use std::mem::swap;

    // Option https://doc.rust-lang.org/std/option/index.html is an optional value
    // Every Option is either Some and contains a value, or None, and does not.
    // Box https://doc.rust-lang.org/std/boxed/index.html creates Node struct on the heap
    type Child<T> = Option<Box<Node<T>>>;

    #[derive(Debug, Eq, PartialEq, PartialOrd, Clone)]
    pub struct Node<T> {
        pub value: T,
        pub left: Child<T>,
        pub right: Child<T>,
    }

    impl<T: PartialOrd + Clone> Node<T> {
        fn new(value: T) -> Node<T> {
            Node {
                value,
                left: None,
                right: None,
            }
        }

        fn value(&self) -> T {
            self.value.clone()
        }

        fn push(&mut self, value: T) {
            if value > self.value {
                match self.right {
                    None => {
                        swap(&mut self.right, &mut Some(Box::from(Node::new(value))));
                    }
                    Some(ref mut n) => n.push(value),
                }
            } else if value < self.value {
                match self.left {
                    None => {
                        swap(&mut self.left, &mut Some(Box::from(Node::new(value))));
                    }
                    Some(ref mut n) => n.push(value),
                }
            }
        }

        fn min(&self) -> T {
            match self.left {
                None => self.value.clone(),
                Some(ref n) => Node::min(&*n),
            }
        }

        fn max(&self) -> T {
            match self.right {
                None => self.value.clone(),
                Some(ref n) => Node::max(&*n),
            }
        }

        fn find(&self, value: &T) -> Option<Box<&Node<T>>> {
            if value > &self.value {
                match self.right {
                    None => None,
                    Some(ref n) => Node::find(&*n, &value),
                }
            } else if value < &self.value {
                match self.left {
                    None => None,
                    Some(ref n) => Node::find(&*n, &value),
                }
            } else {
                Some(Box::from(self))
            }
        }

        fn delete(node: &mut Child<T>, value: &T) {
            let mut all_none = false;
            match *node {
                None => {}
                Some(ref mut r) => {
                    if !r.left.is_some() && !r.right.is_some() {
                        all_none = true;
                    } else {
                        if value > &r.value() {
                            Node::delete(&mut r.right, value);
                        } else if value < &r.value() {
                            Node::delete(&mut r.left, value);
                        } else {
                            if r.right.is_none() {
                                let t = r.left.take();
                                swap(&mut r.value, &mut t.unwrap().value);
                                swap(&mut None, &mut r.left);
                            } else if r.left.is_none() {
                                let t = r.right.take();
                                swap(&mut r.value, &mut t.unwrap().value);
                                swap(&mut None, &mut r.right);
                            } else {
                                let m = r.right.take().unwrap().min();
                                swap(&mut None, &mut r.find(&m));
                                swap(&mut None, &mut r.right);
                            }
                        }
                    }
                }
            }
            if all_none {
                swap(node, &mut None);
            }
        }
    }

    #[derive(Debug, PartialEq, PartialOrd, Clone)]
    pub struct Tree<T> {
        root: Child<T>,
    }

    impl<T: PartialOrd + Clone> Tree<T> {
        pub fn new() -> Tree<T> {
            Tree { root: None }
        }

        pub fn push(&mut self, value: T) -> Box<&mut Tree<T>> {
            match self.root {
                None => {
                    swap(&mut self.root, &mut Some(Box::from(Node::new(value))));
                }
                Some(ref mut n) => {
                    n.push(value);
                }
            }
            Box::new(self)
        }

        pub fn min(&self) -> Option<T> {
            match self.root {
                None => None,
                Some(ref n) => Some(n.min()),
            }
        }

        pub fn max(&self) -> Option<T> {
            match self.root {
                None => None,
                Some(ref n) => Some(n.max()),
            }
        }

        pub fn find(&self, value: &T) -> Option<Box<&Node<T>>> {
            match self.root {
                None => None,
                Some(ref n) => n.find(value),
            }
        }

        pub fn exists(&self, value: &T) -> bool {
            match self.find(value) {
                None => false,
                _ => true,
            }
        }

        pub fn delete(&mut self, value: &T) -> Box<&mut Tree<T>> {
            Node::delete(&mut self.root, value);
            Box::new(self)
        }
    }
}
