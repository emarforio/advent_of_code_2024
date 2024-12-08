#include <stdio.h>
#include <stdlib.h>

typedef struct list list;
struct list {
  int data;
  list *next;
};

// Clear list
void clear_list(list *head) {
  list *traverser = head;
  list *next_item = head;

  while (traverser != NULL) {
    next_item = traverser->next;
    free(traverser);
    traverser = next_item;
  }
}

// void append_to_list(list *head, int data) {
//   list *traverser = head;
//   list *next_item = head;
//
//   // Handle if list is empty
//   if (head == NULL) {
//
//   }
//
//   while (next_item != NULL) {
//     next_item = traverser->next;
//     traverser = next_item;
//   }
// }

int main() {
  FILE *fp;

}
