#include<stdio.h>

// Casting pointer to unsigned char * allows treatment as a byte array
typedef unsigned char *pointer;

void show_bytes(pointer start, size_t len) {
  for (size_t i = 0; i < len; i++)
    printf("%p\t0x%.2x\n", start + i, start[i]);
  printf("/n");
}

int main() {
  double a = 1.1;
  show_bytes((pointer) &a, sizeof(a));

  return 0;
}
