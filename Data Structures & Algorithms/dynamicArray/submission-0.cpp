class DynamicArray {
private:
    int capacity = 0;
    int size = 0;
    int *arr;
public:

    DynamicArray(int capacity) {
        if (capacity > 0) {
            this->capacity = capacity;
        }
        else {
            this->capacity = 1;
        }
        arr = new int[this->capacity];
    }

    int get(int i) {
        return arr[i];
    }

    void set(int i, int n) {
        arr[i] = n;
    }

    void pushback(int n) {
        if (size == capacity) {
            resize();
        }

        arr[size] = n; 
        size++;
    }

    int popback() {
        int last = arr[size-1];
        size--;
        return last;
    }

    void resize() {
        capacity *= 2;
        int *newArr = new int[capacity];

        for (int i=0; i<size; i++) {
            newArr[i] = arr[i];
        }
        delete[] arr;
        arr = newArr;
    }

    int getSize() {
        return size;
    }

    int getCapacity() {
        return capacity;
    }

    ~DynamicArray() {
        delete[] arr;  // Free the memory we borrowed from the heap.
    }
};
