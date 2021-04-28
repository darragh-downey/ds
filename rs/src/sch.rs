mod alg {
    use std::convert::TryFrom;

    pub fn linear_search<T: Ord>(arr: &mut [T], key: T) -> Option<usize> {
        for i in 0..arr.len() {
            if arr[i] == key {
                return Some(i);
            }
        }
        return None;
    }

    pub fn binary_search<T: Ord>(arr: &mut [T], key: T) -> Option<usize> {
        let mut l = 0;
        let mut r = arr.len() - 1;

        while l <= r {
            let m = l + (r - l) / 2;
            if arr[m] == key {
                return Some(m);
            }

            if arr[m] < key {
                // key is larger, ignore left sub-array
                l = m + 1;
            } else {
                // key is smaller, ignore right sub-array
                r = m - 1;
            }
        }
        return None;
    }

    pub fn recursive_binary_search<T: Ord + Copy>(arr: &mut [T], left: u32, right: u32, key: T) -> Option<usize> {
        if right >= left {
            // https://stackoverflow.com/questions/43704758/how-to-idiomatically-convert-between-u32-and-usize/55769098#55769098 
            let mid = left + (right - left) / 2;
            let l_mid = mid - 1;
            let r_mid = mid + 1;
            let mid = usize::try_from(mid).unwrap();

            if key == arr[mid] {
                return Some(mid);
            }

            if arr[mid] > key {
                return recursive_binary_search(arr, left, l_mid, key)
            }

            return recursive_binary_search(arr, r_mid, right, key)

        }
        None
    }
}

#[cfg(test)]
mod test {
    use std::convert::TryInto;

    #[test]
    fn test_linear_search() {
        let mut a: [u32; 5] = [5, 3, 2, 4, 1];
        let idx_1 = crate::sch::alg::linear_search(&mut a, 5);
        let idx_2 = crate::sch::alg::linear_search(&mut a, 4);
        let idx_3 = crate::sch::alg::linear_search(&mut a, 7);

        assert_eq!(idx_1, Some(0));
        assert_eq!(idx_2, Some(3));
        assert_eq!(idx_3, None);
    }

    #[test]
    fn test_binary_search() {
        let mut a: [i32; 10] = [4210, 245, -21, 204, -1234, 24, 45, -195, 1, 0];
        let idx_1 = crate::sch::alg::binary_search(&mut a, 1);
        let idx_2 = crate::sch::alg::binary_search(&mut a, 135);
        assert_eq!(idx_1, Some(8));
        assert_eq!(idx_2, None);
    }

    #[test]
    fn test_recursive_binary_search() {
        let mut a: [i32; 10] = [4210, 245, -21, 204, -1234, 24, 45, -195, 1, 0];
        let idx_1 = crate::sch::alg::recursive_binary_search(&mut a, 0, 9, 1);
        let idx_2 = crate::sch::alg::recursive_binary_search(&mut a, 0, 9, 135);
        assert_eq!(idx_1, Some(8));
        assert_eq!(idx_2, None);
    }
}
