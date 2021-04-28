
mod alg {
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

    pub fn bubble_sort<T: Ord>(arr: &mut [T]) {
        for i in 0..arr.len() {
            for j in 0..arr.len()-i-1 {
                if arr[j] > arr[j+1] {
                    arr.swap(j, j+1);
                }
            }
        }
    }

    pub fn rusty_bubble_sort<T: Ord>(arr: &mut [T]) {
        unimplemented!()
    }

    pub fn insertion_sort<T: Ord>(arr: &mut [T]) {
        for i in 0..arr.len() {
            for j in (1..i + 1).rev() {
                if arr[j - 1] <= arr[j] { break; }
                arr.swap(j-1, j)
            }
        }
    }

    pub fn rusty_insertion_sort<T: Ord>(arr: &mut [T]) {
        unimplemented!()
    }

    pub fn top_down_mergesort<T: Ord + Copy>(arr: &mut [T]) {
        let n = arr.len();
        let mid = n / 2;

        if n <= 1 {
            return;
        }

        top_down_mergesort(&mut arr[0..mid]);
        top_down_mergesort(&mut arr[mid..n]);

        let mut y: Vec<T> = arr.to_vec();
        
        top_down_merge(&arr[0..mid], &arr[mid..n], &mut y[..]);

        arr.copy_from_slice(&y);
    }

    fn top_down_merge<T: PartialOrd + Copy>(temp_left: &[T], temp_right: &[T], arr: &mut [T]) {
        let mut i = 0;
        let mut j = 0;
        let mut k = 0;

        while i < temp_left.len() && j < temp_right.len() {
            if temp_left[i] <= temp_right[j] {
                arr[k] = temp_left[i];
                i += 1;
            } else {
                arr[k] = temp_right[j];
                j += 1;
            }
            k += 1;
        }

        while i < temp_left.len() {
            arr[k] = temp_left[i];
            i += 1;
            k += 1;
        }

        while j < temp_right.len() {
            arr[k] = temp_right[j];
            j += 1;
            k += 1;
        }
    }

    pub fn bottom_up_mergesort<T: Ord>(arr: &mut [T]) {
        unimplemented!()
    }

    fn bottom_up_merge(arr: &mut [u32]) {
        unimplemented!()
    }

    pub fn is_sorted<T: PartialOrd>(arr: &mut [T]) -> bool {
        for i in 1..arr.len() {
            if arr[i-1] > arr[i] {
                return false;
            }
        }
        return true;
    }
}

#[cfg(test)]
mod tests {

    #[test]
    fn test_selection_sort() {
        let mut a: [u32; 5] = [5, 3, 2, 4, 1];
        crate::srt::alg::selection_sort(&mut a);
        assert!(crate::srt::alg::is_sorted(&mut a));

        let mut c: [u32; 5] = [2, 5, 3, 1, 4];
        crate::srt::alg::rusty_selection_sort(&mut c);
        assert!(crate::srt::alg::is_sorted(&mut c));
    }

    #[test]
    fn test_bubble_sort() {
        let mut a: [u32; 5] = [5, 3, 2, 4, 1];
        
        crate::srt::alg::bubble_sort(&mut a);
        assert!(crate::srt::alg::is_sorted(&mut a));

        //crate::srt::alg::rusty_bubble_sort(&mut c);
        //assert_eq!(c, b)
    }

    #[test]
    fn test_insertion_sort() {
        let mut a: [u32; 5] = [5, 3, 2, 4, 1];
        
        crate::srt::alg::insertion_sort(&mut a);
        assert!(crate::srt::alg::is_sorted(&mut a));
        
        //crate::srt::alg::rusty_insertion_sort(&mut c);
        //assert!(crate::srt::alg::is_sorted(&mut c));
    }

    #[test]
    fn test_mergesort() {
        let mut a: [u32; 5] = [5, 3, 2, 4, 1];

        crate::srt::alg::top_down_mergesort(&mut a);
        assert!(crate::srt::alg::is_sorted(&mut a));

        let mut strings = ["beach", "hotel", "airplane", "car", "house", "art"];
        crate::srt::alg::top_down_mergesort(&mut strings);
        assert!(crate::srt::alg::is_sorted(&mut strings));
    }

    #[test]
    fn test_quicksort() {
        let mut a: [u32; 5] = [5, 3, 2, 4, 1];
        let b: [u32; 5] = [1, 2, 3, 4, 5];
        
    }
}