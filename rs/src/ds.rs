use std::mem::swap;


pub mod ds {
    type Child<T> = Option<Box<Node<T>>>;

    #[derive(Debug, Eq, PartialEq, PartialOrd, Clone)]
    struct Node<T> {
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
                        std::mem::swap(&mut self.right, &mut Some(Box::from(Node::new(value))));
                    }
                    Some(ref mut n) => {
                        n.push(value)
                    }
                }
            } else if value < self.value {
                match self.left {
                    None => {
                        std::mem::swap(&mut self.left, &mut Some(Box::from(Node::new(value))));
                    }
                    Some(ref mut n) => {
                        n.push(value)
                    }
                }
            }
        }

        fn min(&self) -> T {
            match self.left {
                None => self.value.clone(),
                Some(ref n) => Node::min(&*n)
            }
        }

        fn max(&self) -> T {
            match self.right {
                None => self.value.clone(),
                Some(ref n) => Node::max(&*n)
            }
        }

        fn find(&self, value: &T) -> Option<Box<&Node<T>>> {
            if value > &self.value {
                match self.right {
                    None => None,
                    Some(ref n) => Node::find(&*n, &value)
                }
            } else if value < &self.value {
                match self.left {
                    None => None,
                    Some(ref n) => Node::find(&*n, &value)
                }
            }else {
                Some(Box::from(self))
            }
        }

    }
}