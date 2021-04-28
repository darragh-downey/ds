

// Ord trait allows any type which implements the Ord trait (<, >, == comparisons)
// to be sorted by selection_sort
pub fn selection_sort<T: Ord>(arr: &mut [T]) {
    // iterate from 0 to array.len()-1
    for i in 0..arr.len() {
        //
        let mut min_idx = i;
        for j in (i + 1)..arr.len() {
            if arr[j] < arr[min_idx] {
                min_idx = j;
            }
        }
        // items swapped using built-in swap method
        arr.swap(min_idx, i)
    }
}

// See References [1] for credit
// list.iter():             [4, 3, 1, 2]
// list.iter().enumerate(): [(0, 4), (1, 3), (2, 1), (3, 2)]
pub fn rusty_selection_sort<T: Ord>(arr: &mut [T]) {
    for i in 0..arr.len() {
        // the if statement pattern matches on Some<Value>. In this case
        // specifically matching the first element of the tuple, which is the 
        // index we want to swap
        if let Some((j, _)) = arr.iter() // provides an iterator over the elements in arr
                                .enumerate() // package up each item in the iterator in a tuple with the item's index
                                .skip(i) // skip the sorted part of the list
                                .min_by_key(|x| x.1) { // takes a closure, allowing you to specify what part of the item you want to find the minimum value in. In this situation we want the element at tuple index 1. (Otherwise you’re finding the minimum by the index in the iterator.) x.1 gives the tuple element at index 1.
            arr.swap(i, j);
        }
    }
}

pub fn insertion_sort<T: Ord>(arr: &mut [T]) {
    for i in 0..arr.len() {
        for j in (1..i + 1).rev() {
            if arr[j - 1] <= arr[j] { break; }
            arr.swap(j-1, j)
        }
    }
}

pub fn top_down_mergesort<T: Ord>(arr: &mut [T]) {}

fn top_down_merge<T: Ord>(arr: &mut [T]) {}

pub fn bottom_up_mergesort<T: Ord>(arr: &mut [T]) {}

fn bottom_up_merge(arr: &mut [u32]) {}


#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    fn test_selection_sort() {
        let mut a: [u32; 5] = [5, 3, 2, 4, 1];
        let mut b: [u32; 5] = [1, 2, 3, 4, 5];
        let mut c: [u32; 5] = [2, 5, 3, 1, 4];
        assert_eq!(crate::alg::selection_sort(&mut a), b);
        assert_eq!(crate::alg::rusty_selection_sort(&mut c), b)
    }

    #[test]
    fn top_down_mergesort() {
        let mut a: [u32; 5] = [5, 3, 2, 4, 1];
        let mut b: [u32; 5] = [1, 2, 3, 4, 5];
        assert_eq!(crate::dsa::alg::top_down_mergesort(&mut a), b);
    }
    #[test]
    fn bottom_up_mergesort() {
        let mut a: [u32; 5] = [5, 3, 2, 4, 1];
        let mut b: [u32; 5] = [1, 2, 3, 4, 5];
        assert_eq!(crate::dsa::alg::bottom_up_mergesort(&mut a), b)
    }
}