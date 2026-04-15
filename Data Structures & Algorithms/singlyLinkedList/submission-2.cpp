class LinkedList {
private:
    struct Node {
        int val;
        Node *next;
    };

    Node *head;
    int size = 0;
public:
    LinkedList() {
        head = NULL;
    }

    ~LinkedList() {
        Node *current = head;
        while (current != NULL) {
            Node *nextOne = current->next;
            delete current;
            current = nextOne;
        }
    }

    int get(int index) {
        if (index < size) {
            Node *tmp = head;

            while (index--) {
                tmp = tmp->next;
            }

            return tmp->val;
        }
        return -1;
    }

    void insertHead(int value) {
        Node *newHead= new Node;
        newHead->val = value;
        newHead->next = head;
        head = newHead;

        size++;
    }
    
    void insertTail(int value) {
        if (head == NULL) {
            insertHead(value);
            return;
        }

        Node *tmp = head;
        while (tmp->next != NULL) {
            tmp = tmp->next;
        }
        
        Node *newNode = new Node;
        newNode->val = value;
        newNode->next = NULL;
        tmp->next = newNode;
        size++;
    }

    bool remove(int index) {
        if (index < 0 || index >= size) { // FIX #2: Check bounds strictly
            return false;
        }

        if (index == 0) {
            Node *temp = head;
            head = head->next;
            delete temp;
            size--;
            return true;
        }

        Node *tmp = head;
        index--;  // decrease index by one since we need to get one node prior to the one to be deleted.
        while (index--) {
            tmp = tmp->next;
        }
        Node *delNode = tmp->next;
        tmp->next = tmp->next->next;

        delete delNode;

        size--;
        return true;
    }

    vector<int> getValues() {
        std::vector<int> values;
        Node *tmp = head;
        while (tmp != NULL) {
            values.push_back(tmp->val);
            tmp = tmp->next;
        }

        return values;
    }
};
