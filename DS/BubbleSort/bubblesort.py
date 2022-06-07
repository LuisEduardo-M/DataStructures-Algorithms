def bubble_sort(arr):
    swap = True
    while swap:
        print("list: " + str(arr))
        swap = False

        for i in range(len(arr)-1):
            if arr[i] > arr[i+1]:
                swap = True
                arr[i], arr[i+1] = arr[i+1], arr[i]

list = [6, 8, 1, 4, 10, 7, 8, 9, 3, 2, 5]
bubble_sort(list)

# iteration 0: 6, 8, 1, 4, 10, 7, 8, 9, 3, 2, 5
# iteration 1: 6, 1, 4, 8, 7, 8, 9, 3, 2, 5, 10
# iteration 2: 1, 4, 6, 7, 8, 8, 3, 2, 5, 9, 10
# sorted list: 1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 10
