#include <stdio.h>
#define ONE() int main() { char *a = "#include <stdio.h>%1$c#define ONE() int main() { char *a = %2$c%3$s%2$c; FILE *f; f = fopen(%2$cGrace_kid.c%2$c, %2$cw+%2$c); if (f) { fprintf(f, a, 10, 34, a); fclose(f); }}%1$c#define TWO 1%1$c#define THREE 2%1$c/*%1$c** FOUR%1$c*/%1$cONE()%1$c"; FILE *f; f = fopen("Grace_kid.c", "w+"); if (f) { fprintf(f, a, 10, 34, a); fclose(f); }}
#define TWO 1
#define THREE 2
/*
** FOUR
*/
ONE()
