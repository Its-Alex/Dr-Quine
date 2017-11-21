#include <stdio.h>
//commentaire1
char *getA() {return "#include <stdio.h>%c//commentaire1%cchar* getA() {return %c%s%c;}%cint main(){%c//commentaire2%cprintf(getA(), 10, 10, 10, 34, getA(), 34, 10, 10, 10);%c}";}
int main(){
//commentaire2
printf(getA(), 10, 10, 34, getA(), 34, 10, 10, 10, 10);
}