# Python3 code to move all zeroes
# at the end of array
 
# Function which pushes all
# zeros to end of an array.
def pushZerosToEnd(arr, n):
    count = 0 # Count of non-zero elements
     
    # Traverse the array. If element 
    # encountered is non-zero, then
    # replace the element at index
    # 'count' with this element
    # print(arr)
    #数一下列表中的非零元素，原来in-place方式 数某类元素的个数是这么操作的
    for i in range(n):
        if arr[i] != 0:
             
            # here count is incremented
            arr[count] = arr[i]
            count+=1
    print(arr)
    # Now all non-zero elements have been
    # shifted to front and 'count' is set
    # as index of first 0. Make all 
    # elements 0 from count to end.
    while count < n:
        arr[count] = 0
        # print(arr)
        count += 1

'''
真的好强，这群人，简简单单的几行代码就解决了这个问题。而且我看了半天都没看懂，似乎看懂了，但是仔细一想，还是没懂...

'''        
# Driver code
arr = [1, 9, 8, 4, 0, 0, 2, 7, 0, 6, 0, 9]
n = len(arr)
pushZerosToEnd(arr, n)
print("Array after pushing all zeros to end of array:")
print(arr)