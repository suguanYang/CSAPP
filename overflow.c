#include<stdio.h>

typedef struct {
  int a[2];
  double d;
} number;

double get_d(int i) {
  volatile number s;
  s.d = 3.14;
  s.a[i] = 1073741824;

  return s.d;
}

int main(void) {
  printf("d: %f", get_d(0)); 
  printf("d: %f", get_d(1)); 
  printf("d: %f", get_d(2)); 
  printf("d: %f", get_d(3)); 
  printf("d: %f", get_d(4)); 
  return 0; 
}
